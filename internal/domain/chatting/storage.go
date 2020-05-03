package chatting

import (
	"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
	//uuid"github.com/satori/go.uuid"
)
type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Chat{})
	return &Connection{db,}
}

func (c Connection) Create(chat *Chat) (*Chat, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(chat).Error; err != nil{
		return nil, err
	}
	return chat, nil
}

func (c Connection) Update(chat *Chat) (*Chat, error) {
	// Update Chat
	if err := c.db.Model(chat).Updates(chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
}

func (c Connection) Delete(chat *Chat) error {
	err := c.db.Where("id = ?", chat.ID).Delete(Chat{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Chat, error) {
	var chats []*Chat
	err := c.db.Find(&chats).Error
	if err != nil{
		return nil, err
	}
	return chats, nil
}

func (c Connection) FindById(chat *Chat) (*Chat, error) {
	err := c.db.Where("id = ?", chat.ID).First(&chat).Error
	if err != nil {
		return chat, err
	}
	return chat, nil
}

