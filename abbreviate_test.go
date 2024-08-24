package abbreviate

import (
	"testing"
)

func ExpectString(t *testing.T, returnedString string, expectedString string) bool {
	if returnedString != expectedString {
		t.Fatalf("Expected '%s', got '%s'", expectedString, returnedString)
		return false
	}
	return true
}
func CheckString(t *testing.T, inputString string, abbreviateLength int, expectedString string) {
	returnedString := Abbreviate(inputString, abbreviateLength)
	if ExpectString(t, returnedString, expectedString) {
		t.Logf("'%s' got abbreviated to '%s' with length %d\n", inputString, returnedString, abbreviateLength)
	}
}

func TestNearLength(t *testing.T) {
	CheckString( t, "Hello", 5, "Hello" )
}

func TestLowercaseFirst(t *testing.T) {
	inputString := "TesT" 
	CheckString( t, inputString, len(inputString)-2, "TT" )
}

func TestString1(t *testing.T) {
	inputString := "test" 
	CheckString( t, inputString, len(inputString)-1, "tst" )
	CheckString( t, inputString, len(inputString)-2, "ts" )
}

func TestString2(t *testing.T) {
	inputString := "Word1 2Drow" 
	CheckString( t, inputString, len(inputString)-1, "Word12Drow" )
	CheckString( t, inputString, len(inputString)-2, "Word12Drw" )
	CheckString( t, inputString, len(inputString)-3, "Wrd12Drw" )
	CheckString( t, inputString, len(inputString)-4, "Wrd12Dr" )
	CheckString( t, inputString, len(inputString)-5, "Wrd12D" )
	CheckString( t, inputString, len(inputString)-6, "Wr12D" )
	CheckString( t, inputString, len(inputString)-7, "W12D" )
	CheckString( t, inputString, len(inputString)-8, "W12" )
	CheckString( t, inputString, len(inputString)-9, "W1" )
	CheckString( t, inputString, len(inputString)-10, "W" )
	CheckString( t, inputString, len(inputString)-11, "W12" )
}

func TestSeparatorRemoval(t *testing.T) {
	inputString := "Word1-2Drow_Word3" 
	CheckString( t, inputString, len(inputString)-1, "Word1-2DrowWord3" )
	CheckString( t, inputString, len(inputString)-2, "Word12DrowWord3" )
}

func TestDigraphsDiblends(t *testing.T) {
	inputString := "so write score" 
	CheckString( t, inputString, 5, "swrsc" )
}

func TestTrigraphsTriblends(t *testing.T) {
	inputString := "total splash" 
	CheckString( t, inputString, 4, "tspl" )
}

func TestTrigraphsTriblends2(t *testing.T) {
	inputString := "Some Important String" 
	CheckString( t, inputString, 8, "SmImpStr" )
}
