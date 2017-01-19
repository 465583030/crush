package squirt

import (
	"testing"
)

func init() {
	/* Start webserver */
}

func TestSquirtNewSquirt(t *testing.T) {
	s := NewSquirt("http://localhost")
	if s.url != "http://localhost/" {
		t.Error("Trailing slash should be added")
	}

	ss := NewSquirt("http://localhost/")
	if ss.url != "http://localhost/" {
		t.Error(ss.url + " != http://localhost/ <-- notice trailing slash")
	}
}
