package client

import "fmt"

type Client interface {
	SendRequest(string) (string, error)
}

func HowTo(client Client, command string, message string) (string, error) {

	formattedString := fmt.Sprintf(`context:You are a Linux Expert speaking about the %s command. 
									The user is going to ask a question about how to do a ceratin task using that command.
									Answer only with the command
									task: %s`, command, message)

	return client.SendRequest(formattedString)

}
