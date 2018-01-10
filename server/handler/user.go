package handler

import (
	"net/http"
	"strconv"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/errorutil"
	"github.com/Hendra-Huang/go-standard-layout/log"
	"github.com/Hendra-Huang/go-standard-layout/router/helper"
	"github.com/Hendra-Huang/go-standard-layout/server/responseutil"
)

type (
	UserHandler struct {
		userService *myapp.UserService
	}
)

func NewUserHandler(us *myapp.UserService) *UserHandler {
	return &UserHandler{us}
}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := uh.userService.FindAll(ctx)
	if err != nil {
		log.Errors(err)
		responseutil.InternalServerError(w)
		return
	}

	responseutil.JSON(w, users, responseutil.WithStatus(http.StatusOK))
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	requestedID := helper.URLParam(r, "id")
	id, err := strconv.ParseInt(requestedID, 10, 64)
	errorutil.CheckAndLog(err)
	if id <= 0 {
		responseutil.NotFound(w)
		return
	}

	user, err := uh.userService.FindByID(ctx, id)
	if err != nil {
		log.Errors(err)
		responseutil.InternalServerError(w)
		return
	}
	if user.ID == 0 {
		responseutil.NotFound(w)
		return
	}

	responseutil.JSON(w, user, responseutil.WithStatus(http.StatusOK))
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
}
