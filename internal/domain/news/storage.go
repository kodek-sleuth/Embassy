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

func (c Connection) Create(news *News) (*News, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(news).Error; err != nil{
		return nil, err
	}
	return news, nil
}

func (c Connection) Update(news *News) (*News, error) {
	// Update News
	if err := c.db.Model(news).Updates(news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

func (c Connection) Delete(news *News) error {
	err := c.db.Where("id = ?", news.ID).Delete(News{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*News, error) {
	var news []*News
	err := c.db.Find(&news).Error
	if err != nil{
		return nil, err
	}
	return news, nil
}

func (c Connection) FindById(news *News) (*News, error) {
	err := c.db.Where("id = ?", news.ID).First(&news).Error
	if err != nil {
		return news, err
	}
	return news, nil
}

