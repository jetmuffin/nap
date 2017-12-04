package master

import (
	"fmt"
	"github.com/JetMuffin/nap/apis"
	"github.com/JetMuffin/nap/apis/router"
	authRouter "github.com/JetMuffin/nap/apis/router/auth"
	consoleRouter "github.com/JetMuffin/nap/apis/router/console"
	mesosRouter "github.com/JetMuffin/nap/apis/router/mesos"
	"github.com/JetMuffin/nap/pkg/config"
	"github.com/JetMuffin/nap/pkg/mesos"
	"net"
)

type Master struct {
	api         *apis.Server
	listener    net.Listener
	mesosClient *mesos.Client

	cfg   *config.MasterConfig
	errCh chan error
}

func New(cfg *config.MasterConfig) (*Master, error) {
	conn, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return nil, err
	}

	apiConfig := &apis.Config{
		LogLevel: cfg.LogLevel,
	}
	apiServer := apis.New(apiConfig)

	return &Master{
		cfg:      cfg,
		listener: conn,
		api:      apiServer,
		errCh:    make(chan error, 1),
	}, nil
}

func (m *Master) Start() error {
	m.mesosClient = mesos.NewClient(m.cfg.MesosAddr, nil)

	m.api.Accept(m.cfg.ListenAddr, m.listener)
	m.initRouter()

	// The serve API routine never exits unless an error occurs
	// We need to start it as a goroutine and wait on
	serveAPIWait := make(chan error)
	go m.api.Wait(serveAPIWait)

	errAPI := <-serveAPIWait
	if errAPI != nil {
		return fmt.Errorf("Shutting down due to ServeAPI error: %v", errAPI)
	}

	return nil
}

func (m *Master) Stop() {
	m.api.Stop()
}

func (m *Master) initRouter() {
	routers := []router.Router{
		consoleRouter.NewRouter(m.mesosClient),
		authRouter.NewRouter(m.cfg.OAuthAddr),
		mesosRouter.NewRouter(),
	}

	m.api.InitRouter(routers...)
}
