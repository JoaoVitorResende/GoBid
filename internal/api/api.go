package api

import (
	"github.com/JoaoVitorResende/GoBid/internal/services"
	"github.com/go-chi/chi/v5"
)

// air --build.cmd "go build -o .\bin\api.exe .\cmd\api" --build.bin ".\bin\api.exe"
type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}
