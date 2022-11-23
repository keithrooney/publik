package internal

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {

	t.Run("TestInvalidRequest", func(t *testing.T) {

		expected := errors.New("invalid parameter(s) supplied")

		values := []string{
			"",
			// "169.2.4.19",
		}

		for _, value := range values {
			_, actual := parse(value)
			if expected.Error() != actual.Error() {
				t.Fail()
			}
		}

	})

	t.Run("TestReturnsAsExpected", func(t *testing.T) {

		expected := Response{
			IP: "169.2.4.19",
		}
		actual, err := parse("169.2.4.19:9080")

		if err != nil {
			t.Error(err)
		}

		if expected != actual {
			t.Error("expected != actual")
		}

	})

}
