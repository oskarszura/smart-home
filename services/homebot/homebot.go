package homebot

import (
    "log"
    "time"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/state"
)

type HomeBot struct {
    store state.IDataFlux
    state state.IState
    mailer state.IMailer
}

func New(store state.IDataFlux, mailer state.IMailer, st state.IState) HomeBot {
    return HomeBot {
        store: store,
        state: st,
        mailer: mailer,
    }
}

func persistData(store state.IDataFlux) func(agent.Agent, map[string]interface{}) {
    return func (agent agent.Agent, data map[string]interface{}) {
        pt, _ := client.NewPoint(
        agent.ID(),
        map[string]string{ "home": agent.Name() },
        data,
        time.Now(),
        )

        err := store.AddData(pt)

        if err != nil {
        log.Println("services", err)
        }
    }
}

func (hb HomeBot) runCommunicationLoop() {
    for range time.Tick(time.Second * 10) {
        if hb.store.IsConnected() == false {
            log.Println("services: cannot fetch packages, Influx is down")
            return
        }

        agents := hb.state.Agents()

        for i := 0; i < len(agents); i++ {
            a := agents[i]
            log.Println("services: fetching from=", a.Name)

            if a.AgentType() == "type1" {
                a.FetchPackage(hb.mailer.BulkEmail, persistData(hb.store), hb.state.IsAlerts())
            }
        }
    }
}

// RunService - setup and run everything
func (hb HomeBot) RunService() {
    hb.runCommunicationLoop()
}

