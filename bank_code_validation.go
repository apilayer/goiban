package goiban

import (
	"database/sql"
)

var (
	SELECT_BY_BANK_CODE = "SELECT 1 FROM BANK_DATA WHERE bankcode = ? and country = ?;"
	SELECT_BY_BANK_CODE_STMT *sql.Stmt
	

)

func ValidateBankCode(iban *Iban, intermediateResult *ValidationResult, db *sql.DB) *ValidationResult {
	if SELECT_BY_BANK_CODE_STMT == nil {
		prepareSelectByBankCodeStatement(db)
	}

	length, ok := COUNTRY_CODE_TO_BANK_CODE_LENGTH[(iban.countryCode)]

	if !ok {
		intermediateResult.CheckResults["bankCode"] = false
		intermediateResult.Messages = append(intermediateResult.Messages, "Cannot validate bank code. No information available.")
		return intermediateResult
	}

	bankCode := iban.bban[0:length]

	var res int
	err := SELECT_BY_BANK_CODE_STMT.QueryRow(bankCode, iban.countryCode).Scan(&res)
	
	switch {
		case err == sql.ErrNoRows:
			intermediateResult.Valid = false
			intermediateResult.Messages = append(intermediateResult.Messages, "Invalid bank code: " + bankCode)
			intermediateResult.CheckResults["bankCode"] = false
			return intermediateResult
	}	
	intermediateResult.Messages = append(intermediateResult.Messages, "Bank code valid: " + bankCode)
	intermediateResult.CheckResults["bankCode"] = true


	return intermediateResult
}

func prepareSelectByBankCodeStatement(db *sql.DB) {
	var err error
	SELECT_BY_BANK_CODE_STMT, err = db.Prepare(SELECT_BY_BANK_CODE)
	if err != nil {
		panic("Couldn't prepare statement: " +  SELECT_BY_BANK_CODE)
	}
}