At [Hexmos](http://hexmos.com), we use a git repository called `APIHub`
for collaborating on API files. Here is a description
of how the workflow functions for us:


1. The new engineer clones `APIHub` repository
1. If necessary, create a folder for organizing the new API (ex: `my_new_service`)
1. Start defining `*.l2` files for each service specific API. 
1. Use `l2 file.l2` to test the newly defined APIs (or execute from VSCode). 
1. Push the API files once ready into APIHub repo
1. On the rare merge conflict, the engineer uses standard git conflict resolution
   mechanisms