# Using with Angular apps

PMCLI can be used with locally-running angular apps to mock out a server that
the app is trying to make requests against.

##  Install Requestly
Requestly is a chrome extension that can cause internal redirects (307) in the 
browser.  This allows us to substitute our locally-running PMCLI server in place 
of a real server on the internet.

## Add a new rule
In this example, we want all requests on the partners.mydomain.co domain to be
directed to localhost

- Choose "Redirect Request Rule"
- Choose "Url" and "Matches (Wildcard)" for the request
- Choose https://partners.mydomain.co/* as the Request URL
- Choose https://localhost:28001/$1 as the Destination URL
    - The port number will depend on the [mockserver.json](./README.md#configuration) of your PMCLI server
        
![image](https://user-images.githubusercontent.com/8963131/114229702-70d11880-9935-11eb-8861-7979bc615bbf.png)

## Common Errors

### Chrome

Unfortunately, if you are trying to make XMLHttpRequests using the `withCredentials`
option, Chrome will block you.  There doesn't seem to be any way around this.

To use pmcli in Chrome, for calling localhost->localhost with the `withCredentials`
(otherwise known as `credentials: include`) option, you will have to start an
insecure Chrome session.

```
chromium --disable-web-security  --user-data-dir=~/chromeTemp
```

### status: 0
```
ERROR HttpErrorResponse {headers: HttpHeaders, status: 0, statusText: "Unknown Error", url: "https://partners.mydomain.co/something/GetSomething", ok: false, …}
```

1) This probably means your browser is blocking the self-signed certificate that
PMCLI uses.  One of the easiest ways to get around this is by enabling insecure 
localhost connections.

   In Chrome, navigate to: [chrome://flags/#allow-insecure-localhost](
chrome://flags/#allow-insecure-localhost)

2) It can also mean there is mismatch between the scheme being used to serve 
the Angular app and the one being used for PMCLI.  

   Change the `useHttps` setting in mockserver.json to match your Angular app.

### origin '___' has been blocked

```
Access to XMLHttpRequest at 'https://localhost:28000/something/DoSomething' 
(redirected from 'https://partners.mydomain.co/something/DoSomething') 
from origin 'null' has been blocked by CORS policy: Response to preflight request 
doesn't pass access control check: The 'Access-Control-Allow-Origin' header has 
a value 'localhost:4000' that is not equal to the supplied origin.
```

Note: `origin 'null' has been blocked ... header has a value 'localhost:4000'`

PMCLI can only accept requests from one origin domain, and it must match the 
domain that your Angular app's requests are reporting.

Change the `allowedOrigin` setting in mockserver.json to match the origin 
mentioned in the error message **exactly**.
