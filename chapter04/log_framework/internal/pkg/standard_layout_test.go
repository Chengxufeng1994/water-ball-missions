package pkg

import (
	"fmt"
	"testing"
)

func TestStandardLayout(t *testing.T) {
	layout := NewStandardLayout()
	log := layout.Format(LevelDebug, "layout.test", "test")
	fmt.Println(log.Message)
}
