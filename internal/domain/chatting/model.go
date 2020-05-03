package chatting

import (
	"Embassy/internal/database"
)

type Chat struct {
	database.Base
	From string
	To string
	Message string
}
