package application

import (
	"go-web-challenge/internal/handler"
	"go-web-challenge/internal/loader"
	"go-web-challenge/internal/repository"
	"go-web-challenge/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	loader := loader.NewLoaderTicketCSV(a.dbFile)
	db, lastId, err := loader.Load()
	rp := repository.NewRepositoryTicketMap(db, lastId)
	// service ...

	sv := service.NewServiceTicketDefault(rp)
	// handler ...

	hd := handler.NewTicketHandler(sv)

	// routes
	a.rt.Get("/health", hd.Health())
	a.rt.Route("/tickets", func(r chi.Router) {
		r.Get("/country/{dest}", hd.GetTicketsByCountry())
		r.Get("/country/percentage/{dest}", hd.GetPercentageTicketsByCountry())
	})

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
