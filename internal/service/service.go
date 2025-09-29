package service

import (
	"github.com/nicitapa/firstProgect/internal/contracts"
)

type Service struct {
	repository contracts.RepositoryI
	cache      contracts.CacheI
}

func NewService(repository contracts.RepositoryI, cache contracts.CacheI) *Service {
	return &Service{
		repository: repository,
		cache:      cache,
	}
}
