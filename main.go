package main

import (
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
		"show hostname",
	}
	confg := []string{
		"hostname LP710",
	}
	result, _ := utils.Show(commands, Client)
	log.Println(result["show hostname | json"]["hostname"])
	err = utils.Config(confg, Client)
	if err!=nil{
		log.Println(err)
	}
	result, _ = utils.Show(commands, Client)
	log.Println(result["show hostname | json"]["hostname"])


}


