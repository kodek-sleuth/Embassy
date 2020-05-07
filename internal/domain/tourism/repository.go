package tourism

type Repository interface {
	Create(tourism *Tourism) (*Tourism, error)
	Update(tourism *Tourism) (*Tourism, error)
	Delete(tourism *Tourism)  error
	FindAll()([]*Tourism, error)
	FindById(tourism *Tourism)(*Tourism, error)
}
