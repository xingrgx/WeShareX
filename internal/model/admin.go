package model

type UpdateProfile struct {
	Id        uint
	Passport  string
	Nickname  string
	Email     string
	Avatar    string
	Status    int
	Gender    int
}
