package utils

func Pagination(page uint32) (int, int) {
	var offset int
	const limit int = 20

	if page == 0 || page == 1 {
		offset = 0
	}

	if page > 1 {
		offset = limit * (int(page) - 1)
	}

	return offset, limit
}
