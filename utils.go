package utilities

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// GetURLID get the id from url
func GetURLID(r *http.Request, name string) (int64, error) {

	if name == "" {
		name = "id"
	}

	idStr, ok := mux.Vars(r)[name]

	if !ok {
		return 0, ErrorID
	}

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return 0, ErrorID
	}

	return id, nil
}

// GetParamsURLLimitAndOffset get limit and offset from url
func GetParamsURLLimitAndOffset(r *http.Request, d int, l, o string) (int, int) {

	if l == "" {
		l = "limit"
	}

	if o == "" {
		o = "offset"
	}

	limitString := r.URL.Query().Get(l)

	limit, err := strconv.Atoi(limitString)

	if err != nil {
		limit = d
	}

	offsetString := r.URL.Query().Get(o)

	offset, err := strconv.Atoi(offsetString)

	if err != nil {
		offset = 0
	}

	return offset, limit
}

// GetUserIDFromContext get current user id from context
func GetUserIDFromContext(ctx context.Context) (int64, error) {

	v := ctx.Value(u.IDCtx)

	id, ok := v.(int64)

	if !ok {
		return 0, ErrorUserID
	}
	return id, nil
}
