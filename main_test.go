package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	for _, tt := range []struct {
		line string
		want string
	}{
		{
			line: "",
			want: "",
		},
		{
			line: "foo",
			want: "foo",
		},
		{
			line: "2019-12-15 20:29:00 +0000 UTC",
			want: "time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC)",
		},
		{
			line: "0001-01-01 00:00:00 +0000 UTC",
			want: "time.Time{}",
		},
		{
			line: "foo 2019-12-15 20:29:00 +0000 UTC bar",
			want: "foo time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) bar",
		},
		{
			line: "2019-12-15 20:29:00 +0000 UTC 2019-12-15 20:29:00 +0000 UTC",
			want: "time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC)",
		},
		{
			line: "foo 2019-12-15 20:29:00 +0000 UTC bar 2019-12-15 20:29:00 +0000 UTC baz",
			want: "foo time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) bar time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) baz",
		},
		{
			line: "2019-12-15 20:29:00 +0000 UTC 2019-12-15 20:30:00 +0000 UTC",
			want: "time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) time.Date(2019, 12, 15, 20, 30, 0, 0, time.UTC)",
		},
		{
			line: "foo 2019-12-15 20:29:00 +0000 UTC bar 2019-12-15 20:30:00 +0000 UTC baz",
			want: "foo time.Date(2019, 12, 15, 20, 29, 0, 0, time.UTC) bar time.Date(2019, 12, 15, 20, 30, 0, 0, time.UTC) baz",
		},
	} {
		have := replace(tt.line)
		assert.Equal(t, tt.want, have)
	}
}
