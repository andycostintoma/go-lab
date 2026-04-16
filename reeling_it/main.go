package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"frontendmasters.com/movies/data"
	"frontendmasters.com/movies/handlers"
	"frontendmasters.com/movies/logger"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	// logInstance.Error("Hello from the Error system", nil)
	if err != nil {
		log.Fatalf("Failed to initialice logger $v", err)
	}
	return logInstance
}

func main() {
	logInstance := initializeLogger()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL not set in environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Data Repository for Movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movierepository")
	}

	// Initialize Account Repository for Users
	accountRepo, err := data.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatal("Failed to initialize account repository")
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	accountHandler := handlers.NewAccountHandler(accountRepo, logInstance)

	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random/", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search/", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie) // api/movies/140
	http.HandleFunc("/api/genres/", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register/", accountHandler.Register)
	http.HandleFunc("/api/account/authenticate/", accountHandler.Authenticate)

	http.Handle("/api/account/favorites/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetFavorites)))

	http.Handle("/api/account/watchlist/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetWatchlist)))

	http.Handle("/api/account/save-to-collection/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.SaveToCollection)))

	// Web Authentication - Passkeys
	// WebAuthn Handlers
	wconfig := &webauthn.Config{
		RPDisplayName: "ReelingIt",
		RPID:          "localhost",
		RPOrigins:     []string{"http://localhost:8080"},
	}

	var webAuthnManager *webauthn.WebAuthn

	if webAuthnManager, err = webauthn.New(wconfig); err != nil {
		logInstance.Error("Error creating WebAuthn", err)
	}

	if err != nil {
		logInstance.Error("Error initialing Passkey engine", err)
	}

	passkeyRepo := data.NewPasskeyRepository(db, *logInstance)
	webAuthnHandler := handlers.NewWebAuthnHandler(passkeyRepo, logInstance, webAuthnManager)
	// Needs User Authentication (for passkey registration)
	http.Handle("/api/passkey/registration-begin",
		accountHandler.AuthMiddleware(http.HandlerFunc(webAuthnHandler.WebAuthnRegistrationBeginHandler)))
	http.Handle("/api/passkey/registration-end",
		accountHandler.AuthMiddleware(http.HandlerFunc(webAuthnHandler.WebAuthnRegistrationEndHandler)))
	// No need for User Authentication before
	http.HandleFunc("/api/passkey/authentication-begin", webAuthnHandler.WebAuthnAuthenticationBeginHandler)
	http.HandleFunc("/api/passkey/authentication-end", webAuthnHandler.WebAuthnAuthenticationEndHandler)

	catchAllClientRoutesHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}

	// SSR for the movie details
	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Count(r.URL.Path, "/") == 2 && strings.HasPrefix(r.URL.Path, "/movies/") {
			handlers.SSRMovieDetailsHandler(movieRepo, logInstance)(w, r)
		} else {
			catchAllClientRoutesHandler(w, r)
		}
	})

	// Catch All
	http.HandleFunc("/movies", catchAllClientRoutesHandler)

	// http.HandleFunc("/movies/", catchAllClientRoutesHandler)

	http.HandleFunc("/account/", catchAllClientRoutesHandler)

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}

}
