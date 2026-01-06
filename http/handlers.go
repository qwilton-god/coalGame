package myhttp

import (
	"coalGame/company"
	myhttp "coalGame/http/in"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Game *company.Company
}

func NewHandler(game *company.Company) *Handlers {
	return &Handlers{
		Game: game,
	}
}

func (h *Handlers) GetStatsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Game.GetStats())
}
func (h *Handlers) StopGame(ctx *gin.Context) {
	result, err := h.Game.EndGame()
	if err != nil {
		ctx.JSON(http.StatusMethodNotAllowed, err)
	}
	ctx.JSON(http.StatusOK, result)
}
func (h *Handlers) GetMinerPriceInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Game.GetMinerPriceInfo())
}

func (h *Handlers) GetActiveMinersByClass(ctx *gin.Context) {
	list, err := h.Game.GetActiveMinersOnClass(ctx.Query("class"))
	fmt.Println(list)
	fmt.Println(ctx.Query("class"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (h *Handlers) BuyMiner(ctx *gin.Context) {
	var req myhttp.BuyMinerReq

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.Game.BuyMiner(req.Class); err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(201, nil)
}

func (h *Handlers) GetEquipmentInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Game.EquipmentInfo())
}

func (h *Handlers) GetBuyedEquipment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Game.EquipmentBuyedInfo())
}

func (h *Handlers) BuyEquipment(ctx *gin.Context) {
	var req myhttp.BuyProductReq

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.Game.BuyProduct(req.Name); err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(201, nil)
}
