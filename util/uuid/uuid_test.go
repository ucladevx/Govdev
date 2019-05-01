package uuid

import (
	"testing"
	"time"
)

func TestUUID(t *testing.T) {
	uuid1 := New()
	time.Sleep(time.Second)
	uuid2 := New()
	uuid3 := New()

	if len(uuid1) != len(uuid2) && len(uuid1) != IDLen {
		t.Errorf("UUIDs generated should have the same length of %d\n", IDLen)
	}

	if uuid1 >= uuid2 {
		t.Errorf("uuids should be sortable based on time\n")
	}

	if uuid2 >= uuid3 {
		t.Errorf("uuids should be sortable based on time\n")
	}
}
