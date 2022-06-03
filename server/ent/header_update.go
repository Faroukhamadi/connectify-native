// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/connectify-native/ent/header"
	"github.com/Faroukhamadi/connectify-native/ent/message"
	"github.com/Faroukhamadi/connectify-native/ent/predicate"
	"github.com/Faroukhamadi/connectify-native/ent/user"
)

// HeaderUpdate is the builder for updating Header entities.
type HeaderUpdate struct {
	config
	hooks    []Hook
	mutation *HeaderMutation
}

// Where appends a list predicates to the HeaderUpdate builder.
func (hu *HeaderUpdate) Where(ps ...predicate.Header) *HeaderUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetStatus sets the "status" field.
func (hu *HeaderUpdate) SetStatus(s string) *HeaderUpdate {
	hu.mutation.SetStatus(s)
	return hu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hu *HeaderUpdate) SetNillableStatus(s *string) *HeaderUpdate {
	if s != nil {
		hu.SetStatus(*s)
	}
	return hu
}

// ClearStatus clears the value of the "status" field.
func (hu *HeaderUpdate) ClearStatus() *HeaderUpdate {
	hu.mutation.ClearStatus()
	return hu
}

// SetReceiverId sets the "receiverId" field.
func (hu *HeaderUpdate) SetReceiverId(i int32) *HeaderUpdate {
	hu.mutation.SetReceiverId(i)
	return hu
}

// SetNillableReceiverId sets the "receiverId" field if the given value is not nil.
func (hu *HeaderUpdate) SetNillableReceiverId(i *int32) *HeaderUpdate {
	if i != nil {
		hu.SetReceiverId(*i)
	}
	return hu
}

// ClearReceiverId clears the value of the "receiverId" field.
func (hu *HeaderUpdate) ClearReceiverId() *HeaderUpdate {
	hu.mutation.ClearReceiverId()
	return hu
}

// SetSenderId sets the "senderId" field.
func (hu *HeaderUpdate) SetSenderId(i int32) *HeaderUpdate {
	hu.mutation.ResetSenderId()
	hu.mutation.SetSenderId(i)
	return hu
}

// SetNillableSenderId sets the "senderId" field if the given value is not nil.
func (hu *HeaderUpdate) SetNillableSenderId(i *int32) *HeaderUpdate {
	if i != nil {
		hu.SetSenderId(*i)
	}
	return hu
}

// AddSenderId adds i to the "senderId" field.
func (hu *HeaderUpdate) AddSenderId(i int32) *HeaderUpdate {
	hu.mutation.AddSenderId(i)
	return hu
}

// ClearSenderId clears the value of the "senderId" field.
func (hu *HeaderUpdate) ClearSenderId() *HeaderUpdate {
	hu.mutation.ClearSenderId()
	return hu
}

// SetCreatedAt sets the "createdAt" field.
func (hu *HeaderUpdate) SetCreatedAt(t time.Time) *HeaderUpdate {
	hu.mutation.SetCreatedAt(t)
	return hu
}

// SetMessageId sets the "messageId" field.
func (hu *HeaderUpdate) SetMessageId(i int32) *HeaderUpdate {
	hu.mutation.ResetMessageId()
	hu.mutation.SetMessageId(i)
	return hu
}

// SetNillableMessageId sets the "messageId" field if the given value is not nil.
func (hu *HeaderUpdate) SetNillableMessageId(i *int32) *HeaderUpdate {
	if i != nil {
		hu.SetMessageId(*i)
	}
	return hu
}

// AddMessageId adds i to the "messageId" field.
func (hu *HeaderUpdate) AddMessageId(i int32) *HeaderUpdate {
	hu.mutation.AddMessageId(i)
	return hu
}

// ClearMessageId clears the value of the "messageId" field.
func (hu *HeaderUpdate) ClearMessageId() *HeaderUpdate {
	hu.mutation.ClearMessageId()
	return hu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hu *HeaderUpdate) SetUserID(id int32) *HeaderUpdate {
	hu.mutation.SetUserID(id)
	return hu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hu *HeaderUpdate) SetNillableUserID(id *int32) *HeaderUpdate {
	if id != nil {
		hu = hu.SetUserID(*id)
	}
	return hu
}

// SetUser sets the "user" edge to the User entity.
func (hu *HeaderUpdate) SetUser(u *User) *HeaderUpdate {
	return hu.SetUserID(u.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (hu *HeaderUpdate) SetMessageID(id int32) *HeaderUpdate {
	hu.mutation.SetMessageID(id)
	return hu
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (hu *HeaderUpdate) SetNillableMessageID(id *int32) *HeaderUpdate {
	if id != nil {
		hu = hu.SetMessageID(*id)
	}
	return hu
}

// SetMessage sets the "message" edge to the Message entity.
func (hu *HeaderUpdate) SetMessage(m *Message) *HeaderUpdate {
	return hu.SetMessageID(m.ID)
}

// Mutation returns the HeaderMutation object of the builder.
func (hu *HeaderUpdate) Mutation() *HeaderMutation {
	return hu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (hu *HeaderUpdate) ClearUser() *HeaderUpdate {
	hu.mutation.ClearUser()
	return hu
}

// ClearMessage clears the "message" edge to the Message entity.
func (hu *HeaderUpdate) ClearMessage() *HeaderUpdate {
	hu.mutation.ClearMessage()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HeaderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HeaderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			if hu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HeaderUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HeaderUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HeaderUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HeaderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   header.Table,
			Columns: header.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: header.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: header.FieldStatus,
		})
	}
	if hu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: header.FieldStatus,
		})
	}
	if value, ok := hu.mutation.SenderId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldSenderId,
		})
	}
	if value, ok := hu.mutation.AddedSenderId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldSenderId,
		})
	}
	if hu.mutation.SenderIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: header.FieldSenderId,
		})
	}
	if value, ok := hu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: header.FieldCreatedAt,
		})
	}
	if value, ok := hu.mutation.MessageId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldMessageId,
		})
	}
	if value, ok := hu.mutation.AddedMessageId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldMessageId,
		})
	}
	if hu.mutation.MessageIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: header.FieldMessageId,
		})
	}
	if hu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.MessageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.MessageIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{header.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// HeaderUpdateOne is the builder for updating a single Header entity.
type HeaderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HeaderMutation
}

// SetStatus sets the "status" field.
func (huo *HeaderUpdateOne) SetStatus(s string) *HeaderUpdateOne {
	huo.mutation.SetStatus(s)
	return huo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableStatus(s *string) *HeaderUpdateOne {
	if s != nil {
		huo.SetStatus(*s)
	}
	return huo
}

// ClearStatus clears the value of the "status" field.
func (huo *HeaderUpdateOne) ClearStatus() *HeaderUpdateOne {
	huo.mutation.ClearStatus()
	return huo
}

// SetReceiverId sets the "receiverId" field.
func (huo *HeaderUpdateOne) SetReceiverId(i int32) *HeaderUpdateOne {
	huo.mutation.SetReceiverId(i)
	return huo
}

// SetNillableReceiverId sets the "receiverId" field if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableReceiverId(i *int32) *HeaderUpdateOne {
	if i != nil {
		huo.SetReceiverId(*i)
	}
	return huo
}

// ClearReceiverId clears the value of the "receiverId" field.
func (huo *HeaderUpdateOne) ClearReceiverId() *HeaderUpdateOne {
	huo.mutation.ClearReceiverId()
	return huo
}

// SetSenderId sets the "senderId" field.
func (huo *HeaderUpdateOne) SetSenderId(i int32) *HeaderUpdateOne {
	huo.mutation.ResetSenderId()
	huo.mutation.SetSenderId(i)
	return huo
}

// SetNillableSenderId sets the "senderId" field if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableSenderId(i *int32) *HeaderUpdateOne {
	if i != nil {
		huo.SetSenderId(*i)
	}
	return huo
}

// AddSenderId adds i to the "senderId" field.
func (huo *HeaderUpdateOne) AddSenderId(i int32) *HeaderUpdateOne {
	huo.mutation.AddSenderId(i)
	return huo
}

// ClearSenderId clears the value of the "senderId" field.
func (huo *HeaderUpdateOne) ClearSenderId() *HeaderUpdateOne {
	huo.mutation.ClearSenderId()
	return huo
}

// SetCreatedAt sets the "createdAt" field.
func (huo *HeaderUpdateOne) SetCreatedAt(t time.Time) *HeaderUpdateOne {
	huo.mutation.SetCreatedAt(t)
	return huo
}

// SetMessageId sets the "messageId" field.
func (huo *HeaderUpdateOne) SetMessageId(i int32) *HeaderUpdateOne {
	huo.mutation.ResetMessageId()
	huo.mutation.SetMessageId(i)
	return huo
}

// SetNillableMessageId sets the "messageId" field if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableMessageId(i *int32) *HeaderUpdateOne {
	if i != nil {
		huo.SetMessageId(*i)
	}
	return huo
}

// AddMessageId adds i to the "messageId" field.
func (huo *HeaderUpdateOne) AddMessageId(i int32) *HeaderUpdateOne {
	huo.mutation.AddMessageId(i)
	return huo
}

// ClearMessageId clears the value of the "messageId" field.
func (huo *HeaderUpdateOne) ClearMessageId() *HeaderUpdateOne {
	huo.mutation.ClearMessageId()
	return huo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (huo *HeaderUpdateOne) SetUserID(id int32) *HeaderUpdateOne {
	huo.mutation.SetUserID(id)
	return huo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableUserID(id *int32) *HeaderUpdateOne {
	if id != nil {
		huo = huo.SetUserID(*id)
	}
	return huo
}

// SetUser sets the "user" edge to the User entity.
func (huo *HeaderUpdateOne) SetUser(u *User) *HeaderUpdateOne {
	return huo.SetUserID(u.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (huo *HeaderUpdateOne) SetMessageID(id int32) *HeaderUpdateOne {
	huo.mutation.SetMessageID(id)
	return huo
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (huo *HeaderUpdateOne) SetNillableMessageID(id *int32) *HeaderUpdateOne {
	if id != nil {
		huo = huo.SetMessageID(*id)
	}
	return huo
}

// SetMessage sets the "message" edge to the Message entity.
func (huo *HeaderUpdateOne) SetMessage(m *Message) *HeaderUpdateOne {
	return huo.SetMessageID(m.ID)
}

// Mutation returns the HeaderMutation object of the builder.
func (huo *HeaderUpdateOne) Mutation() *HeaderMutation {
	return huo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (huo *HeaderUpdateOne) ClearUser() *HeaderUpdateOne {
	huo.mutation.ClearUser()
	return huo
}

// ClearMessage clears the "message" edge to the Message entity.
func (huo *HeaderUpdateOne) ClearMessage() *HeaderUpdateOne {
	huo.mutation.ClearMessage()
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HeaderUpdateOne) Select(field string, fields ...string) *HeaderUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Header entity.
func (huo *HeaderUpdateOne) Save(ctx context.Context) (*Header, error) {
	var (
		err  error
		node *Header
	)
	if len(huo.hooks) == 0 {
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HeaderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			if huo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = huo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, huo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (huo *HeaderUpdateOne) SaveX(ctx context.Context) *Header {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HeaderUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HeaderUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HeaderUpdateOne) sqlSave(ctx context.Context) (_node *Header, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   header.Table,
			Columns: header.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: header.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Header.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, header.FieldID)
		for _, f := range fields {
			if !header.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != header.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: header.FieldStatus,
		})
	}
	if huo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: header.FieldStatus,
		})
	}
	if value, ok := huo.mutation.SenderId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldSenderId,
		})
	}
	if value, ok := huo.mutation.AddedSenderId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldSenderId,
		})
	}
	if huo.mutation.SenderIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: header.FieldSenderId,
		})
	}
	if value, ok := huo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: header.FieldCreatedAt,
		})
	}
	if value, ok := huo.mutation.MessageId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldMessageId,
		})
	}
	if value, ok := huo.mutation.AddedMessageId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: header.FieldMessageId,
		})
	}
	if huo.mutation.MessageIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: header.FieldMessageId,
		})
	}
	if huo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.MessageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.MessageIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Header{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{header.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
