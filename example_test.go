package did_test

import (
	"fmt"
	"github.com/ockam-network/did"
)

func ExampleParse() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"", PathSegments:[]string(nil), Fragment:""}
}

func ExampleParse_withPath() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r/a/b")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"a/b", PathSegments:[]string{"a", "b"}, Fragment:""}
}

func ExampleParse_withFragment() {
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r#keys-1")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	fmt.Printf("%#v", d)
	// Output: &did.DID{Method:"example", ID:"q7ckgxeq1lxmra0r", IDStrings:[]string{"q7ckgxeq1lxmra0r"}, Path:"", PathSegments:[]string(nil), Fragment:"keys-1"}
}
