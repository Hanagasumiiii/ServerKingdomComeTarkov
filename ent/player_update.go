// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/player"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/predicate"
	"github.com/Hanagasumiiii/ServerKingdomComeTarkov/ent/user"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetMana sets the "mana" field.
func (pu *PlayerUpdate) SetMana(i int) *PlayerUpdate {
	pu.mutation.ResetMana()
	pu.mutation.SetMana(i)
	return pu
}

// SetNillableMana sets the "mana" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableMana(i *int) *PlayerUpdate {
	if i != nil {
		pu.SetMana(*i)
	}
	return pu
}

// AddMana adds i to the "mana" field.
func (pu *PlayerUpdate) AddMana(i int) *PlayerUpdate {
	pu.mutation.AddMana(i)
	return pu
}

// SetHp sets the "hp" field.
func (pu *PlayerUpdate) SetHp(i int) *PlayerUpdate {
	pu.mutation.ResetHp()
	pu.mutation.SetHp(i)
	return pu
}

// SetNillableHp sets the "hp" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableHp(i *int) *PlayerUpdate {
	if i != nil {
		pu.SetHp(*i)
	}
	return pu
}

// AddHp adds i to the "hp" field.
func (pu *PlayerUpdate) AddHp(i int) *PlayerUpdate {
	pu.mutation.AddHp(i)
	return pu
}

// SetPositionX sets the "position_x" field.
func (pu *PlayerUpdate) SetPositionX(f float64) *PlayerUpdate {
	pu.mutation.ResetPositionX()
	pu.mutation.SetPositionX(f)
	return pu
}

// SetNillablePositionX sets the "position_x" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillablePositionX(f *float64) *PlayerUpdate {
	if f != nil {
		pu.SetPositionX(*f)
	}
	return pu
}

// AddPositionX adds f to the "position_x" field.
func (pu *PlayerUpdate) AddPositionX(f float64) *PlayerUpdate {
	pu.mutation.AddPositionX(f)
	return pu
}

// SetPositionY sets the "position_y" field.
func (pu *PlayerUpdate) SetPositionY(f float64) *PlayerUpdate {
	pu.mutation.ResetPositionY()
	pu.mutation.SetPositionY(f)
	return pu
}

// SetNillablePositionY sets the "position_y" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillablePositionY(f *float64) *PlayerUpdate {
	if f != nil {
		pu.SetPositionY(*f)
	}
	return pu
}

// AddPositionY adds f to the "position_y" field.
func (pu *PlayerUpdate) AddPositionY(f float64) *PlayerUpdate {
	pu.mutation.AddPositionY(f)
	return pu
}

// SetPositionZ sets the "position_z" field.
func (pu *PlayerUpdate) SetPositionZ(f float64) *PlayerUpdate {
	pu.mutation.ResetPositionZ()
	pu.mutation.SetPositionZ(f)
	return pu
}

// SetNillablePositionZ sets the "position_z" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillablePositionZ(f *float64) *PlayerUpdate {
	if f != nil {
		pu.SetPositionZ(*f)
	}
	return pu
}

// AddPositionZ adds f to the "position_z" field.
func (pu *PlayerUpdate) AddPositionZ(f float64) *PlayerUpdate {
	pu.mutation.AddPositionZ(f)
	return pu
}

// SetInventory sets the "inventory" field.
func (pu *PlayerUpdate) SetInventory(s []string) *PlayerUpdate {
	pu.mutation.SetInventory(s)
	return pu
}

// AppendInventory appends s to the "inventory" field.
func (pu *PlayerUpdate) AppendInventory(s []string) *PlayerUpdate {
	pu.mutation.AppendInventory(s)
	return pu
}

// ClearInventory clears the value of the "inventory" field.
func (pu *PlayerUpdate) ClearInventory() *PlayerUpdate {
	pu.mutation.ClearInventory()
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PlayerUpdate) SetUserID(id int) *PlayerUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PlayerUpdate) SetUser(u *User) *PlayerUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PlayerUpdate) ClearUser() *PlayerUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PlayerUpdate) check() error {
	if pu.mutation.UserCleared() && len(pu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Player.user"`)
	}
	return nil
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Mana(); ok {
		_spec.SetField(player.FieldMana, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedMana(); ok {
		_spec.AddField(player.FieldMana, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Hp(); ok {
		_spec.SetField(player.FieldHp, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedHp(); ok {
		_spec.AddField(player.FieldHp, field.TypeInt, value)
	}
	if value, ok := pu.mutation.PositionX(); ok {
		_spec.SetField(player.FieldPositionX, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPositionX(); ok {
		_spec.AddField(player.FieldPositionX, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.PositionY(); ok {
		_spec.SetField(player.FieldPositionY, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPositionY(); ok {
		_spec.AddField(player.FieldPositionY, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.PositionZ(); ok {
		_spec.SetField(player.FieldPositionZ, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPositionZ(); ok {
		_spec.AddField(player.FieldPositionZ, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Inventory(); ok {
		_spec.SetField(player.FieldInventory, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedInventory(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, player.FieldInventory, value)
		})
	}
	if pu.mutation.InventoryCleared() {
		_spec.ClearField(player.FieldInventory, field.TypeJSON)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   player.UserTable,
			Columns: []string{player.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   player.UserTable,
			Columns: []string{player.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// SetMana sets the "mana" field.
func (puo *PlayerUpdateOne) SetMana(i int) *PlayerUpdateOne {
	puo.mutation.ResetMana()
	puo.mutation.SetMana(i)
	return puo
}

// SetNillableMana sets the "mana" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableMana(i *int) *PlayerUpdateOne {
	if i != nil {
		puo.SetMana(*i)
	}
	return puo
}

// AddMana adds i to the "mana" field.
func (puo *PlayerUpdateOne) AddMana(i int) *PlayerUpdateOne {
	puo.mutation.AddMana(i)
	return puo
}

// SetHp sets the "hp" field.
func (puo *PlayerUpdateOne) SetHp(i int) *PlayerUpdateOne {
	puo.mutation.ResetHp()
	puo.mutation.SetHp(i)
	return puo
}

// SetNillableHp sets the "hp" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableHp(i *int) *PlayerUpdateOne {
	if i != nil {
		puo.SetHp(*i)
	}
	return puo
}

// AddHp adds i to the "hp" field.
func (puo *PlayerUpdateOne) AddHp(i int) *PlayerUpdateOne {
	puo.mutation.AddHp(i)
	return puo
}

// SetPositionX sets the "position_x" field.
func (puo *PlayerUpdateOne) SetPositionX(f float64) *PlayerUpdateOne {
	puo.mutation.ResetPositionX()
	puo.mutation.SetPositionX(f)
	return puo
}

// SetNillablePositionX sets the "position_x" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillablePositionX(f *float64) *PlayerUpdateOne {
	if f != nil {
		puo.SetPositionX(*f)
	}
	return puo
}

// AddPositionX adds f to the "position_x" field.
func (puo *PlayerUpdateOne) AddPositionX(f float64) *PlayerUpdateOne {
	puo.mutation.AddPositionX(f)
	return puo
}

// SetPositionY sets the "position_y" field.
func (puo *PlayerUpdateOne) SetPositionY(f float64) *PlayerUpdateOne {
	puo.mutation.ResetPositionY()
	puo.mutation.SetPositionY(f)
	return puo
}

// SetNillablePositionY sets the "position_y" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillablePositionY(f *float64) *PlayerUpdateOne {
	if f != nil {
		puo.SetPositionY(*f)
	}
	return puo
}

// AddPositionY adds f to the "position_y" field.
func (puo *PlayerUpdateOne) AddPositionY(f float64) *PlayerUpdateOne {
	puo.mutation.AddPositionY(f)
	return puo
}

// SetPositionZ sets the "position_z" field.
func (puo *PlayerUpdateOne) SetPositionZ(f float64) *PlayerUpdateOne {
	puo.mutation.ResetPositionZ()
	puo.mutation.SetPositionZ(f)
	return puo
}

// SetNillablePositionZ sets the "position_z" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillablePositionZ(f *float64) *PlayerUpdateOne {
	if f != nil {
		puo.SetPositionZ(*f)
	}
	return puo
}

// AddPositionZ adds f to the "position_z" field.
func (puo *PlayerUpdateOne) AddPositionZ(f float64) *PlayerUpdateOne {
	puo.mutation.AddPositionZ(f)
	return puo
}

// SetInventory sets the "inventory" field.
func (puo *PlayerUpdateOne) SetInventory(s []string) *PlayerUpdateOne {
	puo.mutation.SetInventory(s)
	return puo
}

// AppendInventory appends s to the "inventory" field.
func (puo *PlayerUpdateOne) AppendInventory(s []string) *PlayerUpdateOne {
	puo.mutation.AppendInventory(s)
	return puo
}

// ClearInventory clears the value of the "inventory" field.
func (puo *PlayerUpdateOne) ClearInventory() *PlayerUpdateOne {
	puo.mutation.ClearInventory()
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PlayerUpdateOne) SetUserID(id int) *PlayerUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PlayerUpdateOne) SetUser(u *User) *PlayerUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PlayerUpdateOne) ClearUser() *PlayerUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PlayerUpdate builder.
func (puo *PlayerUpdateOne) Where(ps ...predicate.Player) *PlayerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PlayerUpdateOne) check() error {
	if puo.mutation.UserCleared() && len(puo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Player.user"`)
	}
	return nil
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Player.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Mana(); ok {
		_spec.SetField(player.FieldMana, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedMana(); ok {
		_spec.AddField(player.FieldMana, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Hp(); ok {
		_spec.SetField(player.FieldHp, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedHp(); ok {
		_spec.AddField(player.FieldHp, field.TypeInt, value)
	}
	if value, ok := puo.mutation.PositionX(); ok {
		_spec.SetField(player.FieldPositionX, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPositionX(); ok {
		_spec.AddField(player.FieldPositionX, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.PositionY(); ok {
		_spec.SetField(player.FieldPositionY, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPositionY(); ok {
		_spec.AddField(player.FieldPositionY, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.PositionZ(); ok {
		_spec.SetField(player.FieldPositionZ, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPositionZ(); ok {
		_spec.AddField(player.FieldPositionZ, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Inventory(); ok {
		_spec.SetField(player.FieldInventory, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedInventory(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, player.FieldInventory, value)
		})
	}
	if puo.mutation.InventoryCleared() {
		_spec.ClearField(player.FieldInventory, field.TypeJSON)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   player.UserTable,
			Columns: []string{player.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   player.UserTable,
			Columns: []string{player.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}