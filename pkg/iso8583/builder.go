package iso8583

import (
	"fmt"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/specs"
)

func BuildFinancialMessage() {
	isomessage := iso8583.NewMessage(specs.Spec87ASCII)

	isomessage.MTI("0200")

	isomessage.Field(2, "4242424242424242")

	isomessage.Field(3, "123456")

	isomessage.Field(4, "100")

	// generate binary representation of the message into rawMessage
	rawMessage, err := isomessage.Pack()

	fmt.Println("fin msg:", rawMessage)
	fmt.Printf("fin msg in hex: %x\n", rawMessage)
	fmt.Println("err:", err)

}
