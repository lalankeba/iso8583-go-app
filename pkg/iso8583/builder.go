package iso8583

import (
	"github.com/moov-io/iso8583"
)

func BuildFinancialMessage() ([]byte, error) {
	isomessage := iso8583.NewMessage(ExtendedSpec)

	isomessage.MTI("0200")

	isomessage.Field(2, "5642570404782927")
	isomessage.Field(3, "011000")
	isomessage.Field(4, "780")
	isomessage.Field(7, "1220145711")
	isomessage.Field(11, "101183")
	isomessage.Field(12, "145711")
	isomessage.Field(13, "1220")
	isomessage.Field(14, "2408")
	isomessage.Field(15, "1220")
	isomessage.Field(18, "6011")
	isomessage.Field(22, "051")
	isomessage.Field(25, "00")
	isomessage.Field(26, "04")
	isomessage.Field(28, "C00000000")
	isomessage.Field(30, "C00000000")
	isomessage.Field(32, "56445700")
	isomessage.Field(37, "567134101183")
	isomessage.Field(41, "N1742   ")
	isomessage.Field(42, "ATM004         ")
	isomessage.Field(43, "45 SR LEDERSHIP DUABANAT NUEVA ECIJAQ PH")
	isomessage.Field(49, "608")
	isomessage.Field(102, "970630181070041")
	isomessage.Field(120, "BRN015301213230443463")

	// generate binary representation of the message into rawMessage
	rawMessage, err := isomessage.Pack()

	return rawMessage, err
}
