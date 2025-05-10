package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/lai0xn/isdb/internal/api/handler"
	"github.com/lai0xn/isdb/internal/api/router"
	"github.com/lai0xn/isdb/internal/app/auth"
	"github.com/lai0xn/isdb/internal/app/documents"
	"github.com/lai0xn/isdb/internal/app/users"
	"github.com/lai0xn/isdb/internal/repository"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	PORT      string
	TwirpPORT string
	Router    *chi.Mux
	Cache     *redis.Client
}

type Config struct {
	PORT      string
	TwirpPORT string
	DB        *gorm.DB
	Router    *chi.Mux
	Cache     *redis.Client
}

func New(cfg *Config) *Server {
	return &Server{
		PORT:      cfg.PORT,
		Router:    cfg.Router,
		TwirpPORT: cfg.TwirpPORT,
		Cache:     cfg.Cache,
	}
}

func (s *Server) Run() {
	log.Info("🚀 Server Starting")

	ctx := context.Background()
  conString := fmt.Sprintf("postgres://%s:%s@%s/%s",
    viper.GetString("database.user"),
    viper.GetString("database.password"),
    viper.GetString("database.host"), // Just container name in Docker network
    viper.GetString("database.name"),
  )
	con, err := pgx.Connect(ctx, conString)

	if err != nil {
		panic(err)
	}

	defer con.Close(ctx)

	repo := repository.New(con)
  verifier := auth.EmailVerifier{
    Rd: *s.Cache,
  }
	authSrv := auth.NewService(repo,verifier)
  userSrv := users.NewService(repo)
  docSrv := documents.NewDocService(repo)
	api := handler.NewApi(authSrv,userSrv,docSrv)

	router.Route(s.Router, *api)
	if err := http.ListenAndServe(s.PORT, s.Router); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
