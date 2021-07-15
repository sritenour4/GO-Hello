package switchback

import (
	"reflect"
	"testing"
	"time"
	
)

func TestSwitchback(t *testing.T) {

	t.Run("hour is even", func(t *testing.T) {
		s := SwitchBack{T: time.Date(2021, 11, 8, 02, 0, 0, 0, time.UTC)}
		want := SwitchBackResponse{"control"}
		got, err := s.EvenOrOdd()

		assertResponse(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("hour is odd", func(t *testing.T) {
		s := SwitchBack{T: time.Date(2021, 11, 8, 01, 0, 0, 0, time.UTC)}
		want := SwitchBackResponse{"treatment"}
		got, err := s.EvenOrOdd()

		assertResponse(t, got, want)
		assertError(t, err, nil)
	})
}

func assertResponse(t testing.TB, got, want SwitchBackResponse) {
	t.Helper()
	
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %v want %v", got, want)
	}
}



// func TestSwitchback(t *testing.T) {

// 	t.Run("hour is even", func(t *testing.T) {
// 		s := SwitchBack{T: time.Date(2021, 11, 8, 02, 0, 0, 0, time.UTC)}
// 		want := SwitchBackResponse{"control"}
// 		if got, err := s.EvenOrOdd(); got != want {
// 			t.Errorf("got %v , want %v %v", got, want, err)
// 		}
// 	})

// 	t.Run("hour is odd", func(t *testing.T) {
// 		s := SwitchBack{T: time.Date(2021, 11, 8, 01, 0, 0, 0, time.UTC)}
// 		want := SwitchBackResponse{"treatment"}
// 		if got, err := s.EvenOrOdd(); got != want {
// 			t.Errorf("got %v , want %v %v", got, want, err)
// 		}
// 	})

// }

