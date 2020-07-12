package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/forms"
	"webServer/helpers"
	"webServer/models"
	"webServer/services"
)

var UserModel = new(models.UserModel)

type UserController struct{}

func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	data.Name = c.PostForm("name")
	data.Email = c.PostForm("email")
	data.Password = c.PostForm("pass")

	if data.Name == "" || data.Email == "" || data.Password == "" {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		return
	}

	result, _ := UserModel.GetUserByEmail(data.Email)

	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		return
	}

	err := UserModel.Signup(data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		return
	}
	c.Redirect(301, "/faq")
}

func (u *UserController) Login(c *gin.Context) {
	var data forms.LoginUserCommand

	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")

	if data.Email == "" || data.Password == "" {
		c.JSON(406, gin.H{"message": "Provide required details"})
		return
	}

	result, err := UserModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		return
	}
	if err != nil {
		c.JSON(401, gin.H{"message": "Problem logging into your account"})
		return
	}

	hashedPassword := []byte(result.Password)
	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		return
	}
	jwtToken := services.CreateToken(data, data.Email, c)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		MaxAge:   1800,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})
	c.Redirect(redirectCode, "/auth/products")

}
