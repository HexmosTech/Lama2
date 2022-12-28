# FAQs/RAQs

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
