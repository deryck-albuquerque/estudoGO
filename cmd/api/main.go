package main

import (
	"github.com/deryck-albuquerque/estudoGO/internal/handlers"
	"github.com/deryck-albuquerque/estudoGO/internal/repositories"
	"github.com/deryck-albuquerque/estudoGO/internal/usecases"
)

func main() {
	repos := repositories.New()

	useCases := usecases.New(repos)

	h := handlers.New(*useCases)

	h.Listen(8080)
}
