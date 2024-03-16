package handler

import (
	"net/http"

	"github.com/Epiq122/dreampic-ai/view/settings"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	return render(r, w, settings.Index(user))

}
