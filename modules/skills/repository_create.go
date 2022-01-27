package skills

type (
	// Creater of skills
	Creater interface {
	}
	// CreateOptions of skills
	CreateOptions struct {
	}
	// CreateOption of skills
	CreateOption func(*CreateOptions)
)

// NewCreater is constructor
func NewCreater(repository Repository) Creater {
	return repository
}

// Create of skills
func (repo Repository) Create() {

}
