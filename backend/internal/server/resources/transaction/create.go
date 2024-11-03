// Code generated by greenstar scripts; DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/arikkfir/greenstar/backend/internal/auth"
	"github.com/arikkfir/greenstar/backend/internal/server/util"
	"github.com/shopspring/decimal"
)

type CreateRequest struct {
	TenantID        string          `json:"-"`
	Amount          decimal.Decimal `json:"amount,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	Date            time.Time       `json:"date,omitempty"`
	Description     *string         `json:"description,omitempty"`
	ReferenceID     string          `json:"referenceId,omitempty"`
	SourceAccountID string          `json:"sourceAccountId,omitempty"`
	TargetAccountID string          `json:"targetAccountId,omitempty"`
	properties      []string
}

func (lr *CreateRequest) HasDescription() bool { return slices.Contains(lr.properties, "description") }
func (lr *CreateRequest) UnmarshalJSON(data []byte) error {
	lr.properties = nil
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}
	for key := range tempMap {
		lr.properties = append(lr.properties, key)
	}
	type typeAlias CreateRequest
	alias := (*typeAlias)(lr)
	if err := json.Unmarshal(data, alias); err != nil {
		return err
	}
	return nil
}

type CreateResponse Transaction

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := util.Logger(ctx)
	if !auth.GetToken(ctx).IsPermittedForTenant(r.PathValue("tenantID"), "transactions:create") {
		util.ServeError(w, r, util.ErrForbidden)
		l.With("tenantID", r.PathValue("TenantPathVariableName")).WarnContext(ctx, "Access denied", "permission", "transactions:create")
		return
	}

	req := CreateRequest{}
	if err := util.UnmarshalBody(r, &req); err != nil {
		util.ServeError(w, r, err)
		return
	}
	req.TenantID = r.PathValue("tenantID")
	if req.TenantID == "" {
		util.ServeError(w, r, fmt.Errorf("%w: invalid tenant ID", util.ErrBadRequest))
		return
	}

	res, err := s.h.Create(ctx, req)
	if err != nil {
		if code := util.ServeError(w, r, err); code >= http.StatusInternalServerError {
			l.ErrorContext(ctx, "Failed creating transaction", "err", err)
		}
		return
	}

	if err := util.Marshal(w, r, http.StatusCreated, res); err != nil {
		l.ErrorContext(ctx, "Failed marshaling transaction", "err", err)
		util.ServeError(w, r, err)
	}
}