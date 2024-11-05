package middleware

import (
	"net/http"
	"strings"

	"github.com/tabinnorway/stupebilder/utils"
)

func CheckCookieMiddleware(cookieName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// We don't require a passkey for the root URL
			// For all other URLS, we need a passkey
			cookie, err := r.Cookie(cookieName)
			if err != nil || cookie.Value != "bstkbilder" {
				if r.RequestURI == "/" || strings.HasPrefix(r.RequestURI, "/style") {
					next.ServeHTTP(w, r)
					return
				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
			// Cookie is OK, if the is a request to the root, pass it on to albums
			if r.RequestURI == "/" {
				http.Redirect(w, r, "/albums", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func UrlSanitizerMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.RequestURI, "..") {
				utils.WriteError(w, http.StatusBadRequest, nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
