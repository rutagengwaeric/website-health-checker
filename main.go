package main

import (
	"fmt"
	"log"
	"os"

	// Importing the cli package from the urfave/cli library
	"github.com/urfave/cli/v2"
)

func main() {
	// Define a new CLI application
	app := &cli.App{
		// Set the name of the application (empty here, but usually should be descriptive)
		Name: "",

		// Set a brief description of what the application does
		Usage: "Tiny tool to check whether a website is running or down",

		// Define the flags that the application can accept
		Flags: []cli.Flag{
			// Define a String flag for the domain name, which is required
			&cli.StringFlag{
				Name:     "domain",               // The flag's name
				Aliases:  []string{"d"},          // Short version of the flag
				Usage:    "Domain name to check", // Description of the flag
				Required: true,                   // Marking the flag as required
			},

			// Define a String flag for the port number, which is optional
			&cli.StringFlag{
				Name:     "port",                 // The flag's name
				Aliases:  []string{"p"},          // Short version of the flag
				Usage:    "Port number to check", // Description of the flag
				Required: false,                  // The flag is not required
			},
		},

		// Define the action to be executed when the CLI is run
		Action: func(c *cli.Context) error {
			// Retrieve the port value from the flag; if empty, default to "80"
			port := c.String("port")
			if c.String("port") == "" {
				port = "80" // Default to port 80 if no port is provided
			}

			// Call the Check function with the domain and port, and store the result
			status := Check(c.String("domain"), port)

			// Print the status of the website to the console
			fmt.Println(status)
			return nil
		},
	}

	// Run the application with the command-line arguments
	err := app.Run(os.Args)
	if err != nil {
		// Log any errors that occur while running the application
		log.Fatal(err)
	}
}
