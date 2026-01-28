package main

import (
	"log"
	"github.com/venkatakoushikp/goEOS/utils"
)

func main (){
	Client, err := utils.Connect("admin", "admin", "10.18.175.160:22", "tcp")
	if err!=nil{
		log.Println("Error ::", err)
	}
	defer Client.Close()

	sess, err := Client.NewSession()
	if err!=nil{
		log.Println("Error ::", err)
	}
	defer sess.Close()
	output, err := sess.CombinedOutput("show version | json")
	if err!=nil{
		log.Println("Error ::", err)
	}
	log.Println(string(output))

}
