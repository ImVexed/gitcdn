# gitcdn
Is a repository that contains a simple `raw.githubusercontent.com` alternative that additionally adds a `Content-Type` header based on the requested file type extension to the response.

The various languages this has been implemented in hold no opinion on what the "correct" packages/dependencies to be used are. They were simply chosen due to their ease of use or maintainability at the time of writing.

## Example
Each program functions the same as `raw.githubusercontent.com` the only change that needs to be made to requests is the base url.

For example
```https://raw.githubusercontent.com/golang/go/master/doc/gopher/fiveyears.jpg```
while debugging locally becomes
```http://localhost/golang/go/master/doc/gopher/fiveyears.jpg```