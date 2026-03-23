package steps

import (
	"github.com/cucumber/godog"

	"github.com/LanternNassi/IMSController/internal/features/contexts"
)

var loginCtx *contexts.LoginContext

func GetLoginContext() *contexts.LoginContext {
	return loginCtx
}

// InitializeLoginSteps registers all login-related step definitions
func InitializeLoginSteps(ctx *godog.ScenarioContext) {
	loginCtx = contexts.NewLoginContext()

	ctx.Step(`^the system has a test database$`, theSystemHasATestDatabase)
	ctx.Step(`^a user exists with email "([^"]*)" and password "([^"]*)"$`, aUserExistsWithEmailAndPassword)
	ctx.Step(`^I have login credentials with email "([^"]*)" and password "([^"]*)"$`, iHaveLoginCredentialsWithEmailAndPassword)
	ctx.Step(`^I send a login request$`, iSendALoginRequest)

}

func theSystemHasATestDatabase() error {
	return loginCtx.Base().SetupTestDatabase()
}

func aUserExistsWithEmailAndPassword(email, password string) error {
	return loginCtx.CreateUser(email, password)
}

func iHaveLoginCredentialsWithEmailAndPassword(email, password string) error {
	loginCtx.SetCredentials(email, password)
	return nil
}

func iSendALoginRequest() error {
	return loginCtx.SendLoginRequest()
}
