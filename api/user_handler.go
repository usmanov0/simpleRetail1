package api

import (
	"golang-project-template/internal/model"
	"golang-project-template/internal/service"
	"net/http"
	"text/template"
)

type UserController struct {
	UserService service.UserService
}

func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	principal, ok := r.Context().Value("user").(string)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	user := c.UserService.FindByUsername(principal)

	tmpl, err := template.New("profile").Parse("profileTemplate")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) ProfilePost(w http.ResponseWriter, r *http.Request) {
	var newUser model.Users
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newUser.Username = r.FormValue("username")
	newUser.FirstName = r.FormValue("firstName")
	newUser.LastName = r.FormValue("lastName")
	newUser.PinFL = r.FormValue("pinfl")
	newUser.Email = r.FormValue("email")
	newUser.Phone = r.FormValue("phone")

	user := c.UserService.FindByUsername(newUser.Username)
	user.Username = newUser.Username
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.PinFL = newUser.PinFL
	user.Email = newUser.Email
	user.Phone = newUser.Phone

	c.UserService.SaveUser(user)

	tmpl, err := template.New("profile").Parse("profileTemplate")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
