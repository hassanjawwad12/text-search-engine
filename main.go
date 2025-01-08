package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	utils "github.com/hassanjawwad12/text-search-engine/utils"
)

func main() {
	var dataPath, query string
	flag.StringVar(&dataPath, "p", "data.xml", "wiki extract dump path")
	flag.StringVar(&query, "q", "small wild cat", "search query")
	flag.Parse()

	fmt.Println("Full text search is in progress...")
	start := time.Now()

	// Load the doc by giving the path
	docs, err := utils.LoadDocument(dataPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Time it took to load the documents
	log.Printf("Loaded %d documents in %v\n", len(docs), time.Since(start))
	start = time.Now()

	// create the inverted index
	idx := make(utils.Index)

	// Add the documents
	idx.Add(docs)
	log.Printf("Indexing of %d documents took %v\n", len(docs), time.Since(start))
	start = time.Now()

	// match the query index with the created inverted index
	matchedIds := idx.Search(query)
	log.Printf("Search found '%d' documents in %v\n", len(matchedIds), time.Since(start))

	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
