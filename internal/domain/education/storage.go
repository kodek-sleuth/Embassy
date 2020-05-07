package education

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
	db.AutoMigrate(&Education{})
	return &Connection{db,}
}

func (c Connection) Create(education *Education) (*Education, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(education).Error; err != nil{
		return nil, err
	}
	return education, nil
}

func (c Connection) Update(education *Education) (*Education, error) {
	// Update Education
	if err := c.db.Model(education).Updates(education).Error; err != nil {
		return nil, err
	}

	return education, nil
}

func (c Connection) Delete(education *Education) error {
	err := c.db.Where("id = ?", education.ID).Delete(Education{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Education, error) {
	var education []*Education
	err := c.db.Find(&education).Error
	if err != nil{
		return nil, err
	}
	return education, nil
}

func (c Connection) FindById(education *Education) (*Education, error) {
	err := c.db.Where("id = ?", education.ID).First(&education).Error
	if err != nil {
		return education, err
	}
	return education, nil
}

