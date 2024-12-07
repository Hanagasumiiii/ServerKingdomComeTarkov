package reg

import (
	"bytes"
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/schema/user"
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

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context, client *ent.Client) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Printf("Invalid registration request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Проверка, существует ли пользователь
	exists, err := client.User.Query().Where(user.UsernameEQ(req.Username)).Exist(context.Background())
	if err != nil {
		logging.Logger.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists {
		logging.Logger.Printf("Username already exists: %s", req.Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Logger.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Создание пользователя
	u, err := client.User.Create().SetUsername(req.Username).SetPassword(string(hashedPassword)).Save(context.Background())
	if err != nil {
		logging.Logger.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Генерация JWT токена
	expirationTime := time.Now().Add(5 * time.Minute)
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

	// Сохранение токена через микросервис redisclient
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

	logging.Logger.Printf("User registered successfully: %s (ID: %d)", u.Username, u.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"token":   tokenString,
	})
}
