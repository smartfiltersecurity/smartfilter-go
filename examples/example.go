package main

import (
	"fmt"
	"github.com/prevoty/smartfilter-go/client"
)

const (
	apiKey = "api key goes here"
	ruleKey = "rule key goes here"
)

func main() {
	input := "the <script>alert('quick brown fox');</script> jumps over the lazy dog & mouse"

	// Create an instance of SmartFilter
	sf := client.NewSmartFilterClient(apiKey)

	// Verify the API key
	verified, verifiedErr := sf.Verify()
	if verified {
		// Get API key information
		info, infoErr := sf.Info()
		if infoErr == nil {
			fmt.Println("Information:", info.Message)
			// Verify rule
			verification, verifyErr := sf.VerifyRule(ruleKey)
			fmt.Println("Verified rule:", verification, verifyErr)
			// Filter XSS
			result, filterErr := sf.Filter(input, ruleKey)
			fmt.Println("Filtered output:", result.Output, filterErr)
		} else {
			fmt.Println("Could not get information")
		}
	} else {
		fmt.Println("API key not verified", verifiedErr)
	}
}
