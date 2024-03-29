package handler

import (
	"log/slog"
	"net/http"

	"github.com/Epiq122/dreampic-ai/pkg/kit/validate"
	"github.com/Epiq122/dreampic-ai/pkg/sb"
	"github.com/Epiq122/dreampic-ai/view/auth"

	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Login())

}

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Signup())

}

func HandleSignupCreate(w http.ResponseWriter, r *http.Request) error {
	params := auth.SignupParams{
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirmPassword"),
	}
	errors := auth.SignupErrors{}
	if ok := validate.New(&params, validate.Fields{
		"Email":           validate.Rules(validate.Email),
		"Password":        validate.Rules(validate.Password),
		"ConfirmPassword": validate.Rules(validate.Equal(params.Password), validate.Message("passwords do not match")),
	}).Validate(&errors); !ok {
		return render(r, w, auth.SignupForm(params, errors))
	}
	user, err := sb.Client.Auth.SignUp(r.Context(), supabase.UserCredentials{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return err
	}
	return render(r, w, auth.SignupSuccess(user.Email))
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("failed to login", "err", err)
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "please check your credentials and try again",
		}))
	}

	setAuthCookie(w, resp.AccessToken)
	return hxRedirect(w, r, "/")

}

func HandleAuthCallBack(w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		return render(r, w, auth.CallbackScript())
	}
	// fmt.Println(accessToken)
	setAuthCookie(w, accessToken)
	http.Redirect(w, r, "/", http.StatusFound)
	return nil

}

func HandleLogoutCreate(w http.ResponseWriter, r *http.Request) error {
	cookie := http.Cookie{
		Name:     "at",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
	return nil
}

func setAuthCookie(w http.ResponseWriter, accessToken string) {
	cookie := &http.Cookie{
		Name:     "at",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

}
