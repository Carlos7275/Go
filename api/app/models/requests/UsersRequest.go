package requests

import "time"

type User struct {
	Email           string
	Nombres         string
	ApellidoPaterno string
	ApellidoMaterno string
	Telefono        string
	FechaNacimiento time.Time
	IDRol           int
	IDGenero        int
}
type UserRequest struct {
	User
	Password string
}
