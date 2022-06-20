package command

type ElasticsearchConfig struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Shards   Shard `json:"_shards"`
	Hits     Hits  `json:"hits"`
}

type Shard struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}

type Hits struct {
	Total    Total  `json:"total"`
	MaxScore *int64 `json:"max_score"`
	Hits     []Hit  `json:"hits"`
}

type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type Hit struct {
	Index     string                 `json:"_index"`
	Type      string                 `json:"_type"`
	Id        string                 `json:"_id"`
	Score     *int64                 `json:"_score"`
	Source    map[string]interface{} `json:"_source"`
	Highlight map[string]interface{} `json:"highlight"`
	Sort      []interface{}          `json:"sort"`
}
