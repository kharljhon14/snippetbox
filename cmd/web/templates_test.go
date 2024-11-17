package main

import (
	"testing"

	"github.com/kharljhon14/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Initialize a new time.time object and pass it to the humanDate function
	tests := []struct {
		name string
		tm   string
		want string
	}{
		{
			name: "UTC",
			tm:   "2024-11-17T10:15:00.000000Z",
			want: "17 Nov 2024 at 10:15",
		},
		{
			name: "Empty",
			tm:   "",
			want: "",
		},
	}

	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
