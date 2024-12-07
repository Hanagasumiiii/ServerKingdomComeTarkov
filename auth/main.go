package auth

import (
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/logging"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	logging.Init()
	defer logging.Close()

	dsn := os.Getenv("DATABASE_URL") //Database source name

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logging.Logger.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer func(client *ent.Client) {
		err = client.Close()
		if err != nil {
			logging.Logger.Fatalf("failed closing client: %v", err)
		}
	}(client)

	if err = client.Schema.Create(context.Background()); err != nil {
		logging.Logger.Fatalf("failed creating schema resources: %v", err)
	}

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		auth.LoginUser(c, client)
	})

	router.GET("/verify", func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		// Проверка токена через redisclient
		redisClientURL := os.Getenv("REDISCLIENT_URL")
		if redisClientURL == "" {
			redisClientURL = "http://localhost:8081"
		}

		getTokenURL := redisClientURL + "/get/" + tokenString
		resp, err := http.Get(getTokenURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
	})

	if err = router.Run(":8080"); err != nil {
		logging.Logger.Fatalf("failed starting server: %v", err)
	}
}
