package usersync

import (
	"testing"
)

func TestLifestreetSyncer(t *testing.T) {
	url := "//ads.lfstmedia.com/idsync/137062?synced=1&ttl=1s&rurl=localhost%2Fsetuid%3Fbidder%3Dlifestreet%26uid%3D%24%24visitor_cookie%24%24"

	info := NewLifestreetSyncer("localhost").GetUsersyncInfo()
	if info.URL != url {
		t.Fatalf("User Sync Info URL '%s' doesn't match '%s'", info.URL, url)
	}
	if info.Type != "redirect" {
		t.Fatalf("should be redirect")
	}
	if info.SupportCORS != false {
		t.Fatalf("should have been false")
	}
}
