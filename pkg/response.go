package pkg

import (
	"encoding/json"
	"net/http"
)

func CreateResponse(v any, w http.ResponseWriter) {
	marshalled, err := json.Marshal(v)

	if err != nil {
		fallbackResponse, _ := json.Marshal(`{"errorMessage": "inconvertible response"}`)
		_, _ = w.Write(fallbackResponse)
	}

	_, _ = w.Write(marshalled)
}
