package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond given a time should return the time after
// a gigasecond (1,000,000,000 seconds) has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
