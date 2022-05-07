package captcha

import "testing"

func TestCaptcha(t *testing.T) {
	cap := New()
	if cap.Text() == "" {
		t.Error("captcha text is empty")
	}

	if cap.Length() != 4 {
		t.Error("captcha length is not 4")
	}

	cap2 := New(&Config{
		Length: 6,
	})

	if cap2.Length() != 6 {
		t.Error("captcha length is not 6")
	}

	cap3 := New(&Config{
		BaseChars: "0123456789",
	})

	if cap3.Text() == "" {
		t.Error("captcha text is empty")
	}

	if cap3.Length() != 4 {
		t.Error("captcha length is not 4")
	}
}
