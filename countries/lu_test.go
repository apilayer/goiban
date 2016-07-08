package countries

import "testing"

func TestCanConvertSliceToLuxembourgBankEntry(t *testing.T) {
	data := []string{"Banque et Caisse d'Epargne de l'Etat, Luxembourg", "001", "BCEE LU LL", ""}
	entry := LuxembourgRowToEntry(data)

	if entry.Bankcode != "001" {
		t.Errorf("expected 001 as bankcode, got %v", entry.Bankcode)
	}

	if entry.Bic != "BCEE LU LL" {
		t.Errorf("expected BCEE LU LL as bic, got %v", entry.Bic)
	}

	if entry.Name != "Banque et Caisse d'Epargne de l'Etat, Luxembourg" {
		t.Errorf("expected Banque et Caisse d'Epargne de l'Etat, Luxembourg as name, got %v", entry.Name)
	}
}
