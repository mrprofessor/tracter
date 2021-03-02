package cli

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mrprofessor/tracter/utils"
)

func Run() {
	var url string

	// Subcommands
	whoCommand := flag.NewFlagSet("who", flag.ExitOnError)

	// who sub-command flags
	whoURL := whoCommand.String("url", "", "The URL you are looking for(Required).")
	whoAll := whoCommand.Bool("all", false, "Get all the information.")
	whoBasic := whoCommand.Bool("basic", true, "Get the information that acutally makes sense.")
	flag.Parse()

	// CHECKS
	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	// Check for subcommands
	switch os.Args[1] {
	case "who":
		whoCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check for each subcommands and their cases
	if whoCommand.Parsed() {
		// Mandatory flags
		if *whoURL == "" {
			whoCommand.PrintDefaults()
			os.Exit(1)
		} else {
			url = *whoURL
		}

		if *whoAll {
			whoIsData(url)
		} else if *whoBasic {
			whoIsData(url)
		}
	}

}

// whoIsData looks up the provided url in icann databases and prints/returns the data
func whoIsData(url string) {

	var lookUpUrl = "https://lookup.icann.org/api/whois?q="
	finalQueryUrl := lookUpUrl + url

	// Call the lookup url
	resp, err := http.Get(finalQueryUrl)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(utils.ParseResponse(resp))
}

// getSSL looks up and finds the ssl certificate data
func sslData(url string) {
	fmt.Println("")
}
