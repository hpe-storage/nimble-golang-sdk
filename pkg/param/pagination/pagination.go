package pagination

type GetPage struct {
	StartRow  *int
	EndRow    *int
	PageSize  *int
	TotalRows *int
}

func (page *GetPage) ContainsMore() bool {

	if *page.EndRow < *page.TotalRows {
		return true
	}
	return false
}
func (page *GetPage) SetStartRow(row int) {
	page.StartRow = &row
}

func (page *GetPage) SetEndRow(row int) {
	page.EndRow = &row
}
func (page *GetPage) SetTotalRows(row int) {
	page.TotalRows = &row
}

func (page *GetPage) SetPageSize(row int) {
	page.PageSize = &row
}
