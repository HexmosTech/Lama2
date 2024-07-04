# How to setup LiveAPI on your API Docs

## Step 1: Add the following script to your project

```html
<script>
  ;(function (
    window,
    document,
    scriptTagName,
    globalName,
    scriptAddress,
    scriptElement,
    otherScriptElement
  ) {
    window["JS-Widget"] = globalName
    window[globalName] =
      window[globalName] ||
      function () {
        ;(window[globalName].q = window[globalName].q || []).push(arguments)
      }
    scriptElement = document.createElement(scriptTagName)
    otherScriptElement = document.getElementsByTagName(scriptTagName)[0]
    scriptElement.id = globalName
    scriptElement.src = scriptAddress
    scriptElement.async = true
    otherScriptElement.parentNode.insertBefore(
      scriptElement,
      otherScriptElement
    )
  })(
    window,
    document,
    "script",
    "liveAPI",
    "https://d2q4vn0gqet98u.cloudfront.net/v0.0.11/liveapi.js"
  )

  if (typeof window.liveAPI === "function") {
    window.liveAPI("liveapi", {
      tagName: "code",
    })
  }
</script>
```

## Step 2

Use the following format in your code block

```
<code>
    GET
    url_to_be_executed
</code>
```

### Example

```
<code>
    GET
    https://httpbin.org/ip
</code>
```

This will render the LiveAPI Widget on your page, like so

![alt text](<liveapibasicexample.png>)



## Additional Configurations

### Logo and link to the site

- You can add your own logo to the widget and provide a link to your website by doing the following
- Towards the end of the script used for importing the script, you can modify it as so

```js
window.liveAPI("liveapi", {
    tagName: "code",
    logoUrl:
        "LINK_TO_YOUR_LOGO_IMAGE",
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


![alt text](<liveapiexample.png>)




