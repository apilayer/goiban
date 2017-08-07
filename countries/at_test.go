package countries

import (
	"testing"
)

func TestCanConvertStringToAustriaBankEntry(t *testing.T) {
	data := "Hauptanstalt;\"10050973\";\"52300\";\"KI\";\"Joint stock banks and private banks\";\"350921k\";\"Addiko Bank AG\";\"Wipplingerstraße 34/4\";\"1010\";\"Wien\";\"\";\"\";\"\";\"\";\"\";\"Wien\";\"050232\";\"050232/3000\";\"holding@addiko.com\";\"HSEEAT2KXXX\";\"www.addiko.com\";\"20130621\";;"
	result := AustriaBankStringToEntry(data)

	if result.Bankcode != "52300" {
		t.Errorf("Couldn't parse bank code.")
	}
	if result.Name != "Addiko Bank AG" {
		t.Errorf("Couldn't parse name.")
	}
	if result.Bic != "HSEEAT2KXXX" {
		t.Errorf("Couldn't parse bic.", result.Bic)
	}
}
