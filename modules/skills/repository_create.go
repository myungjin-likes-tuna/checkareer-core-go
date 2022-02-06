package skills

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type (
	// Creater of skills
	Creater interface {
		Create(id uint64, opts ...CreateOption) (node Node, err error)
	}
	// CreateOptions of skills
	CreateOptions struct {
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
func (repo Repository) Create(id uint64, options ...CreateOption) (node Node, err error) {
	opts := NewCreateOptions(options...)
	session := repo.Neo4jSessionGenerator()
	defer func() {
		err = session.Close()
	}()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "CREATE (:Node {id: $id, name: $name})"
		parameters := map[string]interface{}{
			"id":   id,
			"name": opts.Name,
		}
		return tx.Run(query, parameters)
	})
	if err != nil {
		return Node{}, err
	}
	return Node{ID: id, Name: opts.Name}, nil
}
