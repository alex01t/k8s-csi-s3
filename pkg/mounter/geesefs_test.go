package mounter

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestDerivedCacheDirectory(t *testing.T) {
	t.Setenv("DISK_CACHE_PATH", "")
	if got, want := derivedCacheDirectory("pvc-123"), "/var/lib/geesefs-cache/pvc-123"; got != want {
		t.Fatalf("derivedCacheDirectory() = %q, want %q", got, want)
	}
}

func TestDerivedCacheDirectoryUsesConfiguredRoot(t *testing.T) {
	t.Setenv("DISK_CACHE_PATH", "/mnt/cache")
	if got, want := derivedCacheDirectory("pvc-123"), "/mnt/cache/pvc-123"; got != want {
		t.Fatalf("derivedCacheDirectory() = %q, want %q", got, want)
	}
}

func TestDerivedCacheDirectoryContainsUnsafeVolumeID(t *testing.T) {
	t.Setenv("DISK_CACHE_PATH", "/mnt/cache")
	got := derivedCacheDirectory("../../etc")
	if filepath.Dir(got) != "/mnt/cache" || !strings.HasPrefix(filepath.Base(got), "volume-") {
		t.Fatalf("unsafe volume ID produced cache directory %q", got)
	}
}
