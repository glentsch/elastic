package elastic

type KnnRetriever struct {
	Field         string    `json:"field"`
	Filter        []Query   `json:"filter,omitempty"`
	K             int       `json:"k"`
	NumCandidates int       `json:"num_candidates"`
	QueryVector   []float32 `json:"query_vector,omitempty"`
}

func (r *KnnRetriever) Source() (interface{}, error) {
	src := make(map[string]interface{})
	src["field"] = r.Field
	src["k"] = r.K
	src["num_candidates"] = r.NumCandidates
	if r.QueryVector != nil {
		src["query_vector"] = r.QueryVector
	}
	if len(r.Filter) > 0 {
		filters := make([]interface{}, len(r.Filter))
		for i, filter := range r.Filter {
			filterSrc, err := filter.Source()
			if err != nil {
				return nil, err
			}
			filters[i] = filterSrc
		}
		src["filter"] = filters
	}
	return src, nil
}
