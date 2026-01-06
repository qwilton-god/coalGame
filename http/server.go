package myhttp

import (
	"coalGame/company"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(game *company.Company) {

	myHandlers := NewHandler(game)

	r := gin.Default()
	// Шахтеры
	r.GET("/miners/info", myHandlers.GetMinerPriceInfo) // инфо о ценах
	r.GET("/miners", myHandlers.GetActiveMinersByClass) // список актиных по ?class=
	r.POST("miners", myHandlers.BuyMiner)               // покупка

	// Оборудование
	r.GET("/equipment/info", myHandlers.GetEquipmentInfo) // инфо о ценах
	r.GET("/equipment", myHandlers.GetEquipmentInfo)      // что куплено
	r.POST("/equipment", myHandlers.BuyEquipment)         // покупака

	// Группа игра
	r.GET("/stats", myHandlers.GetStatsHandler) // текущая статистика
	r.POST("/stop", myHandlers.StopGame)        // попробовать остановить игру

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
