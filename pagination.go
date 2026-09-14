package response

// DefaultPageNumber, DefaultPageSize, and MaxPageSize are shared pagination defaults.
// Handlers and repositories SHOULD clamp requested page size to MaxPageSize.
const (
	DefaultPageNumber = 1
	DefaultPageSize   = 10
	MaxPageSize       = 100
)

// CursorPaginationResponse holds the data page for cursor-based list endpoints.
// Pass it to [WriteCursorPaginated] or [CursorPaginated].
type CursorPaginationResponse struct {
	Data       any     `json:"data"`       // The actual data array
	NextCursor *string `json:"nextCursor"` // Cursor for the next page (null if no more pages)
	HasNext    bool    `json:"hasNext"`    // Whether there are more items available
}

// SimplePaginationResponse holds the data page for offset-based list endpoints.
// Pass it to [WriteSimplePaginated] or [SimplePaginated].
type SimplePaginationResponse struct {
	Data       any  `json:"data"`       // The actual data array
	PageNumber int  `json:"pageNumber"` // Current page number
	PageSize   int  `json:"pageSize"`   // Number of items per page
	HasNext    bool `json:"hasNext"`    // Whether there is a next page
	HasPrev    bool `json:"hasPrev"`    // Whether there is a previous page
}
