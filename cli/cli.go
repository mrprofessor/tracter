package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mrprofessor/tracter/utils"
	"github.com/fatih/color"
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

// getSSL looks up and finds the ssl certificate data
func sslData(url string) {
	fmt.Println("")
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

	// Parse json form response body
	parsedJson := utils.ParseResponse(resp)

	// Pretty print in the terminal
	// fmt.Println(utils.PrettyPrintJSON(parsedJson))

	// TODO
	// Pretty print essential info only
	var respStr WhoIsResponse
	// var respStr interface{}
	err = json.Unmarshal(parsedJson, &respStr)
	if err != nil {
		fmt.Println(err)
	}

	prettyPrintWhoIs(respStr)
}

// Print only essential data on the terminal
func prettyPrintWhoIs(resp WhoIsResponse) {

	hr := "-------------------------------------------------------"
	half_hr := "---------------------------"

	y := color.New(color.FgYellow).Add(color.Underline).Add(color.Bold)

	// Registrar info
	fmt.Println(hr)
	y.Println("Registrar Info")
	fmt.Println(hr)
	fmt.Printf("Name: %s\n", resp.Records[0].Registrar.Registrar)
	fmt.Printf("Whois Server: %s\n", resp.Records[0].Whoisserver)
	fmt.Printf("Referral URL: %s\n", resp.Records[0].Referralurl)
	fmt.Printf("Domain Status: %s\n", WhoIsDomainStatus(resp.Records[0].Domainstatus))

	// Dates
	fmt.Println(hr)
	y.Println("Important Dates")
	fmt.Println(hr)
	fmt.Printf("Registered On: %s\n", WhoIsNotProvided(resp.Records[0].Creationdate))
	fmt.Printf("Updated On: %s\n", WhoIsNotProvided(resp.Records[0].Updateddate))
	fmt.Printf("Expires On: %s\n", WhoIsNotProvided(resp.Records[0].Registryexpirydate))

	// Nameservers
	fmt.Println(hr)
	y.Println("Nameservers")
	fmt.Println(hr)
	for _, server := range resp.Records[0].Nameserver {
		fmt.Println(server)
	}

	// Registrant info
	fmt.Println(half_hr)
	y.Println("Registrant Contact Info")
	fmt.Println(half_hr)
	fmt.Printf("Name: %s\n", resp.Records[0].Registrant.Name)
	fmt.Printf("Organization: %s\n", resp.Records[0].Registrant.Organization)
	fmt.Printf("Street: %s\n", resp.Records[0].Registrant.Street)
	fmt.Printf("City: %s\n", resp.Records[0].Registrant.City)
	fmt.Printf("State / Province: %s\n", resp.Records[0].Registrant.StateProvince)
	fmt.Printf("Country: %s\n", resp.Records[0].Registrant.Country)
	fmt.Printf("Phone: %s\n", resp.Records[0].Registrant.Phone.Number)
	fmt.Printf("email: %s\n", resp.Records[0].Registrant.Email)

	fmt.Println(half_hr)
	y.Println("Administrative Contact Info")
	fmt.Println(half_hr)
	fmt.Printf("Name: %s\n", resp.Records[0].Admin.Name)
	fmt.Printf("Organization: %s\n", resp.Records[0].Admin.Organization)
	fmt.Printf("Street: %s\n", resp.Records[0].Admin.Street)
	fmt.Printf("City: %s\n", resp.Records[0].Admin.City)
	fmt.Printf("State / Province: %s\n", resp.Records[0].Admin.StateProvince)
	fmt.Printf("Country: %s\n", resp.Records[0].Admin.Country)
	fmt.Printf("Phone: %s\n", resp.Records[0].Admin.Phone.Number)
	fmt.Printf("email: %s\n", resp.Records[0].Admin.Email)

	fmt.Println(half_hr)
	y.Println("Technical Contact Info")
	fmt.Println(half_hr)
	fmt.Printf("Name: %s\n", resp.Records[0].Tech.Name)
	fmt.Printf("Organization: %s\n", resp.Records[0].Tech.Organization)
	fmt.Printf("Street: %s\n", resp.Records[0].Tech.Street)
	fmt.Printf("City: %s\n", resp.Records[0].Tech.City)
	fmt.Printf("State / Province: %s\n", resp.Records[0].Tech.StateProvince)
	fmt.Printf("Country: %s\n", resp.Records[0].Tech.Country)
	fmt.Printf("Phone: %s\n", resp.Records[0].Tech.Phone.Number)
	fmt.Printf("email: %s\n", resp.Records[0].Tech.Email)

}

func WhoIsNotProvided(data string) string {
	if data != "" {
		return data
	}
	return "NOT FOUND"
}

func WhoIsDomainStatus(rawStatus []string) string {
	resultStatus := ""
	for i, status := range rawStatus {
		resultStatus += strings.Split(status, " ")[0]

		// Don't add comma after the last element.
		if i < len(rawStatus)-1 {
			resultStatus += ", "
		}
	}
	return resultStatus
}
