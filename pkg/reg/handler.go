package reg

import (
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/user"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/pkg/logging"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

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

	logging.Logger.Printf("User registered successfully: %s (ID: %d)", u.Username, u.ID)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
