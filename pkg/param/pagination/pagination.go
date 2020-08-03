package pagination

// Page - This is used for batch processing. User may want to process
// limited objects instead of processing the whole object list . Specially when there are
// 100s of objects exists on Array e.g. volumes, snapshots etc.
type Page struct {
	StartRow  *int
	EndRow    *int
	PageSize  *int
	TotalRows *int
}

// ContainsMore - returns true if there are more pending objects then page size.
func (page *Page) ContainsMore() bool {
	return *page.EndRow < *page.TotalRows

}

// SetStartRow - set the starting row from list of objects
func (page *Page) SetStartRow(row int) {
	page.StartRow = &row
}

// SetEndRow - set the end row from the list of the return objects
func (page *Page) SetEndRow(row int) {
	page.EndRow = &row
}

// SetPageSize - set the page size for batch processing
func (page *Page) SetPageSize(row int) {
	page.PageSize = &row
}
