package utils

type Index map[string][]int

func (idx Index) Add(docs []Document) {
	for i, doc := range docs {

		for _, token := analyze(doc.Text) {
			ids := idx[token]
		}
	}
}
