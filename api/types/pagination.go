package types

type Pagination struct {
	TotalPages   int64 `json:"totalPages"`
	CurrentPage  int64 `json:"currentPage"`
	Limit        int64 `json:"limit"`
	TotalResults int64 `json:"totalResults"`
}
