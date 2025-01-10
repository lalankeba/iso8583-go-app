package main

import (
	"fmt"

	"github.com/lalankeba/iso8583-go-app/pkg/iso8583"
	iso8583moov "github.com/moov-io/iso8583"
)

func main() {
	fmt.Println("ISO8583 Test App")

	finMsgStream := buildAndDisplayMessage()

	parseAndDisplayMessage(finMsgStream)

}

func buildAndDisplayMessage() []byte {
	finMsgStream, err := iso8583.BuildFinancialMessage()
	if err != nil {
		fmt.Printf("Error building financial message: %v\n", err)
		return nil
	}

	fmt.Printf("Built ISO8583 Message (binary stream):\n%v\n", finMsgStream)
	return finMsgStream
}

func parseAndDisplayMessage(finMsgStream []byte) {
	if finMsgStream == nil {
		fmt.Println("No ISO8583 message to parse.")
		return
	}

	isoMsg, err := iso8583.ParseMessage(finMsgStream)
	if err != nil {
		fmt.Printf("Error parsing ISO8583 message: %v\n", err)
		return
	}

	// Display the MTI
	mti, _ := isoMsg.GetMTI()
	fmt.Printf("\nParsed ISO8583 Message:\n")
	fmt.Printf("MTI: %v\n", mti)

	// Print each field
	printFields(isoMsg)
}

func printFields(isoMsg *iso8583moov.Message) {
	fmt.Printf("%s | %s | %-42s | %s\n", "Field", "Length", "Description", "Value")
	fmt.Println("----- | ------ | ------------------------------------------ | ---------------------")

	for i := 2; i <= 128; i++ {
		field := isoMsg.GetField(i)
		if field != nil {
			strValue, err := field.String()
			if err != nil {
				fmt.Printf("Error converting Field %d to string: %v\n", i, err)
				continue
			}

			// Skip empty fields
			if strValue == "" {
				continue
			}

			// Get field details
			desc := field.Spec().Description
			length := field.Spec().Length

			// Print the field
			fmt.Printf("%5d | %6v | %-42s | %s\n", i, length, desc, strValue)
		}
	}
}
