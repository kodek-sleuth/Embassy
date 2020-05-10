package pages

type Repository interface {
	Create(pages *Pages) (*Pages, error)
	Update(pages *Pages) (*Pages, error)
	Delete(pages *Pages)  error
	FindAll()([]*Pages, error)
	FindById(pages *Pages)(*Pages, error)
}
