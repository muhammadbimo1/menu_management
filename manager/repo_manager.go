package manager

import "menu_management/repository"

type RepoManager interface {
	MenuRepo() repository.MenuRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) MenuRepo() repository.MenuRepository {
	return repository.NewMenuRepository(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
