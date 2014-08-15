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
	"fmt"
	co "github.com/fourcube/goiban/countries"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var (
	db, err = sql.Open("mysql", "root:root@/goiban?charset=utf8")
)

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

func TestCanLoadBankInfoFromDatabase(t *testing.T) {
	bankInfo := getBankInformationByCountryAndBankCodeFromDb("DE", "84050000", db)
	fmt.Println(bankInfo)
	if bankInfo == nil {
		t.Errorf("Cannot load data from db. Is it empty?")
	}
}
