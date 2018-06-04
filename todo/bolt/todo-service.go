package bolt

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/Tinee/gophercises/todo/domain"
)

// TodoService is the struct that contains an underlaying client for us to use when we
// want to interact with Todos.
type TodoService struct {
	client *Client
}

// Create attempts to add a new Todo to bolt collection.
func (ts TodoService) Create(t domain.Todo) error {
	tx, err := ts.client.db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	b := tx.Bucket([]byte("Todos"))
	seq, err := b.NextSequence()
	if err != nil {
		return err
	}
	t.ID = int(seq)

	if v, err := json.Marshal(t); err != nil {
		return err
	} else if err := b.Put(itob(t.ID), v); err != nil {
		return err
	}

	return tx.Commit()
}

// All gets all todos in the bucket
func (ts TodoService) All() ([]domain.Todo, error) {
	var todos []domain.Todo

	tx, err := ts.client.db.Begin(true)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	b := tx.Bucket([]byte("Todos"))
	c := b.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		todo := domain.Todo{}
		r := bytes.NewReader(v)
		json.NewDecoder(r).Decode(&todo)

		todos = append(todos, todo)
	}

	// if v, err := json.Marshal(t); err != nil {
	// 	return err
	// } else if err := b.Put([]byte(t.ID), v); err != nil {
	// 	return err
	// }

	return todos, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
