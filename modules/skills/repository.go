package skills

import "checkareer-core/dbms"

type (
	// Repository of skills
	Repository struct {
		dbms.Neo4jSessionGenerator
	}
)

// NewRepository is constructor
func NewRepository(generator dbms.Neo4jSessionGenerator) *Repository {
	return &Repository{generator}
}
