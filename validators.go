package utilities

import "math"

// ValidateID validate an id used to query
func ValidateID(id int64) error {
	if id <= 0 || id > math.MaxInt64 {
		return ErrorID
	}
	return nil
}

// ValidateLimitAndOffset validate offset and limit used for query
func ValidateLimitAndOffset(off, lim, defaultLimit, maxLimit int) (int, int) {

	if lim < defaultLimit {
		lim = defaultLimit
	} else if lim > maxLimit {
		lim = maxLimit
	}

	if off < 0 {
		off = 0
	}

	return off, lim
}
