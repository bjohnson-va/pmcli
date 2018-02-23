SDKGen is a protoc plugin to process protos and generate files from the processed protos.

# Development

After making any changes necessary to the code, up the version in the VERSION file and update the CHANGELOG.md with 
your changes.

To test your changes, run
```
inv sdkgen_test --language={a language}
```
from the root mscli directory. This will generate
files in the `/mscli/tools/sdkgen/tests/` directory.

Once you are happy with your changes, up the `dockerImageVersion` in  
[mscli/pkg/sdk/generate.go](https://github.com/vendasta/mscli/blob/master/pkg/sdk/generate.go)
to reflect the new version given to the sdkGen docker image.

Update the version of MSCLI in
[mscli/cmd/version.go](https://github.com/vendasta/mscli/blob/master/cmd/version.go)
and
[mscli/CHANGELOG.md](https://github.com/vendasta/mscli/blog/master/CHANGELOG.md)
to indicate the new sdkgen version.

### Publishing

Once your changes have been reviews and merged, deploy the new
sdkGen docker image by running:
```
inv sdkgen_dockerfile --version={the sdkgen version}
```
from the root mscli directory.

### Testing

To test the whole process locally, either build a local image of sdkgen or publish an RC version. Then set that version
in the `runGenSDK` function. You can then run `go install github.com/vendasta/mscli` to install your current code
version.
