package model

type UpdateProfile struct {
	Id        uint
	Passport  string
	Password  string
	Nickname  string
	Email     string
	Avatar    string
	Status    int
	Gender    int
}
