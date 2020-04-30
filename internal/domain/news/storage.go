package news

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
	db.AutoMigrate(&News{})
	return &Connection{db,}
}

func (c Connection) Create(food *News) (*News, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(food).Error; err != nil{
		return nil, err
	}
	return food, nil
}

func (c Connection) Update(food *News) (*News, error) {
	// Update News
	if err := c.db.Model(food).Updates(food).Error; err != nil {
		return nil, err
	}

	return food, nil
}

func (c Connection) Delete(food *News) error {
	err := c.db.Where("id = ?", food.ID).Delete(News{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*News, error) {
	var foods []*News
	err := c.db.Find(&foods).Error
	if err != nil{
		return nil, err
	}
	return foods, nil
}

func (c Connection) FindById(food *News) (*News, error) {
	err := c.db.Where("id = ?", food.ID).First(&food).Error
	if err != nil {
		return food, err
	}
	return food, nil
}

