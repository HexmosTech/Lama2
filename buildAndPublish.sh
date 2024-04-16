
## Build and publish to master 
GOOS=js GOARCH=wasm go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o static/main.wasm
# cp /home/sreedeep/Downloads/Lama2/static/main.wasm /home/sreedeep/js-widget/dist/main.wasm
scp /home/sreedeep/workspace/Lama2/static/main.wasm  ubuntu@master:/home/ubuntu/main.wasm
ssh ubuntu@master "sudo mv main.wasm /var/www/lamawasm/"
ssh ubuntu@master "sudo systemctl reload nginx"
