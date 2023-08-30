package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"teal-gopher/initializers"
	"teal-gopher/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("RequireAuth middleware")

	// Get the token
	var header struct {
		Authorization string
	}
	if c.BindHeader(&header) != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	if header.Authorization == "" {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	tokenString := strings.TrimPrefix(header.Authorization, "Bearer ")

	// Decode & validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("APP_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token expired",
			})
		}

		// Find user
		var user models.User
		initializers.DB.First(&user, claims["user_id"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		fmt.Println(err)
	}

}
