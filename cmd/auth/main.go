package main

import (
	"context"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/pkg/auth"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/pkg/logging"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
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

	if err = router.Run(":8081"); err != nil {
		logging.Logger.Fatalf("failed starting server: %v", err)
	}
}
