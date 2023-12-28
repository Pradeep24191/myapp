/*
Copyright © 2023 pchavana

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"bufio"
	"strings"
	"fmt"
)

var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		if prompt == false && (name == "" || greeting == "") {
		cmd.Usage()
		os.Exit(1)
	}

	if debug {
		fmt.Println("Name : ", name)
		fmt.Println("Greeting : ", greeting)
		fmt.Println("Prompt : ", preview)
		fmt.Println("Preview : ", prompt)
		os.Exit(0)
	}

	// Conditionally read from Stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	message := BuildGreeting(name, greeting)

	// Either preview the message or write to the file	
	if preview {
		fmt.Println(message)
	} else {
		//write content
		file, err := os.OpenFile("welcome.txt", os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()

		handleError(err)

		_, err = file.Write([]byte(message))

		handleError(err)
	}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview message for toggle")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "Toggle to Prompt message or not")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to use in message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting message")

	if os.Getenv("debug") != "" {
		debug = true
	}
}

func renderPrompt() (name, greeting string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your Greeting : ")
	greeting, _  = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Println("Your Name : ")
	name, _  = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}

func BuildGreeting (name, greeting string) string {
	return fmt.Sprintf("%s %s", greeting, name)
} 


