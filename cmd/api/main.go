package main

import (
	"fmt"
	"log"

	_ "github.com/lai0xn/isdb/docs"
	_ "github.com/lai0xn/isdb/internal/api/handler"
	"github.com/lai0xn/isdb/internal/infra/redis"
	"github.com/lai0xn/isdb/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	godotenv.Load()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	fmt.Printf("Server configuration loaded successfully. Running on %s:%d\n",
		viper.GetString("server.host"),
		viper.GetInt("server.port"))

	s := server.New(&server.Config{
		PORT:   ":8080",
		Router: chi.NewRouter(),
    Cache: redis.Connect(),

	})
	s.Run()
}
