package notice

type Repository interface {
	Create(notice *Notice) (*Notice, error)
	Update(notice *Notice) (*Notice, error)
	Delete(notice *Notice)  error
	FindAll()([]*Notice, error)
	FindById(notice *Notice)(*Notice, error)
}
