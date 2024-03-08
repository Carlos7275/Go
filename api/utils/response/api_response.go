package utils

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
