// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/arik-kfir/greenstar/backend/admin/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	CreateTenant(ctx context.Context, tenantID *string, tenant model.TenantChanges) (*model.Tenant, error)
	UpdateTenant(ctx context.Context, tenantID string, tenant model.TenantChanges) (*model.Tenant, error)
	DeleteTenant(ctx context.Context, tenantID string) (string, error)
}
type QueryResolver interface {
	Tenants(ctx context.Context) ([]*model.Tenant, error)
	Tenant(ctx context.Context, id string) (*model.Tenant, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_createTenant_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *string
	if tmp, ok := rawArgs["tenantID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tenantID"))
		arg0, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tenantID"] = arg0
	var arg1 model.TenantChanges
	if tmp, ok := rawArgs["tenant"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tenant"))
		arg1, err = ec.unmarshalNTenantChanges2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenantChanges(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tenant"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteTenant_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["tenantID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tenantID"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tenantID"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateTenant_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["tenantID"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tenantID"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tenantID"] = arg0
	var arg1 model.TenantChanges
	if tmp, ok := rawArgs["tenant"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tenant"))
		arg1, err = ec.unmarshalNTenantChanges2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenantChanges(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["tenant"] = arg1
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_tenant_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createTenant(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createTenant(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateTenant(rctx, fc.Args["tenantID"].(*string), fc.Args["tenant"].(model.TenantChanges))
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Tenant)
	fc.Result = res
	return ec.marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createTenant(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tenant_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Tenant_displayName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tenant", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createTenant_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateTenant(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateTenant(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateTenant(rctx, fc.Args["tenantID"].(string), fc.Args["tenant"].(model.TenantChanges))
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Tenant)
	fc.Result = res
	return ec.marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateTenant(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tenant_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Tenant_displayName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tenant", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateTenant_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_deleteTenant(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_deleteTenant(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteTenant(rctx, fc.Args["tenantID"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_deleteTenant(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_deleteTenant_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query_tenants(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_tenants(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Tenants(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Tenant)
	fc.Result = res
	return ec.marshalNTenant2ᚕᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenantᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_tenants(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tenant_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Tenant_displayName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tenant", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_tenant(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_tenant(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Tenant(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Tenant)
	fc.Result = res
	return ec.marshalOTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_tenant(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Tenant_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Tenant_displayName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Tenant", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_tenant_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_id(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_displayName(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_displayName(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DisplayName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_displayName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputTenantChanges(ctx context.Context, obj interface{}) (model.TenantChanges, error) {
	var it model.TenantChanges
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"displayName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "displayName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("displayName"))
			it.DisplayName, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createTenant":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createTenant(ctx, field)
			})

		case "updateTenant":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateTenant(ctx, field)
			})

		case "deleteTenant":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_deleteTenant(ctx, field)
			})

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "tenants":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_tenants(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "tenant":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_tenant(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "__type":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})

		case "__schema":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	return out
}

var tenantImplementors = []string{"Tenant"}

func (ec *executionContext) _Tenant(ctx context.Context, sel ast.SelectionSet, obj *model.Tenant) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, tenantImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Tenant")
		case "id":

			out.Values[i] = ec._Tenant_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "displayName":

			out.Values[i] = ec._Tenant_displayName(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNTenant2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v model.Tenant) graphql.Marshaler {
	return ec._Tenant(ctx, sel, &v)
}

func (ec *executionContext) marshalNTenant2ᚕᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenantᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Tenant) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v *model.Tenant) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Tenant(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTenantChanges2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenantChanges(ctx context.Context, v interface{}) (model.TenantChanges, error) {
	res, err := ec.unmarshalInputTenantChanges(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋadminᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v *model.Tenant) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Tenant(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
