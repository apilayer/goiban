package countries

import (
	"testing"
)

func TestCanConvertStringToAustriaBankEntry(t *testing.T) {
	countryCodeToBankCodeMap := map[string]int{
		"AT": 5,
	}
	
	data := "Hauptanstalt;263575;34322;KI;Raiffeisen;93513w;Raiffeisenbank Mondseeland eGen;Rainerstr. 11;5310;Mondsee;;5310;Mondsee;29;Oberösterreich;06232/3151;06232/315138017;rb@mondseeland.com;RZOOAT2L322;www.mondseeland.com;20.11.1949"

	result := AustriaBankStringToEntry(data, countryCodeToBankCodeMap)

	if result.Bankcode != "34322" {
		t.Errorf("Couldn't parse bank code. %v", result.Bankcode)
	}
	if result.Name != "Raiffeisenbank Mondseeland eGen" {
		t.Errorf("Couldn't parse name. %v", result.Name)
	}
	if result.Bic != "RZOOAT2L322" {
		t.Errorf("Couldn't parse bic. %v", result.Bic)
	}
}
