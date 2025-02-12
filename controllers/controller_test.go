package controllers

import (
	"testing"
)

func TestGet(t *testing.T) {
	if ret, err := Get(3, 5); err != nil {
		t.Error(err)
	} else {
		if len(ret) < 5 {
			t.Error("Cannot get result")
		}
	}
}

func TestGetMostLike(t *testing.T) {
	if ret, err := GetMostLike(6, 5); err != nil {
		t.Error(err)
	} else {
		if len(ret) != 5 {
			t.Error("Cannot get enough result:", len(ret))
		}
		for k, v := range ret {
			if k == 0 {
				continue
			}
			if ret[k-1].MessageCount.Count < v.MessageCount.Count {
				t.Error("Most like list is not in order: previous=", ret[k-1].MessageCount.Count, " current=", v.MessageCount.Count)
			}
		}
	}
}

func TestGetOne(t *testing.T) {
	// Use a sample URL; actual network calls may be involved.
	url := "http://example.com"
	// Not a valid URL. it must fail.
	if _, err := GetOne(url); err == nil {
		t.Error(err)
	} 
}

func TestGetRandom(t *testing.T) {
	count := 3
	if ret, err := GetRandom(count); err != nil {
		t.Error(err)
	} else {
		if len(ret) != count {
			t.Errorf("GetRandom expected %d articles, got %d", count, len(ret))
		}
	}
}

func TestGetKeyword(t *testing.T) {
	// Use a sample keyword; actual search results may vary.
	keyword := "正妹"
	count := 2
	if ret, err := GetKeyword(count, keyword); err != nil {
		t.Error(err)
	} else {
		if len(ret) != count {
			t.Errorf("GetKeyword expected %d articles, got %d", count, len(ret))
		}
	}
}
