package file

import (
	"bytes"
	"errors"
)

func checkHeader(pdf *PdfFile) error {
	desiredReadLength := 9
	readBuffer := make([]byte, desiredReadLength)
	length, err := pdf.file.Read(readBuffer)
	if err != nil {
		return err
	}
	if desiredReadLength != length {
		return errors.New("File is too small to be a valid PDF")
	}
	if err = checkHeaderPrefix(readBuffer); err != nil {
		return err
	}
	if err = checkValidVersion(readBuffer); err != nil {
		return err
	}
	if err = checkHeaderEOL(readBuffer); err != nil {
		return err
	}
	return nil
}

func checkHeaderEOL(readBuffer []byte) error {
	if readBuffer[8] != '\r' && readBuffer[8] != '\n' {
		return errors.New("Incorrect format: File is not ending header line")
	}
	return nil
}

func checkValidVersion(readBuffer []byte) error {
	hasValidMajorVersion := readBuffer[5] == 1
	hasValidMinorVersion := (readBuffer[7] > '0' && readBuffer[7] < '7')
	if !hasValidMajorVersion || !hasValidMinorVersion {
		return errors.New("PDF version not supported")
	}
	return nil
}

func checkHeaderPrefix(readBuffer []byte) error {
	if !bytes.HasPrefix(readBuffer, []byte("%PDF-")) {
		message := "Incorrect format: File is missing the beginning header"
		return errors.New(message)
	}
	return nil
}
