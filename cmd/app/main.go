package main

import (
	"fmt"

	"github.com/lalankeba/iso8583-go-app/pkg/iso8583"
)

func main () {
	fmt.Println("Hello to sample ISO8583 test app")
	iso8583.BuildFinancialMessage()

}