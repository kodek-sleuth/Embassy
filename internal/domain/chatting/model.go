package chatting

import (
	"embassy/internal/database"
)

type Chat struct {
	database.Base
	From string
	To string
	Message string
}
