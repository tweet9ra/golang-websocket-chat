package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type arrayableMap map[uint]*interface{}
func (m arrayableMap) MarshalJSON() ([]byte, error) {
	result := []byte("[")

	counter := 1
	for _, el := range m {
		elJson, _ := json.Marshal(el)
		result = append(result, elJson...)

		if counter < len(m) {
			result = append(result, ',')
		}

		counter++
	}

	return append(result, ']'), nil
}