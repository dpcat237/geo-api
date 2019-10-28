package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/dpcat237/geomicroservices/geolocation"
	"golang.org/x/net/context"

	"gitlab.com/dpcat237/geoapi/controller"
	"gitlab.com/dpcat237/geoapi/logger"
	"gitlab.com/dpcat237/geoapi/model"
)

// apiVersion defines current API version
const apiVersion = "v1"

// Manager is router's manager
type Manager struct {
	locCli geolocation.LocationControllerClient
	logg   logger.Logger
	rtr    *mux.Router
	srv    *http.Server
}

// NewManager initializes router manager
func NewManager(locCli geolocation.LocationControllerClient) Manager {
	return Manager{
		locCli: locCli,
	}
}

// CreateRouter creates router
func (mng *Manager) CreateRouter() {
	mng.rtr = mux.NewRouter().StrictSlash(true)
	mng.addRoutes(mng.rtr, mng.getSysRoutes(), true)
	mng.rtr.Handle("/debug/vars", http.DefaultServeMux)
	mng.addRoutes(mng.rtr.PathPrefix("/"+apiVersion).Subrouter(), mng.getV1Routes(), true)
}

// LunchRouter runs router listener
func (mng *Manager) LunchRouter(port string) {
	mng.srv = &http.Server{Addr: ":" + port, Handler: mng.rtr}
	go func() {
		err := mng.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			mng.logg.Errorf("Error starting http service: %s", err)
		}
	}()
}

// Shutdown close router's connection
func (mng *Manager) Shutdown(port string) {
	if err := mng.srv.Shutdown(context.Background()); err != nil {
		mng.logg.Errorf("Error stopping http service %s", err)
	}
}

// addRoutes set route for router
func (mng *Manager) addRoutes(r *mux.Router, routes []model.Route, useStatusCode bool) {
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

// getSysRoutes sets system routes
func (mng *Manager) getSysRoutes() []model.Route {
	newRoute := func(m, p string, h http.HandlerFunc, n string) model.Route {
		return model.Route{Name: n, Method: m, Pattern: p, HandlerFunc: h}
	}
	return []model.Route{
		newRoute(http.MethodGet, "/services/health", mng.healthCheck, "Check service health"),
	}
}

// getV1Routes sets version 1 routes
func (mng *Manager) getV1Routes() []model.Route {
	// Initialize controllers
	locCnt := controller.NewLocation(mng.locCli)

	newRoute := func(m, p string, h http.HandlerFunc, n string) model.Route {
		return model.Route{Name: n, Method: m, Pattern: p, HandlerFunc: h}
	}

	return []model.Route{
		/** Location **/
		newRoute(http.MethodGet, "/location/{ip}", locCnt.GetLocationByIP, "Get location details by IP address"),
	}
}

// healthCheck checks service health
func (mng *Manager) healthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`{"success": true}`)); err != nil {
		mng.logg.Errorf("Error returning health check %s", err)
	}
}
