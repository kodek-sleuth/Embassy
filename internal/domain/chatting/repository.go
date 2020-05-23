package chatting

type Repository interface {
	Create(chat *Chat) (*Chat, error)
	Update(chat *Chat) (*Chat, error)
	Delete(chat *Chat)  error
	FindAll()([]*Chat, error)
	FindByID(chat *Chat)([]*Chat, error)
}
