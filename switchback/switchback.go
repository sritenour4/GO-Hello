package switchback

import (
	// "errors"
	"time"
)

type SwitchBackInterface interface {
	EvenOrOdd() (SwitchBackResponse, error)
}

type SwitchBackResponse struct {
	Variant string `json:"variant"`
}

type SwitchBack struct {
	T time.Time
}

// If the hour is even, send "control" as variant. If the hour is odd, send "treatment" as variant.
func (s *SwitchBack) EvenOrOdd() (SwitchBackResponse, error) {
	if s.T.Hour()%2 == 0 {
		return SwitchBackResponse{Variant:"control"}, nil
	} else {
		return SwitchBackResponse{Variant:"treatment"}, nil
	} 
	// return SwitchBackResponse{}, errors.New("Error calibrating the switchback element.")
}

