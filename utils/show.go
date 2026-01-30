// Use this only when you wish to run some show commands

// Give a slice of all the commands that you wish to run

// Parse the inputs first and check if you can add json at the end and based on the number of commands, create a new slice of maps ( result) to return

package utils

import (
	"log"
	"golang.org/x/crypto/ssh"
)

//func normalizeCommands(){
//	// Ensure that all the commands return a json output. If | json is not included in the show command, add it!!
//}

func Show( commands []string , Client *ssh.Client) (result map[string]string){
	
	// Run show commands (ONLY!!)
	// For Config commands, we would need to maintain a session and run all the commmands
	// For show commands, use seperate session per command
	result = make(map[string]string)
	for _,command := range commands {
		sess, err := Client.NewSession()
		if err!=nil{
			log.Panicln("Error while creating a new session:", err)
		}
		defer sess.Close()

		output,err := sess.CombinedOutput(command)
		if err!=nil{
			log.Printf("Error while running %s ", command)
			log.Print(err)
		}
		result[command] = string(output)
	}
	return result


}