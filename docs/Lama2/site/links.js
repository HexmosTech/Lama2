function createElementFromHTML(htmlString) {
    var div = document.createElement('div');
    div.innerHTML = htmlString.trim();

    // Change this to div.childNodes to support multiple top-level nodes.
    return div.firstChild;
}

function red() {
    window.location.replace('https://hexmos.com')
    console.log("Redirected to Hexmos home")
    return false;
}
let b = document.getElementsByClassName("md-footer")[0]
console.log(b)
b.prepend(createElementFromHTML(`<div class="logo-wrapper" style="background-color: white">
					<div class="site-logo" style='display:block; float: left; margin: 10px; background: #d7c811; border-radius:5px;'><a class="navbar-brand" href="index.html">
							<img onclick="red()" class="logo-icon me-2" style="width:100px" src="/lama2/hexmoslogo.svg" alt="logo"><span class="logo-text"></span></a>
					</div>
				</div>`))