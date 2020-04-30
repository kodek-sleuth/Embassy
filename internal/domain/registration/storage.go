package registration

import (
	"github.com/jinzhu/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &Connection{db}
}

func (c *Connection) Create(user *Registration) (*Registration, error) {
	if err := c.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Connection) Delete(user *Registration) error {
	err := c.db.Where("id = ?", user.ID).Delete(Registration{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c *Connection) FindBy(user *Registration, mode string) (*Registration, error) {
	switch {
	case mode == "id":
		if err := c.db.Where("id = ?", user.ID).First(&user).Error; err != nil{
			return user, err
		}
	default:
		return user, nil
	}
	return user, nil
}

func (c *Connection) FindAll() ([]*Registration, error) {
	var users []*Registration
	if err := c.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (c *Connection) Update(user *Registration) (*Registration, error) {
	err := c.db.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
