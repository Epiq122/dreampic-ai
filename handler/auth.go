package handler

import (
	"net/http"

	"github.com/Epiq122/dreampic-ai/view/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}
