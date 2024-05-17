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
		return errors.New(FILE_TOO_SMALL_ERROR)
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
		message := "File is not ending header line"
		return errors.New(FILE_FORMAT_ERROR_PREFIX + message)
	}
	return nil
}

func checkValidVersion(readBuffer []byte) error {
	hasValidMajorVersion := readBuffer[5] == 1
	hasValidMinorVersion := (readBuffer[7] > '0' && readBuffer[7] < '7')
	if !hasValidMajorVersion || !hasValidMinorVersion {
		return errors.New(VERSION_NOT_SUPPORTED_ERROR)
	}
	return nil
}

func checkHeaderPrefix(readBuffer []byte) error {
	if !bytes.HasPrefix(readBuffer, []byte(PDF_HEADER_PREFIX)) {
		message := "File is missing the beginning header"
		return errors.New(FILE_FORMAT_ERROR_PREFIX + message)
	}
	return nil
}
