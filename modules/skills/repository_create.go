package skills

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type (
	// Creater of skills
	Creater interface {
		Create(opts ...CreateOption) (node Node, err error)
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
func (repo Repository) Create(options ...CreateOption) (node Node, err error) {
	opts := NewCreateOptions(options...)
	session := repo.Neo4jSessionGenerator()
	defer func() {
		if _err := session.Close(); _err != nil {
			err = _err
		}
	}()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "CREATE (skill: Skill{name: $name}) RETURN id(skill) as id, skill.name as name"
		parameters := map[string]interface{}{
			"name": opts.Name,
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
