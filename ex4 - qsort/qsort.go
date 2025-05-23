package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Функция для разбора текста на отдельные числа
func parseInts(strs []string) ([]int64, error) { // Переименовали параметр с strings на strs
	var result []int64
	for _, s := range strs {
		// Пропускаем пустые строки и пробельные символы
		trimmed := strings.TrimSpace(s) // Теперь strings ссылается на пакет, а не на параметр
		if trimmed == "" {
			continue
		}

		// Конвертация строки в число
		m, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			log.Printf("Warning: Couldn't convert '%s': %v", s, err)
			continue
		}
		result = append(result, m)
	}
	return result, nil
}

// Главная функция
func main() {
	// Шаг 1: Считываем файл
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords) // Считываем каждое слово как отдельное значение
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Шаг 2: Преобразуем строки в числа
	ints, err := parseInts(words)
	if err != nil {
		log.Fatalf("Error parsing integers: %v", err)
	}

	// Шаг 3: Сортируем
	sort.Slice(ints, func(m, n int) bool {
		return ints[m] < ints[n]
	})

	// Шаг 4: Записываем результат в файл
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Cannot create output file: %v", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, val := range ints {
		fmt.Fprintf(writer, "%d ", val)
	}
	writer.Flush()

	// Подтверждение успешного завершения
	fmt.Println("Output written to output.txt.")
}