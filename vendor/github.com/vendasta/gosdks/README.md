# Vendasta Go Sdks
This repository includes all of our internal Golang sdks

The `pb` folder includes Golang code generated from our protos at https://github.com/vendasta/vendastaapis/
These can be rebuilt by running `inv build`, which will do the checkout all over again and regenerate all of the protos.

If you only need to build a single directory of protos, you can use `inv build_dir --path=<directory in pb>`, like `inv build_dir --path=account_group`.

You can run tests by running `inv test`

## Releases
After merging your changes to master, you should create a release via Github to ensure that we can rely on `go dep` (https://github.com/golang/dep) to resolve dependencies.

You can do this by clicking on the `X release(s)` link from the main page of the repository and selecting the `Draft a New Release` button. From here you will need to name the release, and most importantly, set a semantic version, ex `v1.0.1`. If you don't understand semantic versioning, either [read about it](http://semver.org/) or ask someone who is familiar, as getting this right is critical. Also take the time to write the changelog into the description.

We hope to completely automate this process based off of `VERSION` and `CHANGELOG.MD` files committed to the repository in the future.

## Logging

You can read more about the logging module in its [README](https://github.com/vendasta/gosdks/blob/master/logging/README.md)

## Generating Mocks
Install mockery: `go get github.com/vektra/mockery/.../`
Run `./gen_mocks.sh <dirname>` to generate mocks for your sdk
