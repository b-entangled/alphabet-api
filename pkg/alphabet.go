package alphabet

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"
)

// filter and add alphabets to map
func mapLetters(s string) map[string]bool {
	sLower := strings.ToLower(s)
	sMap := make(map[string]bool)
	for _, v := range sLower {
		if unicode.IsLetter(v) {
			sMap[string(v)] = true
		}
	}
	return sMap
}

// generate letters from a-z
func generateAlphabets() []string {
	var letters []string
	// Generate ASCII values from 97 to 122 i.e. a to z
	for i := 97; i < 123; i++ {
		letters = append(letters, string(i))
	}
	return letters
}

// Validate - validates the string
func Validate(s string) bool {
	inputMap := mapLetters(s)
	for _, letter := range generateAlphabets() {
		if ok := inputMap[letter]; !ok {
			return false
		}
	}
	return true
}

// Alphabet - alphabet service
func Alphabet(res http.ResponseWriter, req *http.Request) {
	inputString := req.URL.Query().Get("inputstring")
	if inputString == "" {
		res.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(res).Encode(map[string]string{"Error:": "parameter 'inputstring' is required"})
		return
	}
	// check the input string
	if Validate(inputString) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"inputString": inputString, "valid": "true"})
	} else {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"inputString": inputString, "valid": "false"})
	}
}
