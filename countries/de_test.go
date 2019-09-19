package countries

import (
	"testing"
)

func TestCanConvertStringToBundesbankEntry(t *testing.T) {
	data := "100000001Bundesbank                                                10591Berlin                             BBk Berlin                 20100MARKDEF110009011380U000000000"
	result := BundesbankStringToEntry(data)

	if result.Bankcode != "10000000" {
		t.Error("Couldn't parse bank code.")
	}
	if result.M != 1 {
		t.Error("Couldn't parse M.")
	}
	if result.Name != "Bundesbank" {
		t.Error("Couldn't parse name.")
	}
	if result.Zip != "10591" {
		t.Error("Couldn't parse zip.")
	}
	if result.City != "Berlin" {
		t.Error("Couldn't parse city.")
	}
	if result.ShortName != "BBk Berlin" {
		t.Error("Couldn't parse short name.")
	}
	if result.Pan != 20100 {
		t.Errorf("Couldn't parse pan: %v", result.Pan)
	}
	if result.Bic != "MARKDEF1100" {
		t.Errorf("Couldn't parse bic. %v", result.Bic)
	}
	if result.CheckAlgo != "09" {
		t.Errorf("Couldn't parse check algo.")
	}
	if result.Id != "01138" {
		t.Errorf("Couldn't parse internal id.")
	}
	if result.Change != "U" {
		t.Errorf("Couldn't parse change. %v", result.Change)
	}
	if result.ToBeDeleted != 0 {
		t.Errorf("Couldn't parse to be deleted.")
	}
	if result.NewBankCode != "00000000" {
		t.Errorf("Couldn't parse new bank code.")
	}
}
