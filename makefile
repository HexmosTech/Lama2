all: test lint gofumpt buildme mkdocs

buildme: 
	go mod tidy
	go build -o build/l2 -ldflags "-X main.version=`git tag --sort=-version:refname | head -n 1`" l2.go

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