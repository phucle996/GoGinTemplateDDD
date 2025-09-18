package app

import (
	psql "authen_service/infra/db/postgresql"
	redis "authen_service/infra/db/redis"
	"authen_service/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	Modules *Modules
	Router  *gin.Engine
}

func NewApplication(cfg *config.Config) *App {
	// intialization PostgreSQL
	db := psql.NewPostgres(cfg)

	// intial Redis
	rdb, err := redis.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Passwd, 0)
	if err != nil {
		panic(err)
	}

	// intialization Modules
	modules := NewModules(db, rdb.Client)

	// intialization Router
	router := gin.Default()
	RegisterRoutes(router, modules)

	return &App{
		Modules: modules,
		Router:  router,
	}
}

// start app
func (a *App) Start(cfg *config.Config) error {
	return a.Router.Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}

// close db connection if stop app
func (a *App) Stop() {
	a.Modules.PostgresDB.Close()
	a.Modules.Redis.Close()
}
