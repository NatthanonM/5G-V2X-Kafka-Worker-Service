package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"5g-v2x-data-management-service/internal/config"
	"5g-v2x-data-management-service/internal/services"
	"time"
	"strconv"
	"strings"
)

// Consumer ...
type Consumer struct {
	config           *config.Config
	AccidentServices *services.AccidentService
	DrowsinessServices *services.DrowsinessService
}

// NewConsumer ...
func NewConsumer(config *config.Config, as *services.AccidentService, ds *services.DrowsinessService) *Consumer {
	c := &Consumer{
		config:           config,
		AccidentServices: as,
		DrowsinessServices: ds,
	}
	return c
}

func (c *Consumer) consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r_acc := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.config.Broker1Address, c.config.Broker2Address, c.config.Broker3Address},
		Topic:   c.config.Topic,
		Logger:  l,
	})
	r_dds := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.config.Broker1Address, c.config.Broker2Address, c.config.Broker3Address},
		Topic:   c.config.TopicDDS,
		Logger:  l,
	})
	layout := "2006-01-02T15:04:05.000000Z"
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r_acc.ReadMessage(ctx)
		
		if err != nil {
			panic("could not read message1 " + err.Error())
		}
		// after receiving the message, log its value
		if msg.Value != nil {
			// var result_acc Accident
			var result map[string]interface{}
			json.Unmarshal([]byte(msg.Value), &result)
			
			if  result["condition"] == "ACS" {
				//do something here
				fmt.Println("ACS")
				username :=  fmt.Sprintf("%v", result["username"])
				carID := fmt.Sprintf("%v", result["carID"])
				lat,err := strconv.ParseFloat(fmt.Sprintf("%v", result["lat"]), 64)
				lng,err := strconv.ParseFloat(fmt.Sprintf("%v", result["lng"]), 64)
				t := strings.Split(fmt.Sprintf("%v", result["time"]), " ")
				timeFormat:= t[0]+"T"+t[1]+"Z"
				time, err := time.Parse(layout,timeFormat)
				if err != nil {
					fmt.Println(err)
				}
				go c.AccidentServices.StoreData(username,carID,lat,lng,time)
			
			}
		}
		msg1, err1 := r_dds.ReadMessage(ctx)
		if err1 != nil {
			panic("could not read message2 " + err1.Error())
		}
		if msg1.Value != nil {
			var result1 map[string]interface{}
			json.Unmarshal([]byte(msg1.Value), &result1)
			if  result1["condition"] == "DDS" {
				//do something here
				fmt.Println("DDS")
				username :=  fmt.Sprintf("%v", result1["username"])
				carID := fmt.Sprintf("%v", result1["carID"])
				lat,err := strconv.ParseFloat(fmt.Sprintf("%v", result1["lat"]), 64)
				lng,err := strconv.ParseFloat(fmt.Sprintf("%v", result1["lng"]), 64)
				t := strings.Split(fmt.Sprintf("%v", result1["time"]), " ")
				timeFormat:= t[0]+"T"+t[1]+"Z"
				time, err := time.Parse(layout,timeFormat)
				responseTime, err := strconv.ParseFloat(fmt.Sprintf("%v", result1["response_time"]), 64)
				workingHour := strconv.ParseFloat(fmt.Sprintf("%v", result1["working_time"]),64)
				if err != nil {
					fmt.Println(err)
				}
				go c.DrowsinessServices.StoreData(username,carID,lat,lng,time,responseTime,workingHour)
			
			}
		}	
	}
}

// Start ...
func (c *Consumer) Start() error {
	// create a new context
	ctx := context.Background()
	c.consume(ctx)
	return nil
}
