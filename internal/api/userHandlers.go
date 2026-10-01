package api

import "net/http"

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	panic("Todo handle signup")
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("Todo handle login")
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	panic("Todo handle logout")
}
//sqlc generate -f ./internal/store/pgstore/sqlc.yml
//create table go run ./cmd/terndotenv 
// create class sql to create table tern new create_user_table
//go to place where migrations are cd internal/store/pgstore/migrations 