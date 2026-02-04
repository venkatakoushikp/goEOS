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
	"strings"

	"golang.org/x/crypto/ssh"
)

func singleString(commands []string) (string) {

	// Append enable and Config commands to the final string 
	// ssh.CombinedOutput() expects a string. However we need to maintain the session to be able to run multiple commands.
	// Hence we append everything into a single strinf and then return it to be used as command.


	var result string = "enable\nconfig\n"
	return result+strings.Join(commands, "\n")
}


func Config(Commands []string, Client *ssh.Client) (error){

	

	session, err := Client.NewSession()
	if err !=nil {
		log.Println("Error while creating a new session :", err)
	}

	_, err = session.CombinedOutput(singleString(Commands))
	if err!=nil{
		return err
	}

	return nil


}


func CreateConfigSession ():
/*

Creates a new config session where the config commands are run.
If there are no errors, commit the session
If there are error, throw those errors 

*/