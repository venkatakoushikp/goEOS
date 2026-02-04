/*

Use this package only when runnign some config commands
Creates a new config session and runs the commands and commits the changes
If there are any show commands that are supposed to be run, create a seperate session and use it

Do not Mention enable, Config again.
Directly specify the config that has to be pushed.

*/

package utils

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

const (
	CommandsNeeded = "enable\nconfigure terminal\n+configure session goEOS-1493101506"
)


func commandParser(commands []string) (string) {

	// ssh.CombinedOutput() expects a string. However we need to maintain the session to be able to run multiple commands.
	// Hence we append everything into a single string and then return it to be used as command.
	// Adds CommandsNeeded Const for the header to ensure Config mode is entered.
	// Also makes sure a new config session is created and the command is run within that session.
	// If any errors are thrown, just exit the config session 
	// If not errors, commit the config session.

	configSession := "configure session goEOS-" + strconv.Itoa(rand.Intn( int(math.Pow(2,31)-1) ))
	log.Println("Config Session Created :", configSession)
	commands = append(commands, configSession)
	return CommandsNeeded+strings.Join(commands, "\n")

}


func Config(Commands []string, Client *ssh.Client) (error){

	session, err := Client.NewSession()
	if err !=nil {
		log.Println("Error while creating a new session :", err)
	}

	_, err = session.CombinedOutput(commandParser(Commands))
	if err!=nil{
		return err
	}

	return nil


}

func validateSession() {

}