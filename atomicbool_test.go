package atomicbool

import "testing"

func TestNew(t *testing.T) {
	a := New(false)
	if int32(*a) != 0 {
		t.Errorf("expected 0, got %d", *a)
	}

	a = New(true)
	if int32(*a) != 1 {
		t.Errorf("expected 1, got %d", *a)
	}
}

func TestAtomicBool_Set(t *testing.T) {
	a := New(false)
	a.Set(true)
	if int32(*a) != 1 {
		t.Errorf("expected 1, got %d", *a)
	}

	a.Set(false)
	if int32(*a) != 0 {
		t.Errorf("expected 0, got %d", *a)
	}
}

func TestAtomicBool_Enabled(t *testing.T) {
	a := New(false)
	if a.Enabled() == true {
		t.Error("expected false, got true")
	}

	a.Set(true)
	if a.Enabled() == false {
		t.Error("expected true, got false")
	}
}

func TestAtomicBool_Disabled(t *testing.T) {
	a := New(false)
	if a.Disabled() == false {
		t.Error("expected true, got false")
	}

	a.Set(true)
	if a.Disabled() == true {
		t.Error("expected false, got true")
	}
}
