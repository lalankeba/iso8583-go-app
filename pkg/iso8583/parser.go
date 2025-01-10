package iso8583

import (
	"github.com/moov-io/iso8583"
)

func ParseMessage(rawIso []byte) (*iso8583.Message, error) {
	isomessage := iso8583.NewMessage(ExtendedSpec)
	err := isomessage.Unpack(rawIso)

	if err != nil {
		return nil, err
	}

	return isomessage, nil
}
