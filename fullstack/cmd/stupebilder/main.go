package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	mw "github.com/tabinnorway/stupebilder/middleware"
	"github.com/tabinnorway/stupebilder/services/albums"
	"github.com/tabinnorway/stupebilder/services/folders"
	"github.com/tabinnorway/stupebilder/services/home"
	"github.com/tabinnorway/stupebilder/services/images"
	"github.com/tabinnorway/stupebilder/services/thumbs"
)

var PORT = ":3001"
var IMG_ROOT = "/mnt/familyshare/images"
var HOME_BASE = "/home/tberg/dev.p/bstk/stupebilder.no/fullstack"

func main() {
	toDir := HOME_BASE
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Printf("Executable is %s", ex)
	if strings.Contains(ex, "/bin/") {
		index := strings.Index(ex, "/bin/")
		if index != -1 {
			// Slice the string up to the start of the substring plus its length
			toDir = ex[:index+len("/bin/")]
		}
		toDir = toDir[:index]
	}
	log.Printf("Changing directory to %s", toDir)
	err = os.Chdir(toDir)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
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

	r.Route("/", home.NewHandler().RegisterRoutes)
	r.Route("/albums", albums.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/folders", folders.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/thumbs", thumbs.NewHandler(IMG_ROOT).RegisterRoutes)
	r.Route("/images", images.NewHandler(IMG_ROOT).RegisterRoutes)

	log.Printf("Listening to %s", PORT)
	err = http.ListenAndServe(PORT, r)
	if err != nil {
		log.Panic(err)
	}
}
