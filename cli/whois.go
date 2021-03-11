package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/mrprofessor/tracter/utils"
)

// whoIsData looks up the provided url in in icann databases and prints/returns
// the data.
func WhoIsData(url string, all bool) {

	var lookUpUrl = "https://lookup.icann.org/api/whois?q="
	finalQueryUrl := lookUpUrl + url

	// Call the lookup url
	resp, err := http.Get(finalQueryUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse json form response body
	parsedJson := utils.ParseResponse(resp)

	if all {
		// Pretty print json response in the terminal
		fmt.Println(utils.PrettyPrintJSON(parsedJson))
	} else {

		var respStr WhoIsResponse
		// var respStr interface{}
		err = json.Unmarshal(parsedJson, &respStr)
		if err != nil {
			fmt.Println(err)
		}

		prettyPrintWhoIs(respStr)
	}
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
