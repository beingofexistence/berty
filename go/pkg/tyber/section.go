package tyber

import (
	"context"

	"go.uber.org/zap"
)

func Section(ctx context.Context, logger *zap.Logger, name string) (context.Context, bool, func(error, string, ...StepMutator)) {
	ctx, newTrace := ContextWithTraceID(ctx)
	if newTrace {
		LogTraceStart(ctx, logger, name)
	} else {
		LogStep(ctx, logger, name)
	}
	return ctx, newTrace, func(err error, rename string, muts ...StepMutator) {
		if err != nil {
			errorName := "Error while " + name
			if rename != "" {
				errorName = rename
			}
			if newTrace {
				_ = LogFatalError(ctx, logger, errorName, err, muts...)
			} else {
				_ = LogError(ctx, logger, errorName, err, muts...)
			}
		} else {
			successName := name + " succeeded"
			if rename != "" {
				successName = rename
			}
			if newTrace {
				LogTraceEnd(ctx, logger, successName, muts...)
			} else {
				LogStep(ctx, logger, successName, muts...)
			}
		}
	}
}
