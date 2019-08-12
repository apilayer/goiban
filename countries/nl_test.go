package countries

import "testing"

func TestCanConvertSliceToNetherlandsBankEntry(t *testing.T) {
	data := []string{"FVLBNL22", "FVLB", "F.VAN LANSCHOT BANKIERS N.V.", ""}
	entry := NetherlandsRowToEntry(data)

	if entry.Bankcode != "FVLB" {
		t.Errorf("expected FVLB as bankcode, got %v", entry.Bankcode)
	}

	if entry.Bic != "FVLBNL22" {
		t.Errorf("expected FVLBNL22 as bic, got %v", entry.Bic)
	}

	if entry.Name != "F.VAN LANSCHOT BANKIERS N.V." {
		t.Errorf("expected F.VAN LANSCHOT BANKIERS N.V. as name, got %v", entry.Name)
	}
}
