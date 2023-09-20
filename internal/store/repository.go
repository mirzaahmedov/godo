package store

type UserRepository interface {
	GetAll() error
	GetByID() error
	Delete() error
	Create() error
	Update() error
}

type TodoRepository interface {
	GetAll() error
	GetByID() error
	Delete() error
	Create() error
	Update() error
}
