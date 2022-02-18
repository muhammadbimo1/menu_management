package usecase

import (
	"menu_management/entity"
	"menu_management/repository"
)

type MenuUseCase interface {
	GetAllMenu() ([]entity.Menu, error)
	SearchMenuByName(menuName string) ([]entity.Menu, error)
	SearchMenuById(menuId string) (*entity.Menu, error)
}

type menuUseCase struct {
	menuRepo repository.MenuRepository
}

func NewMenuUseCase(menuRepo repository.MenuRepository) MenuUseCase {
	return &menuUseCase{
		menuRepo,
	}
}

func (m *menuUseCase) GetAllMenu() ([]entity.Menu, error) {
	return m.menuRepo.GetAll()
}

func (m *menuUseCase) SearchMenuByName(menuName string) ([]entity.Menu, error) {
	return m.menuRepo.GetByName(menuName)
}

func (m *menuUseCase) SearchMenuById(menuId string) (*entity.Menu, error) {
	return m.menuRepo.GetById(menuId)
}
