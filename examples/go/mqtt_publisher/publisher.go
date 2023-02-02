package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/tidwall/gjson"
)

func main() {
	wg := sync.WaitGroup{}

	// paho mqtt client options
	opts := paho.NewClientOptions()
	opts.SetCleanSession(true)
	opts.SetConnectRetry(false)
	opts.SetAutoReconnect(false)
	opts.SetProtocolVersion(4)
	opts.SetClientID("dash-producer")
	opts.AddBroker("127.0.0.1:4057")
	opts.SetKeepAlive(3 * time.Second)
	opts.SetUsername("user")
	opts.SetPassword("pass")

	// connect to server with paho mqtt client
	client := paho.NewClient(opts)
	connectToken := client.Connect()
	connectToken.WaitTimeout(1 * time.Second)
	if connectToken.Error() != nil {
		panic(connectToken.Error())
	}

	checkTableDone := false
	tableExists := false

	client.Subscribe("db/reply", 1, func(_ paho.Client, msg paho.Message) {
		defer wg.Done()

		buff := msg.Payload()
		str := string(buff)
		fmt.Println("RECV:", str)
		vSuccess := gjson.Get(str, "success")
		if !vSuccess.Bool() {
			return
		}
		if !checkTableDone {
			checkTableDone = true
			vCount := gjson.Get(str, "data.rows.0.0")
			tableExists = vCount.Int() == 1
		}
	})

	// check table existence
	jsonStr := `{ "q": "select count(*) from M$SYS_TABLES where name = 'TAGDATA'" }`
	wg.Add(1)
	client.Publish("db/query", 1, false, []byte(jsonStr))

	wg.Wait()
	if tableExists {
		// drop table
		jsonStr = `{ "q": "drop table TAGDATA" }`
		wg.Add(1)
		client.Publish("db/query", 1, false, []byte(jsonStr))
		wg.Wait()
	}

	// create table
	jsonStr = `{
		"q": "create tag table TAGDATA (name varchar(200) primary key, time datetime basetime, value double summarized, jsondata json)"
	}`
	wg.Add(1)
	client.Publish("db/query", 1, false, []byte(jsonStr))
	wg.Wait()

	quitChan := make(chan os.Signal)
	alive := true

	// start generator
	wg.Add(1)
	go func() {
		rand.Seed(time.Now().Unix())
		for alive {
			jsonStr = fmt.Sprintf(`[ "series-1", %d, %.5f, null ]`, time.Now().UTC().UnixNano(), rand.Float32())
			rt := client.Publish("db/append/TAGDATA", 1, false, []byte(jsonStr))

			// if publish was not successful
			if !rt.WaitTimeout(1 * time.Second) {
				fmt.Println("no reponse from server")
				quitChan <- os.Interrupt
			} else if err := rt.Error(); err != nil {
				fmt.Println("ERR", err.Error())
				quitChan <- os.Interrupt
			}
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	// wait signal
	signal.Notify(quitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan

	// wait generator to finish
	alive = false
	wg.Wait()

	// disconnect mqtt connection
	client.Disconnect(100)
}
