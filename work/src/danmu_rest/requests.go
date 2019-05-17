package danmu_rest

//SimilarWordRequest
type SimilarWordRequest struct {
	Word string `json:"word"` //to make JSON not have uppercased Keys
}

//SimilarWordRequests ...a slice of SimilarWordRequest
type SimilarWordRequests []SimilarWordRequest

//SimilarWordsResponse
type SimilarWordsResponse struct {
	SimilarWords string `json:"similarwords"`
}

//SimilarWordsResponses ...a slice of SimilarWordResponse
type SimilarWordsResponses []SimilarWordsResponse

//RecentTopicRequest
type RecentTopicRequest struct {
	Interval int `json:"interval"`
}

//RecentTopicResponse
type RecentTopicResponse struct {
	RecentTopic string `json:"recenttopic"`
}
