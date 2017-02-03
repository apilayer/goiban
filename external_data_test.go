/*
The MIT License (MIT)

Copyright (c) 2014 Chris Grieger

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

package goiban

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	co "github.com/stefan93/goiban/countries"
)

var (
	db, err = sql.Open("mysql", "root:root@/goiban?charset=utf8")
)

func TestCanReadFromAustriaFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/austria.csv", &co.AustriaBankFileEntry{}, ch)

	peek := (<-ch).(*co.AustriaBankFileEntry)
	if peek.Name == "" {
		t.Errorf("Failed to read file.")
	}
}

func TestCannotReadFromNonExistingAustriaFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/austria_blablablabla.csv", &co.AustriaBankFileEntry{}, ch)
	result := <-ch
	if result != nil {
		t.Errorf("Failed to read file.")
	}
}
func TestCanReadFromBundesbankFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/bundesbank.txt", &co.BundesbankFileEntry{}, ch)

	peek := (<-ch).(*co.BundesbankFileEntry)
	if peek.Name == "" {
		t.Errorf("Failed to read file.")
	}
}

func TestCannotReadFromNonExistingBundesbankFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/bundesbank_doesntexist.txt", &co.BundesbankFileEntry{}, ch)
	result := <-ch
	if result != nil {
		t.Errorf("Failed to read file.")
	}
}

/*func TestCanLoadBankInfoFromDatabase(t *testing.T) {
	bankInfo := getBankInformationByCountryAndBankCodeFromDb("DE", "84050000", db)
	fmt.Println(bankInfo)
	if bankInfo == nil {
		t.Errorf("Cannot load data from db. Is it empty?")
	}
}*/

func TestCanReadFromBelgiumXLSX(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/belgium.xlsx", &co.BelgiumFileEntry{}, ch)

	peek := (<-ch).([]co.BelgiumFileEntry)
	if peek[0].Name != "bpost bank" {
		t.Errorf("Failed to read file.")
	}
}

func TestCanReadFromNetherlandsXLSX(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/netherlands.xlsx", &co.NetherlandsFileEntry{}, ch)

	peek := (<-ch).(co.NetherlandsFileEntry)
	if peek.Name != "ABN AMRO BANK N.V" {
		t.Errorf("Failed to read file.")
	}
}
func TestCanReadFromSwitzerlandFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/switzerland.txt", &co.SwitzerlandBankFileEntry{}, ch)

	peek := (<-ch).(*co.SwitzerlandBankFileEntry)
	if peek.Bic != "SNBZCHZZXXX" {
		t.Errorf("Failed to read file.")
	}
}

func TestCannotReadFromNonExistingSwitzerlandFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/switzerland_blablablabla.txt", &co.SwitzerlandBankFileEntry{}, ch)
	result := <-ch
	if result != nil {
		t.Errorf("Failed to read file.")
	}
}
func TestCanReadFromLiechtensteinXLSX(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/liechtenstein.xlsx", &co.LiechtensteinFileEntry{}, ch)
	peek := (<-ch).(co.LiechtensteinFileEntry)
	if peek.Bic != "BALPLI22" {
		t.Errorf("Failed to read file." + peek.Bic)
	}
}
func TestCannotReadFromNonExistingLiechtensteinFile(t *testing.T) {
	ch := make(chan interface{})
	go ReadFileToEntries("test/lliechtenstein_blablablabla.xlsx", &co.LiechtensteinFileEntry{}, ch)
	result := <-ch
	if result != nil {
		t.Errorf("Failed to read file.")
	}
}
