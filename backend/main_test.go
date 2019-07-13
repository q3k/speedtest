package main

// Copyright 2019 Serge Bazanski
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
// SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION
// OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN
// CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

import (
	"net/http"
	"testing"
)

func TestRemoteIP(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"1.2.3.4", "1.2.3.4"},
		{" 1.2.3.4	", "1.2.3.4"},
		{"1.2.3.4.5", ""},
		{"1.2.3.4:8080", "1.2.3.4"},
		{"fe80::2", "fe80::2"},
		{"[fe80::2]:3213", "fe80::2"},
		{"fe80::2:43214", ""},
	}

	for i, c := range cases {
		r := http.Request{
			RemoteAddr: c.in,
		}
		res := remoteIP(&r)
		got := ""
		if res != nil {
			got = res.String()
		}

		if c.out != got {
			t.Errorf("Test case %d: got %q, want %q", i, got, c.out)
		}
	}
}
