var http = require("http");

// const auth = 'Basic ' + Buffer.from('proxyServer:proxy22523146server').toString('base64')
const username = 'proxyServer';
const password = 'proxy22523146server';
const auth = 'Basic ' + btoa(username + ':' + password);

var options = {
    host: "proxyserver.hexmos.com",
    port: 8080,
    path: "https://proxyserver.hexmos.com",
    headers: {
        Host: "proxyserver.hexmos.com",
        'Proxy-Authorization': auth
    }
};

http.get(options, function(res) {
    console.log(res);
    res.pipe(process.stdout);
    
});

module.exports = http
