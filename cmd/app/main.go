package main

import (
	"fmt"

	"github.com/lalankeba/iso8583-go-app/pkg/iso8583"
)

func main() {
	fmt.Println("Hello to sample ISO8583 test app")

	// Build a financial ISO8583 message
	finMsgStream, err := iso8583.BuildFinancialMessage()

	if err != nil {
		fmt.Printf("Error building financial message: %v\n", err)
		return
	}

	fmt.Printf("Built ISO8583 Message (binary stream):\n%v\n", finMsgStream)

	// Parse the financial ISO8583 message
	isoMsg, err := iso8583.ParseMessage(finMsgStream)

	if err != nil {
		fmt.Printf("Error parsing ISO8583 message: %v\n", err)
		return
	}

	fmt.Println("Parsed ISO8583 Message:")
	for i := 2; i <= 128; i++ {
		field := isoMsg.GetField(i)
		if field != nil {
			strValue, err := field.String()
			if err != nil {
				fmt.Printf("Error converting Field %d to string: %v\n", i, err)
				continue
			}
			fmt.Printf("Field %d: %s\n", i, strValue)
		}

	}

}
