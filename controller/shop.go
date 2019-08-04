package controller

import (
	"e-commerce/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateShop(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("shop-id"))
	shop_id := uint(id)
	goods_id, _ := strconv.Atoi(c.PostForm("goods-id"))
	number, _ := strconv.Atoi(c.PostForm("number"))

	shop := &models.Shop{
		BaseModel: models.BaseModel{
			ID: shop_id,
		},
		Goods_id: goods_id,
		Number:   number,
	}
	err := shop.Insert()
	if err != nil {
		log.Println("fail to insert", err)
	}
}
