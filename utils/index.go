package utils

type Index map[string][]int

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// dont add the same id twice
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func intersect(a, b []int) []int {
	maxlen := len(a)
	if len(b) > maxlen {
		maxlen = len(b)
	}
	r := make([]int, 0, maxlen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			r = append(r, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return r
}

func (idx Index) Search(query string) []int {
	var r []int
	for _, token := range analyze(query) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				// if we get disjoint set
				r = intersect(r, ids)
			}
		} else {
			// token does not exist
			return nil
		}
	}
	return r
}
