package intern

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

func serveProfile(c echo.Context) error {
	return nil
}

type PeeringPeer struct {
	Miner string `json:"miner"`
	Epoch string `json:"epoch"`
}

func GetPower(c echo.Context) error {
	//
	//var params PeeringPeer
	//if err := c.Bind(&params); err != nil {
	//	return err
	//}

	urlApi := os.Getenv("FIL_API_CONFIG")
	urlToken := os.Getenv("FIL_API_TOKEN")

	if urlApi == "" {
		return c.JSON(http.StatusOK, "urlApi is empty")
	}
	if urlToken == "" {
		return c.JSON(http.StatusOK, "urlToken is empty")
	}

	miner := c.QueryParam("miner")
	heigh := c.QueryParam("epoch")
	ctx := context.Background()

	log.Errorf("completed mineOne %+v", miner)
	log.Errorf("completed mineOne %+v", heigh)

	return c.JSON(http.StatusOK, ctx)
}

func ServeAPI() error {
	e := echo.New()
	e.GET("/debug", serveProfile)
	e.GET("/get-power", GetPower)
	return e.Start(":3004")
}
