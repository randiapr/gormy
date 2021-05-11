package gormy

import (
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

// TotalPages ..
func TotalPages(count, limit int64) int64 {
	return int64(math.Ceil(float64(count) / float64(limit)))
}

// NextPage ..
func NextPage(page, totalPages int64) int64 {
	if page == totalPages {
		return page
	}
	return page + 1
}

// PrevPage ..
func PrevPage(page int64) int64 {
	if page > 1 {
		return page - 1
	}
	return page
}

// SortedBy ..
func SortedBy(sort []string) []string {
	var sorted []string
	for _, v := range sort {
		split := strings.Split(v, ",")
		sorted = append(sorted, fmt.Sprintf("%s %s", split[0], split[1]))
	}
	return sorted
}

// Offset ..
func Offset(page, limit int64) int64 {
	var offset int64
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return offset
}

func Paginate(page, size int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(int(offset)).Limit(int(size))
	}
}

// GenPagination generate pagination meta data
func GenPagination(data interface{}, page, size, count int64) *Pagination {
	totalPages := TotalPages(count, size)
	offset := Offset(page, size)
	return &Pagination{
		TotalRecords: count,
		TotalPages:   totalPages,
		Data:         data,
		Offset:       offset,
		Limit:        size,
		Page:         page,
		PrevPage:     PrevPage(page),
		NextPage:     NextPage(page, totalPages),
	}
}
