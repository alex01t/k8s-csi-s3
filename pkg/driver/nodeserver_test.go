package driver

import "testing"

func TestGetMetaDiskCache(t *testing.T) {
	if getMeta("bucket", "prefix", nil).DiskCache {
		t.Fatal("disk cache must be disabled by default")
	}
	if !getMeta("bucket", "prefix", map[string]string{"diskCache": "true"}).DiskCache {
		t.Fatal("disk cache was not enabled")
	}
	if getMeta("bucket", "prefix", map[string]string{"diskCache": "invalid"}).DiskCache {
		t.Fatal("invalid diskCache value must not enable disk caching")
	}
}
