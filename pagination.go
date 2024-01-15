package horizon

// Pagination
type SearchResults[T any] struct {
	Results   []T  `json:"results"`
	PageIndex int  `json:"pageIndex"`
	PageSize  int  `json:"pageSize"`
	Count     int  `json:"count,omitempty"`
	HasMore   bool `json:"hasMore"`
}

type SortOrder string

const (
	Ascendant     SortOrder = "Asc"
	Descendant    SortOrder = "Desc"
	KeyAscendant  SortOrder = "KeyAsc"
	KeyDescendant SortOrder = "KeyDesc"
)

type SortFields struct {
	Element string    `json:"element"`
	Order   SortOrder `json:"order"`
}
