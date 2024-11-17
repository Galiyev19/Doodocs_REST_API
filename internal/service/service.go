package service

import "doodocs_rest_api/internal/entity"

type Service struct {
	ArchiveInfo entity.ArchiveInfoService
}

func NewService() *Service {
	return &Service{
		ArchiveInfo: NewArchiveInfoService(),
	}
}
