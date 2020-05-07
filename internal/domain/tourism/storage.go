package tourism

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
	db.AutoMigrate(&Tourism{})
	return &Connection{db,}
}

func (c Connection) Create(tourism *Tourism) (*Tourism, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(tourism).Error; err != nil{
		return nil, err
	}
	return tourism, nil
}

func (c Connection) Update(tourism *Tourism) (*Tourism, error) {
	// Update Tourism
	if err := c.db.Model(tourism).Updates(tourism).Error; err != nil {
		return nil, err
	}

	return tourism, nil
}

func (c Connection) Delete(tourism *Tourism) error {
	err := c.db.Where("id = ?", tourism.ID).Delete(Tourism{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Tourism, error) {
	var tourism []*Tourism
	err := c.db.Find(&tourism).Error
	if err != nil{
		return nil, err
	}
	return tourism, nil
}

func (c Connection) FindById(tourism *Tourism) (*Tourism, error) {
	err := c.db.Where("id = ?", tourism.ID).First(&tourism).Error
	if err != nil {
		return tourism, err
	}
	return tourism, nil
}

