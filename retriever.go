package elastic

type Retriever struct {
	Knn      *KnnRetriever      `json:"knn,omitempty"`
	Rrf      *RrfRetriever      `json:"rrf,omitempty"`
	Standard *StandardRetriever `json:"standard,omitempty"`
}

func (r *Retriever) Source() (interface{}, error) {
	source := make(map[string]interface{})

	if r.Knn != nil {
		src, err := r.Knn.Source()
		if err != nil {
			return nil, err
		}
		source["knn"] = src
	}

	if r.Rrf != nil {
		src, err := r.Rrf.Source()
		if err != nil {
			return nil, err
		}
		source["rrf"] = src
	}

	if r.Standard != nil {
		src, err := r.Standard.Source()
		if err != nil {
			return nil, err
		}
		source["standard"] = src
	}

	return source, nil
}
