// Code generated by greenstar scripts; DO NOT EDIT.

package tenant

import (
	"github.com/arikkfir/greenstar/backend/internal/server/util"
	"github.com/arikkfir/greenstar/backend/internal/util/observability"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = time.Time{}
)

type GetRequest struct {
	ID string `json:"id"`
}

func (lr *GetRequest) UnmarshalFromRequest(r *http.Request) error {
	lr.ID = r.PathValue("id")
	if lr.ID == "" {
		return util.ErrBadRequest
	}
	return nil
}

type GetResponse Tenant

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := observability.GetLogger(ctx)
	if err := util.VerifyPermissions(ctx, "tenants:get"); err != nil {
		util.ServeError(w, r, err)
		return
	}

	req := GetRequest{}
	if err := req.UnmarshalFromRequest(r); err != nil {
		util.ServeError(w, r, err)
		return
	}

	res, err := s.h.Get(ctx, req)
	if err != nil {
		if code := util.ServeError(w, r, err); code >= http.StatusInternalServerError {
			l.ErrorContext(ctx, "Failed getting tenant", "err", err)
		}
		return
	} else if res == nil {
		util.ServeError(w, r, util.ErrNotFound)
		return
	}

	if err := util.Marshal(w, r, http.StatusOK, res); err != nil {
		l.ErrorContext(ctx, "Failed marshaling tenant", "err", err)
		util.ServeError(w, r, err)
	}
}
