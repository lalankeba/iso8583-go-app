package iso8583

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/specs"
)

var ExtendedSpec *iso8583.MessageSpec = func() *iso8583.MessageSpec {
	// Copy the original Spec87ASCII
	extendedSpec := *specs.Spec87ASCII

	// Add or modify fields
	extendedSpec.Fields[3] = field.NewString(&field.Spec{
		Length:      6,
		Description: "Processing Code (Custom as String)",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
	})

	extendedSpec.Fields[70] = field.NewString(&field.Spec{
		Length:      3,
		Description: "Network Management Information Code",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
	})

	extendedSpec.Fields[102] = field.NewString(&field.Spec{
		Length:      16,
		Description: "Account Identification 1",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.LL,
	})

	extendedSpec.Fields[120] = field.NewString(&field.Spec{
		Length:      30,
		Description: "Reserved for Private Use",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.LLL,
	})

	return &extendedSpec
}()
