package DJJPromise

import "testing"

func TestBasicRun(t *testing.T) {
	p := NewTask(func() (string, error) {
		return "123", nil
	})

	value, err := p.Await()
	if err != nil {
		t.Error(err)
	}

	if *value != "123" {
		//print(value)
		t.Error("value is not 123")
	}
}
