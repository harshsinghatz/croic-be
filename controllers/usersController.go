package controllers

import (
	"croic/initializers"
	"croic/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	bindErr := c.Bind(&body)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload for signup",
		})

		return
	}

	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while securing your password",
		})

		return
	}

	user := models.User{Email: body.Email, Password: string(hashedPass)}
	res := initializers.DB.Create(&user)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while registering user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	bindErr := c.Bind(&body)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload for login",
		})

		return
	}

	var user models.User

	initializers.DB.Find(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email or Password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"expires_at": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, signErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if signErr != nil {

		fmt.Println(signErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while logging user in",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully logged in",
	})
}

func Validate(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user object in ctx not found",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
