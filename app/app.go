package app

import (
	"github.com/kataras/iris/v12"
	"github.com/iris-contrib/middleware/cors"
	"github.com/vibin18/bot_dash/controllers"
)

func StartApp() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true})
	app.Use(crs)
	v1 := app.Party("/api")
	{
		kobukiApi := v1.Party("/kobuki")
		{
			kobukiApi.Get("/battery/voltage", controllers.GetDiagBatteryVoltage)
			kobukiApi.Get("/battery/percentage", controllers.GetDiagBatteryPercentage)
			kobukiApi.Get("/battery/capacity", controllers.GetDiagBatteryCapacity)
			kobukiApi.Get("/battery/state", controllers.GetDiagBatteryState)
			kobukiApi.Get("/odom", controllers.GetOdom)
		}
	}
	// GET: http://localhost:2000/api/kobuki/battery/state
	app.Listen(":2000")


}