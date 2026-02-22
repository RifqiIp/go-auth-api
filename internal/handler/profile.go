package handler

import "net/http"

// Profile
// endpoint protected (harus pakai JWT)
func Profile(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, "profile access granted", nil)
}