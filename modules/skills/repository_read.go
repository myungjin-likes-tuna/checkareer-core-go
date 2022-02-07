package skills

import (
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type (
	// Reader of skills
	Reader interface {
		Read(options ...ReadOption) (node Node, err error)
	}
	// ReadOptions of skills
	ReadOptions struct {
		ID int64
	}
	// ReadOption of skills
	ReadOption func(*ReadOptions)
)

// NewReadOptions constructor
func NewReadOptions(opts ...ReadOption) *ReadOptions {
	o := &ReadOptions{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithID is id option
func WithID(id int64) ReadOption {
	return func(opts *ReadOptions) {
		opts.ID = id
	}
}

// NewReader is constructor
func NewReader(repository *Repository) Reader {
	return repository
}

// ErrNotFound is not found error
var ErrNotFound = errors.New("not found")

// Read of skills
func (repo Repository) Read(options ...ReadOption) (node Node, err error) {
	opts := NewReadOptions(options...)
	session := repo.Neo4jSessionGenerator()
	defer func() {
		err = session.Close()
	}()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (n:Skill) WHERE id(n) = $id RETURN id(n) as id, n.name as name"
		parameters := map[string]interface{}{
			"id": opts.ID,
		}
		result, err := tx.Run(query, parameters)
		if err != nil {
			return nil, err
		}
		record, err := result.Single()
		if err != nil {
			return nil, err
		}
		id, exist := record.Get("id")
		if !exist {
			return nil, ErrNotFound
		}
		name, exist := record.Get("name")
		if err != nil {
			return nil, ErrNotFound
		}
		return Node{ID: id.(int64), Name: name.(string)}, nil
	})
	if err != nil {
		return Node{}, err
	}

	return result.(Node), nil
}
