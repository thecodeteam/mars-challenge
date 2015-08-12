package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"golang.org/x/net/websocket"
	"time"
	"encoding/json"
	"net/http"
	"bytes"
	
	"strconv"
	
)
type JSONTime time.Time
type Message struct {

	Solarflare	bool   	`json:"solarFlare"`
	Temperature float64 `json:"temperature"`
	Radiation   int32	`json:"radiation"`
	Stamp		time.Time
}


func readSensorMessages(ws *websocket.Conn, incomingMessages chan string, sensor int) {
	for {
		var in [] byte
			if err := websocket.Message.Receive(ws, &in); err != nil {
				return
			}
			incomingMessages <- string(in)
	}
}

func postAvgReading(msg Message,gc string){
		log.Println("=========================================")
		admin_password := "1234"
	    log.Println("game console ulr:",gc)
		log.Println("Aadmin Pass:",admin_password)
		log.Println("AVG Solar Flare:",msg.Solarflare)
		log.Println("AVG Radation:",msg.Radiation)
		log.Println("AVG Temperature:",msg.Temperature)
		log.Println("TimeStamp:",msg.Stamp)
		
		log.Println("=========================================")
		b, err := json.Marshal(msg)
		fmt.Printf("Marshal Json: %s\n",b)
		var jsonStr = []byte(`{"solarFlare":`+strconv.FormatBool(msg.Solarflare)+`,"temperature":`+strconv.FormatFloat(msg.Temperature,'f', -1, 64)+`,"radiation":`+string(msg.Radiation)+`}`)
		//var jsonStr=[]byte(b)
		
		fmt.Printf("Marshal Json2: %s\n",jsonStr)
		
	    req, err := http.NewRequest("POST", gc, bytes.NewBuffer(jsonStr))
    	req.Header.Set("X-Auth-Token", admin_password)
    	req.Header.Set("Content-Type", "application/json")
		
		client := &http.Client{}
		resp, err := client.Do(req)
		 if err != nil {
		      panic(err)
		}
		defer resp.Body.Close()
		
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		
		time.Sleep(time.Second)
}

func getAvgReading(messagelist [5]string,gc string){

	var msg Message
	var avgMsg Message
	avgMsg.Radiation=0
	avgMsg.Temperature=0
	var avgFlare float32=0
	
	var c int = 0
	for i := range messagelist {
		 
		   fmt.Println("sendor #",i,":",messagelist[i])
		   b:=[]byte(messagelist[i])
		   
	 	   err := json.Unmarshal(b, &msg)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				//return
			}else{
				c++;
			/*	
				log.Println("Current Solar Flare:",msg.Solarflare)
				log.Println("Current Radation:",msg.Radiation)
				log.Println("current Temperature:",msg.Temperature)
				log.Println("current TimeStamp:",msg.Stamp)
				*/
				avgMsg.Radiation=avgMsg.Radiation+msg.Radiation
				avgMsg.Temperature=avgMsg.Temperature+msg.Temperature
				if msg.Solarflare{
					avgFlare++
				}
				avgMsg.Stamp=msg.Stamp
			}
			
			
	}
	if c >0 {
		fmt.Println("counter#:",c)
		log.Println("Total Solar Flare:",avgMsg.Solarflare)
		log.Println("Total Radation:",avgMsg.Radiation)
		log.Println("Total Temperature:",avgMsg.Temperature)
		log.Println("Last TimeStamp:",avgMsg.Stamp)
		
		avgMsg.Radiation=avgMsg.Radiation/int32(c)
		avgMsg.Temperature=avgMsg.Temperature/float64(c)
		avgFlare=avgFlare/float32(c)
		if avgFlare>=0.5{
				avgMsg.Solarflare = true
		}
		 postAvgReading(avgMsg,gc)
		
		
	}
}

func main() {

	//Read Env Variables
	sensor_endpoints_str :="104.40.93.11:8080,104.40.93.11:8081,104.40.93.11:8082,104.40.93.11:8083"
	
	
	//sensor_endpoints_str:=os.Getenv("SENSOR_ENDPOINT")
	controller_endpoint := os.Getenv("GC_ENDPOINT")
	controller_endpoint ="104.40.89.227:8080"
	
	log.Printf("Current Sensor EndPoints '%s'", sensor_endpoints_str)
	log.Printf("Current Controller '%s'", controller_endpoint)	
	
	//Prepare Game Controller URL
	controller_url := "http://" + controller_endpoint + "/api/readings"
	log.Printf("Current Sensor EndPoints %s ",controller_url)

	sensor_endpoint_list := strings.Split(sensor_endpoints_str, ",")
	for i := 0; i < len(sensor_endpoint_list); i++ {
		sensor_endpoint_list[i] = "ws://" + sensor_endpoint_list[i] +"/ws"
		log.Printf("Sensor EndPoint '%s'", sensor_endpoint_list[i])
	}

	//var message Message
	var chans [5]chan string
	var incomingMessages [5]string
	
	for i := range chans {
	   chans[i] = make(chan string)
	}

	for i := 0; i < len(sensor_endpoint_list); i++ {
			log.Println("Process Socket:", sensor_endpoint_list[i])
			ws, err := websocket.Dial(sensor_endpoint_list[i], "", sensor_endpoint_list[i])
			if err != nil {
				log.Printf("Dial failed: %s\n", err.Error())
	
			}else{

				log.Println("Print mesg:", chans[i])
				go readSensorMessages(ws, chans[i],i)	
	
			}
	}
	
	for {
		select {
	
		case sensor0 := <- chans[0]:
			//fmt.Println("Message from IncommingMessage1:", sensor0)
			incomingMessages[0]=sensor0
		
		case sensor1 := <- chans[1]:
			//fmt.Println("Message from IncommingMessage2:", sensor1)
			incomingMessages[1]=sensor1
			
		case sensor2 := <- chans[2]:
			//fmt.Println("Message from IncommingMessage3:", sensor2)
			incomingMessages[2]=sensor2
			
		case sensor3 := <- chans[3]:
			//fmt.Println("Message from IncommingMessage4:", sensor3)
			incomingMessages[3]=sensor3
			
		case sensor4 := <- chans[4]:
			//fmt.Println("Message from IncommingMessage4:", sensor4)
			incomingMessages[4]=sensor4
			
		default:
			getAvgReading(incomingMessages,controller_url)
		}
	}
	
	
	
	

}
