package main

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

func TestHandler(t *testing.T) {
	v := viper.New()
	v.Set("foo.bar", "foobar_value")

	h := &handler{v: v}
	got := h.load("/foo/bar/")
	want := []byte("foobar_value")

	if !bytes.Equal(got, want) {
		t.Errorf("config value got=%v, want=%v", got, want)
	}
}
