package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jindrichskupa/supervisord-statuspage/app/handler"
	"github.com/jindrichskupa/supervisord-statuspage/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	RPCURL string
}

// Initialize application with predefined configuration
func (a *App) Initialize(config *config.Config) {
	a.RPCURL = config.RPCURL
	a.Router = mux.NewRouter()
	a.setRouters()
	log.Println("Supervisord statuspage started")
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/healtz", a.GetHealtStatus)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// GetHealtStatus retuns application status info
func (a *App) GetHealtStatus(w http.ResponseWriter, r *http.Request) {
	handler.GetHealtStatus(a.RPCURL, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
