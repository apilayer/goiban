package goiban

import (
	"testing"
	//"fmt"
)


func TestCanInstantiateIbanFromValidString(t *testing.T) {
	input:="GB29 NWBK 6016 1331 9268 19"
	result:=ParseToIban(input)
	if(result == nil) {
		t.Errorf("Failed to extract BBAN.")
	}
}

func TestCanCheckIbanValidity(t *testing.T) {
	input:="GB29 NWBK 6016 1331 9268 19"
	parsedIban := ParseToIban(input)
	result:=parsedIban.Validate()
	if(result.Valid != true) {
		t.Errorf("Failed to validate BBAN.")
	}
}

func TestCanCheckIbanParseable(t *testing.T) {
	input:="GB29 NWBK 6016 1331 9268 19"
	result:=IsParseable(input)	
	if(result != true) {
		t.Errorf("Failed to validate BBAN.")
	}

	input="GB29"
	result=IsParseable(input)	
	if(result == true) {
		t.Errorf("Failed to validate BBAN.")
	}

	input=""
	result=IsParseable(input)	
	if(result == true) {
		t.Errorf("Failed to validate BBAN.")
	}
}

func TestShouldNotInstantiateIbanFromInvalidString(t *testing.T) {
	input:="GB NWBK 6016 1331 9268 19"
	result:=ParseToIban(input)
	if(result != nil) {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input="1112 NWBK 6016 1331 9268 19"
	result=ParseToIban(input)
	if(result != nil) {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input="DE12"
	result=ParseToIban(input)
	if(result != nil) {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input="DE"
	result=ParseToIban(input)
	if(result != nil) {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input=""
	result=ParseToIban(input)
	if(result != nil) {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
}

func TestExtractValidBBAN(t *testing.T) {
	input:="GB29 NWBK 6016 1331 9268 19"
	result:=extractBBAN(input)
	if(result != "NWBK60161331926819") {
		t.Errorf("Failed to extract BBAN.")
	}
}

func TestExtractInvalidBBAN(t *testing.T) {
	input:="GB29"
	if(extractBBAN(input) != "") {
		t.Errorf("Extracted invalid BBAN.")
	}

	input="GB29 NWBK 6016 1331 9268 19GB29 NWBK 6016 1331 9268 19GB29 NWBK 6016 1331 9268 19"
	if(extractBBAN(input) != "") {
		t.Errorf("Extracted invalid BBAN.")
	}

	input="GB29 NWBK 6016 1331 9268 19 DURR"
	result:=extractBBAN(input)
	if(result != "") {
		t.Errorf("Extracted invalid BBAN. (length restriction by map)")
	}
}

func TestExtractValidCountryCode(t *testing.T) {
	input:="DE"
	if(extractCountryCode(input) != input) {
		t.Errorf("Failed to extract country code.")
	}
}

func TestExtractInvalidCountryCode(t *testing.T) {
	input:="D1"
	if(extractCountryCode(input) == input) {
		t.Errorf("Extracted invalid country code.")
	}

	input = ""
	if(extractCountryCode(input) != "") {
		t.Errorf("Extracted invalid country code.")
	}
}

func TestExtractCheckDigit(t *testing.T) {
	input:="DE12"
	if(extractCheckDigit(input) != "12") {
		t.Errorf("Failed to extract check digit.")
	}
}

func TestExtractInvalidCheckDigit(t *testing.T) {
	input:="D1"
	if(extractCheckDigit(input) == input) {
		t.Errorf("Extracted invalid check digit.")
	}

	input = ""
	if(extractCheckDigit(input) != "") {
		t.Errorf("Extracted invalid check digit.")
	}
}

func TestCountryCodeToNumericString(t *testing.T) {
	input:="DE"
	result:=countryCodeToNumericString(input)
	if(result != "1314") {
		t.Errorf("Can't convert country code to numeric string: " + result)
	}
}

func TestInvalidCountryCodeToNumericString(t *testing.T) {
	input:=""
	result:=countryCodeToNumericString(input)
	if(result != "") {
		t.Errorf("Converted invalid code!")
	}

	input="DEE"
	result=countryCodeToNumericString(input)
	if(result != "") {
		t.Errorf("Converted invalid code!")
	}
}