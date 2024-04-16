#change the file name to push to the site index.html 

scp /home/sreedeep/workspace/Lama2/ztestserver/index.html  ubuntu@master:/home/ubuntu/index.html
ssh ubuntu@master "sudo mv index.html /var/www/lamawasm/"
ssh ubuntu@master "sudo systemctl reload nginx"

scp /home/sreedeep/workspace/Lama2/ztestserver/testJs.js  ubuntu@master:/home/ubuntu/testJs.js
ssh ubuntu@master "sudo mv testJs.js /var/www/lamawasm/"
ssh ubuntu@master "sudo systemctl reload nginx"

