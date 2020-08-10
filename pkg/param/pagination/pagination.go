package pagination

// Page controls enable iterative retrieval of a large data set. This
// structure is used both on input, to specify the boundaries of a
// page to retrieve, and on output, to indicate the items actually
// retrieved. The TotalRows value is ignored on input
type Page struct {
	StartRow  *int
	EndRow    *int
	TotalRows *int
}

// SetStartRow - Specify the first row to retrieve in a page of results, using a zero-based index.
func (page *Page) SetStartRow(row int) {
	page.StartRow = &row
}

// SetEndRow - Specify the first row NOT to retrieve in a page of results, using a zero-based index
// value one larger than the last row desired.
func (page *Page) SetEndRow(row int) {
	page.EndRow = &row
}

// SetPageSize - set the page size for batch processing
func (page *Page) SetPageSize(psize int) {
	zero := 0
	page.StartRow = &zero
	page.EndRow = &psize
}

// NextPage - Returns true if objects are more than pagesize.
func (page *Page) NextPage() bool {
	if page.StartRow == nil || page.EndRow == nil || page.TotalRows == nil ||
		*page.EndRow >= *page.TotalRows || *page.StartRow >= *page.EndRow {
		return false
	}
	pageSize := *page.EndRow - *page.StartRow
	page.StartRow = page.EndRow
	// overflow check
	if (*page.StartRow + pageSize) > *page.TotalRows {
		page.SetEndRow(*page.TotalRows)
	} else {
		page.SetEndRow(*page.StartRow + pageSize)
	}
	return true
}
