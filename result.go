package goiban

type ValidationResult struct {
	Valid bool `json:"valid"`
	Message string `json:"message"`
	Iban string `json:"iban"`

}

func NewValidationResult(valid bool, message string, iban string) *ValidationResult {
	return &ValidationResult{valid,message, iban}
}