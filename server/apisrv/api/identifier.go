package api

import (
	"encoding/hex"
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

var identifierPattern = regexp.MustCompile("^[0-9a-f]{32}$")

// Identifier represents a 16 byte universally unique identifier
type Identifier struct {
	bytes [16]byte
}

// NewIdentifier creates a new universally unique identifier
func NewIdentifier() (id Identifier) {
	copy(id.bytes[:], uuid.NewV4().Bytes())
	return id
}

// String returns the textual representation of the identifier
func (id *Identifier) String() string {
	return hex.EncodeToString(id.bytes[:])
}

// FromString parses the identifier from a hex encoded 32 char string
func (id *Identifier) FromString(str string) (err error) {
	if !identifierPattern.MatchString(str) {
		return errors.Errorf(
			"invalid identifier string representation: '%s'",
			str,
		)
	}
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return errors.New("couldn't decode hex string to bytes")
	}
	copy(id.bytes[:], bytes)
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (id Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(id.bytes[:]))
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (id *Identifier) UnmarshalJSON(bytes []byte) (err error) {
	if len(bytes) != 34 {
		return errors.Errorf(
			"invalid identifier (wrong length: %d)",
			len(bytes),
		)
	}
	return id.FromString(string(bytes[1:33]))
}
