package service

type Handlers interface {
	HomePage()
}

type Service struct {
	Handlers
}
