package user

import (
	"net/http"

	"github.com/VadimShara/rest-api-first/internal/apperror"
	"github.com/VadimShara/rest-api-first/internal/handlers"
	"github.com/VadimShara/rest-api-first/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger : logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error{
	return apperror.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error{
	return fmt.Errorf("this is API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error{
	return apperror.NewAppError(nil, "test", "test", "t13")
}
