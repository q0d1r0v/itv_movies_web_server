package controllers

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// GetUsers - Retrieve all users
//
//	@Summary		Retrieve all users
//	@Description	This endpoint returns a list of all registered users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		500	{object}	map[string]string
//	@Router			/load/users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// CreateUser - Register a new user
//
//	@Summary		Register a new user
//	@Description	This endpoint allows creating a new user with Name, Email, and Password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User registration details"
//	@Success		201		{object}	models.User
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/register/user [post]
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name, Email, and Password are required"})
		return
	}

	err := c.userService.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// Login - User authentication
//
//	@Summary		Authenticate a user
//	@Description	This endpoint allows users to log in using their email and password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		models.User	true	"User login credentials"
//	@Success		200			{object}	map[string]string
//	@Failure		400			{object}	map[string]string
//	@Failure		401			{object}	map[string]string
//	@Router			/auth/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Required email and password!"})
		return
	}

	token, err := uc.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
