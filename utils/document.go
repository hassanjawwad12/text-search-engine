package utils

import (
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// Load document from given path into a slice of Document struct
func LoadDocument(path string) ([]Document, error) {

	// Open the file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// process the XML content
	dec := xml.NewDecoder(f)

	// Define XML structure
	// Map XML elements named <doc> to the Documents field
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	// A loop iterates over each Document to assign a unique ID to each element based on its index in the slice.
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
