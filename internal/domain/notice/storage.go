package notice

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
	db.AutoMigrate(&Notice{})
	return &Connection{db,}
}

func (c Connection) Create(notice *Notice) (*Notice, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(notice).Error; err != nil{
		return nil, err
	}
	return notice, nil
}

func (c Connection) Update(notice *Notice) (*Notice, error) {
	// Update Notice
	if err := c.db.Model(notice).Updates(notice).Error; err != nil {
		return nil, err
	}

	return notice, nil
}

func (c Connection) Delete(notice *Notice) error {
	err := c.db.Where("id = ?", notice.ID).Delete(Notice{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Notice, error) {
	var notice []*Notice
	err := c.db.Find(&notice).Error
	if err != nil{
		return nil, err
	}
	return notice, nil
}

func (c Connection) FindById(notice *Notice) (*Notice, error) {
	err := c.db.Where("id = ?", notice.ID).First(&notice).Error
	if err != nil {
		return notice, err
	}
	return notice, nil
}

