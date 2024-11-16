package service

type ArchiveInfo interface{}

type Service struct {
	ArchiveInfo
}

func NewService() *Service {
	return &Service{}
}
