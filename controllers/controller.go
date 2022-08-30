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
	c.JSON(201, gin.H{
		"POST": post,
	})
}

func PostIndex(c *gin.Context) {
	var post []models.Product
	initializers.DB.Find(&post)

	c.JSON(200, gin.H{
		"products":post,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Product
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"product":post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		ID uint
		COdeProduct string
		NamaProduct string
		HargaProduct int64
		UOM string
	}
	c.Bind(&body)
	var post models.Product
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Product{
		COdeProduct: body.COdeProduct, 
		NamaProduct: body.NamaProduct, 
		HargaProduct: body.HargaProduct, 
		UOM: body.UOM, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	c.JSON(201, gin.H{
		"post":post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Product{}, id)
	c.Status(200)
}