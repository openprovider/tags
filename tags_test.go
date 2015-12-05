// Copyright 2015 Openprovider Authors. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package tags

import (
	"strings"
	"testing"
)

func TestTags(t *testing.T) {
	testData := []struct {
		data   Tags
		query  Tags
		result bool
	}{
		{
			Tags{"popular", "new", "free"},
			Tags{"popular"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"new"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"popular", "new"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"popular", "old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"+popular", "old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"+popular", "+new"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"+popular", "-old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"popular", "-old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"-used", "-old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"popular", "-new"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"-popular", "new"},
			true,
		},
		{
			Tags{"popular"},
			Tags{"+popular"},
			true,
		},
		{
			Tags{"popular"},
			Tags{"-old"},
			true,
		},
		{
			Tags{"popular"},
			Tags{"+popular", "-old"},
			true,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"+popular", "-new"},
			false,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"-popular", "+new"},
			false,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"-popular", "-new"},
			false,
		},
		{
			Tags{"popular", "new", "free"},
			Tags{"old"},
			false,
		},
		{
			Tags{"popular"},
			Tags{"old"},
			false,
		},
		{
			Tags{"popular"},
			Tags{"+old"},
			false,
		},
		{
			Tags{"popular"},
			Tags{"+popular", "+old"},
			false,
		},
		{
			Tags{},
			Tags{"old"},
			false,
		},
	}

	tag := Tags{}
	for _, td := range testData {
		tag = td.data
		if tag.IsTagged(td.query) != td.result {
			answer := "NOT be"
			if td.result {
				answer = "be"
			}
			t.Error(
				"Expected data: [", strings.Join(td.data, ","),
				"] should", answer, "tagged by [", strings.Join(td.query, ","), "]",
			)
		}
	}
}
