package mounter

import (
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
	if got, want := derivedCacheDirectory("bucket/../volume"), "/mnt/cache/bucket%2F..%2Fvolume"; got != want {
		t.Fatalf("derivedCacheDirectory() = %q, want %q", got, want)
	}
}
