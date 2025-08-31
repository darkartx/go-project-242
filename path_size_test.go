package code

import "testing"

func TestGetPathSize(t *testing.T) {
	actual, err := GetPathSize("testdata", false, false)
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

func TestGetPathSizeWithRecurcive(t *testing.T) {
	actual, err := GetPathSize("testdata", true, false)
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

func TestGetPathSizeWithAll(t *testing.T) {
	actual, err := GetPathSize("testdata", false, true)
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

func TestGetPathSizeWithAllAndRecurcive(t *testing.T) {
	actual, err := GetPathSize("testdata", true, true)
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

func TestFormatSize(t *testing.T) {
	cases := map[int64]string{
		1000:      "1000 B",
		1000000:   "1000000 B",
		123456789: "123456789 B",
	}

	for value, expected := range cases {
		actual := FormatSize(value, false)

		if actual != expected {
			t.Errorf("Actual %s not equal to expected %s", actual, expected)
			return
		}
	}
}

func TestFormatSizeWithHuman(t *testing.T) {
	cases := map[int64]string{
		1000:      "1.0 kB",
		1000000:   "1.0 MB",
		123456789: "124 MB",
	}

	for value, expected := range cases {
		actual := FormatSize(value, true)

		if actual != expected {
			t.Errorf("Actual %s not equal to expected %s", actual, expected)
			return
		}
	}
}
