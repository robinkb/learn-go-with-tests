package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"}},

		{"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "Vermont"},
			[]string{"Chris", "Vermont"}},

		{"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 20},
			[]string{"Chris"}},

		{"struct with nested fields",
			Person{
				"Chris",
				Profile{20, "Vermont"},
			},
			[]string{"Chris", "Vermont"}},

		{"pointer to a struct with nested fields",
			&Person{
				"Chris",
				Profile{20, "Vermont"},
			},
			[]string{"Chris", "Vermont"}},

		{"slice of structs",
			[]Profile{
				{10, "London"},
				{22, "Boston"},
			},
			[]string{"London", "Boston"}},

		{"array of structs",
			[2]Profile{
				{10, "London"},
				{22, "Boston"},
			},
			[]string{"London", "Boston"}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string

			walk(test.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}

	t.Run("map of strings", func(t *testing.T) {
		input := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		var got []string
		walk(input, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	for _, x := range haystack {
		if needle == x {
			return
		}
	}

	t.Errorf("expected %+v to contain %q", haystack, needle)
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
