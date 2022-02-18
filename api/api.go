package api

import (
	"database/sql"
	"log"
	"menu_management/config"
	"menu_management/delivery"
	"menu_management/entity"
	"menu_management/manager"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func (s *server) Run() {
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, err := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Panicln(err)
			}
		}(db)
		s.InitRouter()
		s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Panicln(err)
		}
	} else {
		db := s.infra.SqlDb()
		err := db.AutoMigrate(&entity.Menu{})
		db.Unscoped().Where("id like ?", "%%").Delete(entity.Menu{})
		db.Model(&entity.Menu{}).Save([]entity.Menu{
			{
				ID:       "0001",
				MenuName: "sayur",
				Price:    2000,
			},
			{
				ID:       "0002",
				MenuName: "Mayur",
				Price:    3000,
			},
			{
				ID:       "0003",
				MenuName: "Perkedel",
				Price:    5000,
			},
		})
		if err != nil {
			log.Panicln(err)
		}
	}

}

func (s *server) InitRouter() {
	publicRoute := s.config.RouterEngine.Group("/api")
	delivery.NewMenuApi(publicRoute, s.usecase.MenuUseCase())
}

func NewApiServer() Server {
	appconfig := config.NewConfig()
	infra := manager.NewInfra(appconfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManager(repo)

	return &server{
		config:  appconfig,
		infra:   infra,
		usecase: usecase,
	}
}
