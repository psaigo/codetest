package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svr *Service) GetCurrency(c *gin.Context) {
	symbol := c.Param("symbol")

	if len(symbol) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "please provide the symbol")
	}

	currency, err := svr.Client.GetCurrency(symbol)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, currency)
}

func (svr *Service) GetAllCurrency(c *gin.Context) {
	currency, err := svr.Client.GetAllCurrency()

	fmt.Println(currency)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, currency)
}
