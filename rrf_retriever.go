package elastic

type RrfRetriever struct {
	Filter         Query       `json:"filter,omitempty"`
	RankConstant   *int        `json:"rank_constant,omitempty"`
	RankWindowSize *int        `json:"rank_window_size,omitempty"`
	Retrievers     []Retriever `json:"retrievers"`
}

func (r *RrfRetriever) Source() (interface{}, error) {
	src := make(map[string]interface{})
	if r.RankConstant != nil {
		src["rank_constant"] = *r.RankConstant
	}
	if r.RankWindowSize != nil {
		src["rank_window_size"] = *r.RankWindowSize
	}
	if r.Filter != nil {
		filter, err := r.Filter.Source()
		if err != nil {
			return nil, err
		}
		src["filter"] = filter
	}
	if r.Retrievers != nil {
		retrievers := make([]interface{}, len(r.Retrievers))
		for i, retriever := range r.Retrievers {
			retrieverSrc, err := retriever.Source()
			if err != nil {
				return nil, err
			}
			retrievers[i] = retrieverSrc
		}
		src["retrievers"] = retrievers
	}

	return src, nil
}
