package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond to return the moment after a gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
