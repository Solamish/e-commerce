package middlewares
//
//import (
//	"e-commerce/models"
//	"encoding/json"
//	"github.com/streadway/amqp"
//	"log"
//	"math/rand"
//	"time"
//)
//
//func handleError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}
//
//func Queue(form *models.Form)  {
//
//
//		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//		handleError(err, "can't connect to MQ")
//		defer conn.Close()
//
//		amqpChannel, err := conn.Channel()
//		handleError(err, "can't create a amqpChannel")
//		defer amqpChannel.Close()
//
//		queue, err := amqpChannel.QueueDeclare("formList", true, false, false, false, nil)
//		handleError(err, "Could not declare queue")
//
//		rand.Seed(time.Now().UnixNano())
//
//		body, err := json.Marshal(form)
//		if err != nil {
//			handleError(err, "Error encoding JSON")
//		}
//		err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
//			DeliveryMode: amqp.Persistent,
//			ContentType:  "text/plain",
//			Body:         body,})
//
//		if err != nil {
//			log.Fatalf("Error publishing message: %s", err)
//		}
//		log.Printf("AddGood: %s", string(body))
//	}
