package jobs

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var Lf = log.Printf
var L = log.Println

func ConvertAsNeeded(file_path string) error {
	const converted_path = "tmp/converted.xml"

	file, err := OpenFile(file_path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc, err := DetermineEncoding(file)
	if err != nil {
		return err
	}
	L("Encoding is", enc)

	if enc == "iso-8859-1" {
		new_file, err := os.Create(converted_path)
		if err != nil {
			L("Unable to create temporary file for conversion:", err)
			return err
		}
		defer new_file.Close()
		file_path = converted_path

		err = ISO8859toUTF8(file, new_file)
		if err != nil {
			return err
		}
	}
	return err
}

func ISO8859toUTF8(file *os.File, new_file *os.File) error {
	L("Converting 8859-1 to UTF-8...")
	file.Seek(0, 0) // make sure we are at the top of the file

	// Wrap an ISO8859-1 decoder
	utf8_reader := charmap.ISO8859_1.NewDecoder().Reader(file)

	// Read the first few bytes
	first_bytes, err := ReadFromStart(utf8_reader, 200)
	if err != nil {
		return err
	}

	// Replace the encoding attribute so it says "utf-8"
	b_utf8 := []byte("utf-8")
	// At the low level here, we have to do our own casing
	// And you've got to handle upper and lower too
	b_8859_lwr := []byte("iso-8859-1")
	new_bytes := bytes.Replace(first_bytes, b_8859_lwr, b_utf8, 1)
	b_8859_upr := []byte("ISO-8859-1")
	new_bytes = bytes.Replace(new_bytes, b_8859_upr, b_utf8, 1)
	Lf("Replaced buffer is:\n%s", new_bytes)

	// Write the change
	_, err = new_file.Write(new_bytes)
	if err != nil {
		log.Fatal("Error writing to output file", new_file)
	}
	// Make sure the file pointer is updated to the new offset
	new_file.Seek(int64(len(new_bytes)), 0)

	// Copy the rest
	_, err = io.Copy(new_file, utf8_reader)
	if err != nil {
		log.Fatal("Error writing to output file", new_file)
	}

	return err
}

func DetermineEncoding(file *os.File) (string, error) {
	var enc string

	first_bytes, err := ReadFromStart(file, 300)
	if err != nil {
		return enc, err
	}
	first_part := strings.ToLower(string(first_bytes))

	switch {
	case strings.Contains(first_part, "iso-8859-1"):
		enc = "iso-8859-1"
	case strings.Contains(first_part, "utf-8"):
		enc = "utf-8"
	default:
		enc = "utf-8" // default to utf-8 for now
	}
	return enc, err
}

func OpenFile(file_path string) (*os.File, error) {
	var err error
	file, err := os.Open(file_path)
	if err != nil {
		L("Could not open input file:", file_path)
		return file, err
	}
	return file, err
}

func ReadFromStart(reader io.Reader, size int) ([]byte, error) {
	start_bytes := make([]byte, size)
	_, err := reader.Read(start_bytes)
	return start_bytes, err
}
