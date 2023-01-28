package controllers

import (
	"errors"
	"net/http"

	"github.com/berz8/pulpmovies-backend-api/config"
	"github.com/berz8/pulpmovies-backend-api/models"
	"github.com/berz8/pulpmovies-backend-api/validators"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserController struct{}

func (ctrl UserController) Register(c *gin.Context) {
  var registerValidator validators.Register

  if validationErr := c.ShouldBindJSON(&registerValidator); validationErr != nil {
    c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid data"})
    return
  }

  user := models.User{Username: registerValidator.Username, Password: registerValidator.Password, Email: registerValidator.Email, ProfilePath: registerValidator.ProfilePath}
  error := config.DB.Select("Username","Email","Password","ProfilePath").Create(&user).Error
  if error != nil {
    var duplicateKeyError = &pgconn.PgError{Code: "23505"}
    if errors.As(error, &duplicateKeyError) {
      c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Username or Email already in use"})
      return
    }
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "New user successfully registered", "data": user})
}
func GetUsers(c *gin.Context) {
  var users []models.User
  config.DB.Find(&users)

  c.JSON(http.StatusOK, gin.H{"data": users})
}
