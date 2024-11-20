package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/tabinnorway/stupebilder/interfaces"
	"golang.org/x/oauth2"
)

var cookieStore = sessions.NewCookieStore([]byte("your-very-secret-key"))

type Handler struct {
	store       interfaces.AuthStore
	oauthConfig *oauth2.Config
}

func NewHandler(oauthConfig *oauth2.Config, store interfaces.AuthStore) *Handler {
	return &Handler{
		store:       store,
		oauthConfig: oauthConfig,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/login", h.loginHandler)
	r.Get("/callback", h.callbackHandler)
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ClientID    : %s\n", h.oauthConfig.ClientID)
	fmt.Printf("ClientSecret: %s\n", h.oauthConfig.ClientSecret)
	fmt.Printf("RedirectURL : %s\n", h.oauthConfig.RedirectURL)

	url := h.oauthConfig.AuthCodeURL("randomstate")

	fmt.Printf("%s\n", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	// Exchange code for token
	token, err := h.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Fetch user information
	userInfo, err := h.fetchUserInfo(context.Background(), token)
	if err != nil {
		http.Error(w, "Failed to fetch user info", http.StatusInternalServerError)
		return
	}

	// Retrieve or create a session
	session, err := cookieStore.Get(r, "auth-session")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	// Save only essential user info fields
	session.Values["email"] = userInfo["email"]
	session.Values["name"] = userInfo["name"]
	session.Values["picture"] = userInfo["picture"]
	session.Values["access-token"] = token.AccessToken

	// Save session
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	// Redirect to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) fetchUserInfo(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error) {
	client := h.oauthConfig.Client(ctx, token)

	// Request user information
	resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
