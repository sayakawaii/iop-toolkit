package models

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoID() string {
	id, err := gonanoid.New(10) // default len 21
	if err != nil {
		return "fallback_id"
	}
	return id
}

func GenerateRequestID() string {
	id, _ := gonanoid.New(10)
	return "req_" + id
}
