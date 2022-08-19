package gday

import (
	"fmt"
	"testing"
	"time"
)

func TestStartOfDay(t *testing.T) {
	expected := StartOfDay(time.Now())

	t.Run("StartOfDay", func(t *testing.T) {
		fmt.Println(expected)
	})
}
