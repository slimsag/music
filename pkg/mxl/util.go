package mxl

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// UnmarshalFile is a helper which unmarshals the named file.
func UnmarshalFile(file string) (*Document, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var doc *Document
	if err := xml.NewDecoder(f).Decode(&doc); err != nil {
		return nil, err
	}
	return doc, nil
}

// MarshalFile is a helper which marshals the named file.
func MarshalFile(file string, doc *Document) error {
	data, err := MarshalIndent(doc, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0666)
}

// Marshal is a helper which wraps xml.Marshal and prepends the Header.
func Marshal(v interface{}) ([]byte, error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return []byte(Header + string(data)), nil
}

// MarshalIndent is a helper which wraps xml.MarshalIndent and prepends the Header.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	data, err := xml.MarshalIndent(v, prefix, indent)
	if err != nil {
		return nil, err
	}
	return []byte(Header + string(data)), nil
}
