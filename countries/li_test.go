package countries

import "testing"

func TestCanConvertSliceToLiechtensteinBankEntry(t *testing.T) {
	data := []string{"Bank Alpinum AG", "BALPLI22", "8801"}
	entry := LiechtensteinRowToEntry(data)

	if entry.Bankcode != "8801" {
		t.Errorf("expected 8801 as bankcode, got %v", entry.Bankcode)
	}

	if entry.Bic != "BALPLI22" {
		t.Errorf("expected BALPLI22 as bic, got %v", entry.Bic)
	}

	if entry.Name != "Bank Alpinum AG" {
		t.Errorf("expected Bank Alpinum AG as name, got %v", entry.Name)
	}
}
