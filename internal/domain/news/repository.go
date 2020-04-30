package news

type Repository interface {
	Create(news *News) (*News, error)
	Update(news *News) (*News, error)
	Delete(news *News)  error
	FindAll()([]*News, error)
	FindById(news *News)(*News, error)
}
