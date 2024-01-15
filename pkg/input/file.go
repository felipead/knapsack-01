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

func LoadProblemFromFile(filename string) (*Problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	numItems, err := readInt(scanner)
	if err != nil {
		return nil, fmt.Errorf("unable to read first line as Number of Items: %w", err)
	}

	capacity, err := readDecimal(scanner)
	if err != nil {
		return nil, fmt.Errorf("unable to read second line as Capacity: %w", err)
	}

	items := make([]model.Item, 0, numItems)
	for i := 0; i < numItems; i++ {
		index := i + 1
		item, err := readItem(scanner, index)
		if err != nil {
			return nil, fmt.Errorf("unable to read Item #%v: %w", index, err)
		}
		items = append(items, *item)
	}

	return &Problem{
		Capacity: capacity,
		Items:    items,
	}, nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
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

func readDecimal(scanner *bufio.Scanner) (decimal.Decimal, error) {
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

func readItem(scanner *bufio.Scanner, index int) (*model.Item, error) {
	line, err := readLine(scanner)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, fmt.Errorf("line was expected to have 2 fields, got %v", len(fields))
	}

	item := model.Item{
		Index: index,
	}
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
