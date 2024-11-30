GOOS=js GOARCH=wasm go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o static/main.wasm
cp /home/rtp/Projects/Lama2/static/main.wasm /home/rtp/Projects/liveapi/dist/main.wasm