package file

func (pdf *PdfFile) Close() error {
	return pdf.file.Close()
}
