package countries

import (
	"testing"
)

func TestCanConvertStringToSwitzerlandBankEntry(t *testing.T) {
	data := "088792 0000     0879298792 120091124301BNP PAR SEC SERBNP PARIBAS SECURITIES SERVICES                             Selnaustrasse 16                   Postfach 2119                      8022      Zürich                             058 212 63 00     058 212 63 60            87-230511-3 PARBCHZZXXX   "
	result := SwitzerlandBankStringToEntry(data)

	if result.BankCode != "8792" {
		t.Errorf("Couldn't parse bank code: ", result.BankCode)
	}
	if result.NewBankCode != "" {
		t.Errorf("Couldn't parse new bank code: ", result.NewBankCode)
	}
	if result.Address != "Selnaustrasse 16" {
		t.Errorf("Couldn't parse address: ", result.Address)
	}
	if result.Zip != "8022" {
		t.Errorf("Couldn't parse zip: ", result.Zip)
	}
	if result.Place != "Zürich" {
		t.Errorf("Couldn't parse place: ", result.Place)
	}
	if result.ShortName != "BNP PAR SEC SER" {
		t.Errorf("Couldn't parse short name: ", result.ShortName)
	}
	if result.BankName != "BNP PARIBAS SECURITIES SERVICES" {
		t.Errorf("Couldn't parse bank name: ", result.BankName)
	}
	if result.Bic != "PARBCHZZXXX" {
		t.Errorf("Couldn't parse bic.", result.Bic)
	}
}
