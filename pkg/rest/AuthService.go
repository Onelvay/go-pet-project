package service

type Handlers interface {
	GetBooks()
	GetBookById()
	GetBookByName()
	DeleteBookById()
}

type Service struct {
	Handlers
}
