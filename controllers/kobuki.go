package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/vibin18/bot_dash/kobuki_provider"
	"github.com/vibin18/bot_dash/test"
)

func GetDiagBatteryVoltage(ctx iris.Context) {

	message := <-kobuki_provider.Messages
	ctx.JSON(message.Status[0].Values[0])
}

func GetDiagBatteryPercentage(ctx iris.Context) {

	message := <-kobuki_provider.Messages
	ctx.JSON(message.Status[0].Values[1])
}

func GetDiagBatteryCapacity(ctx iris.Context) {

	message := <-kobuki_provider.Messages
	ctx.JSON(message.Status[0].Values[3])
}

func GetDiagBatteryState(ctx iris.Context) {

	message := <-kobuki_provider.Messages
	ctx.JSON(message.Status[0].Values[2])
}

func GetOdom(ctx iris.Context) {
	//	go test.OdomUpdate()
	message := test.OdomMessages
	ctx.JSON(message)
}
