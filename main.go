package main

import (
	"AI/internal/handler"
	"AI/internal/repository"
	"AI/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize repository
	gameRepo := repository.NewGameRepository()

	// Initialize service
	gameService := service.NewGameService(gameRepo)

	// Initialize handlers
	gameHandler := handler.NewGameHandler(gameService)
	webHandler := handler.NewWebHandler(gameService, gameRepo)

	// API routes
	http.HandleFunc("/api/games", gameHandler.CreateGame)
	http.HandleFunc("/api/games/move", gameHandler.MakeMove)

	// Web routes (register specific routes first)
	http.HandleFunc("/web/create-game", webHandler.CreateGame)
	http.HandleFunc("/web/move", webHandler.MakeMove)
	http.HandleFunc("/", webHandler.Index) // Root path last

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Printf("Open http://localhost:%s in your browser!\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
