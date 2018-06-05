package bolt

import (
	"encoding/binary"
	"time"

	"github.com/Tinee/gophercises/todo/domain"

	"github.com/boltdb/bolt"
)

// Client represents a client to the underlying BoltDB data store.
type Client struct {
	Path        string
	db          *bolt.DB
	todoService TodoService
}

// NewClient configure everything and returns a pointer to the Client.
func NewClient(path string) *Client {
	c := &Client{Path: path}
	c.todoService.client = c
	return c
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return err
	}

	c.db = db

	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("Todos")); err != nil {
		return err
	}

	return tx.Commit()
}

// Close closes then underlying BoltDB database.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// TodoService returns the todo service associated with the client.
func (c *Client) TodoService() domain.TodoService {
	return &c.todoService
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
