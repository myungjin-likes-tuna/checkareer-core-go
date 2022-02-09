package skills

import (
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type (
	// Reader of skills
	Reader interface {
		Read(options ...ReadOption) (nodes []Node, err error)
	}
	// ReadOptions of skills
	ReadOptions struct {
		ID    int64
		Limit uint
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

// WithLimit is limit option
func WithLimit(limit uint) ReadOption {
	return func(opts *ReadOptions) {
		opts.Limit = limit
	}
}

// NewReader is constructor
func NewReader(repository *Repository) Reader {
	return repository
}

// ErrNotFound is not found error
var ErrNotFound = errors.New("not found")

// Read of skills
func (repo Repository) Read(options ...ReadOption) (nodes []Node, err error) {
	opts := NewReadOptions(options...)
	session := repo.Neo4jSessionGenerator()
	defer func() {
		err = session.Close()
	}()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		parameters := map[string]interface{}{}
		query := "MATCH (n:Skill) "
		if opts.ID != 0 {
			query += "WHERE id(n) = $id "
			parameters["id"] = opts.ID
		}
		query += "RETURN id(n) as id, n.name as name "
		if opts.Limit != 0 {
			query += "LIMIT $limit"
			parameters["limit"] = opts.Limit
		}
		result, err := tx.Run(query, parameters)
		if err != nil {
			return nil, err
		}
		records, err := result.Collect()
		if err != nil {
			return nil, err
		}
		nodes = make([]Node, len(records))
		for i, record := range records {
			id, exist := record.Get("id")
			if !exist {
				return nil, ErrNotFound
			}
			name, exist := record.Get("name")
			if err != nil {
				return nil, ErrNotFound
			}
			nodes[i] = Node{ID: id.(int64), Name: name.(string)}
		}
		return nodes, nil
	})
	if err != nil {
		return nodes, err
	}

	return result.([]Node), nil
}
