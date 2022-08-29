package controllers

import (
	"store/initializers"
	"store/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		COdeProduct string
		NamaProduct string
		HargaProduct int64
		UOM string
	}
	c.Bind(&body)
	post := models.Product{COdeProduct: body.COdeProduct, NamaProduct: body.NamaProduct, HargaProduct: body.HargaProduct, UOM: body.UOM, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	
	result := initializers.DB.Create(&post)

	if  result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"POST": post,
	})
}