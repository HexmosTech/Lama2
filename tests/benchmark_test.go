package tests

import (
	"os"
	"testing"

	controller "github.com/HexmosTech/lama2/controller"
)

func BenchmarkBasicGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Args = []string{"l2", "-n", "../examples/0009_processor_basic/0009_processor_basic.l2"}
		// os.Args = []string{"l2", "../elfparser/ElfTestSuite/y_0017_processor.l2"}
		// os.Args = []string{"l2", "../elfparser/ElfTestSuite/y_0018_processor_post.l2"}
		controller.Process("")
	}
}

func BenchmarkBasicPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Args = []string{"l2", "../examples/0001_sample_post_varjson.l2"}
		controller.Process("")
	}
}
