package file

import "os"

type PdfFile struct {
	path string
	file *os.File
	size int64
}

func NewPdfFile(path string) PdfFile {
	return PdfFile{
		path: path,
	}
}
