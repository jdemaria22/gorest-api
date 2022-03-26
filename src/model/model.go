package model

type Employee struct {
	Id          int32
	FirstName   string
	LastName    string
	BadgeNumber int32
}

type UserInsert struct {
	ID        int    `validate:"isdefault"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type Status struct {
	Status      int
	Description string
}
