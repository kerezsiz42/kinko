package csv_database

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/kerezsiz42/kinko/internal/database"
	"github.com/kerezsiz42/kinko/internal/utils"
)

type CsvDatabase struct {
	DbFilePath string
	mu         sync.Mutex
}

func NewCsvDatabase(filePath string) *CsvDatabase {
	db := &CsvDatabase{DbFilePath: filePath}
	if _, err := db.Load(); err != nil {
		panic(err)
	}

	return db
}

func (db *CsvDatabase) Load() ([]database.Record, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Open(db.DbFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open csv file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not parse csv records: %w", err)
	}

	var parsedRecords []database.Record
	for line, record := range records {
		if line == 0 && !utils.AreSlicesIdentical(record, database.RecordFields) {
			return nil, fmt.Errorf("invalid header, should be exactly %s", database.RecordFields)
		}

		if len(record) != len(database.RecordFields) {
			return nil, fmt.Errorf("line %d has %d rows instead of %d", line, len(record), len(database.RecordFields))
		}

		parsedRecords = append(parsedRecords, database.RecordFromCSV(record))
	}

	return parsedRecords, nil
}

func (db *CsvDatabase) Persist(records []database.Record) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Open(db.DbFilePath)
	if err != nil {
		return fmt.Errorf("could not open csv file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(database.RecordFields); err != nil {
		return fmt.Errorf("could not write header: %w", err)
	}

	for _, record := range records {
		if err := writer.Write(record.ToCSV()); err != nil {
			return fmt.Errorf("could not write record: %w", err)
		}
	}

	return nil
}
