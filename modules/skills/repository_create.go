package skills

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type (
	// Creater of skills
	Creater interface {
		Create(opts ...CreateOption) (node Node, err error)
	}
	// CreateOptions of skills
	CreateOptions struct {
		ID   uint64
		Name string
	}
	// CreateOption of skills
	CreateOption func(*CreateOptions)
)

// NewCreateOptions constructor
func NewCreateOptions(opts ...CreateOption) *CreateOptions {
	o := &CreateOptions{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithID is id option
func WithID(id uint64) CreateOption {
	return func(opts *CreateOptions) {
		opts.ID = id
	}
}

// WithName is name option
func WithName(name string) CreateOption {
	return func(opts *CreateOptions) {
		opts.Name = name
	}
}

// NewCreater is constructor
func NewCreater(repository *Repository) Creater {
	return repository
}

// Create of skills
func (repo Repository) Create(options ...CreateOption) (node Node, err error) {
	opts := NewCreateOptions(options...)
	session := repo.Neo4jSessionGenerator()
	defer func() {
		err = session.Close()
	}()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "CREATE (:Node {id: $id, name: $name})"
		parameters := map[string]interface{}{
			"id":   opts.ID,
			"name": opts.Name,
		}
		return tx.Run(query, parameters)
	})
	if err != nil {
		return Node{}, err
	}
	return Node{ID: opts.ID, Name: opts.Name}, nil
}
