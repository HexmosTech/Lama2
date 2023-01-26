all: test lint gofumpt buildme mkdocs

buildme: 
	go mod tidy
	go build -o build/l2 -ldflags "-X main.version=`git tag --sort=-version:refname | head -n 1`" l2.go

debug:
	go mod tidy
	go build -gcflags=all="-N -l" -ldflags=-w -ldflags=-compressdwarf=false
	echo "\033[31m(1) At localhost:8081, set gdbserver address as localhost:8082;\n(2) Also run "monitor exit" from gdb client to terminate the server\033[0m"
	-pkill -9 gdbgui && true
	gdbgui -p 8081&
	gdbserver localhost:8082 ./lama2 -o blah.json -vvv elfparser/ElfTestSuite/n_0003_404.l2


lint:
	golangci-lint run --disable-all -E revive ./...

gofumpt:
	gofumpt -w .

test:
	go test ./tests/ -count=1

benchmark:
	go test -bench=. -count 1 -run=^# -benchmem ./tests/

mkdocs:
	cd docs/Lama2; ./build.sh

serve:
	cd docs/Lama2; ./serve.sh

rnotes:
	./releasenotes.sh `git tag --sort=-version:refname | sed -n '2p'` `git tag --sort=-version:refname | head -n 1`

clean:
	rm -rf build/*
