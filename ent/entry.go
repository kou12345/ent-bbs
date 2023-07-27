// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kou12345/ent-bbs/ent/entry"
)

// Entry is the model entity for the Entry schema.
type Entry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Entry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entry.FieldID:
			values[i] = new(sql.NullInt64)
		case entry.FieldContent:
			values[i] = new(sql.NullString)
		case entry.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Entry fields.
func (e *Entry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case entry.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				e.Content = value.String
			}
		case entry.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Entry.
// This includes values selected through modifiers, order, etc.
func (e *Entry) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Entry.
// Note that you need to call Entry.Unwrap() before calling this method if this Entry
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Entry) Update() *EntryUpdateOne {
	return NewEntryClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Entry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Entry) Unwrap() *Entry {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Entry is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Entry) String() string {
	var builder strings.Builder
	builder.WriteString("Entry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("content=")
	builder.WriteString(e.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Entries is a parsable slice of Entry.
type Entries []*Entry
