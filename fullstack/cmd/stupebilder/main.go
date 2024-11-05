package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	mw "github.com/tabinnorway/stupebilder/middleware"
	"github.com/tabinnorway/stupebilder/services/albums"
	"github.com/tabinnorway/stupebilder/services/folders"
	"github.com/tabinnorway/stupebilder/services/home"
	"github.com/tabinnorway/stupebilder/services/images"
	"github.com/tabinnorway/stupebilder/services/thumbs"
	"github.com/tabinnorway/stupebilder/services/users"
)

var IMG_ROOT = "/mnt/familyshare/images"
var HOME_BASE = "/home/tberg/dev.p/bstk/stupebilder.no/fullstack"

func createConnectionString() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	return connStr
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", createConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.UrlSanitizerMiddleware())
	r.Use(mw.CheckCookieMiddleware("bstkpasskey"))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", home.NewHandler(db).RegisterRoutes)
	r.Route("/albums", albums.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/folders", folders.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/thumbs", thumbs.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/images", images.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/users", users.NewHandler(db).RegisterRoutes)

	listenPort := fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT"))

	log.Printf("Listening to %s", listenPort)
	err = http.ListenAndServe(listenPort, r)
	if err != nil {
		log.Panic(err)
	}
}
