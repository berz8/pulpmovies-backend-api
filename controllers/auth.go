package controllers

import (
	"net/http"

	"github.com/berz8/pulpmovies-backend-api/validators"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ctrl AuthController) Login(c *gin.Context) {
  var loginValidator validators.Login

  if validationErr := c.ShouldBindJSON(&loginValidator); validationErr != nil {
    c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid credentials"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}
