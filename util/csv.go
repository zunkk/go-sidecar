package util

import (
	"encoding/csv"
	"os"

	"github.com/jszwec/csvutil"
	"github.com/pkg/errors"
)

func WriteStructsToCSVFile[T any](csvFilePath string, structs []T) error {
	analyzeFile, err := os.OpenFile(csvFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return errors.Wrapf(err, "failed to create csv file: %s", csvFilePath)
	}
	csvWriter := csv.NewWriter(analyzeFile)
	csvEncoder := csvutil.NewEncoder(csvWriter)
	err = csvEncoder.Encode(structs)
	if err != nil {
		return errors.Wrapf(err, "failed to write csv file: %v", structs)
	}
	csvWriter.Flush()
	_ = analyzeFile.Close()
	return nil
}
