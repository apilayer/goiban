package countries

import (
	"testing"
)

func TestCanConvertStringToBundesbankEntry(t *testing.T) {
	data := "100000001Bundesbank                                                10591Berlin                             BBk Berlin                 20100MARKDEF110009011380U000000000"
	result := BundesbankStringToEntry(data)

	if result.Bankcode != "10000000" {
		t.Errorf("Couldn't parse bank code.")
	}
	if result.M != 1 {
		t.Errorf("Couldn't parse M.")
	}
	if result.Name != "Bundesbank" {
		t.Errorf("Couldn't parse name.")
	}
	if result.Zip != "10591" {
		t.Errorf("Couldn't parse zip.")
	}
	if result.City != "Berlin" {
		t.Errorf("Couldn't parse city.")
	}
	if result.ShortName != "BBk Berlin" {
		t.Errorf("Couldn't parse short name.")
	}
	if result.Pan != 20100 {
		t.Errorf("Couldn't parse pan: ", result.Pan)
	}
	if result.Bic != "MARKDEF1100" {
		t.Errorf("Couldn't parse bic.", result.Bic)
	}
	if result.CheckAlgo != "09" {
		t.Errorf("Couldn't parse check algo.")
	}
	if result.Id != "01138" {
		t.Errorf("Couldn't parse internal id.")
	}
	if result.Change != "U" {
		t.Errorf("Couldn't parse change.", result.Change)
	}
	if result.ToBeDeleted != 0 {
		t.Errorf("Couldn't parse to be deleted.")
	}
	if result.NewBankCode != "00000000" {
		t.Errorf("Couldn't parse new bank code.")
	}
}
