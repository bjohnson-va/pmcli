# Changelog

### Don't forget to set the version in `cmd/version.go` so `mscli version` is correct

5.5.1
- Added support for public routes
- update gosdks to 13.x.x

5.4.0
- Change to python sdkgen that makes the generated code a bit easier to work with
    - Made it so that the internal/client functions generate with keyword args

5.3.0
- Added Graphs to Datadog Generation to monitor Horizontal Pod Autoscaler (HPA), and Unavailable Pods.

5.2.0
- Added `mscli app zone image [zone name] --env=(test|demo|prod) (--tag-only)` to get the current image
running in a zone.

5.1.0
- Added:
   - `mscli app zone list (--enabled)` list the possible zones a microservice can be deployed to.
   The optional `--enabled` flag and be provided to list what zones are serving traffic.
   - `mscli app zone traffic enable [zone name] --env=(test|demo|prod)` enables traffic to be served
   to the given zone for the given environment.
   - `mscli app zone traffic disable [zone name] --env=(test|demo|prod)` disables traffic from being
   served to the given zone for the given environment.

5.0.4
- Cleanup unused import and invalid semi colon in newly bootstrapped project

5.0.3
- Max pod unavailability on test and demo will be 0% instead of 25%

5.0.2
- Use the vendored protos for deploying endpoints, same as sdk generation

5.0.1
- Fix to typescript SDK gen to include all services in index and module

5.0.0
- BREAKING CHANGE
 - Typescript SDKGen will now generate a separate file per service in your protos
 - See [sdkgen changelog](tools/sdkgen/CHANGELOG.md) for details

4.7.0
- Horizontal Pod Autoscalers will be created for zones
- Prod max unavailable will be 25% instead of 0%

4.6.0
- Adds `--zone` flag to the `mscli app deploy` command. This allows deployment images to be restricted
to a subset of zones specified in a microservice's microservice.yaml.

4.5.0
- Adds the `mscli app datadog` command to create a timeboard in datadog

4.4.2
- Missed using variable instead of restarting package discovery in python sdk gen.

4.4.1
- Make the python sdk generation create a setup.py with the correct package discovery

4.4.0
- Add HttpHost to the typescript HostService.

4.3.1
- Change to sdkGen 0.16.1 to fix python appengine import issues.

4.3.0
- Fix all the python sdk generation work. Does require a new version of the sdkGen container

4.2.2
- Update sdkgen Docker image

4.2.1
- Fix initial microservice gen to account for breaking changes on auth and refactors

4.2.0
- Auto-update local endpoints version in docker-compose file

4.1.19
- Fix broken auto-update of local Docker image tag in docker-compose.yaml

4.1.18
- Use the newest SdkGen image

4.1.17
- Fix microservice.yaml generation

4.1.16
- Typescript SDKGen fix: since int64 is JSON encoded as string, need to use parseInt

4.1.15
- Python SDKGen: Allow python SDKs to be generated using the full proto include paths.

4.1.14
- SDKGen: Fix quotes and semicolon in Typescript template

4.1.13
- Temporarily pin to gcloud 183.0.0 to get around build errors

4.1.12
- Recover mscli app buildLocal command (which will use local docker daemon instead of shipping to cloud builder)

4.1.11
- Add tzdata to apt install in dockerfile

4.1.10
- SDKGen: Fix generation for vendasta protofile paths which would contain google in the name

4.1.9
- Fix lint error in generated Typescript SDK

4.1.8
- Fix bug with deployment name

4.1.7
- Fix typescript SDK generation duplicate imports in services

4.1.6
- SDK Gen will now use the root proto path instead of the individual paths for each proto being generated.
This should allow for SDKs to be generated from protos which include other protos outside their directory.

4.1.5
- Update endpoints container image tag to 1.12.0 as 1.13.0 is not officially released, but 1.12.0 includes what we need for now anyway.

4.1.4
- Update endpoints container image tag to 1.13.0

4.1.3
- Update protocContainer path for python sdk gen to gcr.io/repcore-prod

4.1.2
- Only writes `options: machineType: N1_HIGHCPU_8` in the cloudbuild.yaml when on jenkins as it doesn't work running locally

4.1.1
- Add help to the `mscli app` command

4.1.0
- Protos are now compiled with the import directory set to the base vendastaapis directory. This means all imports
  should be relative to that directory and not the current directory of the proto. i.e. if two protos api.proto and addons.proto are
  in accounts/v1. The import path must be changed from `import "addons.proto";` to `import "accounts/v1/addons.proto";`

4.0.0
- Commands that rely on a `microservice.yaml` are now invoked using `mscli app [subcommand]`
  - Examples include: `mscli app test`, `mscli app deploy`, `mscli app endpoints`
- Commands that do not rely on a `microservice.yaml` remain unchanged
  - Examples include: `mscli bootstrap`, `mscli tools`, `mscli version`
- The global "version" option to specify an image version is now called the "tag" option, and it only applies to subcommands of `mscli app`
  - `mscli deploy -v mytag123` is now `mscli app deploy -t mytag123`
- The project code has been completely reorganized
  - Command definitions use the https://github.com/spf13/cobra framework and live in the `cmd` folder
  - Domain logic and the implementation of those commands lives in the `pkg` folder

3.9.0
- Change other container usage to use gcr.io repository.

3.8.0
- Change the container for python sdk generation to use gcr.io.

3.7.0
- New version of typscript SDK-Gen that switches api services to use the angular HttpClient

3.6.0
- Revert ApplyService no longer supports update
- Add UpdateService

3.5.0
- ApplyService now supports updating a service

3.4.0
- Deployments `env` label has been changed to `environment`

3.3.0
- Tag k8 deployments with their zone.

3.2.0
- Increase resources for cloud builder

3.1.1
- Remove github repo url

3.1.0
- Add affinity to a specific zone for our deployments
- Add labels to the deployment

2.30.0
- Add `traffic-enabled` label to deployments

2.29.2
- set server_names_hash_bucket_size in custom endpoints nginx config to 128

2.29.1
- fix and tidy apigateway secret creation

2.29.0
- updated config generation to also generate demo/prod environments,
include `GOOGLE_APPLICATION_CREDENTIALS` and `SERVICE_ACCOUNT` pod
environment variables, service key secret, and secondary SSL config.

2.28.0
- updated `provision` command to create apigateway tls secret,
and not create vendasta-internal-tls secret suspected to be a duplicated
of vendasta-internal-secret which is set automatically in deploys

2.27.0
- new `provision` command which will
  - create a gcloud service-account
  - provision a service account key 
  - create k8s namespace
  - create service account key secret
  - create vendasta internal tls secret 

2.26.0
- mscli dns commands will now test and configure DNS records for the whitelabelled GRPC host at `SecondarySSLConfig.Host`

2.25.0
- New `grpcLoadBalancerIp` and `httpsLoadBalancerIp` fields
- mscli dns commands will now test and configure against these IPs
- Deprecated existing `loadBalancerIp` field of microservice configurations, it is now the fallback for `grpcLoadBalancerIp`

2.24.3
- New version of SDK-Gen to fix missing imports in TypeScript service file
generations.

2.24.2 [2017-11-10]
- Fix empty message generation in Typescript SDK to simply return empty object on toApiJson

2.24.1 [2017-09-15]
- fixed docker-compose.debug.yaml template paths to properly include '/' character'.

2.24.0 [2017-10-12]
- Upgrade build to use golang 1.9
- Remove usage of golang.org/x/net/context.  golang 1.9 supports type aliases so "context" is now acceptable everywhere :)

2.23.3 [2017-10-03]
- Don't generate enum index file if there aren't any enum files

2.23.2 [2017-09-28]
- Minikube is no longer required

2.23.1 [2017-09-28]
- Fix python import path

2.23.0 [2017-09-28]
- Breaking: Python SDK Generation - from_proto will now be handled in the generated layer, and no longer needs to be done in the hand rolled layer.
This means that any SDK generations for existing SDKs will need to have from_proto removed.

2.22.0 [2017-09-25]
- Added new default way to generate the Go stubs for the microservice. Old style stubs can be generated using the --stub empty flag in the setupAll and protoc commands.

2.21.2 [2017-09-25]
- Fix bug with Typescript SDK enums

2.21.1 [2017-09-21]
 - Typescript SDK Gen's preferred host will fallback to the grpchost if secondary ssl is unavailable

2.20.1 [2017-09-21]
 - Typescript SDK Gen no longer relies on @vendasta/core index to import session/environment

2.20.0 [2017-09-19]
- FIx lint errors created during sdkgen for typescript

2.19.1 [2017-09-18]
- Minor refactor in SDKGen

2.19.0 [2017-09-15]
- Use the correct proto namespace for golang imports of the proto files
- Import the proto files from gosdks instead of /pb 

2.18.0 [2017-09-15]
 - Typescript SDK Generator fixed to handle maps as objects rather than list of generated entries

2.17.0 [2017-09-15]
- Change generation to assume the SDK is being added to the @vendasta/core package of frontend

2.16.1 [2017-09-14]
- Fix proto path ordering bug where proto path "." comes first.

2.16.0 [2017-09-14]
- Allow for multiple proto paths to be set in microservice.yaml using `protoPaths` instead of `protoPath`.

2.15.0 []
- Add change /pkg to /internal for all new projects

2.14.2 [2017-09-13]
- Fix protopath formatting to replace `-` with `_` in service name.

2.14.1 [2017-09-12]
- Fix typescript host service generation to use secondary SSL host if available, and to scheme in urls to avoid relative calls

2.14.0 [2017-09-10]
- unwind the slave connection from the build container

2.13.0 [2017-09-08]
- Increase resources for endpoints container

2.12.0 [2017-09-08]
- Log startup errors in the server template as Criticalf so that they are flushed to cloud logging before the pod restarts.

2.11.1 [2017-09-07]
- Support updating docker tags in .yml files when running `mscli buildLocal`

2.11.0 [2017-09-06]
- Tweak resources for endpoints container

2.10.0 [2017-08-30]
- Typescript generation will now create a module, environment, and api files
- Improvements to typescript generation

2.9.0 [2017-08-28]
- Add support for `mscli coverage`

2.8.0 [2017-08-21]
- `mscli setupAll` and `mscli bootstrap` now uses `dep` as its Go dependency manager.

2.7.0 [2017-08-18]
- Add support for Wrapper WKTs (ex. BoolValue) for Python and Typescript

2.6.0 [2017-08-16]
- Java sdkgen no longer sets the value of messages to the default instance, instead it just omits those values

2.5.0 [2017-08-09]
- `mscli endpoints` now accepts the `--skip-generation` flag to skip yaml generation. This is to work around
 the change introduced in 2.2.0 in projects that have made custom modifications to their yaml files.

2.4.1 [2017-08-01]
- use the latest release of go instead of a random git hash

2.4.0 [2017-07-27]
- `mscli jwt` will now infer scope from the microservice.yaml and the --env argument if 
 no scope argument is provided.

2.3.0 [2017-07-26]
- Change rolling update settings for prod apps to be safer
- pull if not present for our applications

2.2.0 [2017-07-18]
- Always regen the endpoints yaml file

2.1.0 [2017-07-11]
- Add Delve to local dockerfile for debugging

2.0.0 [2017-07-08]
- Move the `test-dns` command into a subcommand `dns test`
- Add `dns configure` command for setting up DNS settings for an environment

1.57.0 [2017-07-07]
- Add support for Java code generation

1.56.0 [2017-07-07]
- Add `jenkins` subcommand for nesting commands that are intended to be run inside of jenkins
- Add `jenkins integration-test` command (requires more supporting work) to run integration tests inside jenkins/kubernetes

1.55.0 [2017-07-05]
- Allow apps to optionally configure additional container ports.

1.54.0 [2017-07-04]
- Add Pod IP to the application containers environment.

1.53.0 [2017-07-05]
- Add Typescript to available generateSDK language options

1.52.3 [2017-06-28]
- Use the proper cloud endpoints gcr repo.

1.52.2 [2017-06-28]
- Fix bug where deployment failed on empty secondarySSLConfig

1.52.1 [2017-06-27]
- Add source information to the output descriptor.pb
1.52.0 [2017-06-23]
- Generate a docker compose file during init
- Generate the local-service-account.json for local endpoints
1.51.0 [2017-06-22]
- Update cloud endpoints to the new gcr.io repository.

1.50.0 [2017-06-21]
- Allow a custom Dockerfile to be specified in `microservice.yaml`

1.49.0 [2017-06-15]
- Update sdkgen image to 0.2.0

1.48.0 [2017-06-06]
- Add support for additional SSL certificate

1.47.0 [2017-06-05]
- Use protos from vendored gosdks for generating descriptor for endpoints 

1.46.0 [2017-05-25]
- Add a force flag to the endpoints deploy command

1.45.0 [2017-05-24]
- Only check major/minor go version when determining compatibility with `mscli tools`

1.44.0 [2017-05-24]
- Set CORS headers on all status codes
- Update config maps if they already exist.

1.43.0 [2017-05-24]
- Automatically update vbootstrap-compose.yaml when running `buildLocal`

1.42.2 [2017-05-23]
- Fix bug where endpoints yaml file was being overwritten if it already existed

1.42.1 [2017-05-17]
- Bug fixes for python sdk generator (see sdkgen changelog)

1.42.0 [2017-05-15]
- Update the k8s HPA on deploy

1.41.0 [2017-05-14]
- Add miscroservice name tag to the cloud build

1.40.0 [2017-05-08]
- General improvements to python sdk generator

1.39.0 [2017-05-04]
- Added `buildLocal` command to build the image and return the version to add to your vbootstrap-compose.yaml file

1.38.0 [2017-05-01]
- Refactor sdkgen and change file structure of generated files

1.37.0 [2017-05-01]
- SDKGEN: generate a config.py file for environment variables

1.36.0 [2017-04-26]
- Fix 2 bugs in code generation (missing GRPC import, using var and := in same line)
- Set the default namespace of the test, demo, and prod envs to <app>-env, not to sandbox
- Use a temp directory for any files needed during commands, clean it up after

1.35.0 [2017-04-25]
- update vendasta internal tls cert and key

1.34.0 [2017-04-25]
- separate build and deploy commands so Jenkins doesn't build 3 times when deploying to prod

1.32.0 [2017-04-22]
- Add support for CORS requests to endpoint APIs
- Use IAM auth interceptor to server/main.go
- Add support for IAM auth with cloud endpoints

1.31.0 [2017-04-19]
- Added dns-test option to verify DNS settings
  - Added ability to parse/update local /etc/hosts static NS settings
  - Added code to confirm correct A-Record settings are available over DNS
- Removed the ability to specify a protoc path on the commanline, it now uses
  the `protoPath` from YAML as relative path to the vendastaapis repo under the $GOPATH

1.30.0 [2017-04-12]
- Use cloud container builder for building containers
- Don't run docker commands in cloud ci

1.29.0 [2017-04-04]
- Fix 2 bugs in code generation [duplicate import and mis-named structure declaration]
- Verify SERVICE_ACCOUNT and GOOGLE_APPLICATION_CREDENTIALS are set for non-local deploy commands
- Added support to utils lib for using API_KEY based identity rather than JWT-based ServiceAccounts
- Adjusted codegen for service to use new Authentication Interceptors
- Added support for static token identity (no GCP service account)

1.28.0 [2017-04-03]
- Increase CPU limits/requests for redis on prod.

1.27.0 [2017-04-03]
- Improved generated jenkinsfile to be docker in docker and not docker in docker in docker.

1.26.0 [2017-03-30]
- Improved proto stub generation (with tests!)
  - Experimental server streaming support, not tested yet
- ```mscli endpoints``` now updates the microservice.yaml file with the version
- Add support for custom JWTs (ie. not the default one)

1.25.0 [2017-03-28]
- Fixed local environment HTTPSHost and GRPCHost names to use vendasta-local.com
- Update endpoints template to use allow_unregisterd_callers option (no more api_key url parameter required)
- Update the required gcloud version for endpoints use of allow_unregisterd_callers option
- Pin the cloud.google.com/go package in glide.yaml to ensure correct new pubsub is grabbed
- Update the protoc-stubs.go code to actually parse the protoc generated representation of the protos
- Fix some bugs in the generated api code using new protoc-stubs code (note: multi-file is not supported yet)
- Fix inserted code into the server to correctly use "pb" and "api" module on the inserted code
- Fixed inserted imports into the server to include "api" module
- Fix utils.EnsureDirExists to behave like `mkdir -p` so all missing directories in the path are created
- Fixed endpoints directory check to build "endpoints/<environment>"
- Removed descriptor proto file and go, found in "github.com/golang/protobuf/protoc-gen-go/descriptor"

1.24.0 [2017-03-26]
- Added the `tools` command

1.23.1 [2017-03-20]
- Use the correct context import for auth

1.23.0 [2017-03-20]
- Access Management for MSCLI

1.22.1 [2017-03-15]
- Rollback scratch image work since it doesn't work on jenkins.

1.22.0 [2017-03-07]
- Allow third party apps like redis to be configured on an environment by environment basis

1.21.0 [2017-03-01]
- Add HTTP port to local proxy service.
- Fix issue with partner microservice being hardcoded in docker image.

1.20.0 [2017-03-01]
- Added sdk gen for python

1.19.4 [2017-03-01]
- Build using a golang-2.8 image which produces a binary (buildmode=exe)
- Run using a scratch (vendasta/ca-certs) image instead of golang

1.19.3 [2017-02-28]
- automatically deploy endpoints using gcloud
  - and write the endpoints version back to the microservice.yaml
- (added code for asserting gcloud version)
- Removed GRPCPort and HTTPSPort from the microservice.yaml (hardcoded)

1.19.2 [2017-02-28]
- genConfig command now prompts for the service name
- endpoints config now properly works with annotated protos
- microservice.yaml now contains a top level syntax version
- microservice.yaml now contains the proto path, relative to the vendastaapis repo root

1.19.1 [2017-02-24]
- `microservice.Services`: Don't overwrite `GRPCService` config with `HTTPService` config

1.19.0 [2017-02-22]
- Lower default resource request

1.18.1 [2017-02-17]
- Fix jenkinsfile template

1.18.0 [2017-02-17]
- Refactor mscli to have modules/interfaces.

1.17.0 [2017-02-10]
- give newly bootstrapped apps a reference to a vstore client

1.16.2 [2017-02-09]
- use the grpc host for local proxy to route the traffic to

1.16.1 [2017-02-08]
- Fix outdated main.go template to initialize the healthz server correctly and generate the right golang import paths for generated protos.

1.16.0 [2017-02-08]
- `mscli bootstrap` now also creates a Jenkinsfile and a Dockerfile for jenkins.

1.15.1 [2017-02-01]
- Leverage endpoints' healthz option when available
- Upgrade to endpoints 1.0

1.15.0 [2017-02-01]
- Add readiness/liveness checks to our app container

1.14.0 [2017-01-31]
- Generate a version based on a timestamp if version is not explicitly provided to the deploy command.

1.13.0 [2017-01-28]
- Add jwt command to easily generate a JWT for calling local cloud endpoints APIs.

1.12.2 [2017-01-30]
- Don't use nodeports for HTTP and GRPC services on local

1.12.1 [2017-01-24]
- Detect if we are running on jenkins

1.12.0 [2017-01-24]
- Adding redis app config

1.11.0 [2017-01-23]
- Added local proxy to local deployment

1.10.2 [2017-01-23]
- Build image with minikube env
- Customize domain per ms

1.10.1 [2017-01-23]
- fix image pull secrets

1.10.0 [2017-01-23]
- Only app local app creds on local containers
- Add annotations for local proxy
- Add image pull secrets

1.9.3 [2017-01-19]
- Increase endpoints CPU on prod

1.9.2 [2017-01-19]
 - Dont add local creds to prod containers

1.9.1 [2017-01-19]
 - Handle missing EndpointsVersion (ie. Don't run the container)

1.9.0 [2017-01-17]
 - Add structure to the microservice.yml
 
1.8.0
- Adding support for local ssl story on vendasta-local.com
- Fixing endpoints to run locally.

1.7.8 [2017-01-16]
 - Add dependency check for goreturns

1.7.7 [2017-01-16]
- Update docker in mscli dockerfile

1.7.4 [2017-01-13]
 - Added Horizonal pod autoscalers
 - Various fixes around service/pod configuration.

0.0.5 Add support for GKE and Endpoints [2016-12-21]
 - break up the kubernetes.go file to be more readable
 - Add support for load balancer IP specification
 - Endpoints container now runs

0.0.4 Update to remove WOMCOM issues [2016-12-20]
 - (Not tagged)
 - HACK-239 was the branch and PR
 - Added readme and figured out glide incosistencies
 - minikube deployment working
 - No endpoints container support
 - No GKE deployment

0.0.3 Initial Release [2016-12-08]
 - (Not tagged)
 - First attempt to automate the creation of microservices
 - Creation of the microservice.yaml format

