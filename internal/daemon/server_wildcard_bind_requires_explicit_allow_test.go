package daemon

import (
	"testing"
)

func TestWildcardBindRequiresExplicitAllow(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.Host = "0.0.0.0"
	if _, err := New(config); err == nil {
		t.Fatal("expected wildcard bind to require explicit allow")
	}
}
