package Reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{24, "Seoul"}
			aChannel <- Profile{26, "New york"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Seoul", "New york"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{24, "Seoul"}, Profile{26, "New york"}
		}

		var got []string
		want := []string{"Seoul", "New york"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Sonia"},
			[]string{"Sonia"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Sonia", "Seoul"},
			[]string{"Sonia", "Seoul"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sonia", 24},
			[]string{"Sonia"},
		},
		{
			"Nested fields",
			Person{
				"Sonia",
				Profile{24, "Seoul"},
			},
			[]string{"Sonia", "Seoul"},
		},
		{
			"Pointers to things",
			&Person{
				"Sonia",
				Profile{24, "Seoul"},
			},
			[]string{"Sonia", "Seoul"},
		},
		{
			"Slices",
			[]Profile{
				{24, "Seoul"},
				{26, "New york"},
			},
			[]string{"Seoul", "New york"},
		},
		{
			"Arrays",
			[2]Profile{
				{24, "Seoul"},
				{26, "New york"},
			},
			[]string{"Seoul", "New york"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
