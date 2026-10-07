package api

import (
	"errors"
	"net/http"

	"github.com/JoaoVitorResende/GoBid/internal/jsonutils"
	"github.com/JoaoVitorResende/GoBid/internal/services"
	"github.com/JoaoVitorResende/GoBid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUserName) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {

	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)

	if err != nil{
		if errors.Is(err, services.ErrInvalidCredentials){
			jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"error": "unexpected internal server error",
			})
			return
		}
	}
	err = api.Sessions.RenewToken(r.Context())

	if err != nil{
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error":"unexpected internal server rerror",
		})
		return
	}

	err = api.Sessions.RenewToken(r.Context())

	if err != nil{
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticateUserID", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged in, successfully",
	})
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())

	if err != nil{
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticateUserID")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged out, successfully",
	})
}

//sqlc generate -f ./internal/store/pgstore/sqlc.yml
//create table go run ./cmd/terndotenv
// create class sql to create table tern new create_user_table
//go to place where migrations are cd internal/store/pgstore/migrations
