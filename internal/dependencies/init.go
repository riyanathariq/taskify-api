package dependencies

import (
	"sync"
)

var (
	once sync.Once
	deps Dependency
)

func LoadDependencies(options ...Options) {
	once.Do(func() {
		for _, opt := range options {
			opt(&deps)
		}
	})
}

// New mengembalikan dependency yang sudah di-load
func New() Dependency {
	return deps
}
