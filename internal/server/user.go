package server

import (
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/LanternNassi/IMSController/internal/utils"
	"github.com/labstack/echo"
)

func (s *EchoServer) AddUser(ctx echo.Context) error {
	user := new(models.User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if user.Username == "" || user.Password == "" {
		return ctx.JSON(http.StatusBadRequest, "Username or Password no supplied")
	}

	related_users, err := s.DB.GetUsers(ctx.Request().Context(), &models.User{Email: user.Email})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if len(related_users) > 0 {
		return ctx.JSON(http.StatusBadRequest, "Email already exists")
	}

	generated_harsh, err := utils.HarshPassword(user.Password)

	if err != nil {
		return ctx.JSON(http.StatusFailedDependency, err)
	}

	user.Password = generated_harsh

	created_user, err := s.DB.AddUser(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, created_user)
}

func (s *EchoServer) GetUsers(ctx echo.Context) error {
	user_filters := new(models.User)

	if err := ctx.Bind(user_filters); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	users, err := s.DB.GetUsers(ctx.Request().Context(), user_filters)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, users)

}

func (s *EchoServer) GetUserById(ctx echo.Context) error {
	user_id := ctx.Param("id")

	user, err := s.DB.GetUserById(ctx.Request().Context(), user_id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (s *EchoServer) DeleteUser(ctx echo.Context) error {
	user_id := ctx.Param("id")

	err := s.DB.DeleteUser(ctx.Request().Context(), user_id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusAccepted, "User deleted successfully")
}

func (s *EchoServer) Login(ctx echo.Context) error {

	// type LoginDetails struct {
	// 	email    string
	// 	password string
	// }

	details := new(models.User)

	if err := ctx.Bind(details); err != nil {

		return ctx.JSON(http.StatusBadRequest, err)
	}

	user, err := s.DB.GetUsers(ctx.Request().Context(), &models.User{
		Email: details.Email,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	is_valid, err := utils.ValidPassword(details.Password, user[0].Password)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if !is_valid {
		return ctx.JSON(http.StatusUnauthorized, "Invalid password or username")
	}

	generated_token, err := utils.GenerateToken(user[0].ID, user[0].Email)

	if err != nil {
		return ctx.JSON(http.StatusFailedDependency, "Failed to generate token")
	}

	type LoginResponse struct {
		Username string
		Token    string
		Email    string
	}

	return ctx.JSON(http.StatusAccepted, LoginResponse{
		Username: user[0].Username,
		Token:    generated_token,
		Email:    user[0].Email,
	})
}
