package gocsv

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/tgoikawa/excelcsvjp"
)

// Unmarshal parses the CSV from the reader in the interface.
func Unmarshal(r io.Reader, out interface{}) error {
	return gocsv.UnmarshalCSV(excelcsvjp.NewReader(r), out)
}

// Marshal returns  CSV in writer from the interface.
func Marshal(w io.Writer, in interface{}) error {
	return gocsv.MarshalCSV(in, excelcsvjp.NewSafeCSVWriter(csvjp.NewWriter(w)))
}
