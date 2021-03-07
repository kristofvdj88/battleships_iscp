package admapi

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/webapi/routes"
	"github.com/iotaledger/wasp/plugins/gracefulshutdown"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

func addShutdownEndpoint(adm echoswagger.ApiGroup) {
	adm.GET(routes.Shutdown(), handleShutdown).
		SetSummary("Shut down the node")
}

func handleShutdown(c echo.Context) error {
	log.Info("Received a shutdown request from WebAPI.")
	gracefulshutdown.Shutdown()
	return c.String(http.StatusOK, "Shutting down...")
}
