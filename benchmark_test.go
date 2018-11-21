package did_test

import (
	"net/url"
	"testing"

	"github.com/ockam-network/did"
)

var parsed *did.DID

func BenchmarkParse(b *testing.B) {
	var p *did.DID
	for n := 0; n < b.N; n++ {
		p, _ = did.Parse("did:ockam:amzbjdl8etgpgwoe841sfi6fc4q9yh82m6pkmkw5pteabvtzm7p6qe106ysiawmo")
	}
	parsed = p
}

func BenchmarkParseWithPath(b *testing.B) {
	var p *did.DID
	for n := 0; n < b.N; n++ {
		p, _ = did.Parse("did:ockam:amzbjdl8etgpgwoe841sfi6fc4q9yh82/6pkmkw5pteabvtzm7p6qe106ysiawmo")
	}
	parsed = p
}

func BenchmarkParseWithQuery(b *testing.B) {
	var p *did.DID
	for n := 0; n < b.N; n++ {
		p, _ = did.Parse("did:ockam:amzbjdl8etgpgwoe841sfi6fc4q9yh82?6pkmkw5pteabvtzm7p6qe106ysiawmo")
	}
	parsed = p
}

func BenchmarkParseWithFragment(b *testing.B) {
	var p *did.DID
	for n := 0; n < b.N; n++ {
		p, _ = did.Parse("did:ockam:amzbjdl8etgpgwoe841sfi6fc4q9yh82#6pkmkw5pteabvtzm7p6qe106ysiawmo")
	}
	parsed = p
}

// Sanity check against Go's URL parsing to make sure we're in the same order of magnitude

var parsedURL *url.URL

func BenchmarkUrlParse(b *testing.B) {
	var u *url.URL
	for n := 0; n < b.N; n++ {
		u, _ = url.Parse("http://amzbjdl8etgpgwoe841sfi6fc4q9yh82m6pkmkw5pteabvtzm7p6qe106ysiawm.com")
	}
	parsedURL = u
}

func BenchmarkUrlParseWithPath(b *testing.B) {
	var u *url.URL
	for n := 0; n < b.N; n++ {
		u, _ = url.Parse("http://amzbjdl8etgpgwoe841sfi6fc4q9yh82.com/6pkmkw5pteabvtzm7p6qe106ysiawm")
	}
	parsedURL = u
}

func BenchmarkUrlParseWithQuery(b *testing.B) {
	var u *url.URL
	for n := 0; n < b.N; n++ {
		u, _ = url.Parse("http://amzbjdl8etgpgwoe841sfi6fc4q9yh82.com?6pkmkw5pteabvtzm7p6qe106ysiawm")
	}
	parsedURL = u
}

func BenchmarkUrlParseWithFragment(b *testing.B) {
	var u *url.URL
	for n := 0; n < b.N; n++ {
		u, _ = url.Parse("http://amzbjdl8etgpgwoe841sfi6fc4q9yh82.com#6pkmkw5pteabvtzm7p6qe106ysiawm")
	}
	parsedURL = u
}
