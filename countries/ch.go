/*
The MIT License (MIT)


Copyright (c) 2016 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package countries

import (
	"fmt"
)


type SwitzerlandBankFileEntry struct {
	BankCode    string //2-6
	NewBankCode string //11-15
	ShortName   string //39-54
	BankName    string //54-114
	Address     string //116-151
	Zip         string //178-188
	Place       string //194-229
	Bic         string //290-304
}

func SwitzerlandBankStringToEntry(val string) *SwitzerlandBankFileEntry {
	runeVal := []rune(val)

	return &SwitzerlandBankFileEntry{
		fmt.Sprintf("%05s", toTrimmedString(runeVal[2:6])),
		fmt.Sprintf("%05s", toTrimmedString(runeVal[11:15])),
		toTrimmedString(runeVal[39:54]),
		toTrimmedString(runeVal[54:114]),
		toTrimmedString(runeVal[114:149]),
		toTrimmedString(runeVal[178:188]),
		toTrimmedString(runeVal[194:229]),
		toTrimmedString(runeVal[284:298]),
	}
}
