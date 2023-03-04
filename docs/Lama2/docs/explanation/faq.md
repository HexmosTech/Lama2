# FAQs/RAQs

### Why pick Javascript as the scripting/glue language?

Lama2's [design philosophy](/reference/philosophy.md) advocates "delegate to mature and preferably open tools". 

After analysis, we finally picked JS due to following reasons:

- Most people working with APIs probably already know JS
- Easy to support XPath/JSONPath or even JQ clones right into JS
- Even if you don't know any of (2), simple JS object notation + good old loops will take you far
- Native support for JSON (and dom manipulation is common too, for xml type responses)
- Good amount of power for implementation effort exerted
- For simple use cases, one barely has to understand any serious JS. Object/map notation is quite intuitive.

Ultimately, we believe JS passes our "Intern Test" for usability.

### What is the *Intern Test*?

*Lama2* API files must remain easy for interns to get used with minimal handholding from more experienced engineers. If it doesn't work for interns, then it doesn't work for our teams at Hexmos as well.


### Why did you create *Lama2*?

At [Hexmos](https://hexmos.com), our engineering infrastructure
is split into dozens of self-contained software services.
We deal with 100s of internal APIs, and so felt a need for a robust workflow for defining, sharing and updating APIs within our teams.

Traditional
solutions such as Postman/Insomnia  implement the collaboration features within their applications, and also tend to charge
a fee for collaboration. We felt using `git` is the right way to 
collaborate on APIs, rather than any custom built solution. So,
in a matter of 2-days we got a regex-based prototype DSL language
to store API files. 

Lots of issues cropped up over time, but we kept
making improvements to *Lama2* as needs arose. Hexmos accumulated 100s of APIS
over time. Then, we decided that the tool deserves to be out there,
benefitting teams that want to collaborate on APIs over `git`. So,
to make it happen, first we invested into formalizing the grammar,
and implementing the DSL as a hand-written recursive descent parser.
Then we invested into helpful documentation, demos and so on. Once
we had the basics, we released *Lama2* into the world.
