package countries

import (
	"strconv"
)

/* 
	Describes the field width of the fixed width text file
	which can be found at
	http://www.bundesbank.de/Redaktion/DE/Standardartikel/Kerngeschaeftsfelder/Unbarer_Zahlungsverkehr/bankleitzahlen_download.html?searchArchive=0&submit=Suchen&searchIssued=0&templateQueryString=Bankleitzahlen
	
	Example:
	bankcodemname######                                                zip##city                               short#####                 pan##bic#########z#id###cdnext####
	100000001Bundesbank                                                10591Berlin                             BBk Berlin                 20100MARKDEF110009011380U000000000
*/

type BundesbankFileEntry struct{
	Bankcode string		// 8
	M int				// 1
	Name string			// 58
	Zip string			// 5
	City string			// 35
	ShortName string	// 27
	Pan int 			// 5
	Bic string			// 12
	CheckAlgo string	// 2 	enumerates some checksum algorithms
						// 		described in http://www.bundesbank.de/Redaktion/DE/Downloads/Kerngeschaeftsfelder/Unbarer_Zahlungsverkehr/pruefzifferberechnungsmethoden.pdf?__blob=publicationFile
	Id int 				// 5	internal id
	Change string		// 1
	ToBeDeleted int 	// 1
	NewBankCode	string	// 8							
}

func BundesbankStringToEntry(val string) *BundesbankFileEntry {
	runeVal := []rune(val)
	m,_:=strconv.Atoi(string(runeVal[8:9]))
	pan,_:=strconv.Atoi(string(runeVal[134:139]))
	id,_:=strconv.Atoi(string(runeVal[153:158]))
	toBeDeleted,_:=strconv.Atoi(string(runeVal[159:160]))

	return &BundesbankFileEntry{
		string(runeVal[0:8]),
		m,
		string(runeVal[9:67]),
		string(runeVal[67:72]),
		string(runeVal[72:107]),
		string(runeVal[107:134]),
		pan,
		string(runeVal[139:151]),
		string(runeVal[151:153]),
		id,
		string(runeVal[158:159]),
		toBeDeleted,
		string(runeVal[160:168]),
	}
}