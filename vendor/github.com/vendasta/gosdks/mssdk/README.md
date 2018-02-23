Presense Builder SDK
=================
An sdk to communicate with Presence Builder/Microsite's apis via Go.

### 1.0.0
Initial Release

### Instantiating a Client
You need to provide a valid apiUser and apiKey that is registered with Core Services. You must also specify the environment you want to connect to using one of the constants from the gosdks config package.

```
client := FakeClient("myApiUser", "myApiKey", config.Test)
```


