package base

import "math"

const (
	DEFAULT_PAGE_NO = 1
	MAX_PAGE_SIZE   = 500
)

func PageOffset(pageNo int, pageSize int) int64 {
	pageNo = handlePageNo(pageNo)
	pageSize = handlePageSize(pageSize)
	return int64((pageNo - 1) * pageSize)
}

func CreatePage(pageNo int, pageSize int, total int64, data interface{}) *Page {
	var page Page
	page.First = handleFirst(pageNo)
	page.PageNo = handlePageNo(pageNo)
	page.PageSize = handlePageSize(pageSize)
	page.Total = handleTotal(total)
	page.TotalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	page.Last = handleLast(pageNo, page.TotalPage)
	page.Data = data
	return &page
}

func handleFirst(pageNo int) bool {
	if pageNo == 1 {
		return true
	} else {
		return false
	}
}

func handleLast(pageNo int, totalPage int) bool {
	if pageNo == totalPage {
		return true
	} else {
		return false
	}
}

func handlePageNo(pageNo int) int {
	if pageNo <= 0 {
		return DEFAULT_PAGE_NO
	} else {
		return pageNo
	}
}

func handlePageSize(pageSize int) int {
	if pageSize <= 0 || pageSize > MAX_PAGE_SIZE {
		return MAX_PAGE_SIZE
	} else {
		return pageSize
	}
}

func handleTotal(total int64) int64 {
	if total <= 0 {
		return 0
	} else {
		return total
	}
}
