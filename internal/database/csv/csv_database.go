package csv_database

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/kerezsiz42/kinko/internal/database"
)

type CsvDatabase struct {
	FilePath string
}

func NewCsvDatabase(filePath string) *CsvDatabase {
	db := &CsvDatabase{FilePath: filePath}
	if _, err := db.Load(); err != nil {
		panic(err)
	}

	return db
}

func (db *CsvDatabase) Load() ([]database.Record, error) {
	file, err := os.Open(db.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var parsedRecords []database.Record
	for lineIndex, record := range records {
		// Skip header
		if lineIndex == 0 {
			continue
		}

		if len(record) != database.RECORD_LEN {
			return nil, fmt.Errorf("line %d has %d rows instead of %d", lineIndex, len(record), database.RECORD_LEN)
		}

		parsedRecords = append(parsedRecords, database.Record{
			Key:       record[0],
			Name:      record[1],
			Public:    record[2],
			Private:   record[3],
			Info:      record[4],
			UpdatedAt: record[5],
			CreatedAt: record[6],
		})
	}

	return parsedRecords, nil
}

func (db *CsvDatabase) Persist([]database.Record) error {
	return errors.New("not implemented")
}
