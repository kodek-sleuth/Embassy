package validations

import (
	uuid "github.com/satori/go.uuid"
)

type Chat struct {
	ID uuid.UUID
	From string
	To string
	Message string
}

type User struct {
	ID uuid.UUID
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	IsAdmin bool `json:"is_admin"`
	IsVerified bool `json:"is_verified"`
	RegistrationDetails Registration
}

type Registration struct {
	UserID uuid.UUID
	Gender string `json:"gender"`
	FirstName string `json:"first_name"`
	Surname string `json:"surname"`
	Photo string `json:"photo"`
	ProofOfResidence string `json:"proof_of_residence"`
	PassportNumber string `json:"passport_number"`
	City string `json:"city"`
	Address string `json:"address"`
	Marriage string `json:"marriage"`
	KinName string `json:"kin_name"`
	KinContact string `json:"kin_contact"`
	KinRelationship string `json:"kin_relationship"`
	OriginArea string `json:"origin_area"`
	ArrivalDate string `json:"arrival_date"`
	Comment string `json:"comment"`
	Code string `json:"code"`
}

type News struct {
	ID uuid.UUID
	Image string `json:"image"`
	Title string `json:"title"`
	Body string `json:"body"`
	UserID uuid.UUID
}
