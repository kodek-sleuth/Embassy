package registration

import (
	"Embassy/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Registration struct {
	database.Base
	UserID uuid.UUID `gorm:"type:uuid;not_null;unique_index;"`
	Gender string `gorm:"type:varchar(250);"`
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
	ArrivalDate string `gorm:"type:varchar(100);not_null"`
	Comment string
	Code string `gorm:"type:varchar(100);not_null"`
}
