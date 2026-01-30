package main

import (
	_ "encoding/json"
	"log"

	"github.com/venkatakoushikp/goEOS/utils"
)

func main (){
	Client, err := utils.Connect("admin", "", "10.85.128.51:22", "tcp")
	if err!=nil{
		log.Fatalln("Error ::", err)
	}
	defer Client.Close()
	commands := []string{
		"show version",
		"show lldp Neighbor",
	}
	result := utils.Show(commands, Client)

	for k,v := range result {
		log.Println(k)
		log.Println(v)
		log.Println("====================")
	}

//	var data map[string]interface{}
//
//	err = json.Unmarshal(output, &data)
//	if err!=nil{
//		log.Fatalln(err)
//	}
//
//	log.Println(data["modelName"])



}


