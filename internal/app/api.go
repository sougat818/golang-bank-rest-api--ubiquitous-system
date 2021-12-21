package app

import (
	"encoding/json"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getHealthMessage("ok"))
}

func getHealthMessage(key string) map[string]bool {
	return map[string]bool{key: true}
}
