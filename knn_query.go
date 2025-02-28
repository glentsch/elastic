package elastic

// KNNQuery allows to define KNN as filters.
type KNNQuery struct {
	query         Query
	field         string
	value         []float32
	k             int
	boost         float32
	numCandidates int
	similarity    float32
}

// NewKNNQuery creates and initializes a new KNNQuery.
func NewKNNQuery(
	query Query,
	field string,
	value []float32,
	boost float32,
	k int,
	numCandidates int,
	similarity float32,
) *KNNQuery {
	return &KNNQuery{
		query:         query,
		field:         field,
		value:         value,
		boost:         boost,
		k:             k,
		numCandidates: numCandidates,
		similarity:    similarity,
	}
}

func (q *KNNQuery) Source() (interface{}, error) {
	source := make(map[string]map[string]interface{})
	field := make(map[string]interface{})

	src, err := q.query.Source()
	if err != nil {
		return nil, err
	}
	field["query_vector"] = q.value
	field["k"] = q.k
	field["filter"] = src
	field["boost"] = q.boost
	field["field"] = q.field
	field["num_candidates"] = q.numCandidates

	source["knn"] = field

	return source, nil
}
