package handler

import (
	"fmt"
	"net/http"

	"github.com/foolin/gomap"
	"github.com/foolin/gosupervisor"
	"github.com/kolo/xmlrpc"
)

// HealthCheck struct
type HealthCheck struct {
	Status    string          `json:"status"`
	Processes []ProcessStatus `json:"processes"`
}

// ProcessStatus struct
type ProcessStatus struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Pid         string `json:"pid"`
}

// GetHealtStatus retuns application status info
func GetHealtStatus(rpcURL string, w http.ResponseWriter, r *http.Request) {
	rpc := gosupervisor.New(rpcURL)

	list := make([]gomap.Mapx, 0)

	client, err := xmlrpc.NewClient(rpc.Url, nil)
	if err != nil {
		health := HealthCheck{Status: "ERROR", Processes: nil}
		respondJSON(w, http.StatusInternalServerError, health)
		return
	}
	err = client.Call("supervisor.getAllProcessInfo", nil, &list)
	if err != nil {
		health := HealthCheck{Status: "ERROR", Processes: nil}
		respondJSON(w, http.StatusInternalServerError, health)
		return
	}

	var processes []ProcessStatus
	var status string = "OK"

	for _, v := range list {
		processStatus := fmt.Sprintf("%v", v.Get("statename"))

		if !(processStatus == "STARTING" || processStatus == "RUNNING") {
			status = "ERROR"
		}

		process := ProcessStatus{
			Name:        fmt.Sprintf("%v", v.Get("name")),
			Status:      processStatus,
			Pid:         fmt.Sprintf("%v", v.Get("pid")),
			Description: fmt.Sprintf("%v", v.Get("description")),
		}
		processes = append(processes, process)
	}

	health := HealthCheck{Status: status, Processes: processes}
	returnStatus := http.StatusOK
	if status == "ERROR" {
		returnStatus = http.StatusInternalServerError
	}

	respondJSON(w, returnStatus, health)
}
