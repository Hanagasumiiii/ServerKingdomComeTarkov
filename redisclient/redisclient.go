package main

import (
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/logging"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		logging.Logger.Println("Не удалось загрузить файл .env, используем системные переменные окружения")
	}

	// Инициализация логирования
	logging.Init()
	defer logging.Close()

	// Получение переменных окружения
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379" // Значение по умолчанию
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	// Если пароль не установлен, оставляем пустым

	// Подключение к Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword, // Используем пароль из переменных окружения
		DB:       0,             // Номер базы данных
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		logging.Logger.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	// Настройка маршрутизатора Gin
	router := gin.Default()

	// Маршруты
	router.POST("/set", SetToken)
	router.GET("/get/:token", GetToken)

	// Запуск сервера
	port := os.Getenv("REDISCLIENT_PORT")
	if port == "" {
		port = "8081"
	}
	if err := router.Run(":" + port); err != nil {
		logging.Logger.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

// SetToken сохраняет токен в Redis с временем жизни
func SetToken(c *gin.Context) {
	var req struct {
		Token     string        `json:"token" binding:"required"`
		UserID    int           `json:"user_id" binding:"required"`
		ExpiresIn time.Duration `json:"expires_in" binding:"required"` // в секундах
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Printf("Неверный запрос на сохранение токена: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := redisClient.Set(ctx, req.Token, req.UserID, req.ExpiresIn*time.Second).Err()
	if err != nil {
		logging.Logger.Printf("Ошибка при сохранении токена в Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Токен успешно сохранён"})
}

// GetToken получает userID по токену
func GetToken(c *gin.Context) {
	token := c.Param("token")
	userID, err := redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Токен не найден"})
		return
	} else if err != nil {
		logging.Logger.Printf("Ошибка при получении токена из Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}
