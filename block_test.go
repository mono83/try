package try

import "testing"

func TestBlock_Do(t *testing.T) {
	index := 0
	Block{
		Try: func() {
			if index != 0 {
				t.Fatal("Try is supposed to be first")
			}
			index++
			Throw("some")
		},
		Catch: func(e Exception) {
			if index != 1 {
				t.Fatal("Catch is supposed to be second")
			}
			if s, ok := e.(string); ok {
				if s != "some" {
					t.Fatal("Unexpected exception ", s)
				}
			} else {
				t.Fatal("Unexpected exception", e)
			}
			index++
		},
		Finally: func() {
			if index != 2 {
				t.Fatal("Finally is supposed to be third")
			}
		},
	}.Do()
}

func TestBlock_Do_WithoutError(t *testing.T) {
	index := 0
	Block{
		Try: func() {
			if index != 0 {
				t.Fatal("Try is supposed to be first")
			}
			index++
		},
		Catch: func(e Exception) {
			t.Fatal("Catch block should not be invoked")
		},
		Finally: func() {
			if index != 1 {
				t.Fatal("Finally is supposed to be second")
			}
		},
	}.Do()
}
