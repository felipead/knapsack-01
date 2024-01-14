package input

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/felipead/knapsack-01/pkg/model"
)

func LoadProblemFromFile(filename string) (Problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	numItems, err := readLineAsInt(scanner)
	if err != nil {
		return Problem{}, fmt.Errorf("unable to read first line as Number of Items: %w", err)
	}

	capacity, err := readLineAsDecimal(scanner)
	if err != nil {
		return Problem{}, fmt.Errorf("unable to read second line as Capacity: %w", err)
	}

	items := make([]model.Item, 0, numItems)
	for i := 0; i < numItems; i++ {
		item, err := readLineAsItem(scanner)
		if err != nil {
			return Problem{}, fmt.Errorf("unable to read Item #%v: %w", i+1, err)
		}
		items = append(items, *item)
	}

	return Problem{
		Capacity: capacity,
		Items:    items,
	}, nil
}

func readLineAsInt(scanner *bufio.Scanner) (int, error) {
	line, err := readLine(scanner)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func readLineAsDecimal(scanner *bufio.Scanner) (decimal.Decimal, error) {
	line, err := readLine(scanner)
	if err != nil {
		return decimal.Zero, err
	}
	value, err := decimal.NewFromString(line)
	if err != nil {
		return decimal.Zero, err
	}
	return value, nil
}

func readLineAsItem(scanner *bufio.Scanner) (*model.Item, error) {
	line, err := readLine(scanner)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, fmt.Errorf("line was expected to have 2 fields, got %v", len(fields))
	}

	item := model.Item{}
	item.Value, err = decimal.NewFromString(fields[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Value field: %w", err)
	}
	item.Cost, err = decimal.NewFromString(fields[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Cost field: %w", err)
	}

	return &item, nil
}

func readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if scanner.Err() == nil {
			return "", errors.New("file stream ended with EOF")
		}
		return "", fmt.Errorf("file stream ended with: %w", scanner.Err())
	}
	return scanner.Text(), nil
}
