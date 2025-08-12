package path_size

import "testing"

func TestGetSize(t *testing.T) {
	actual, err := GetSize("testdata", false, false)
	// expected := int64(1826)
	expected := int64(4235)

	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	if actual != expected {
		t.Errorf("Actual %d not equal to expected %d", actual, expected)
	}
}

func TestGetSizeWithRecurcive(t *testing.T) {
	actual, err := GetSize("testdata", true, false)
	// expected := int64(5968)
	expected := int64(8373)

	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	if actual != expected {
		t.Errorf("Actual %d not equal to expected %d", actual, expected)
	}
}

func TestGetSizeWithAll(t *testing.T) {
	actual, err := GetSize("testdata", false, true)
	// expected := int64(4239)
	expected := int64(4235)

	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	if actual != expected {
		t.Errorf("Actual %d not equal to expected %d", actual, expected)
	}
}

func TestGetSizeWithAllAndRecurcive(t *testing.T) {
	actual, err := GetSize("testdata", true, true)
	// expected := int64(8381)
	expected := int64(8373)

	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	if actual != expected {
		t.Errorf("Actual %d not equal to expected %d", actual, expected)
	}
}
