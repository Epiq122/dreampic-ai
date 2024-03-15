package handler

import (
	"net/http"

	"github.com/Epiq122/dreampic-ai/pkg/util"
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

	if !util.IsValidEmail(credentials.Email) {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Email: "please enter a valid email address",
		}))
	}

	if reason, ok := util.ValidatePassword(credentials.Password); !ok {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Password: reason,
		}))
	}

	// resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	// if err != nil {
	// 	return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
	// 		InvalidCredentials: "please check your credentials and try again",
	// 	}))
	// }

	return nil
}
