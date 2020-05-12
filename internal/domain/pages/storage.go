package pages

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
	db.AutoMigrate(&Pages{})
	return &Connection{db,}
}

func (c Connection) Create(pages *Pages) (*Pages, error) {
	// Find type
	err := c.db.Where("type = ?", pages.Type).First(&pages).Error
	if err != nil{
		// Insert into menu if not ready there
		if err := c.db.Create(pages).Error; err != nil{
			return nil, err
		}
		return pages, nil
	}

	if err := c.db.Where("type = ?", pages.Type).Updates(pages).Error; err != nil {
		return nil, err
	}

	return pages, nil
}

func (c Connection) Update(pages *Pages) (*Pages, error) {
	// Update Pages
	if err := c.db.Where("type = ?", pages.Type).Updates(pages).Error; err != nil {
		return nil, err
	}

	return pages, nil
}

func (c Connection) Delete(pages *Pages) error {
	err := c.db.Where("type = ?", pages.Type).Delete(Pages{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Pages, error) {
	var pages []*Pages
	err := c.db.Find(&pages).Error
	if err != nil{
		return nil, err
	}
	return pages, nil
}

func (c Connection) FindById(pages *Pages) (*Pages, error) {
	err := c.db.Where("type = ?", pages.Type).First(&pages).Error
	if err != nil {
		return pages, err
	}
	return pages, nil
}

