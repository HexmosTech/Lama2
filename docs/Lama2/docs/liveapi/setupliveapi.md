# How to setup LiveAPI on your API Docs

## Step 1: Add the following script to your project

```html
<script>
  ((w, d, s) => {
    w.liveAPI = (...args) => (w.liveAPI.q = w.liveAPI.q || []).push(args);
    d.head.appendChild(Object.assign(d.createElement("script"), { src: s, async: true }));
  })(window, document, `https://d2q4vn0gqet98u.cloudfront.net/latest/liveapi.js?foo=${Math.random().toString(36).slice(2)}`);
  liveAPI("loadwidget", {
    selectors: ['#liveapi-element'], // Specify selectors to control where you want the widget to be rendered
  });
</script>
```

## Step 2

Use the following format in your code block

```html
<code id="liveapi-element">
    GET
    url_to_be_executed
</code>
```

### Example

```html
<code id="liveapi-element">
    GET
    https://httpbin.org/ip
</code>
```

This will render the LiveAPI Widget on your page, like so

<pre id="liveapi-element">
GET
https://httpbin.org/ip
</pre>



## Additional Configurations

### Logo and link to the site

- You can add your own logo to the widget and provide a link to your website by doing the following
- Towards the end of the script used for importing the script, you can modify it as so

```js
window.liveAPI("liveapi", {
    selectors: ['#liveapi-element'],
    logoUrl: "LINK_TO_YOUR_LOGO_IMAGE",
    siteLink: "LINK_TO_YOUR_WEBSITE",
})
```

### Adding Metadata to the widgets

- You can add Metadata to each widget, like providing the API name and the sample response
- Adding metadata is done like so

```
GET
url_to_be_executed
---

_METADATA = {"sample_response":"Sample_Response_JSON",
"api_name":"Name of your API"
}
```


Here is a real example.

```
GET
https://httpbin.org/ip

---

_METADATA = {"sample_response":{
    "origin": "IP_ADDRESS"
},
"api_name":"Get the IP"
}
```

This will render the widget like so.

<pre id="liveapi-element">
GET
https://httpbin.org/ip

---

_METADATA = {"sample_response":{
    "origin": "IP_ADDRESS"
},
"api_name":"Get the IP"
}
</pre>





