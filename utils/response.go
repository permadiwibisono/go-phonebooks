package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status int, msg string) map[string]interface{} {
	return map[string]interface{}{"status_code": status, "message": msg}
}

func MessageWithData(status int, msg string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"status_code": status, "message": msg, "data": data}
}

func Respond(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, statusCode int, msg string, errData map[string]interface{}) {
	errData["status_code"] = statusCode
	errData["message"] = msg
	Respond(w, statusCode, errData)
}
