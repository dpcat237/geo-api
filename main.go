package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/dpcat237/geomicroservices/geolocation"
	"google.golang.org/grpc"

	"gitlab.com/dpcat237/geoapi/config"
	"gitlab.com/dpcat237/geoapi/logger"
	"gitlab.com/dpcat237/geoapi/router"
)

func main() {
	cfg := config.LoadConfigData()
	logg := logger.New()

	connLocation, err := grpc.Dial(cfg.LocAddr, grpc.WithInsecure())
	if err != nil {
		logg.Errorf("Can't connect to location GRPC with error %s", err)
		return
	}
	locCli := geolocation.NewLocationControllerClient(connLocation)

	rtrMng := router.NewManager(locCli)
	rtrMng.CreateRouter()
	rtrMng.LunchRouter(cfg.HTTPport)
	logg.Infof("Router started at on port %s", cfg.HTTPport)

	gracefulStop(cfg, logg, rtrMng)
}

// gracefulStop stops router after receiving system or key interruption
func gracefulStop(cfg config.Config, logg logger.Logger, rtrMng router.Manager) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	close(c)

	rtrMng.Shutdown(cfg.HTTPport)
	logg.Info("Service stopped")
}
