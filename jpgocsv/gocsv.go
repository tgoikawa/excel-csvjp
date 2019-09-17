package jpgocsv

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/tgoikawa/jpcsv"
)

// Unmarshal parses the CSV from the reader in the interface.
func Unmarshal(r io.Reader, out interface{}) error {
	return gocsv.UnmarshalCSV(jpcsv.NewReader(r), out)
}

// Marshal returns  CSV in writer from the interface.
func Marshal(w io.Writer, in interface{}) error {
	return gocsv.MarshalCSV(in, gocsv.NewSafeCSVWriter(jpcsv.NewWriter(w)))
}
