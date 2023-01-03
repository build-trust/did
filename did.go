// Package did is a set of tools to work with Decentralized Identifiers (DIDs) as described
// in the DID spec https://w3c.github.io/did-core/
package did

import (
	"errors"
	"strings"

	didlib "github.com/pascaldekloe/did"
)

// A DID represents a parsed DID or a DID URL
type DID struct {
	// DID Method
	// https://w3c.github.io/did-core/#method-specific-syntax
	Method string

	// The method-specific-id component of a DID
	// method-specific-id = *idchar *( ":" *idchar )
	ID string

	// method-specific-id may be composed of multiple `:` separated idstrings
	IDStrings []string

	// DID Path, the portion of a DID reference that follows the first forward slash character.
	// https://w3c.github.io/did-core/#path
	Path string

	// Path may be composed of multiple `/` separated segments
	// path-abempty  = *( "/" segment )
	PathSegments []string

	// DID Query
	// https://w3c.github.io/did-core/#query
	// query = *( pchar / "/" / "?" )
	Query string

	// DID Fragment, the portion of a DID reference that follows the first hash sign character ("#")
	// https://w3c.github.io/did-core/#fragment
	Fragment string
}

// IsURL returns true if a DID has a Path, a Query or a Fragment
// https://w3c-ccg.github.io/did-spec/#dfn-did-reference
func (d *DID) IsURL() bool {
	return (d.Path != "" || len(d.PathSegments) > 0 || d.Query != "" || d.Fragment != "")
}

// String encodes a DID struct into a valid DID string.
// nolint: gocyclo
func (d *DID) String() string {
	if d.Method == "" {
		// if there is no Method, return an empty string
		return ""
	}

	var u didlib.URL
	u.Method = d.Method

	if d.ID != "" {
		u.SpecID = d.ID
	} else if len(d.IDStrings) > 0 {
		u.SpecID = strings.Join(d.IDStrings[:], ":")
	} else {
		// if there is no ID, return an empty string
		return ""
	}

	if d.Path != "" {
		u.RawPath = "/" + d.Path
	} else if len(d.PathSegments) > 0 {
		u.SetPathSegments(d.PathSegments...)
	}

	if d.Query != "" {
		u.RawQuery = "?" + d.Query
	}

	if d.Fragment != "" {
		u.RawFragment = "#" + d.Fragment
	}

	return u.String()
}

// Parse parses the input string into a DID structure.
func Parse(input string) (*DID, error) {
	u, err := didlib.ParseURL(input)
	if err != nil {
		return nil, err
	}
	if u.IsRelative() {
		return nil, errors.New("relative URL denied")
	}

	d := DID{
		Method:       u.Method,
		ID:           u.SpecID,
		IDStrings:    strings.Split(u.SpecID, ":"),
		Path:         u.RawPath,
		PathSegments: u.PathSegments(),
		Query:        u.RawQuery,
		Fragment:     u.RawFragment,
	}

	// trim leading characters
	if d.Path != "" {
		d.Path = d.Path[1:]
	}
	if d.Query != "" {
		d.Query = d.Query[1:]
	}
	if d.Fragment != "" {
		d.Fragment = d.Fragment[1:]
	}

	return &d, nil
}
