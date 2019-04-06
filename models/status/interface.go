package status

import "time"

type Updater interface {
	ApplyEffect(s *Status) error
	CancelEffect(s *Status) error
	Validate(t time.Time) bool
	Describe() string
}
