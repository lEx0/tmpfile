package tmpfile

import (
	"io"
	"os"
	"sync"
)

const tmpFilePattern = "*.go.tmp"

// File is a wrapper around a temporary file
type File struct {
	file   *os.File
	isRead bool
	m      sync.Mutex
}

var _ io.ReadSeekCloser = &File{}

// NewFromReader creates a temporary file from a reader
func NewFromReader(reader io.Reader) (*File, error) {
	file, err := os.CreateTemp("", tmpFilePattern)

	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(file, reader); err != nil {
		return nil, err
	}

	// rewind the pointer to the beginning of the file
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return &File{file: file}, nil
}

// New creates an empty temporary file
func New() (*File, error) {
	file, err := os.CreateTemp("", tmpFilePattern)

	if err != nil {
		return nil, err
	}

	return &File{file: file}, nil
}

// Read reads from the temporary file
func (f *File) Read(p []byte) (n int, err error) {
	f.m.Lock()
	defer f.m.Unlock()

	// if the file has not been read yet, move the pointer to the beginning
	if !f.isRead {
		if _, err = f.file.Seek(0, io.SeekStart); err != nil {
			return 0, err
		}

		f.isRead = true
	}

	return f.file.Read(p)
}

// Write writes to the temporary file
func (f *File) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}

// Seek sets the offset for the next read or write
func (f *File) Seek(offset int64, whence int) (int64, error) {
	f.m.Lock()
	defer f.m.Unlock()

	f.isRead = true

	return f.file.Seek(offset, whence)
}

// Close closes the temporary file descriptor and deletes the file
func (f *File) Close() error {
	filename := f.file.Name()

	if err := f.file.Close(); err != nil {
		return err
	}

	return os.Remove(filename)
}
