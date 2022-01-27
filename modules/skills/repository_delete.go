package skills

type (
	// Deleter of skills
	Deleter interface {
	}
	// DeleteOptions of skills
	DeleteOptions struct {
	}
	// DeleteOption of skills
	DeleteOption func(*DeleteOptions)
)

// NewDeleter is constructor
func NewDeleter(repository Repository) Deleter {
	return repository
}

// Delete of skills
func (repo Repository) Delete() {

}
