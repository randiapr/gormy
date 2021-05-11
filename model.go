package gormy

type Pagination struct {
	TotalRecords int64       `json:"totalRecords"`
	TotalPages   int64       `json:"totalPages"`
	Data         interface{} `json:"data"`
	Offset       int64       `json:"offset"`
	Limit        int64       `json:"limit"`
	Page         int64       `json:"page"`
	PrevPage     int64       `json:"prevPage"`
	NextPage     int64       `json:"nextPage"`
}
