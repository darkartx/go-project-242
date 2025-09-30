package code

import "testing"

func TestGetSize(t *testing.T) {
	actual, err := GetSize("testdata", false, false)
	expected := int64(1)

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
	expected := int64(3)

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
	expected := int64(2)

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
	expected := int64(4)

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
		1000:      "1000B",
		1000000:   "1000000B",
		123456789: "123456789B",
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
		1024:       "1.0KB",
		1234567:    "1.2MB",
		1234567890: "1.1GB",
	}

	for value, expected := range cases {
		actual := FormatSize(value, true)

		if actual != expected {
			t.Errorf("Actual %s not equal to expected %s", actual, expected)
			return
		}
	}
}
