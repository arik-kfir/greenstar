// Code generated by greenstar scripts; DO NOT EDIT.

package transaction

import (
	"github.com/arikkfir/greenstar/backend/internal/server/util"
	"github.com/arikkfir/greenstar/backend/internal/util/observability"
	"net/http"
)

type DeleteRequest struct {
	ID string `json:"id"`
}

func (lr *DeleteRequest) UnmarshalFromRequest(r *http.Request) error {
	lr.ID = r.PathValue("id")
	if lr.ID == "" {
		return util.ErrBadRequest
	}
	return nil
}

type DeleteResponse struct{}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := observability.GetLogger(ctx)
	if err := util.VerifyPermissions(ctx, "transactions:update"); err != nil {
		util.ServeError(w, r, err)
		return
	}

	req := DeleteRequest{}
	if err := req.UnmarshalFromRequest(r); err != nil {
		util.ServeError(w, r, err)
		return
	}

	err := s.h.Delete(ctx, req)
	if err != nil {
		if code := util.ServeError(w, r, err); code >= http.StatusInternalServerError {
			l.ErrorContext(ctx, "Failed deleting transaction", "err", err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
