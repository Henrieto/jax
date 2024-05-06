package jax

import (
	"encoding/json"
	"net/http"
)

func Json(writer http.ResponseWriter, response any, status int) (err error) {
	writer.WriteHeader(status)
	Encoder := json.NewEncoder(writer)
	err = Encoder.Encode(response)
	return
}

func Render() {}

func Write(writer http.ResponseWriter, response []byte) (err error) {
	_, err = writer.Write(response)
	return
}
