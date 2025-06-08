package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/g-villarinho/hexagonal-demo/config"
	http_adapter "github.com/g-villarinho/hexagonal-demo/internal/adapter/handler/http"
	"github.com/g-villarinho/hexagonal-demo/internal/adapter/repository/postgres"
	"github.com/g-villarinho/hexagonal-demo/internal/adapter/token/paseto"
	"github.com/g-villarinho/hexagonal-demo/internal/core/service"
	_ "github.com/lib/pq"
)

func main() {

	// =========================================================================
	// INICIALIZAÇÃO E INJEÇÃO DE DEPENDÊNCIA (A responsabilidade do main)
	// =========================================================================

	cfg, err := config.Load(".")
	if err != nil {
		log.Fatalf("loading configuration: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connecting to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("pinging the database: %v", err)
	}

	tokenMaker, err := paseto.NewPasetoMaker(cfg.PasetoSymmKey)
	if err != nil {
		log.Fatalf("creating token maker: %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo, tokenMaker)
	userHandler := http_adapter.NewUserHandler(userService)

	// =========================================================================
	// CONFIGURAÇÃO E INÍCIO DO SERVIDOR (A responsabilidade da camada HTTP)
	// =========================================================================

	router := http_adapter.ConfigureRoutes(userHandler)
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor HTTP iniciando em %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("não foi possível iniciar o servidor http: %v", err)
	}
}
