package clock

import "time"

type Clock interface {
	Now() time.Time
}

type Realtime struct {
}
func (r *Realtime) Now() time.Time {
	return time.Now()
}

type Static struct {
	saved time.Time
}
func (s *Static) Now() time.Time {
	return s.saved
}
func (s *Static) Set(t time.Time) {
	s.saved = t
}
