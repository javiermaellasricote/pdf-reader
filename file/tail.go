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
	if err = checkTailCrossReference(readBuffer); err != nil {
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

func checkTailCrossReference(readBuffer []byte) error {
	tagLength := len(START_XREF_TAG)
	tagStartIdx, err := findLastCrossReference(readBuffer)
	if err != nil {
		return nil
	}
	trailingBuffer := readBuffer[(tagLength + tagStartIdx):]
	trimmedBuffer := bytes.TrimLeft(trailingBuffer, "\n\r\t")
	if !checkByteIsNumber(trimmedBuffer[0]) {
		message := "Cross reference must start with integer"
		return errors.New(FILE_FORMAT_ERROR_PREFIX + message)
	}
	return nil
}

// Finds last cross-reference tag between two EOL characters
func findLastCrossReference(readBuffer []byte) (int, error) {
	xRefTag := []byte(START_XREF_TAG)
	tagLength := len(xRefTag)
	bufferLength := len(readBuffer)
	maxIdx := bufferLength
	maxIters := bufferLength / tagLength
	for iter := 0; iter < maxIters; iter++ {
		isMaxIdxValid := maxIdx > 0
		tagFitsInBuffer := maxIdx+tagLength < bufferLength
		if !isMaxIdxValid || !tagFitsInBuffer {
			return 0, errors.New("Tag not found in buffer")
		}
		idx := bytes.LastIndex(readBuffer[:maxIdx], xRefTag)
		eolPreceedsTag := checkByteIsEOL(readBuffer[idx-1])
		eolFollowsTag := checkByteIsEOL(readBuffer[idx+tagLength])
		if eolPreceedsTag && eolFollowsTag {
			return int(idx), nil
		}
		maxIdx = idx
	}
	return 0, errors.New("Tag not found in buffer")
}
