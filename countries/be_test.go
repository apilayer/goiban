package countries

import "testing"

func TestCanConvertSliceToBelgiuqueNationalBankEntry(t *testing.T) {
	data := []string{"000", "000", "BPOT BE B1", "bpost bank", "", "", ""}
	res := BelgiumRowToEntry(data)

	if len(res) != 1 {
		t.Errorf("expected result with one entry")
	}

	entry := res[0]
	if entry.Bankcode != "000" {
		t.Errorf("expected 000 as bankcode, got %v", entry.Bankcode)
	}

	if entry.Bic != "BPOT BE B1" {
		t.Errorf("expected BPOT BE B1 as bic, got %v", entry.Bic)
	}

	if entry.Name != "bpost bank" {
		t.Errorf("expected bpost bank as name, got %v", entry.Name)
	}
}
