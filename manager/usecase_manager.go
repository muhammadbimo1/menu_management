package manager

import "menu_management/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(uc.repo.MenuRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repoManager,
	}
}
