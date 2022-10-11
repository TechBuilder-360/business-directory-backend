package types

func (q *Query) CleanUp() {
	if q.Page <= 0 {
		q.Page = 1
	}

	// Set min page size to 10
	if q.PageSize <= 0 {
		q.PageSize = 10
	}

	// Set max page size to 20
	if q.PageSize > 20 {
		q.PageSize = 20
	}
}

func (q *Query) PaginationOffset() int {
	return (q.PageSize * q.Page) - q.PageSize
}
