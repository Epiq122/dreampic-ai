package handler

import (
	"net/http"

	"github.com/Epiq122/dreampic-ai/view/auth"
	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Login())

}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// call supabase
	return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
		Email:              "email is incorrect",
		Password:           "password is incorrect",
		InvalidCredentials: "please check your credentials and try again",
	}))

	return nil
}
