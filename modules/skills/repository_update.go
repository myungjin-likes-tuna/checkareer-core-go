package skills

type (
	// Updater of skills
	Updater interface {
	}
	// UpdateOptions of skills
	UpdateOptions struct {
	}
	// UpdateOption of skills
	UpdateOption func(*UpdateOptions)
)

// NewUpdater is constructor
func NewUpdater(repository *Repository) Updater {
	return repository
}

// Update of skills
func (repo Repository) Update() {

}
