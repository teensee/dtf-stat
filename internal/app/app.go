package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"recomends/internal/config"
	"recomends/internal/domain/user/handlers"
	"recomends/internal/domain/user/model"
)

type App struct {
	config *config.Config
	router *httprouter.Router
	db     *gorm.DB
	DI     *ServiceLocator
	server *Server
}

func NewKernel(config *config.Config) App {
	return App{
		config: config,
		DI:     NewServiceLocator(),
		server: NewServer(),
	}
}

func (a *App) End() {
	a.server.ConfigureAndRun(a.config.Listen.BindIP, a.config.Listen.Port, a.config.AppDebug, a.router)
}

func (a *App) ConfigureDatabase() *App {
	log.Println("Start configure database connection")
	db, err := gorm.Open(postgres.Open(a.config.GetDsn()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	a.db = db

	return a
}

func (a *App) ConfigureServiceLocator() *App {
	a.DI.Set(handlers.DtfHandlerTag, handlers.NewDtfUserHandler(a.db))

	return a
}

func (a *App) ConfigureRoutes() *App {
	log.Print("Configure routes")
	r := httprouter.New()

	r.GET("/users", a.DI.Get(handlers.DtfHandlerTag).(*handlers.DtfUserHandler).LoadDtfUsers)

	a.router = r

	return a
}

func (a *App) AfterInitializationEvents() *App {
	log.Println("Run after initialization events")

	err := a.db.AutoMigrate(
		model.DtfUser{},
		model.UserRequestLog{},
	)
	if err != nil {
		log.Fatal(err)
	}

	a.DI.Clear()

	return a
}

//ConfigureCli Run after ConfigureHandlers method
func (a *App) ConfigureCli() {
	app := cli.NewApp()
	app.Usage = "Init usage"

	dtfUserHandler := a.DI.Get(handlers.DtfHandlerTag).(*handlers.DtfUserHandler)

	app.Commands = []*cli.Command{
		{
			Name: "load_dtf_users",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "from",
					Aliases: []string{"f"},
					Value:   1,
					Usage:   "Start import users from...",
				},
			},
			Usage:  "This command runs and load dtf users",
			Action: dtfUserHandler.LoadDtfUsersInCli,
		},
	}

	a.DI.Set("cli", app)
}
