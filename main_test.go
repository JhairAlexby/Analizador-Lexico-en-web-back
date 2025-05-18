package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"juan 890 juan", []string{"juan", "890", "juan"}},
		{"9890LUpchiapas", []string{"9890LUpchiapas"}},
		{"identificador 123 otro_identificador", []string{"identificador", "123", "otro_identificador"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := tokenize(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Para '%s', se esperaba %v pero se obtuvo %v", test.input, test.expected, result)
		}
	}
}

func TestClassifyToken(t *testing.T) {
	tests := []struct {
		input    string
		expected Token
	}{
		{"juan", Token{Value: "juan", Type: "identificador"}},
		{"890", Token{Value: "890", Type: "numero"}},
		{"9890LUpchiapas", Token{Value: "9890LUpchiapas", Type: "error"}},
		{"Upchiapas", Token{Value: "Upchiapas", Type: "identificador"}},
	}

	for _, test := range tests {
		result := classifyToken(test.input)
		if result.Value != test.expected.Value || result.Type != test.expected.Type {
			t.Errorf("Para '%s', se esperaba %v pero se obtuvo %v",
				test.input, test.expected, result)
		}
	}
}

func TestAnalyzeText(t *testing.T) {
	tests := []struct {
		input                 string
		expectedTokenCount    int
		expectedIdentifiers   int
		expectedUniqueIdent   int
		expectedNumbers       int
		expectedUniqueNumbers int
		expectedErrors        int
	}{
		{
			"juan 890 juan",
			3, // total tokens
			2, // identificadores
			1, // identificadores únicos
			1, // números
			1, // números únicos
			0, // errores
		},
		{
			"juan 890 9890LUpchiapas",
			3, // total tokens
			1, // identificadores
			1, // identificadores únicos
			1, // números
			1, // números únicos
			1, // errores
		},
	}

	for _, test := range tests {
		result := analyzeText(test.input)

		if len(result.Tokens) != test.expectedTokenCount {
			t.Errorf("Para '%s', se esperaban %d tokens pero se obtuvieron %d",
				test.input, test.expectedTokenCount, len(result.Tokens))
		}

		if result.TotalsByType["identificador"] != test.expectedIdentifiers {
			t.Errorf("Para '%s', se esperaban %d identificadores pero se obtuvieron %d",
				test.input, test.expectedIdentifiers, result.TotalsByType["identificador"])
		}

		if result.TotalsByType["identificador_unicos"] != test.expectedUniqueIdent {
			t.Errorf("Para '%s', se esperaban %d identificadores únicos pero se obtuvieron %d",
				test.input, test.expectedUniqueIdent, result.TotalsByType["identificador_unicos"])
		}

		if result.TotalsByType["numero"] != test.expectedNumbers {
			t.Errorf("Para '%s', se esperaban %d números pero se obtuvieron %d",
				test.input, test.expectedNumbers, result.TotalsByType["numero"])
		}

		if result.TotalsByType["numero_unicos"] != test.expectedUniqueNumbers {
			t.Errorf("Para '%s', se esperaban %d números únicos pero se obtuvieron %d",
				test.input, test.expectedUniqueNumbers, result.TotalsByType["numero_unicos"])
		}

		if result.TotalsByType["error"] != test.expectedErrors {
			t.Errorf("Para '%s', se esperaban %d errores pero se obtuvieron %d",
				test.input, test.expectedErrors, result.TotalsByType["error"])
		}
	}
}
