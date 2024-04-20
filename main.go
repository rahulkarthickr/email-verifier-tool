package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read input from user
	scanner := bufio.NewScanner(os.Stdin)
	// Print headers for the output
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	// Loop through each line of input
	for scanner.Scan() {
		// Call the checkDomain function with the text from the input
		checkDomain(scanner.Text())
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		// Print an error message if there was an error reading from input
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

// Function to check DNS records for a given domain
func checkDomain(domain string) {
	// Declare variables to store the results
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Lookup MX records for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		// Print an error message if there was an error looking up MX records
		log.Printf("Error: %v\n", err)
	}
	// If MX records were found, set hasMX to true
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Lookup TXT records for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		// Print an error message if there was an error looking up TXT records
		log.Printf("Error: %v\n", err)
	}

	// Loop through each TXT record
	for _, record := range txtRecords {
		// Check if the record starts with "v=spf1"
		if strings.HasPrefix(record, "v=spf1") {
			// If so, set hasSPF to true and store the SPF record
			hasSPF = true
			spfRecord = record
			// Exit the loop since we found the SPF record
			break
		}
	}

	// Lookup DMARC records for the domain
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		// Print an error message if there was an error looking up DMARC records
		log.Printf("Error: %v\n", err)
	}

	// Loop through each DMARC record
	for _, record := range dmarcRecords {
		// Check if the record starts with "v=DMARC1"
		if strings.HasPrefix(record, "v=DMARC1") {
			// If so, set hasDMARC to true and store the DMARC record
			hasDMARC = true
			dmarcRecord = record
			// Exit the loop since we found the DMARC record
			break
		}
	}

	// Print the results for the domain
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
