package analyze

import "testing"

func TestCount(t *testing.T) {
    counts := Count([]string{"a", "b", "a"})
    if counts["a"] != 2 {
        t.Fatal("expected 2")
    }
}
