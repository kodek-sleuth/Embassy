package education

type Repository interface {
	Create(education *Education) (*Education, error)
	Update(education *Education) (*Education, error)
	Delete(education *Education)  error
	FindAll()([]*Education, error)
	FindById(education *Education)(*Education, error)
}
