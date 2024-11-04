package middleware

import (
	"net/http"
)

func CheckCookieMiddleware(cookieName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// We don't require a passkey for the root URL
			// For all other URLS, we need a passkey
			cookie, err := r.Cookie(cookieName)
			if err != nil || cookie.Value != "bstkbilder" {
				if err == http.ErrNoCookie {
					http.Error(w, "No cookie or bad cookie", http.StatusUnauthorized)
					return
				}
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Cookie is OK, if the is a request to the root, pass it on to albums
			if r.RequestURI == "/" {
				http.Redirect(w, r, "/albums", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
		})
	}
}
