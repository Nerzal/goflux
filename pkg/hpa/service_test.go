package hpa

import "testing"

func TestHPA_Create(t *testing.T) {
	service := NewService()

	err := service.Create("kycnow-api", "clarilab", 4, 8, "../../test/hpa")
	if err != nil {
		t.Error(err)
	}
}
