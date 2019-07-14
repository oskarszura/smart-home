package webserver

import (
	"errors"
	"github.com/coda-it/gowebserver"
	"github.com/smart-evolution/smarthome/datasources/dataflux"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/api"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/api/agentconfigs"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/api/agents"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/sapi"
	"github.com/smart-evolution/smarthome/utils"
	"golang.org/x/net/websocket"
	"net/http"
)

// WebServer - adapter for gowebserver instance
type WebServer struct {
	server *gowebserver.WebServer
}

func getServerAddress(port string) (string, error) {
	if port == "" {
		return "", errors.New("Port not set")
	}
	return ":" + port, nil
}

// New - creates new WebServer instance
func New(port string, store dataflux.IDataFlux, persistence persistence.IPersistance, s state.IState) *WebServer {
	addr, err := getServerAddress(port)

	if err != nil {
		utils.Log(err)
	}

	serverOptions := gowebserver.WebServerOptions{
		Port:           addr,
		StaticFilesUrl: "/static/",
		StaticFilesDir: "public",
	}

	server := gowebserver.New(serverOptions, controllers.NotFound)

	server.Router.AddRoute("/agent/{agent}", controllers.CtrDashboard)
	server.Router.AddRoute("/agent/{agent}/edit", controllers.CtrAgentEdit)
	server.Router.AddRoute("/login/register", controllers.Register)
	server.Router.AddRoute("/login/logout", controllers.AuthenticateLogout)
	server.Router.AddRoute("/login", controllers.Authenticate)
	server.Router.AddRoute("/", controllers.CtrDashboard)
	server.Router.AddRoute("/api/agents", agents.CtrAgents)
	server.Router.AddRoute("/api/agents/{agent}", agents.CtrAgents)
	server.Router.AddRoute("/api/agent-configs/{agent}", agentconfigs.CtrAgentConfig)
	server.Router.AddRoute("/api/alerts", api.CtrAlerts)
	server.Router.AddRoute("/api/sendalert", api.CtrSendAlert)
	server.Router.AddRoute("/api/sniffagents", api.CtrSniffAgents)
	http.Handle("/sapi", websocket.Handler(sapi.AgentStreaming))

	server.AddDataSource("dataflux", store)
	server.AddDataSource("persistence", persistence)
	server.AddDataSource("state", s)

	return &WebServer{
		server: server,
	}
}

// RunService - runs WebServer process
func (ws *WebServer) RunService() {
	ws.server.Run()
}
