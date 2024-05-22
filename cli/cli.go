// main.go
package cli

import (
	"fmt"
	"os"
	"strings"
	"tw/client"

	"github.com/spf13/cobra"
)

func Start() {
	var rootCmd = &cobra.Command{
		Use:   "tw",
		Short: "tw is a simple CLI application that is used to help developers use command line programs",
		Long:  `tw is a simple CLI application that is used to help developers use command line programs`,
	}

	var howtoCCmd = &cobra.Command{
		Use:   "howtoc [command] [message]",
		Short: "Gives howto on how use a specfic command",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			command := args[0]
			message := strings.Join(args[1:], " ")

			openAPIClient := client.OpenAPIClient{}
			response, err := client.HowToC(openAPIClient, command, message)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(response)

		},
	}

	var howtoCmd = &cobra.Command{
		Use:   "howto [message]",
		Short: "Gives howto on what commands to run to perform a certain task ",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			message := strings.Join(args[0:], " ")

			openAPIClient := client.OpenAPIClient{}
			response, err := client.HowTo(openAPIClient, message)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(response)

		},
	}

	rootCmd.AddCommand(howtoCCmd, howtoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
