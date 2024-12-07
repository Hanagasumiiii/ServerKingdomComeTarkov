package auth

import (
	"bytes"
	"context"
	"encoding/json"
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

var jwtKey []byte

func init() {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		logging.Logger.Println("JWT_SECRET_KEY не установлен, используется ключ по умолчанию для тестирования")
		jwtSecret = "default_test_secret_key"
	}
	jwtKey = []byte(jwtSecret)
}

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

	expirationTime := time.Now().Add(5 * time.Minute) //token generation
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

	redisClientURL := os.Getenv("REDISCLIENT_URL")
	if redisClientURL == "" {
		redisClientURL = "http://localhost:8081"
	}

	setTokenURL := redisClientURL + "/set"
	payload := map[string]interface{}{
		"token":      tokenString,
		"user_id":    u.ID,
		"expires_in": int64(5 * 60), // 5 минут в секундах
	}
	payloadBytes, _ := json.Marshal(payload)
	resp, err := http.Post(setTokenURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil || resp.StatusCode != http.StatusOK {
		logging.Logger.Printf("Error saving token via redisclient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving token"})
		return
	}

	logging.Logger.Printf("Successfully generated token for user %s", u.Username)
	c.JSON(http.StatusOK, TokenResponse{Token: tokenString}) //token response
}
