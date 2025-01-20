# Word Frequency Analysis

## Description:

This project is a Go-based command-line utility designed to analyze a text file and display the most frequently occurring words.
The program reads the file content, counts the occurrences of each word, and displays the top 10 most common words.

## Features:

### - Reading lines from a file

### - Cleaning words from non-alphanumeric characters

### - Counting word frequency

### - Sorting words in descending order of frequency

### - Displaying the top 10 most frequently occurring words

## Usage:

Place a text file in the project's root directory and name it test.txt.

Run the program:

```
go run main.go
```

The program will output the 10 most frequently occurring words along with their count.

## Main Code Components:

ReadFile(filename string) ([]string, error) — Reads lines from a file.

CountWords(lines []string) map[string]int — Counts the occurrences of each word, converting them to lowercase and removing extra characters.

GetTopWords(wordCount map[string]int, n int) WordFrequencies — Forms a list of the N most frequently occurring words.

print() — Outputs the list of words with their frequency to the console.

## Requirements:

Go 1.18+

## License:

This project is distributed under the MIT license.

# Анализ частоты слов в файле

## Описание:

Этот проект представляет собой консольную утилиту на языке Go, предназначенную для анализа текстового файла и вывода самых часто встречающихся слов. 
Программа читает содержимое файла, подсчитывает вхождения каждого слова и отображает топ-10 самых частых слов.

## Функциональность:

### -Чтение строк из файла

### -Очистка слов от небуквенно-цифровых символов

### -Подсчет частоты встречаемости слов

### -Сортировка слов по убыванию частоты

### -Вывод топ-10 наиболее часто встречающихся слов

## Использование:

Поместите текстовый файл в корневую директорию проекта и назовите его test.txt.

Запустите программу:

```
go run main.go
```

Программа выведет 10 самых часто встречающихся слов с их количеством вхождений.

Основные компоненты кода

ReadFile(filename string) ([]string, error) — считывает строки из файла.

CountWords(lines []string) map[string]int — подсчитывает количество вхождений каждого слова, приводя их к нижнему регистру и удаляя лишние символы.

GetTopWords(wordCount map[string]int, n int) WordFrequencies — формирует список из N самых часто встречающихся слов.

print() — выводит список слов с их частотой в консоль.

Требования

Go 1.18+

Лицензия

Этот проект распространяется под лицензией MIT.

