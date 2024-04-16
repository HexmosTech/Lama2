document.addEventListener("DOMContentLoaded", function() {

    function makeRequestThroughProxy(){
        const username = 'proxyServer';
        const password = 'proxy22523146server';
        const auth = 'Basic ' + btoa(username + ':' + password);

        // var options = {
        // host: "proxyserver.hexmos.com",
        // path: "https://www.google.com",
        // headers: {
        //     Host: "www.google.com",
        //     'Proxy-Authorization': auth   
        // },
        // // mode: 'cors'
        // };
        // fetch('https://www.google.com', options)
        // .then(response => {
        //     if (!response.ok) {
        //         throw new Error('Network response was not ok');
        //     }
        //     return response.text();
        // })
        // .then(data => {
        //     console.log(data);
        // })
        // .catch(error => {
        //     console.error('There was a problem with your fetch operation:', error);
        // });

        const proxyUrl = 'https://proxyserver.hexmos.com/';
        const targetUrl = 'https://www.google.com';
    
        var xhr = new XMLHttpRequest();
        var url = proxyUrl + targetUrl;
    
        xhr.open('GET', url, true);
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest'); // Set an allowed header
        xhr.setRequestHeader('Proxy-Authorization', auth);
        xhr.onreadystatechange = function() {
            console.log('Ready state:', xhr.readyState);
            console.log('Status:', xhr.status);
            console.log('Status text:', xhr.statusText);

            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    console.log(xhr.responseText);
                } else {
                    console.error('Error:', xhr.status);
                }
            }
        };
        xhr.send();
    }
    // Add event listener to the button
    document.getElementById('proxyButton').addEventListener('click', makeRequestThroughProxy);
});
