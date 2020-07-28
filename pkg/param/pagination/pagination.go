package pagination

// GetPage - pagination
type GetPage struct {
	StartRow  *int
	EndRow    *int
	PageSize  *int
	TotalRows *int
}

// ContainsMore - returns true if there are more pending objects then page size.
func (page *GetPage) ContainsMore() bool {

	if *page.EndRow < *page.TotalRows {
		return true
	}
	return false
}

// SetStartRow - set the starting row from list of objects
func (page *GetPage) SetStartRow(row int) {
	page.StartRow = &row
}

// SetEndRow - set the end row from the list of the return objects
func (page *GetPage) SetEndRow(row int) {
	page.EndRow = &row
}

// SetPageSize - set the page size for batch processing
func (page *GetPage) SetPageSize(row int) {
	page.PageSize = &row
}
