/*

Invoke this whenever you want to connect to the switch.
Returns a Client connection pointer

*/

package utils

import (
	"golang.org/x/crypto/ssh"
)

func Connect(
	Username string,
	Password string,
	IP_port string, // Include port number as well here
	Conn_type string,
) (Client *ssh.Client, err error){
	conf :=& ssh.ClientConfig{
		User: Username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(
				func(name, instruction string, questions []string, echos []bool) (answers []string, err error) {
					answers = make([]string, len(questions))
					for i := range questions{
						answers[i] = Password
					}
					return answers, nil
				},
			),
			
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	Conn, err := ssh.Dial(
		Conn_type ,   
		IP_port, 			// Make sure that port number is added here.
		conf,
	)

	if err!=nil {
		return nil, err
	}

	return Conn, nil

}
