package auth

import (
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/user"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/logging"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func LoginUser(c *gin.Context, client *ent.Client) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Printf("invalid login request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := client.User.Query().Where(user.UsernameEQ(req.Username)).Only(context.Background())
	if err != nil {
		logging.Logger.Printf("User not found: %s, error: %v", req.Username, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		logging.Logger.Printf("User invalid password: %s, error: %v", u.Username, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
	}

	expirationTime := time.Now().Add(24 * time.Hour) //token generation
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(u.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logging.Logger.Printf("Error generating token for user %s: %v", u.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	logging.Logger.Printf("Successfully generated token for user %s", u.Username)
	c.JSON(http.StatusOK, TokenResponse{Token: tokenString}) //token response
}
