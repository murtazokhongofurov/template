package auth

import "net/http"

type Handlers interface {
	Register() http.Handler
	Login() http.Handler
	Logout() http.Handler
	Update() http.Handler
	Delete() http.Handler
	GetUserByID() http.Handler
	FindByName() http.Handler
	GetUsers() http.Handler
	GetMe() http.Handler
	UploadAvatar() http.Handler
	GetCSRFToken() http.Handler
}
