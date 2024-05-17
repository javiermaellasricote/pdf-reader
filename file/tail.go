package file

import "errors"

func checkTail(pdf *PdfFile) error {
	var desiredReadLength int64 = 100
	offset := pdf.size - desiredReadLength
	if offset <= 9 {
		return errors.New(FILE_TOO_SMALL_ERROR)
	}
	readBuffer := make([]byte, desiredReadLength)
	length, err := pdf.file.ReadAt(readBuffer, pdf.size)
	if err != nil {
		return err
	}
	if int64(length) != desiredReadLength {
		return errors.New(FILE_TOO_SMALL_ERROR)
	}
	return nil
}
