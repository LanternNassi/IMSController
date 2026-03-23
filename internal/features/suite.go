package features

import (
	"context"

	"github.com/cucumber/godog"

	"github.com/LanternNassi/IMSController/internal/features/contexts"
	"github.com/LanternNassi/IMSController/internal/features/steps"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	baseCtx := contexts.GetBaseContext()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Reset contexts before each scenario
		if loginCtx := steps.GetLoginContext(); loginCtx != nil {
			loginCtx.Reset()
		}
		if clientCtx := steps.GetClientContext(); clientCtx != nil {
			clientCtx.Reset()
		}
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		baseCtx.Cleanup()
		return ctx, nil
	})

	steps.InitializeLoginSteps(ctx)
	steps.InitializeCommonSteps(ctx)
	steps.InitializeClientSteps(ctx)
}

