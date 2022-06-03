// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/connectify-native/ent/header"
	"github.com/Faroukhamadi/connectify-native/ent/message"
	"github.com/Faroukhamadi/connectify-native/ent/user"
)

// HeaderCreate is the builder for creating a Header entity.
type HeaderCreate struct {
	config
	mutation *HeaderMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (hc *HeaderCreate) SetStatus(s string) *HeaderCreate {
	hc.mutation.SetStatus(s)
	return hc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hc *HeaderCreate) SetNillableStatus(s *string) *HeaderCreate {
	if s != nil {
		hc.SetStatus(*s)
	}
	return hc
}

// SetReceiverId sets the "receiverId" field.
func (hc *HeaderCreate) SetReceiverId(i int32) *HeaderCreate {
	hc.mutation.SetReceiverId(i)
	return hc
}

// SetNillableReceiverId sets the "receiverId" field if the given value is not nil.
func (hc *HeaderCreate) SetNillableReceiverId(i *int32) *HeaderCreate {
	if i != nil {
		hc.SetReceiverId(*i)
	}
	return hc
}

// SetSenderId sets the "senderId" field.
func (hc *HeaderCreate) SetSenderId(i int32) *HeaderCreate {
	hc.mutation.SetSenderId(i)
	return hc
}

// SetNillableSenderId sets the "senderId" field if the given value is not nil.
func (hc *HeaderCreate) SetNillableSenderId(i *int32) *HeaderCreate {
	if i != nil {
		hc.SetSenderId(*i)
	}
	return hc
}

// SetCreatedAt sets the "createdAt" field.
func (hc *HeaderCreate) SetCreatedAt(t time.Time) *HeaderCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetMessageId sets the "messageId" field.
func (hc *HeaderCreate) SetMessageId(i int32) *HeaderCreate {
	hc.mutation.SetMessageId(i)
	return hc
}

// SetNillableMessageId sets the "messageId" field if the given value is not nil.
func (hc *HeaderCreate) SetNillableMessageId(i *int32) *HeaderCreate {
	if i != nil {
		hc.SetMessageId(*i)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HeaderCreate) SetID(i int32) *HeaderCreate {
	hc.mutation.SetID(i)
	return hc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hc *HeaderCreate) SetUserID(id int32) *HeaderCreate {
	hc.mutation.SetUserID(id)
	return hc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hc *HeaderCreate) SetNillableUserID(id *int32) *HeaderCreate {
	if id != nil {
		hc = hc.SetUserID(*id)
	}
	return hc
}

// SetUser sets the "user" edge to the User entity.
func (hc *HeaderCreate) SetUser(u *User) *HeaderCreate {
	return hc.SetUserID(u.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (hc *HeaderCreate) SetMessageID(id int32) *HeaderCreate {
	hc.mutation.SetMessageID(id)
	return hc
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (hc *HeaderCreate) SetNillableMessageID(id *int32) *HeaderCreate {
	if id != nil {
		hc = hc.SetMessageID(*id)
	}
	return hc
}

// SetMessage sets the "message" edge to the Message entity.
func (hc *HeaderCreate) SetMessage(m *Message) *HeaderCreate {
	return hc.SetMessageID(m.ID)
}

// Mutation returns the HeaderMutation object of the builder.
func (hc *HeaderCreate) Mutation() *HeaderMutation {
	return hc.mutation
}

// Save creates the Header in the database.
func (hc *HeaderCreate) Save(ctx context.Context) (*Header, error) {
	var (
		err  error
		node *Header
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HeaderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Header)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HeaderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HeaderCreate) SaveX(ctx context.Context) *Header {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HeaderCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HeaderCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HeaderCreate) check() error {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Header.createdAt"`)}
	}
	return nil
}

func (hc *HeaderCreate) sqlSave(ctx context.Context) (*Header, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (hc *HeaderCreate) createSpec() (*Header, *sqlgraph.CreateSpec) {
	var (
		_node = &Header{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: header.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: header.FieldID,
			},
		}
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: header.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := hc.mutation.SenderId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldSenderId,
		})
		_node.SenderId = value
	}
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: header.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := hc.mutation.MessageId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldMessageId,
		})
		_node.MessageId = value
	}
	if nodes := hc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   header.UserTable,
			Columns: []string{header.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReceiverId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   header.MessageTable,
			Columns: []string{header.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HeaderCreateBulk is the builder for creating many Header entities in bulk.
type HeaderCreateBulk struct {
	config
	builders []*HeaderCreate
}

// Save creates the Header entities in the database.
func (hcb *HeaderCreateBulk) Save(ctx context.Context) ([]*Header, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Header, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HeaderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HeaderCreateBulk) SaveX(ctx context.Context) []*Header {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HeaderCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HeaderCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
