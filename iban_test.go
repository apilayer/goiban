/*
The MIT License (MIT)

Copyright (c) 2014 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package goiban

import (
	"testing"
	//"fmt"
)

func TestCanInstantiateIbanFromValidString(t *testing.T) {
	input := "GB29 NWBK 6016 1331 9268 19"
	result := ParseToIban(input)
	if result == nil {
		t.Errorf("Failed to extract BBAN.")
	}
}

func TestCanCheckIbanValidity(t *testing.T) {
	input := "GB29 NWBK 6016 1331 9268 19"
	parsedIban := ParseToIban(input)
	result := parsedIban.Validate()
	if result.Valid != true {
		t.Errorf("Failed to validate BBAN.")
	}
}

func TestCanCheckIbanParseable(t *testing.T) {
	input := "GB29 NWBK 6016 1331 9268 19"
	result := IsParseable(input)
	if result.Valid != true {
		t.Errorf("Failed to validate BBAN.")
	}

	input = "GB29"
	result = IsParseable(input)
	if result.Valid {
		t.Errorf("Failed to validate BBAN.")
	}

	input = ""
	result = IsParseable(input)
	if result.Valid {
		t.Errorf("Failed to validate BBAN.")
	}
}

func TestShouldNotInstantiateIbanFromInvalidString(t *testing.T) {
	input := "GB NWBK 6016 1331 9268 19"
	result := ParseToIban(input)
	if result != nil {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input = "1112 NWBK 6016 1331 9268 19"
	result = ParseToIban(input)
	if result != nil {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input = "DE12"
	result = ParseToIban(input)
	if result != nil {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input = "DE"
	result = ParseToIban(input)
	if result != nil {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
	input = ""
	result = ParseToIban(input)
	if result != nil {
		t.Errorf("Instantiated IBAN from invalid string.")
	}
}

func TestExtractValidBBAN(t *testing.T) {
	input := "GB29 NWBK 6016 1331 9268 19"
	result, _ := extractBBAN(input)
	if result.Data != "NWBK60161331926819" {
		t.Errorf("Failed to extract BBAN.")
	}
}

func TestExtractInvalidBBAN(t *testing.T) {
	input := "GB29"
	var ok bool
	_, ok = extractBBAN(input)

	if ok {
		t.Errorf("Extracted invalid BBAN.")
	}

	input = "GB29 NWBK 6016 1331 9268 19GB29 NWBK 6016 1331 9268 19GB29 NWBK 6016 1331 9268 19"
	_, ok = extractBBAN(input)
	if ok {
		t.Errorf("Extracted invalid BBAN.")
	}

	input = "GB29 NWBK 6016 1331 9268 19 DURR"
	_, ok = extractBBAN(input)
	if ok {
		t.Errorf("Extracted invalid BBAN. (length restriction by map)")
	}
}

func TestExtractValidCountryCode(t *testing.T) {
	input := "DE"
	if ExtractCountryCode(input) != input {
		t.Errorf("Failed to extract country code.")
	}
}

func TestExtractInvalidCountryCode(t *testing.T) {
	input := "D1"
	if ExtractCountryCode(input) == input {
		t.Errorf("Extracted invalid country code.")
	}

	input = ""
	if ExtractCountryCode(input) != "" {
		t.Errorf("Extracted invalid country code.")
	}
}

func TestExtractCheckDigit(t *testing.T) {
	input := "DE12"
	if extractCheckDigit(input) != "12" {
		t.Errorf("Failed to extract check digit.")
	}
}

func TestExtractInvalidCheckDigit(t *testing.T) {
	input := "D1"
	if extractCheckDigit(input) == input {
		t.Errorf("Extracted invalid check digit.")
	}

	input = ""
	if extractCheckDigit(input) != "" {
		t.Errorf("Extracted invalid check digit.")
	}
}

func TestCountryCodeToNumericString(t *testing.T) {
	input := "DE"
	result := countryCodeToNumericString(input)
	if result != "1314" {
		t.Errorf("Can't convert country code to numeric string: " + result)
	}
}

func TestInvalidCountryCodeToNumericString(t *testing.T) {
	input := ""
	result := countryCodeToNumericString(input)
	if result != "" {
		t.Errorf("Converted invalid code!")
	}

	input = "DEE"
	result = countryCodeToNumericString(input)
	if result != "" {
		t.Errorf("Converted invalid code!")
	}
}

func TestCanGenerateIBANFromValidData(t *testing.T) {
	result := CalculateIBAN("BE", "539", "007547034")
	if !result.Valid {
		t.Errorf("expected result to be valid")
	}

	if result.Data != "BE68539007547034" {
		t.Errorf("expected result iban to match %v != %v", result.Data, "BE68539007547034")
	}
}

func TestCanGenerateIBANFromValidDataEnsureCheckdigitsLeadingZeroPresent(t *testing.T) {
	result := CalculateIBAN("DE", "1", "9")
	if !result.Valid {
		t.Errorf("expected result to be valid")
	}

	if result.Data != "DE0819" {
		t.Errorf("expected result iban to match %v != %v", result.Data, "DE0819")
	}
}
