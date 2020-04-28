package user

import (
	"embassy/internal/database"
	"time"
)

type User struct {
	database.Base
	Email string `gorm:"type:varchar(100);unique_index;not_null"`
	Password string `gorm:"type:varchar(250);not_null"`
	FirstName string `gorm:"type:varchar(100);not_null"`
	Surname string `gorm:"type:varchar(100);not_null"`
	Photo string `gorm:"type:varchar(100);not_null"`
	ProofOfResidence string `gorm:"type:varchar(100);not_null"`
	PassportNumber string `gorm:"type:varchar(100);not_null"`
	City string `gorm:"type:varchar(100);not_null"`
	Address string `gorm:"type:varchar(100);not_null"`
	IsMarried bool `gorm:"type:boolean;not_null"`
	KinName string `gorm:"type:varchar(100);not_null"`
	KinContact string `gorm:"type:varchar(100);not_null"`
	KinRelationship string `gorm:"type:varchar(100);not_null"`
	OriginArea string `gorm:"type:varchar(100);not_null"`
	ArrivalDate time.Time
	Comment string
	Code string `gorm:"type:varchar(100);not_null"`
	IP string `gorm:"type:varchar(250);not_null"`
	IsAdmin bool `gorm:"type:boolean;not_null"`
}
