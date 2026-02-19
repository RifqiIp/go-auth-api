package handler

import (
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ini profile, harusnya butuh login",
	})
}
