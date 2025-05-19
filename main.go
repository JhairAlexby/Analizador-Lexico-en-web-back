package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// TokenAnalysis contiene el resultado del análisis léxico
type TokenAnalysis struct {
	Tokens       []Token          `json:"tokens"`
	TotalsByType map[string]int   `json:"totalsByType"`
	UniqueTokens map[string]Token `json:"uniqueTokens"`
}

// Token representa un token identificado en el texto
type Token struct {
	Value string `json:"value"`
	Type  string `json:"type"` // "identificador", "numero", "error"
}

func main() {
	http.HandleFunc("/analyze", corsMiddleware(analyzeHandler))
	http.HandleFunc("/health", corsMiddleware(healthCheckHandler))

	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "8080" // Puerto por defecto si no se especifica
	}

	log.Printf("Servidor iniciado en el puerto %s", portEnv)
	log.Fatal(http.ListenAndServe(":"+portEnv, nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// healthCheckHandler responde a las solicitudes de verificación de estado
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Responder con un estado 200 OK y un mensaje simple
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// analyzeHandler procesa la solicitud de análisis léxico
func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	// Solo permitir método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar la solicitud JSON
	var request struct {
		Text string `json:"text"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Realizar análisis léxico
	analysis := analyzeText(request.Text)

	// Enviar respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

// analyzeText realiza el análisis léxico del texto
func analyzeText(text string) TokenAnalysis {
	// Dividir el texto en tokens (palabras y números)
	words := tokenize(text)

	// Analizar cada token
	var tokens []Token
	uniqueTokens := make(map[string]Token)
	totalsByType := make(map[string]int)

	for _, word := range words {
		if word == "" {
			continue
		}

		// Clasificar el token
		token := classifyToken(word)
		tokens = append(tokens, token)

		// Contar por tipo
		totalsByType[token.Type]++

		// Registrar tokens únicos
		if token.Type != "error" {
			// Solo registramos como únicos los tokens válidos
			uniqueTokens[token.Value] = token
		}
	}

	// Actualizar conteo de tokens únicos
	totalsByType["identificador_unicos"] = 0
	totalsByType["numero_unicos"] = 0

	for _, token := range uniqueTokens {
		if token.Type == "identificador" {
			totalsByType["identificador_unicos"]++
		} else if token.Type == "numero" {
			totalsByType["numero_unicos"]++
		}
	}

	return TokenAnalysis{
		Tokens:       tokens,
		TotalsByType: totalsByType,
		UniqueTokens: uniqueTokens,
	}
}

// tokenize divide el texto en tokens individuales
func tokenize(text string) []string {
	// Eliminar caracteres especiales y reemplazarlos por espacios
	regex := regexp.MustCompile(`[^\w\s]`)
	text = regex.ReplaceAllString(text, " ")

	// Dividir por espacios
	return strings.Fields(text)
}

// classifyToken clasifica un token según las reglas especificadas
func classifyToken(token string) Token {
	// Verificar si es un número válido (solo dígitos)
	isNumber := regexp.MustCompile(`^\d+$`).MatchString(token)
	if isNumber {
		return Token{Value: token, Type: "numero"}
	}

	// Verificar si es un identificador válido (solo letras)
	isIdentifier := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(token)
	if isIdentifier {
		return Token{Value: token, Type: "identificador"}
	}

	// Si contiene mezcla de letras y números, es un error
	return Token{Value: token, Type: "error"}
}
