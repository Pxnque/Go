package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "ingrese el nombre de su csv, nombre.csv. Funciona con csv del formato pregunta y respuesta")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Println("No se pudo abrir el csv:", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	preguntas, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error al leer el csv:", err)
		os.Exit(1)
	}
	contador := 0
	var useranswer string
	for _, pregunta := range preguntas {

		fmt.Println("Pregunta:", pregunta[0])
		fmt.Print("Ingresa tu respuesta: ")
		fmt.Scanln(&useranswer)
		useranswer = strings.ToLower(useranswer)
		useranswer = strings.TrimSpace(useranswer)
		if useranswer == pregunta[1] {
			contador++
		}
		fmt.Println()
	}
	fmt.Printf("Obtuviste %v buenas de %v", contador, len(preguntas))

}
