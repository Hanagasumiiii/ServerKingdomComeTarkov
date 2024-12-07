package reg

import (
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/logging"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	// Инициализация логирования
	logging.Init()
	defer logging.Close()

	// Подключение к базе данных
	dsn := os.Getenv("DATABASE_URL")

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logging.Logger.Fatalf("failed opening connection to database: %v", err)
	}
	defer func(client *ent.Client) {
		err = client.Close()
		if err != nil {
			logging.Logger.Fatalf("failed closing client: %v", err)
		}
	}(client)

	// Миграция схемы базы данных (только при первом запуске или при изменениях)
	if err := client.Schema.Create(context.Background()); err != nil {
		logging.Logger.Fatalf("failed creating schema resources: %v", err)
	}

	// Настройка маршрутизатора Gin
	router := gin.Default()

	// Регистрация маршрутов
	router.POST("/register", func(c *gin.Context) {
		reg.RegisterUser(c, client)
	})

	// Запуск сервера на другом порту, например, 8081
	if err := router.Run(":8080"); err != nil {
		logging.Logger.Fatalf("failed to run server: %v", err)
	}
}
