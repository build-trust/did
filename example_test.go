package did_test

import (
	"fmt"
	"log"

	"github.com/ockam-network/did"
)

func ExampleParse() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method - %s, ID - %s", d.Method, d.ID)
	// Output: Method - example, ID - q7ckgxeq1lxmra0r
}

func ExampleParse_withPath() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r/a/b")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method - %s, ID - %s, Path - %s", d.Method, d.ID, d.Path)
	// Output: Method - example, ID - q7ckgxeq1lxmra0r, Path - a/b
}

func ExampleParse_withQuery() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r?dskjsdjj")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method - %s, ID - %s, Query - %s", d.Method, d.ID, d.Query)
	// Output: Method - example, ID - q7ckgxeq1lxmra0r, Query - dskjsdjj
}

func ExampleParse_withFragment() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r#keys-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method - %s, ID - %s, Fragment - %s", d.Method, d.ID, d.Fragment)
	// Output: Method - example, ID - q7ckgxeq1lxmra0r, Fragment - keys-1
}

func ExampleDID_String() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r
}

func ExampleDID_String_withPath() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Path: "a/b"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r/a/b
}

func ExampleDID_String_withPathSegments() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", PathSegments: []string{"a", "b"}}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r/a/b
}

func ExampleDID_String_withQuery() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Query: "abc"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r?abc
}

func ExampleDID_String_withFragment() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Fragment: "keys-1"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r#keys-1
}

func ExampleDID_IsReference_withPath() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Path: "a/b"}
	fmt.Println(d.IsURL())
	// Output: true
}

func ExampleDID_IsReference_withFragment() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Fragment: "keys-1"}
	fmt.Println(d.IsURL())
	// Output: true
}

func ExampleDID_IsReference_noPathOrFragment() {
	d := &did.DID{Method: "example", ID: "q7ckgxeq1lxmra0r"}
	fmt.Println(d.IsURL())
	// Output: false
}
