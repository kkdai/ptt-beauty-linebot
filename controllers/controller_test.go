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
