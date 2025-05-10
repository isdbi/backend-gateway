package router

import (
	"github.com/lai0xn/isdb/internal/api/handler"
	"github.com/lai0xn/isdb/internal/api/middleware"
	"github.com/go-chi/chi/v5"
  "github.com/swaggo/http-swagger/v2"
)

func Route(r *chi.Mux, api handler.API) {
	r.Use(middleware.Json)


	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/api/v1", func(mux chi.Router) {
		mux.Route("/auth", func(mux chi.Router) {
			mux.Post("/login", api.Login)
			mux.Post("/signup", api.Signup)
      mux.Post("/verify",api.AcivateUser)
      mux.Post("/send-verification",api.SendVerificationEmail)
      mux.Get("/google",api.GoogleLogin)
      mux.Get("/google/callback",api.GoogleCallback)
    })
    mux.Route("/doc",func(mux chi.Router) {
      mux.Post("/upload",api.ParseDocument)
    })
	})
}
