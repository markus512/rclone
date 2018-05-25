// Package mockobject provides a mock object which can be created from a string
package mockobject

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/markus512/rclone/fs"
	"github.com/markus512/rclone/fs/hash"
)

var errNotImpl = errors.New("not implemented")

// Object is a mock fs.Object useful for testing
type Object string

// New returns mock fs.Object useful for testing
func New(name string) Object {
	return Object(name)
}

// String returns a description of the Object
func (o Object) String() string {
	return string(o)
}

// Fs returns read only access to the Fs that this object is part of
func (o Object) Fs() fs.Info {
	return nil
}

// Remote returns the remote path
func (o Object) Remote() string {
	return string(o)
}

// Hash returns the selected checksum of the file
// If no checksum is available it returns ""
func (o Object) Hash(hash.Type) (string, error) {
	return "", errNotImpl
}

// ModTime returns the modification date of the file
// It should return a best guess if one isn't available
func (o Object) ModTime() (t time.Time) {
	return t
}

// Size returns the size of the file
func (o Object) Size() int64 { return 0 }

// Storable says whether this object can be stored
func (o Object) Storable() bool {
	return true
}

// SetModTime sets the metadata on the object to set the modification date
func (o Object) SetModTime(time.Time) error {
	return errNotImpl
}

// Open opens the file for read.  Call Close() on the returned io.ReadCloser
func (o Object) Open(options ...fs.OpenOption) (io.ReadCloser, error) {
	return nil, errNotImpl
}

// Update in to the object with the modTime given of the given size
func (o Object) Update(in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) error {
	return errNotImpl
}

// Remove this object
func (o Object) Remove() error {
	return errNotImpl
}

// SeekMode specifies the optional Seek interface for the ReadCloser returned by Open
type SeekMode int

const (
	// SeekModeNone specifies no seek interface
	SeekModeNone SeekMode = iota
	// SeekModeRegular specifies the regular io.Seek interface
	SeekModeRegular
	// SeekModeRange specifies the fs.RangeSeek interface
	SeekModeRange
)

// SeekModes contains all valid SeekMode's
var SeekModes = []SeekMode{SeekModeNone, SeekModeRegular, SeekModeRange}

type contentMockObject struct {
	Object
	content  []byte
	seekMode SeekMode
}

// WithContent returns a fs.Object with the given content.
func (o Object) WithContent(content []byte, mode SeekMode) fs.Object {
	return &contentMockObject{
		Object:   o,
		content:  content,
		seekMode: mode,
	}
}

func (o *contentMockObject) Open(options ...fs.OpenOption) (io.ReadCloser, error) {
	var offset, limit int64 = 0, -1
	for _, option := range options {
		switch x := option.(type) {
		case *fs.SeekOption:
			offset = x.Offset
		case *fs.RangeOption:
			offset, limit = x.Decode(o.Size())
		default:
			if option.Mandatory() {
				return nil, fmt.Errorf("Unsupported mandatory option: %v", option)
			}
		}
	}
	if limit == -1 || offset+limit > o.Size() {
		limit = o.Size() - offset
	}

	var r *bytes.Reader
	if o.seekMode == SeekModeNone {
		r = bytes.NewReader(o.content[offset : offset+limit])
	} else {
		r = bytes.NewReader(o.content)
		_, err := r.Seek(offset, io.SeekStart)
		if err != nil {
			return nil, err
		}
	}
	switch o.seekMode {
	case SeekModeNone:
		return &readCloser{r}, nil
	case SeekModeRegular:
		return &readSeekCloser{r}, nil
	case SeekModeRange:
		return &readRangeSeekCloser{r}, nil
	default:
		return nil, errors.New(o.seekMode.String())
	}
}
func (o *contentMockObject) Size() int64 {
	return int64(len(o.content))
}

type readCloser struct{ io.Reader }

func (r *readCloser) Close() error { return nil }

type readSeekCloser struct{ io.ReadSeeker }

func (r *readSeekCloser) Close() error { return nil }

type readRangeSeekCloser struct{ io.ReadSeeker }

func (r *readRangeSeekCloser) RangeSeek(offset int64, whence int, length int64) (int64, error) {
	return r.ReadSeeker.Seek(offset, whence)
}

func (r *readRangeSeekCloser) Close() error { return nil }

func (m SeekMode) String() string {
	switch m {
	case SeekModeNone:
		return "SeekModeNone"
	case SeekModeRegular:
		return "SeekModeRegular"
	case SeekModeRange:
		return "SeekModeRange"
	default:
		return fmt.Sprintf("SeekModeInvalid(%d)", m)
	}
}
