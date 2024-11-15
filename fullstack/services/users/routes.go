package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/dtos/users"
	"github.com/tabinnorway/stupebilder/interfaces"
	"github.com/tabinnorway/stupebilder/utils"
)

type Handler struct {
	store interfaces.UserStore
}

func NewHandler(s interfaces.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.getAll)
	r.Post("/", h.createUser)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	// users, err := h.store.GetAll()
	// if err != nil {
	// 	log.Printf("error getting users: %s", err.Error())
	// }
	// views.Users(&users).Render(r.Context(), w)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user users.UserCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	newUser, err := h.store.Create(*user.ToModel())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteJSON(w, http.StatusOK, newUser)
}
