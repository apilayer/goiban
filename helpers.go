package goiban

import "strconv"

func isValidChar(ch uint8) bool {
	return (ch > 64 && ch < 91)
}

func isValidNum(ch uint8) bool {
	return (ch > 47 && ch < 58)
}


func toNumericString(val string) string {
	numericVal := ""
	for _,ch := range val {		
		// if it's neither a number nor a char
		// fail
		intCh := uint8(ch)
		if(!isValidNum(intCh) &&
		   !isValidChar(intCh)) {
			return "";
		}
		if(isValidChar(intCh)) {
			numericVal += strconv.Itoa(int(ch)-55)
		} else {
			numericVal += string(ch)
		}
	}	

	return numericVal
}
