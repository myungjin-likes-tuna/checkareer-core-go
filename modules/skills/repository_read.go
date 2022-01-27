package skills

type (
	// Reader of skills
	Reader interface {
	}
	// ReadOptions of skills
	ReadOptions struct {
	}
	// ReadOption of skills
	ReadOption func(*ReadOptions)
)

// NewReader is constructor
func NewReader(repository *Repository) Reader {
	return repository
}

func (repo Repository) Read() {

}
