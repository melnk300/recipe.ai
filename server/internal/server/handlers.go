package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/melnk300/recipe.ai/server/internal/configs"
)

func InitRoute(config configs.ServerConfig) *chi.Mux {
	r := chi.NewRouter()

	return r
}
