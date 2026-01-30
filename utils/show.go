// Use this only when you wish to run some show commands

// Give a slice of all the commands that you wish to run

// Parse the inputs first and check if you can add json at the end and based on the number of commands, create a new slice of maps ( result) to return

package utils

import (
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func normalizeCommands( commands []string) ( commands_parsed []string){
	// Ensure that all the commands return a json output. If | json is not included in the show command, add it!!
	for _, cmd := range commands {
		if strings.HasPrefix( cmd, "show"){
			if strings.HasSuffix(cmd, "| json") || strings.HasSuffix(cmd, "|json"){
				commands_parsed = append(commands_parsed, cmd)
			}else{
				commands_parsed = append(commands_parsed, cmd+" | json")
			}
		}else{
			log.Printf(" Ignoring command : %s, Not a Show Command", cmd)
		}
	}
	return
}

func Show( commands []string , Client *ssh.Client) (map[string]map[string]interface{}){
	
	// Run show commands (ONLY!!)
	// For Config commands, we would need to maintain a session and run all the commmands
	// For show commands, use seperate session per command
	var(
		result = make(map[string]map[string]interface{})
	)
	commands_to_exec := normalizeCommands(commands)
	for _,command := range commands_to_exec {
		var (
			data map[string]interface{}
		)
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
		err = json.Unmarshal(output, &data)
		if err!=nil{
			log.Fatalln("Error while Converting the output to json :", err)
		}
		result[command] = data		

	}
	return result


}