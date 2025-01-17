package component

import (
	"context"

	"github.com/pkg/errors"
	"github.com/projectdiscovery/nuclei/v3/pkg/fuzz/dataformat"
	"github.com/projectdiscovery/retryablehttp-go"
)

// Path is a component for a request Path
type Path struct {
	value *Value

	req *retryablehttp.Request
}

var _ Component = &Path{}

// NewPath creates a new URL component
func NewPath() *Path {
	return &Path{}
}

// Name returns the name of the component
func (q *Path) Name() string {
	return RequestPathComponent
}

// Parse parses the component and returns the
// parsed component
func (q *Path) Parse(req *retryablehttp.Request) (bool, error) {
	q.req = req
	q.value = NewValue(req.URL.Path)

	parsed, err := dataformat.Get(dataformat.RawDataFormat).Decode(q.value.String())
	if err != nil {
		return false, err
	}
	q.value.SetParsed(parsed, dataformat.RawDataFormat)
	return true, nil
}

// Iterate iterates through the component
func (q *Path) Iterate(callback func(key string, value interface{}) error) error {
	for key, value := range q.value.Parsed() {
		if err := callback(key, value); err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets a value in the component
// for a key
func (q *Path) SetValue(key string, value string) error {
	if !q.value.SetParsedValue(key, value) {
		return ErrSetValue
	}
	return nil
}

// Delete deletes a key from the component
func (q *Path) Delete(key string) error {
	if !q.value.Delete(key) {
		return ErrKeyNotFound
	}
	return nil
}

// Rebuild returns a new request with the
// component rebuilt
func (q *Path) Rebuild() (*retryablehttp.Request, error) {
	encoded, err := q.value.Encode()
	if err != nil {
		return nil, errors.Wrap(err, "could not encode query")
	}
	cloned := q.req.Clone(context.Background())
	if err := cloned.UpdateRelPath(encoded, true); err != nil {
		cloned.URL.RawPath = encoded
	}
	return cloned, nil
}
