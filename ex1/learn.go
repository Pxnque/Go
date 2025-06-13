package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "ingrese el nombre de su csv, nombre.csv. Funciona con csv del formato pregunta y respuesta")
	timeLimit := flag.Int("timer", 30, "Lo que va a durar el quiz (default 30)")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for _, pregunta := range preguntas {
		fmt.Println("Pregunta:", pregunta[0])
		fmt.Print("Ingresa tu respuesta: ")
		answerCh := make(chan string)
		var useranswer string
		go func() {
			fmt.Scanln(&useranswer)
			useranswer = strings.ToLower(useranswer)
			useranswer = strings.TrimSpace(useranswer)
			answerCh <- useranswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n Se termino el tiempo!")
			fmt.Printf("\n Obtuviste %v buenas de %v", contador, len(preguntas))
			return
		case answer := <-answerCh:
			if answer == useranswer {
				contador++
			}
		}

	}

}
