package did

import (
	"fmt"
	"log"
)

func ExampleParse() {
	d, err := Parse("did:example:q7ckgxeq1lxmra0r")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"", PathSegments:[]string(nil), Fragment:""}
}

func ExampleParse_withPath() {
	d, err := Parse("did:example:q7ckgxeq1lxmra0r/a/b")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"a/b", PathSegments:[]string{"a", "b"}, Fragment:""}
}

func ExampleParse_withFragment() {
	d, err := Parse("did:example:q7ckgxeq1lxmra0r#keys-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"", PathSegments:[]string(nil), Fragment:"keys-1"}
}

func ExampleDID_String() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r
}

func ExampleDID_String_withPath() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Path: "a/b"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r/a/b
}

func ExampleDID_String_withPathSegments() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r", PathSegments: []string{"a", "b"}}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r/a/b
}

func ExampleDID_String_withFragment() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Fragment: "keys-1"}
	fmt.Println(d.String())
	// Output: did:example:q7ckgxeq1lxmra0r#keys-1
}

func ExampleDID_IsReference_withPath() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Path: "a/b"}
	fmt.Println(d.IsReference())
	// Output: true
}

func ExampleDID_IsReference_withFragment() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r", Fragment: "keys-1"}
	fmt.Println(d.IsReference())
	// Output: true
}

func ExampleDID_IsReference_noPathOrFragment() {
	d := &DID{Method: "example", ID: "q7ckgxeq1lxmra0r"}
	fmt.Println(d.IsReference())
	// Output: false
}
