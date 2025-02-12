package utils

import (
	"os"
	"testing"
)

func TestGetLogger(t *testing.T) {
	// Test with nil file
	logger := GetLogger(nil)
	if logger == nil {
		t.Error("Expected logger, got nil")
	}

	// Test with an actual file
	f, err := os.CreateTemp("", "logger_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	logger2 := GetLogger(f)
	if logger2 == nil {
		t.Error("Expected logger with file, got nil")
	}
}

func TestCheckTitleWithBeauty(t *testing.T) {
	validTitle := "[正妹] 美麗照片"
	invalidTitle := "[新聞] 今日新聞"
	if !CheckTitleWithBeauty(validTitle) {
		t.Errorf("Expected true for title: %s", validTitle)
	}
	if CheckTitleWithBeauty(invalidTitle) {
		t.Errorf("Expected false for title: %s", invalidTitle)
	}
}

func TestGetPttIDFromURL(t *testing.T) {
	url := "https://www.ptt.cc/bbs/Beauty/12345.html"
	expected := "12345"
	if res := GetPttIDFromURL(url); res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestGetRandomIntSet(t *testing.T) {
	max := 100
	count := 10
	set := GetRandomIntSet(max, count)
	if len(set) != count {
		t.Errorf("Expected %d elements, got %d", count, len(set))
	}
	seen := make(map[int]bool)
	for _, num := range set {
		if num < 0 || num >= max {
			t.Errorf("Value %d out of range [0, %d)", num, max)
		}
		if seen[num] {
			t.Errorf("Found duplicate value %d", num)
		}
		seen[num] = true
	}
}

func TestInArray(t *testing.T) {
	// Test with slice of ints
	nums := []int{1, 2, 3, 4, 5}
	exists, idx := InArray(3, nums)
	if !exists || idx != 2 {
		t.Errorf("Expected to find 3 at index 2, got exists=%v, index=%d", exists, idx)
	}
	exists, _ = InArray(10, nums)
	if exists {
		t.Error("Did not expect to find 10 in the slice")
	}

	// Test with slice of strings
	strs := []string{"a", "b", "c"}
	exists, idx = InArray("b", strs)
	if !exists || idx != 1 {
		t.Errorf("Expected to find 'b' at index 1, got exists=%v, index=%d", exists, idx)
	}
	exists, _ = InArray("z", strs)
	if exists {
		t.Error("Did not expect to find 'z' in the slice")
	}
}

func TestRemoveStringItem(t *testing.T) {
	original := []string{"a", "b", "c", "d"}
	updated := RemoveStringItem(original, 2)
	expected := []string{"a", "b", "d"}
	if len(updated) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(updated))
	}
	for i, v := range expected {
		if updated[i] != v {
			t.Errorf("At index %d, expected %s, got %s", i, v, updated[i])
		}
	}
}

func TestCheckPttURL(t *testing.T) {
	validURL := "https://www.ptt.cc/bbs/Beauty/ABC.html"
	invalidURL := "http://www.example.com"
	if !CheckPttURL(validURL) {
		t.Errorf("Expected valid PTT URL: %s", validURL)
	}
	if CheckPttURL(invalidURL) {
		t.Errorf("Expected invalid PTT URL: %s", invalidURL)
	}
}
