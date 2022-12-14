package read_pdf

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func ReadPdf(path string) (string, error) {
	file, reader, err := pdf.Open(path)

	defer file.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := reader.GetPlainText()

	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	return buf.String(), nil
}
