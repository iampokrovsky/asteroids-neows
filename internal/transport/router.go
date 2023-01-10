package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
	"github.com/pokrovsky-io/neows-asteroids/internal/interfaces"
	"net/http"
)

type Body struct {
	Data []entity.AsteroidsReport `json:"neo_counts"`
}

func NewRouter(router *gin.Engine, uc interfaces.Asteroids) {
	router.GET("/neo/count", func(c *gin.Context) {
		dates := c.QueryArray("dates")
		// TODO: Проверить, является ли массива пустым

		data, err := uc.Get(dates...)
		if err != nil {
			// TODO: Обработать ошибку
		}

		c.JSON(http.StatusOK, data)
	})

	router.POST("/neo/count", func(c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := uc.Create(body.Data...); err != nil {
			// TODO: Обработать ошибку
		}

		// TODO: Подумать, какой лучше сделать ответ
		c.JSON(http.StatusAccepted, body.Data)
	})
}
