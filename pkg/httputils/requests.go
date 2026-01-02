package httputils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonGet[TOut any](url string) (TOut, error) {
	var output TOut
	resp, err := http.Get(url)
	if err != nil {
		return output, err
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, nil
}

func JsonPost[TIn, TOut any](url string, body TIn) (TOut, error) {
	var output TOut
	content, err := json.Marshal(body)
	if err != nil {
		return output, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return output, err
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, nil
}
