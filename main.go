package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// KeyValue связывает слово с количеством вхождений.
type KeyValue struct {
	word  string
	count int
}

// WordFrequencies — это слайс KeyValue, к которому прикручен метод print.
type WordFrequencies []KeyValue

func main() {
	filename := "test.txt"

	// Считываем строки из файла
	lines, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Подсчитываем вхождения слов
	wordCount := CountWords(lines)

	// Получаем топ-N слов, в данном случае топ-10
	topWords := GetTopWords(wordCount, 10)

	// Выводим результат
	topWords.print()
}

// ReadFile читает все строки из файла и возвращает их как слайс string.
func ReadFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Регулярное выражение для “очистки” от не-буквенно-цифровых символов.
// \p{L} — любая буква (Unicode)
// \p{N} — любая цифра (Unicode)
// [^\p{L}\p{N}]+ — все символы, не являющиеся буквами или цифрами
var re = regexp.MustCompile(`[^\p{L}\p{N}]+`)

// CountWords подсчитывает вхождения каждого уникального слова, приводя к нижнему регистру
// и удаляя лишние символы регуляркой.
func CountWords(lines []string) map[string]int {
	wordCount := make(map[string]int)

	for _, line := range lines {
		// Разбиваем строку на “предварительные” слова
		rawWords := strings.Fields(line)
		for _, w := range rawWords {
			// Используем регулярку: удаляем все не-буквенно-цифровые символы
			cleanWord := re.ReplaceAllString(strings.ToLower(w), "")
			if cleanWord != "" {
				wordCount[cleanWord]++
			}
		}
	}
	return wordCount
}

// GetTopWords возвращает слайс из N наиболее часто встречающихся слов.
func GetTopWords(wordCount map[string]int, n int) WordFrequencies {
	var freq WordFrequencies
	for k, v := range wordCount {
		freq = append(freq, KeyValue{k, v})
	}

	// Сортируем по убыванию частоты
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].count > freq[j].count
	})

	// Проверяем, не меньше ли длина слайса, чем N
	if n > len(freq) {
		n = len(freq)
	}

	return freq[:n]
}

// print — метод для вывода списка слов с их частотой.
func (wf *WordFrequencies) print() {
	for _, kv := range *wf {
		fmt.Printf("%-20s %d\n", kv.word, kv.count)
	}
}
