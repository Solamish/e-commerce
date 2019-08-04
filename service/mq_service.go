package main

import (
	"e-commerce/models"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	/**
	第一次写的时候这里没有连接数据库，然后会抛出诡异的异常
	panic: runtime error: invalid memory address or nil pointer dereference
	 */
	db, err := models.InitDB()
	if err != nil {
		log.Println("err open databases", err)
		return
	}
	defer db.Close()


	//connect MQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failError(err, "can't connect to MQ")
	defer conn.Close()

	//create amqpChannel
	amqpChannel, err := conn.Channel()
	failError(err, "can't create a amqpChannel")
	defer amqpChannel.Close()

	//configure Qos
	queue, err := amqpChannel.QueueDeclare("formList", true, false, false, false, nil)
	failError(err, "Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	failError(err, "can't configure Qos")

	//register consumer
	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failError(err, "can't cregister consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", string(d.Body))
			form := &models.Form{}

			err = json.Unmarshal(d.Body, form)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			//将表单信息写入数据库
			err := form.Insert()
			if err != nil {
				log.Println("fail to write to DB", err)
			}

			//更新商铺信息
			shop, _ := models.GetShopByID(form.Shop_id)
			fmt.Print("库存: ", shop.Number)
			fmt.Println(" ,订单购买数量: ", form.Buy_num)
			err = models.Update(form.Goods_id, shop.Number, form.Buy_num)
			if err == nil {
				log.Println("success to update")
			}

			log.Printf("Good: %s", string(d.Body))
			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	//stop the channel
	<-stopChan
}
