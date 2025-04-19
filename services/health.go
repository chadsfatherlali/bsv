package services

import "net/http"

func ResponseOK() (int, string) {
	return http.StatusOK, "Hello Wordld"
}
