package main

import (
	"fmt"
	"io"
	"testing"
)

func TestAlternateReader(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		want        int
		hasErr      bool
		expectedErr error
	}{
		{
			name:   "Success TestCase",
			input:  "hello world",
			want:   11,
			hasErr: false,
		},
		{
			name:   "Blank String Test",
			input:  " ",
			want:   1,
			hasErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := []byte(tc.input)
			ar := NewAlternateReader(sb)
			nb := make([]byte, 1024)
			actual, err := ar.Read(nb)
			if tc.hasErr {
				if err == nil {
					t.Fatalf("expected error, received nil")
				}
			} else if !tc.hasErr {
				if err != nil {
					t.Fatalf("expected no error, but received an error")
				} else if err != tc.expectedErr {
					t.Fatalf("expected %v, got  %v", tc.expectedErr, err)
				}
			}
			if actual != tc.want {
				t.Fatalf("expected %d, got %d", tc.want, actual)
			}
		})
	}
}

func TestEofError(t *testing.T) {
	s := "hello world"
	bb := []byte(s)
	nb := make([]byte, 1024)
	ar := NewAlternateReader(bb)
	ab, _ := ar.Read(nb)
	fmt.Println(ab)
	_, err := ar.Read(nb) // read again
	if err != io.EOF {
		t.Fatalf("expected io.EOF error, received: %v", err)
	}
}
