package transport

import (
	"asteroids-neows/internal/entity"
	"asteroids-neows/internal/interfaces"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Data []entity.AsteroidsReport `json:"neo_counts"`
}

func NewRouter(router *gin.Engine, uc interfaces.Asteroids) {
	router.GET("/neo/count", func(c *gin.Context) {
		dates := c.QueryArray("dates")

		data, err := uc.Get(dates)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}

		jsonData, _ := json.MarshalIndent(data, "", "  ")

		c.String(http.StatusOK, string(jsonData))
	})

	router.POST("/neo/count", func(c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := uc.Create(body.Data); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.String(http.StatusOK, "Данные успешно сохранены!")
	})

	router.GET("/neo/count/clear", func(c *gin.Context) {
		if err := uc.Clear(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.String(http.StatusOK, "Все данные удалены!")
	})
}
