package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/JoaoVitorResende/GoBid/internal/api"
	"github.com/JoaoVitorResende/GoBid/internal/services"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// go build -o ./bin/api.exe ./cmd/api
// .\bin\api.exe
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	/*
		%s → string

		%d → número inteiro

		%f → número decimal

		%v → valor em formato padrão (muito usado quando você não quer especificar o tipo)
	*/
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBID_DATABASE_USER"),
		os.Getenv("GOBID_DATABASE_PASSWORD"),
		os.Getenv("GOBID_DATABASE_HOST"),
		os.Getenv("GOBID_DATABASE_PORT"),
		os.Getenv("GOBID_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	
	api := api.Api{
		Router:      chi.NewMux(),
		UserService: services.NewUserService(pool),
		Sessions:    s,
	}

	api.BindRoutes()

	fmt.Println("starting server on pont :3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}
