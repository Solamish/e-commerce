package controller

import (
	"e-commerce/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)
func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var (
	err  error
	res  = gin.H{}
	form = &models.Form{}
)

//下单
func PostForm(c *gin.Context) {

	user_id, _ := strconv.Atoi(c.PostForm("user_id"))
	goods_id, _ := strconv.Atoi(c.PostForm("goods_id"))
	shop_id, _ := strconv.Atoi(c.PostForm("shop_id"))
	number, _ := strconv.Atoi(c.PostForm("number"))
	form = &models.Form{
		User_id:  user_id,
		Goods_id: goods_id,
		Shop_id:  shop_id,
		Buy_num:  number,
		Buy_time: time.Now().Format("2006-01-02 15:04:05"),
	}

	//connect MQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "can't connect to MQ")
	defer conn.Close()

	//create amqpChannel
	amqpChannel, err := conn.Channel()
	handleError(err, "can't create a amqpChannel")
	defer amqpChannel.Close()

	//configure Qos
	queue, err := amqpChannel.QueueDeclare("formList", true, false, false, false, nil)
	handleError(err, "Could not declare queue")

	rand.Seed(time.Now().UnixNano())

	body, err := json.Marshal(form)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}
	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	log.Printf("Ordering: %s", string(body))


	res = gin.H{
		"user-id":  user_id,
		"goods-id": goods_id,
		"shop_id":  shop_id,
		"number":   number,
	}
}
