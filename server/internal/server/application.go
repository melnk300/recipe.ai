package server

import (
	"github.com/melnk300/recipe.ai/server/internal/configs"
	"github.com/rs/cors"
	"net/http"
)

func InitServer(config configs.ServerConfig) error {
	r := InitRoute(config)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	return http.ListenAndServe(
		config.Host+":"+config.Port,
		c.Handler(r),
	)
}
