package did

import (
	"testing"
)

func TestIsReference(t *testing.T) {

	t.Run("returns false if no Path or Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123"}
		if d.IsReference() {
			t.Errorf("returned true")
		}
	})

	t.Run("returns true if Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b"}
		if !d.IsReference() {
			t.Errorf("returned false")
		}
	})

	t.Run("returns true if PathSegements", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", PathSegments: []string{"a", "b"}}
		if !d.IsReference() {
			t.Errorf("returned false")
		}
	})

	t.Run("returns true if Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "00000"}
		if !d.IsReference() {
			t.Errorf("returned false")
		}
	})

	t.Run("returns true if Path and Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b", Fragment: "00000"}
		if !d.IsReference() {
			t.Errorf("returned false")
		}
	})
}

// nolint
func TestString(t *testing.T) {

	t.Run("assembles a DID", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123"}
		output := d.String()
		expected := "did:example:123"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("assembles a DID from IDStrings", func(t *testing.T) {
		d := &DID{Method: "example", IDStrings: []string{"123", "456"}}
		output := d.String()
		expected := "did:example:123:456"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("returns empty string if no method", func(t *testing.T) {
		d := &DID{ID: "123"}
		output := d.String()
		expected := ""
		if output != expected {
			t.Errorf("output: %s, expected: empty string", output)
		}
	})

	t.Run("returns empty string in no ID or IDStrings", func(t *testing.T) {
		d := &DID{Method: "example"}
		output := d.String()
		expected := ""
		if output != expected {
			t.Errorf("output: %s, expected: empty string", output)
		}
	})

	t.Run("includes Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b"}
		output := d.String()
		expected := "did:example:123/a/b"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("includes Path assembled from PathSegements", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", PathSegments: []string{"a", "b"}}
		output := d.String()
		expected := "did:example:123/a/b"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("includes Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "00000"}
		output := d.String()
		expected := "did:example:123#00000"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("does not include Fragment if Path is present", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b", Fragment: "00000"}
		output := d.String()
		expected := "did:example:123/a/b"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})

	t.Run("does not include Fragment if PathSegments is present", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", PathSegments: []string{"a", "b"}, Fragment: "00000"}
		output := d.String()
		expected := "did:example:123/a/b"
		if output != expected {
			t.Errorf("output: %s, expected: %s", output, expected)
		}
	})
}

// nolint
func TestParse(t *testing.T) {

	t.Run("returns error if input is empty", func(t *testing.T) {
		_, err := Parse("")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returns error if input length is less than length 7", func(t *testing.T) {
		_, err := Parse("did:")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returns error if input does not have a second : to mark end of method", func(t *testing.T) {
		_, err := Parse("did:aaaaaaaaaaa")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returns error if method is empty", func(t *testing.T) {
		_, err := Parse("did::aaaaaaaaaaa")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returns error if idstring is empty", func(t *testing.T) {
		_, err := Parse("did:a::123:456")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123::456")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456:")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:/abc")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:#abc")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returns error if input does not begin with did: scheme", func(t *testing.T) {
		_, err := Parse("a:12345")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("returned value is nil if input does not begin with did: scheme", func(t *testing.T) {
		d, _ := Parse("a:12345")
		if d != nil {
			t.Errorf("output value is not nil - %+v", d)
		}
	})

	t.Run("succeeds if it has did prefix and length is greater than 7", func(t *testing.T) {
		d, err := Parse("did:a:1")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}
		if d == nil {
			t.Errorf("return value is nil")
		}
	})

	t.Run("succeeds to extract method", func(t *testing.T) {
		d, err := Parse("did:a:1")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		method := d.Method
		if method != "a" {
			t.Errorf("method is %s, expected: a", method)
		}

		d, err = Parse("did:abcdef:11111")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}
		method = d.Method
		if method != "abcdef" {
			t.Errorf("method is %s, expected: abcdef", method)
		}
	})

	t.Run("returns error if method has any other char than 0-9 or a-z", func(t *testing.T) {
		_, err := Parse("did:aA:1")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:aa-aa:1")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("succeeds to extract id", func(t *testing.T) {
		d, err := Parse("did:a:1")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		id := d.ID
		if id != "1" {
			t.Errorf("id is %s, expected: 1", id)
		}
	})

	t.Run("succeeds to extract id parts", func(t *testing.T) {
		d, err := Parse("did:a:123:456")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		parts := d.IDStrings
		if parts[0] != "123" || parts[1] != "456" {
			t.Errorf("parts is %s, expected: [123 456]", parts)
		}
	})

	t.Run("fails if ID has an invalid char", func(t *testing.T) {
		_, err := Parse("did:a:1&&111")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("succeeds to extract path", func(t *testing.T) {
		d, err := Parse("did:a:123:456/someService")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		path := d.Path
		if path != "someService" {
			t.Errorf("path is %s, expected: someService", path)
		}
	})

	t.Run("succeeds to extract path segements", func(t *testing.T) {
		d, err := Parse("did:a:123:456/a/b")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		segments := d.PathSegments
		if segments[0] != "a" || segments[1] != "b" {
			t.Errorf("segments is %s, expected: [a b]", segments)
		}
	})

	t.Run("succeeds with percent encoded chars in path", func(t *testing.T) {
		d, err := Parse("did:a:123:456/a/%20a")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		path := d.Path
		if path != "a/%20a" {
			t.Errorf("path is %s, expected: a/%%20a", path)
		}
	})

	t.Run("fails if % in path is not followed by 2 hex chars", func(t *testing.T) {
		_, err := Parse("did:a:123:456/%")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456/%a")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456/%!*")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456/%!A")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456/%A!")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456/%A%")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("fails if path is empty but there is a slash", func(t *testing.T) {
		_, err := Parse("did:a:123:456/")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("fails if first path segment is empty", func(t *testing.T) {
		_, err := Parse("did:a:123:456//abc")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("does not fail if second path segment is empty", func(t *testing.T) {
		_, err := Parse("did:a:123:456/abc//pqr")
		if err != nil {
			t.Errorf("error is not nil")
		}
	})

	t.Run("fails if path has invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456/ssss^sss")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("does not fail if path has atleast one segment and a trailing slash", func(t *testing.T) {
		_, err := Parse("did:a:123:456/a/b/")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}
	})

	t.Run("succeeds to extract fragment", func(t *testing.T) {
		d, err := Parse("did:a:123:456#keys-1")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		f := d.Fragment
		if f != "keys-1" {
			t.Errorf("fragment is %s, expected: keys-1", f)
		}
	})

	t.Run("succeeds with percent encoded chars in fragment", func(t *testing.T) {
		d, err := Parse("did:a:123:456#aaaaaa%20a")
		if err != nil {
			t.Errorf("error is not nil - %+v", err)
		}

		path := d.Fragment
		if path != "aaaaaa%20a" {
			t.Errorf("path is %s, expected: aaaaaa%%20a", path)
		}
	})

	t.Run("fails if % in fragment is not followed by 2 hex chars", func(t *testing.T) {
		_, err := Parse("did:a:123:456#%")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456#%a")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456#%!*")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456#%!A")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456#%A!")
		if err == nil {
			t.Errorf("error is nil")
		}

		_, err = Parse("did:a:123:456#%A%")
		if err == nil {
			t.Errorf("error is nil")
		}
	})

	t.Run("fails if fragment has invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456#ssss^sss")
		if err == nil {
			t.Errorf("error is nil")
		}
	})
}

func Test_errorf(t *testing.T) {
	p := &parser{}
	p.errorf(10, "%s,%s", "a", "b")

	if p.currentIndex != 10 {
		t.Errorf("did not set currentIndex")
	}

	e := p.err.Error()
	if e != "a,b" {
		t.Errorf("err message is: '%s' expected: 'a,b'", e)
	}
}

func Test_isNotValidIDChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'.', '-'}
	for _, c := range a {
		if isNotValidIDChar(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'%', '^', '#', ' ', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '@', '/', '?'}
	for _, c := range a {
		if !isNotValidIDChar(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotValidFragmentChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=',
		':', '@',
		'/', '?'}
	for _, c := range a {
		if isNotValidFragmentChar(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'%', '^', '#', ' '}
	for _, c := range a {
		if !isNotValidFragmentChar(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotValidPathChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=',
		':', '@'}
	for _, c := range a {
		if isNotValidPathChar(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'%', '/', '?'}
	for _, c := range a {
		if !isNotValidPathChar(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotUnreservedOrSubdelim(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '='}
	for _, c := range a {
		if isNotUnreservedOrSubdelim(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'%', ':', '@', '/', '?'}
	for _, c := range a {
		if !isNotUnreservedOrSubdelim(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotHexDigit(t *testing.T) {
	a := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f'}
	for _, c := range a {
		if isNotHexDigit(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'G', 'g', '%', '\x40', '\x47', '\x60', '\x67'}
	for _, c := range a {
		if !isNotHexDigit(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotDigit(t *testing.T) {
	a := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, c := range a {
		if isNotDigit(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'A', 'a', '\x29', '\x40', '/'}
	for _, c := range a {
		if !isNotDigit(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotAlpha(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z'}
	for _, c := range a {
		if isNotAlpha(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'\x40', '\x5B', '\x60', '\x7B', '0', '9', '-', '%'}
	for _, c := range a {
		if !isNotAlpha(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotBigLetter(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Z', 'Y', 'Z'}
	for _, c := range a {
		if isNotBigLetter(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'\x40', '\x5B', 'a', 'z', '1', '9', '-', '%'}
	for _, c := range a {
		if !isNotBigLetter(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}

func Test_isNotSmallLetter(t *testing.T) {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'z', 'y', 'z'}
	for _, c := range a {
		if isNotSmallLetter(c) {
			t.Errorf("should be false but returned true: %v", c)
		}
	}

	a = []byte{'\x60', '\x7B', 'A', 'Z', '1', '9', '-', '%'}
	for _, c := range a {
		if !isNotSmallLetter(c) {
			t.Errorf("should be true but returned false: %v", c)
		}
	}
}
