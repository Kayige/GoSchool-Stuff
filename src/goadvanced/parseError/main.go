package main

import "fmt"

// ParseError Struct
type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on ln Line %d, Char %d"

	return fmt.Sprintf(format, p.Message, p.Line, p.Char)

}

// ParseLine function
func parseLine(s1 string) (string, error) {
	errorMsg := "Parse error occured"
	invalidLine := 5
	invalidChar := 4
	hasError := true
	if hasError {
		return "", &ParseError{errorMsg, invalidLine, invalidChar}
	}
	return "Parse Success", nil

}

func main() {
	status, err := parseLine("dummyString")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status)
	}
}
