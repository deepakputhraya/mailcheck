package libs

import "testing"

func shouldBeTrue(t *testing.T, value bool, msg string) {
	if !value {
		t.Fatal(msg)
	}
}

func shouldBeFalse(t *testing.T, value bool, msg string) {
	if value {
		t.Fatal(msg)
	}
}

func TestGetEmailDetails(t *testing.T) {
	details, err := GetEmailDetails("test@gmail.com")
	if err != nil {
		t.Fatalf("Error should be nil %s", err)
	}
	shouldBeTrue(t, details.IsValid, "Email should be valid")
	shouldBeTrue(t, details.IsFree, "Email should be free")
	shouldBeTrue(t, details.IsRoleBased, "Email should be role based")
	shouldBeFalse(t, details.IsDisposable, "Email should not be disposable")

	details, err = GetEmailDetails("elonmusk@tesla.com")
	if err != nil {
		t.Errorf("Error should be nil")
	}
	shouldBeTrue(t, details.IsValid, "Email should be valid")
	shouldBeFalse(t, details.IsFree, "Email should not be free")
	shouldBeFalse(t, details.IsRoleBased, "Email should not be role based")
	shouldBeFalse(t, details.IsDisposable, "Email should not be disposable")

	details, err = GetEmailDetails("elonmusk@mailinator.com")
	if err != nil {
		t.Errorf("Error should be nil")
	}
	shouldBeTrue(t, details.IsValid, "Email should be valid")
	shouldBeTrue(t, details.IsFree, "Email should be free")
	shouldBeFalse(t, details.IsRoleBased, "Email should not be role based")
	shouldBeTrue(t, details.IsDisposable, "Email should be disposable")

	details, err = GetEmailDetails("hello-world")
	if err != nil {
		t.Errorf("Error should be nil")
	}
	shouldBeFalse(t, details.IsValid, "Email should not be valid")
}
