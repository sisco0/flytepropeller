package task

import (
	"context"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/catalog"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/io"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/ioutils"

	"github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"github.com/flyteorg/flytepropeller/pkg/controller/nodes/handler"
	errors2 "github.com/flyteorg/flytepropeller/pkg/controller/nodes/errors"

	"github.com/flyteorg/flytestdlib/contextutils"
	"github.com/flyteorg/flytestdlib/logger"
)

func (t *Handler) GetCatalogKey(ctx context.Context, nCtx handler.NodeExecutionContext) (catalog.Key, error) {
	// read task template
	taskTemplatePath, err := ioutils.GetTaskTemplatePath(ctx, nCtx.DataStore(), nCtx.NodeStatus().GetDataDir())
	if err != nil {
		return catalog.Key{}, err
	}

	taskReader := ioutils.NewLazyUploadingTaskReader(nCtx.TaskReader(), taskTemplatePath, nCtx.DataStore())
	taskTemplate, err := taskReader.Read(ctx)
	if err != nil {
		logger.Errorf(ctx, "failed to read TaskTemplate, error :%s", err.Error())
		return catalog.Key{}, err
	}

	return catalog.Key{
		Identifier:     *taskTemplate.Id,
		CacheVersion:   taskTemplate.Metadata.DiscoveryVersion,
		TypedInterface: *taskTemplate.Interface,
		InputReader:    nCtx.InputReader(),
	}, nil
}

func (t *Handler) IsCacheable(ctx context.Context, nCtx handler.NodeExecutionContext) (bool, error) {
	// check if plugin has caching disabled
	ttype := nCtx.TaskReader().GetTaskType()
	ctx = contextutils.WithTaskType(ctx, ttype)
	p, err := t.ResolvePlugin(ctx, ttype, nCtx.ExecutionContext().GetExecutionConfig())
	if err != nil {
		return false, errors2.Wrapf(errors2.UnsupportedTaskTypeError, nCtx.NodeID(), err, "unable to resolve plugin")
	}

	checkCatalog := !p.GetProperties().DisableNodeLevelCaching
	if !checkCatalog {
		logger.Infof(ctx, "Node level caching is disabled. Skipping catalog read.")
	}

	// read task template
	taskTemplatePath, err := ioutils.GetTaskTemplatePath(ctx, nCtx.DataStore(), nCtx.NodeStatus().GetDataDir())
	if err != nil {
		return false, err
	}

	taskReader := ioutils.NewLazyUploadingTaskReader(nCtx.TaskReader(), taskTemplatePath, nCtx.DataStore())
	taskTemplate, err := taskReader.Read(ctx)
	if err != nil {
		logger.Errorf(ctx, "failed to read TaskTemplate, error :%s", err.Error())
		return false, err
	}

	return taskTemplate.Metadata.Discoverable, nil
}

func (t *Handler) IsCacheSerializable(ctx context.Context, nCtx handler.NodeExecutionContext) (bool, error) {
	// check if plugin has caching disabled
	ttype := nCtx.TaskReader().GetTaskType()
	ctx = contextutils.WithTaskType(ctx, ttype)
	p, err := t.ResolvePlugin(ctx, ttype, nCtx.ExecutionContext().GetExecutionConfig())
	if err != nil {
		return false, errors2.Wrapf(errors2.UnsupportedTaskTypeError, nCtx.NodeID(), err, "unable to resolve plugin")
	}

	checkCatalog := !p.GetProperties().DisableNodeLevelCaching
	if !checkCatalog {
		logger.Infof(ctx, "Node level caching is disabled. Skipping catalog read.")
	}

	// read task template
	taskTemplatePath, err := ioutils.GetTaskTemplatePath(ctx, nCtx.DataStore(), nCtx.NodeStatus().GetDataDir())
	if err != nil {
		return false, err
	}

	taskReader := ioutils.NewLazyUploadingTaskReader(nCtx.TaskReader(), taskTemplatePath, nCtx.DataStore())
	taskTemplate, err := taskReader.Read(ctx)
	if err != nil {
		logger.Errorf(ctx, "failed to read TaskTemplate, error :%s", err.Error())
		return false, err
	}

	return taskTemplate.Metadata.Discoverable && taskTemplate.Metadata.CacheSerializable, nil
}

func (t *Handler) ValidateOutput(ctx context.Context, nodeID v1alpha1.NodeID, i io.InputReader,
	r io.OutputReader, outputCommitter io.OutputWriter, executionConfig v1alpha1.ExecutionConfig,
	tr ioutils.SimpleTaskReader) (*io.ExecutionError, error) {

	tk, err := tr.Read(ctx)
	if err != nil {
		logger.Errorf(ctx, "Failed to read TaskTemplate, error :%s", err.Error())
		return nil, err
	}

	iface := tk.Interface
	outputsDeclared := iface != nil && iface.Outputs != nil && len(iface.Outputs.Variables) > 0

	if r == nil {
		if outputsDeclared {
			// Whack! plugin did not return any outputs for this task
			// Also When an error is observed, cache is automatically disabled
			return &io.ExecutionError{
				ExecutionError: &core.ExecutionError{
					Code:    "OutputsNotGenerated",
					Message: "Output Reader was nil. Plugin/Platform problem.",
				},
				IsRecoverable: true,
			}, nil
		}
		return nil, nil
	}
	// Reader exists, we can check for error, even if this task may not have any outputs declared
	y, err := r.IsError(ctx)
	if err != nil {
		return nil, err
	}
	if y {
		taskErr, err := r.ReadError(ctx)
		if err != nil {
			return nil, err
		}

		if taskErr.ExecutionError == nil {
			taskErr.ExecutionError = &core.ExecutionError{Kind: core.ExecutionError_UNKNOWN, Code: "Unknown", Message: "Unknown"}
		}
		return &taskErr, nil
	}

	// Do this if we have outputs declared for the Handler interface!
	if !outputsDeclared {
		return nil, nil
	}
	ok, err := r.Exists(ctx)
	if err != nil {
		logger.Errorf(ctx, "Failed to check if the output file exists. Error: %s", err.Error())
		return nil, err
	}

	if !ok {
		// Does not exist
		return &io.ExecutionError{
				ExecutionError: &core.ExecutionError{
					Code:    "OutputsNotFound",
					Message: "Outputs not generated by task execution",
				},
				IsRecoverable: true,
			}, nil
	}

	if !r.IsFile(ctx) {
		// Read output and write to file
		// No need to check for Execution Error here as we have done so above this block.
		err = outputCommitter.Put(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "Failed to commit output to remote location. Error: %v", err)
			return nil, err
		}
	}

	return nil, nil
}
