package list

import "testing"

func TestArrayList_Contains(t *testing.T) {
	l := NewArrayList[int]()
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
	l := NewArrayList[int]()
	l.Add(1)
	l.Add(2)

	l.Clear()

	if l.IsEmpty() != true {
		t.Errorf("list is not empty")
		t.FailNow()
	}
}

func TestArrayList_AddAll(t *testing.T) {
	l := NewArrayList[int]()
	l.Add(1)
	l.Add(2)

	l2 := NewArrayList[int]()

	l2.AddAll(l)

	if l2.IsEmpty() != false {
		t.Errorf("list is empty")
		t.FailNow()
	}

	t.Log(l2)
}

func TestArrayList_Add(t *testing.T) {
	l := NewArrayList[int]()
	l.Add(1)
	l.Add(2)

	if l.IsEmpty() != false {
		t.Errorf("list is empty")
		t.FailNow()
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	l := NewArrayList[int]()

	if l.IsEmpty() != true {
		t.Errorf("list is not empty; list size: %v", l.Size())
		t.FailNow()
	}
}

func TestArrayList_Size(t *testing.T) {
	l := NewArrayList[int]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("list size incorrect; list size: %v", l.Size())
		t.FailNow()
	}
}
