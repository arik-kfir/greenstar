package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/arikkfir/greenstar/api/gql"
	"github.com/arikkfir/greenstar/api/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Roots is the resolver for the roots field.
func (r *queryResolver) Roots(ctx context.Context) ([]*model.Account, error) {
	session := r.getNeo4jSession(ctx, neo4j.AccessModeRead)
	defer session.Close(ctx)

	v, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		//language=Cypher
		const getRootsCypher = `// Get root accounts
MATCH (root:Account) 
WHERE NOT exists ((root)-[:ChildOf]->(:Account)) 
RETURN root`

		result, err := tx.Run(ctx, getRootsCypher, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to collect records: %w", err)
		}

		accounts := make([]*model.Account, 0)
		for _, rec := range records {
			accounts = append(accounts, r.readAccount(rec.Values[0].(neo4j.Node)))
		}
		return accounts, nil
	})
	return v.([]*model.Account), err
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	session := r.getNeo4jSession(ctx, neo4j.AccessModeRead)
	defer session.Close(ctx)

	v, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		//language=Cypher
		const getAccountCypher = `// Get account by ID
MATCH (account:Account {accountID: $accountID})
RETURN account`

		result, err := tx.Run(ctx, getAccountCypher, map[string]any{"accountID": id})
		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to collect records: %w", err)
		}

		if len(records) == 0 {
			return nil, nil
		}

		if len(records) > 1 {
			return nil, fmt.Errorf("too many records")
		}

		return r.readAccount(records[0].Values[0].(neo4j.Node)), nil
	})
	return v.(*model.Account), err
}

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
