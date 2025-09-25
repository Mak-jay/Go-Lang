package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)
func main()  {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf("%-25s %-6s %-6s %-50s %-8s %-50s\n","Domain", "MX", "SPF", "SPF Record", "DMARC", "DMARCRecord")
 
    for scanner.Scan(){
        verifyDomain(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error: could not read from input: %v\n", err)	
    }
}

func verifyDomain(domain string){
    var hasMX, hasSPF, hasDMARC bool
    var spfRecord, dmarcRecord string

    // checking for MX records
    mxRecords, err := net.LookupMX(domain)
    if err != nil {
        log.Printf("Error looking up MX records for %s: %v\n", domain, err)
    }
    if len(mxRecords) > 0 {
        hasMX = true
    }

    // checking for SPF records
    txtRecords, err := net.LookupTXT(domain)
    if err != nil {
        log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
    }
    for _, record := range txtRecords {
        if strings.HasPrefix(record, "v=spf1") {
            hasSPF = true
            spfRecord = record
            break
        }
    }

    // checking for DMARC records
    dmarcRecords, err := net.LookupTXT("dmarc." + domain)
    if err != nil {
        log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
    }
    for _, record := range dmarcRecords {
        if strings.HasPrefix(record, "v=dmarc1") {
            hasDMARC = true
            dmarcRecord = record
            break
        }
    }

    // Refined output: tab-separated, one line per domain
	fmt.Printf("%-25s %-6t %-6t %-50s %-8t %-50s\n",domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}