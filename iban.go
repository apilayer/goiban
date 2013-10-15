package goiban

import (
	"strings"
	//"fmt"
	"math/big"
)

// describes the structure of an IBAN
type Iban struct {
	countryCode string
	checkDigit string
	bban string
	original string
}

var (
	successIndicator = big.NewInt(1)
	ibanMod = big.NewInt(97)
)
/*
	Returns a pointer to an Iban instance or nil on structural errors

	Can return NIL on structural error
*/
func ParseToIban(val string) *Iban {
	// Init empty Iban object
	cc := extractCountryCode(val)
	checkDigit := extractCheckDigit(val)
	bban := extractBBAN(val)

	if(len(cc)==0 || len(checkDigit) == 0 || len(bban) == 0) {
		return nil;
	}

	iban := &Iban{cc,checkDigit,bban,val}
	return iban
}

/*
	Returns true for a valid IBAN
*/
func (iban *Iban) Validate() *ValidationResult {
	var ok bool
	validateableString:= toNumericString(iban.bban) + countryCodeToNumericString(iban.countryCode) + iban.checkDigit

	intBuf:=big.NewInt(0)
	
	intBuf, ok = intBuf.SetString(validateableString,10, )
	if !ok {
		return NewValidationResult(false,"Could not parse IBAN number.",iban.original);
	}
	result := intBuf.Mod(intBuf,ibanMod) 
	
	if result.Cmp(successIndicator) == 0 {
		return NewValidationResult(true,"",iban.original);
	}

	return NewValidationResult(false,"Validation failed.",iban.original);
}


/*
	Returns true if the string val can be parsed to an Iban Struct
*/
func IsParseable(val string) bool {
	// Init empty Iban object
	cc := extractCountryCode(val)
	checkDigit := extractCheckDigit(val)
	bban := extractBBAN(val)

	if(len(cc)==0 || len(checkDigit) == 0 || len(bban) == 0) {
		return false;
	} 
	return true
}

func extractCountryCode(val string) string {
	// has to be at least two digits long
	if(len(val) < 2) {
		return "";
	}

	possibleCountryCode := strings.ToUpper(val[0:2])

	if(!isValidChar(possibleCountryCode[0]) || 
	   !isValidChar(possibleCountryCode[1])) {
		return "";
	}

	return possibleCountryCode
}

func extractCheckDigit(val string) string {
	// starts at position 2 and is 2 digits long
	if(len(val) < 4) {
		return "";
	}

	possibleCheckDigit := strings.ToUpper(val[2:4])

	if(!isValidNum(possibleCheckDigit[0]) || 
	   !isValidNum(possibleCheckDigit[1])) {
		return "";
	}

	return possibleCheckDigit
}

func extractBBAN(val string) string {
	// replace all spaces
	val = strings.Replace(val," ","",-1);
	// starts at position 4 in the string	
	if(len(val) < 4 || len(val) > 34) {	
		return "";
	}

	// we can do a more accurate check for some countries
	// see static_data.go
	allowedLength := getAllowedLength(extractCountryCode(val))
	if(allowedLength > 0 && (len(val) > allowedLength)) {
		return "";
	}

	bban := strings.ToUpper(val[4:len(val)])

	// only alphanumeric chars may be used
	for _,ch := range bban {
		if(!isValidNum(uint8(ch)) && !isValidChar(uint8(ch))) {
			return ""
		}
	}

	return bban
}

/*
	Returns a numeric representation of a two-character country code

	e.g. DE -> 1314

	The char value is by 55 and it's integer representation is concatenated to a string.
*/
func countryCodeToNumericString(countryCode string) string {
	if(len(countryCode) > 2 || len(countryCode) < 2) {
		return ""		
	}	
	return toNumericString(countryCode)
}
