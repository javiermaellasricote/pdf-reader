package file

import (
	"os"
)

func (pdf *PdfFile) Open() error {
	file, err := os.Open(pdf.path)
	if err != nil {
		return err
	}
	pdf.file = file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	pdf.size = fileInfo.Size()
	if err = checkHeader(pdf); err != nil {
		return err
	}
	if err = checkTail(pdf); err != nil {
		return err
	}
	return nil
}
