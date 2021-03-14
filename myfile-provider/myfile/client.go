package myfile

import (
	"io/ioutil"
	"os"
)

// FileClient is a simple client to interact with files
type FileClient struct {
	// Encoding. Not used. Defined here for the example
	Encoding string
}

// Create file
func (c FileClient) Create(path string, contents string) error {
	var err error
	var f *os.File
	if f, err = os.Create(path); err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(contents); err != nil {
		return err
	}
	return nil
}

// Read file
func (c FileClient) Read(path string) (string, error) {
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(path); err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Update file
func (c FileClient) Update(path string, contents string) error {
	var err error
	if err = os.Truncate(path, 0); err != nil {
		return err
	}
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(contents); err != nil {
		return err
	}
	return nil
}

// Delete file
func (c FileClient) Delete(path string) error {
	err := os.Remove(path)
	return err
}

// Owner of the file
func (c FileClient) Owner(path string) (string, error) {
	return "me", nil //stub
}
