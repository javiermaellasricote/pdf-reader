package file

import (
	"bytes"
	"errors"
)

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
	if err = checkTailEOF(readBuffer); err != nil {
		return err
	}
	return nil
}

func checkTailEOF(readBuffer []byte) error {
	trimmedBuffer := bytes.TrimRight(readBuffer, "\r\n\t")
	if !bytes.HasSuffix(trimmedBuffer, []byte(EOF_TAG)) {
		message := "Missing EOF tag at the end of file"
		return errors.New(FILE_FORMAT_ERROR_PREFIX + message)
	}
	return nil
}

// TODO: unfinished func
func findLastCrossReference(readBuffer []byte) (int64, error) {
	xRefTag := []byte(START_XREF_TAG)
	tagLength := len(xRefTag)
	bufferLength := len(readBuffer)
	maxIdx := bufferLength
	maxIters := bufferLength / tagLength
	for iter := 0; iter < maxIters; iter++ {
		isMaxIdxValid := maxIdx > 0
		tagFitsInBuffer := maxIdx+tagLength < bufferLength
		if !isMaxIdxValid || !tagFitsInBuffer {
			return 0, errors.New("")
		}
	}
	return 0, nil
}
