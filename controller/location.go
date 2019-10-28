package controller

import (
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/context"

	"gitlab.com/dpcat237/geoapi/model"
	"gitlab.com/dpcat237/geomicroservices/geolocation"
)

// LocationController is an interface for Location controller methods
type LocationController interface {
	GetLocationByIP(w http.ResponseWriter, r *http.Request)
}

// locationController is a Location controller
type locationController struct {
	locCli geolocation.LocationControllerClient
}

// NewLocation initializes the Location controller
func NewLocation(locCli geolocation.LocationControllerClient) *locationController {
	return &locationController{locCli: locCli}
}

// GetLocationByIP returns last geo location details by IP address
func (ctr *locationController) GetLocationByIP(w http.ResponseWriter, r *http.Request) {
	ipAddStr := getVariable(r, "ip")
	if ipAddStr == "" {
		returnBadRequest(w, model.NewErrorBadRequest("IP address is required"))
		return
	}
	ipAdd := net.ParseIP(ipAddStr)
	if ipAdd == nil {
		returnBadRequest(w, model.NewErrorBadRequest("IP address is not valid"))
		return
	}

	req := geolocation.GetLocationByIDRequest{
		IpAddress: ipAdd.String(),
	}
	resp, err := ctr.locCli.GetLocationByID(context.Background(), &req)
	if err != nil {
		returnFailed(w, model.NewErrorServer("Error to get location details").WithError(err))
		return
	}
	if int(resp.Code) != http.StatusOK {
		returnFailed(w, model.NewErrorServer(fmt.Sprintf("%s. Status %d", resp.Message, resp.Code)))
		return
	}
	if resp.Location.Id == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	returnJson(w, model.LocationFromGRPC(resp.Location))
}
