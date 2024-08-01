package utility

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestSetBasicOperations(t *testing.T) {
	set := NewSet()

	// Test Add
	set.Add("item1")
	if !set.Contains("item1") {
		t.Errorf("Expected set to contain 'item1'")
	}

	// Test Add duplicate element
	set.Add("item1")
	if !set.Contains("item1") {
		t.Errorf("Expected set to contain 'item1'")
	}
	if set.Size() != 1 {
		t.Errorf("Expected set size to be 1, got %d", set.Size())
	}

	// Test Remove
	set.Remove("item1")
	if set.Contains("item1") {
		t.Errorf("Expected set to not contain 'item1' after removal")
	}

	// Test Contains
	set.Add("item2")
	if !set.Contains("item2") {
		t.Errorf("Expected set to contain 'item2'")
	}

	// Test Size
	set.Add("item3")
	if set.Size() != 2 {
		t.Errorf("Expected set size to be 2, got %d", set.Size())
	}

	// Test Clear
	set.Clear()
	if set.Size() != 0 {
		t.Errorf("Expected set size to be 0 after clear, got %d", set.Size())
	}

	// Test ToSlice
	set.Add(1)
	set.Add("test")
	set.Add(1.0)
	set.Add(time.Time{})

	expected := []interface{}{1, "test", 1.0, time.Time{}}
	actual := set.ToSlice()

	for _, actElem := range actual {
		found := false
		for _, expElem := range expected {
			if reflect.DeepEqual(actElem, expElem) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Element %v in actual slice is not found in expected slice", actElem)
		}
	}
}

func TestSetConcurrency(t *testing.T) {
	set := NewSet()
	var wg sync.WaitGroup
	numGoroutines := 100
	numElements := 1000

	// Test concurrent Add
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < numElements; j++ {
				set.Add(goroutineID*1000 + j)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numElements
	if set.Size() != expectedSize {
		t.Errorf("Expected set size to be %d, got %d", expectedSize, set.Size())
	}

	// Test concurrent Contains
	wg.Add(numGoroutines)
	containsCount := 0
	var mu sync.Mutex
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < numElements; j++ {
				if set.Contains(goroutineID*1000 + j) {
					mu.Lock()
					containsCount++
					mu.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()

	if containsCount != expectedSize {
		t.Errorf("Expected contains count to be %d, got %d", expectedSize, containsCount)
	}

	// Test concurrent Remove
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < numElements; j++ {
				set.Remove(goroutineID*1000 + j)
			}
		}(i)
	}
	wg.Wait()

	if set.Size() != 0 {
		t.Errorf("Expected set size to be 0 after concurrent removals, got %d", set.Size())
	}
}
