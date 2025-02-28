package elastic

type StandardRetriever struct {
	Filter []Query `json:"filter,omitempty"`
	Query  Query   `json:"query,omitempty"`
}

func (r *StandardRetriever) Source() (interface{}, error) {
	src := make(map[string]interface{})
	if r.Query != nil {
		query, err := r.Query.Source()
		if err != nil {
			return nil, err
		}
		src["query"] = query
	}
	if r.Filter != nil {
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
