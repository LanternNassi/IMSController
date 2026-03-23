package contexts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/LanternNassi/IMSController/internal/utils"
	"github.com/labstack/echo"
)

type LoginContext struct {
	base         *BaseContext
	email        string
	password     string
	response     *httptest.ResponseRecorder
	statusCode   int
	responseBody map[string]interface{}
}

func NewLoginContext() *LoginContext {
	return &LoginContext{
		base: GetBaseContext(),
	}
}

func (l *LoginContext) Reset() {
	l.email = ""
	l.password = ""
	l.response = nil
	l.statusCode = 0
	l.responseBody = nil
}

func (l *LoginContext) SetCredentials(email, password string) {
	l.email = email
	l.password = password
}

func (l *LoginContext) GetEmail() string {
	return l.email
}

func (l *LoginContext) GetPassword() string {
	return l.password
}

func (l *LoginContext) CreateUser(email, password string) error {
	if err := l.base.SetupTestDatabase(); err != nil {
		return err
	}


	hashedPassword, err := utils.HarshPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user := &models.User{
		Username: "testuser",
		Email:    email,
		Password: hashedPassword,
		Verified: true,
	}

	_, err = l.base.DBClient.AddUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (l *LoginContext) SendLoginRequest() error {
	if err := l.base.SetupTestDatabase(); err != nil {
		return err
	}

	loginData := map[string]string{
		"Email":    l.email,
		"Password": l.password,
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return fmt.Errorf("error marshaling login data: %w", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/Users/login", bytes.NewBuffer(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	l.base.Server.ServeHTTP(rec, req)
	l.response = rec
	l.statusCode = rec.Code

	// Parse response body
	if rec.Body.Len() > 0 {
		var body map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err == nil {
			l.responseBody = body
		}
	}

	return nil
}

func (l *LoginContext) GetStatusCode() int {
	return l.statusCode
}

func (l *LoginContext) GetResponseBody() map[string]interface{} {
	return l.responseBody
}

func (l *LoginContext) GetResponse() *httptest.ResponseRecorder {
	return l.response
}

func (l *LoginContext) Base() *BaseContext {
	return l.base
}
