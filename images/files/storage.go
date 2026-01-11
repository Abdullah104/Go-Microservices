package files

import "io"

// Storage defines the behavior for file operations
// Implementations may be of the time local disk, or cloud storage, et
type Storage interface {
	Save(path string, file io.Reader) error
}
