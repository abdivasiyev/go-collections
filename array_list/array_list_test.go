package array_list

import "testing"

func TestArrayList_Contains(t *testing.T) {
	l := New[int]()
	l.Add(1)
	l.Add(2)

	if l.Contains(2) != true {
		t.Errorf("element not found in list")
		t.FailNow()
	}

	if l.Contains(3) != false {
		t.Errorf("element found in list")
		t.FailNow()
	}
}

func TestArrayList_Clear(t *testing.T) {
	l := New[int]()
	l.Add(1)
	l.Add(2)

	l.Clear()

	if l.IsEmpty() != true {
		t.Errorf("list is not empty")
		t.FailNow()
	}
}

func TestArrayList_AddAll(t *testing.T) {
	l := New[int]()
	l.Add(1)
	l.Add(2)

	l2 := New[int]()

	l2.AddAll(l)

	if l2.IsEmpty() != false {
		t.Errorf("list is empty")
		t.FailNow()
	}

	t.Log(l2)
}

func TestArrayList_Add(t *testing.T) {
	l := New[int]()
	l.Add(1)
	l.Add(2)

	if l.IsEmpty() != false {
		t.Errorf("list is empty")
		t.FailNow()
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	l := New[int]()

	if l.IsEmpty() != true {
		t.Errorf("list is not empty; list size: %v", l.Size())
		t.FailNow()
	}
}

func TestArrayList_Size(t *testing.T) {
	l := New[int]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("list size incorrect; list size: %v", l.Size())
		t.FailNow()
	}
}

func TestArrayList_Enumerate(t *testing.T) {
	l := New[int]()

	for k := 0; k < 10; k++ {
		l.Add(k)
	}

	l = l.ForEach().Enumerate(func(i int, item int) int {
		return item
	}).Collect()

	t.Logf("result: %v", l)
}

func TestArrayList_Reduce(t *testing.T) {
	l := New[int]()

	for k := 0; k < 10; k++ {
		l.Add(k)
	}

	sum := l.ForEach().Reduce(func(prev int, item int) int {
		return prev + item
	})

	if sum != 45 {
		t.Fatalf("invalid sum: %v, expected: %v", sum, 45)
	}

	l2 := New[string]()

	l2.Add("Hello,")
	l2.Add("World!")

	s := l2.ForEach().Reduce(func(prev string, item string) string {
		return prev + " " + item
	})

	if s != " Hello, World!" {
		t.Fatalf("invalid join: %v, expected: %v", s, " Hello, World!")
	}
}
