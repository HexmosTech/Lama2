package tests

import (
	"os"
	"testing"

	controller "github.com/HexmosTech/lama2/controller"
)

func BenchmarkBasicGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Args = []string{"l2", "../examples/0000_sample_get.l2"}
		controller.Process("")
	}
}

func BenchmarkBasicPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Args = []string{"l2", "../examples/0001_sample_post_varjson.l2"}
		controller.Process("")
	}
}
