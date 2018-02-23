# Overview

MSCLI is a command line interface for bootstrapping production ready microservices on the
Vendasta chosen technology stack.

The concept is that for each microservice project, there is a large amount of boilerplate
code that only has a relatively small amount of project specific parameters, and thus
this code can be generated for you as a template to get started.

MSCLI does this by abstracting these parameters into a config file at the root
of your project, called `microservice.yaml`.

Some (but not all) of the things MSCLI takes care of for you are:
- CI (Jenkins, cloud contatiner builder, etc) config
- Generating a skeleton of a main golang file, initializing things like logging, authentication, and GRPC servers
- GRPC service implementation stubs
- Testing, linting, vetting
- SDK generation
- Cloud Endpoints config
- DNS settings testing
- Datadog Timeboard

# Projects in the Single GOPATH
All the instructions below __assume__ you are following the "Single Go Path" pattern. This means all our projects would fall under `$HOME/go`. Your folders would look like:

```
/User/<USERNAME>/go/src/github.com/vendasta/AA
/User/<USERNAME>/go/src/github.com/vendasta/ARM
/User/<USERNAME>/go/src/github.com/vendasta/BS
...
```

To set this up use the following steps:

1. Install golang from `https://golang.org/dl/`
1. Make the project folder `mkdir -p ~/go/src/github.com/vendasta/`
1. Add the following to your `~/.bash_profile`
  1. `export GOPATH=$HOME/go`
  1. `export GOBIN=$GOPATH/bin`
  1. `export PATH=$GOBIN:$PATH`
1. Source the `~/.bash_profile` (ie restart terminal or whatever)

# Installation

Later the binary might be distributed, for now you have to build it yourself. 

```
go get github.com/vendasta/mscli
```

Install

```
go install github.com/vendasta/mscli
```

Run `mscli` to see a list of commands to verify the tool is on your path, and then...
1. Run `mscli tools` to install the go tool dependencies, such as dep, golint, etc.

# Getting a Microservice Started

*NOTE*: Some of the following instructions are specific only to golang
microservices. See the next section for non-golang microservice instructions.

## Setup protobuf definitions

### Define your proto file
See https://github.com/vendasta/vendastaapis for details & examples.

### Build your proto
Once your proto is defined and pushed to a branch on vendastaapis, we need to build it

1. clone gosdks into `~/go/src/github.com/vendasta/`
1. `cd ~/go/src/github.com/vendasta/gosdks/`
1. Checkout & pull master
1. `invoke build -b <vendastaapis-branch-name>`
1. Assuming no errors, commit and push to your remote gosdks branch


## Setup your microservice

1. `mkdir ~/go/src/github.com/vendasta/<project-name>`
1. `cd ~/go/src/github.com/vendasta/<project-name>`
1. `mscli bootstrap --name=<project-name>`
  * Prerequisites for this is to have your branch checked out for vendastaapis as well as the build protobuf branch in gosdks

You will now have some files generated
* `microservice.yaml`
  * This has the configuration for deploying your microservice. It will generate with configurations for all environments, however, resource limit configuration should be updated to fit the microservice's needs once known.
* `server/`
  * `main.go`
    * The entry point for your application. If you've generated your protos, those services should already be registered in here
    * You probably don't need to touch this for a basic setup
* `internal/[proto file name].go`
    * If you've generated your protos you will have one generated file for each proto file. This is where you should add your implementation

## Initialize and push to your github repo
1. Talk to a team lead to create your repo in github
1. Have the team lead add the `Vendasta Developers (push-pull)` team to the repo collaborators
1. `git init`
1. `git remote add origin https://github.com/vendasta/<project-name>.git`
1. `git add -A`
1. `git commit -m "initial generation of <project-name>"`
1. `git push origin master`

These steps will automatically trigger integrations with Jenkins and Pivotal so that they know about your new microservice.

## Change the version of the vendored gosks to your remote branch
1. Open Gopkg.toml
1. Under `name = "github.com/vendasta/gosdks"`, set one of:
    * `version = "#.#.#"`
    * `branch = "your-branch"`
    * `revision = "your git hash"`
3. Run `dep ensure`

## Building a stub GRPC Service
Let's say you named your proto file metrics.proto. After running setupAll, you will now have a `internal` directory with the file `metrics.go`.
Open `metrics.go` and add your implementation.

mscli will also generate a directory called `pb` which includes Go `proto buffer` code. 
This is legacy code, so please delete this directory and use the code in `github.com/vendasta/gosdks/pb`

## Building and Deploying your container (locally)

Microservices are run locally through docker-compose. You'll need a `docker-compose.yaml` file in the base dir of your project (previously named `vbootstrap-compose.yaml`). You can copy this and modify it from a previous project.

### To build it locally to test it out

If your project uses `docker-compose` locally, you can update the version of your code that is running:

1. `cd <path to where microservice.yaml is located>`
1. `mscli app build -e local`
    1. ~This will output a version number and should automatically update your docker-compose.yaml~
    2. This builds a docker image that you want to set in your docker-compose.yaml
    3. Alternatively, you can specify the version tag with `-t` and the version wont be a timestamp
    
If you have issues running `mscli app build -e local` make sure you have the latest version of mscli installed `mscli version`
If you have issues pulling the docker image you will need to run `docker login -u oauth2accesstoken -p "$(gcloud auth print-access-token)" https://gcr.io` first

To use it, you can start it with `docker-compose up`

## Debugging locally

This will require mscli >= 1.54.0, but you can do manual setup work to get it working with a legacy vbootstrap-compose.yml like in https://github.com/vendasta/snapshot-widget/pull/116.

1. Using vscode, you will need to add a new launch configuration for debugging. You can do this by going to `.vscode > launch.json` (create the file if needed) and add the following (make sure to replace `your-port-number-no-quotes` with the port in your `docker-compose.debug.yml` and `your-repo-name` with the name of your repo, ex. `snapshot-widget`):
```json
{
    "version": "0.2.0",
    "configurations": [
       {
            "name": "Remote Docker",
            "type": "go",
            "request": "launch",
            "mode": "remote",
            "program": "${workspaceRoot}",
            "env": {},
            "args": [],
            "remotePath": "/go/src/github.com/vendasta/your-repo-name",
            "port": your-port-number-no-quotes,
            "host": "127.0.0.1"
        }
    ]
}
```
2. Add breakpoints/watchpoints into your microservice with vscode
2. Start your local docker image (ie. `docker-compose -f docker-compose.yaml -f docker-compose.debug.yaml up`, we put this behind `inv debug`). You will see `API server listening at: [::]:your-port-number` in your docker logs to know Delve is listening
2. Start the debugger in vscode (choose the debugger tab on the side and press play). Make sure you have your new configuration selected! In this case, `Remote Docker`
2. Debug!

## Building and Deploying your container (not locally)
 
### Provisioning

In order to deploy, you first need to provision a few things.
This can be performed for each environment using the `provision` command:
```
mscli provision --env=<env>
```

Provisioning the microservice will:
- Create a Kubernetes namespace
- Create a service account and it's key
- Create secrets in the namespace for the service account key and apigateway.co SSL certificate

### First Deploy

Triggering a Jenkin's build after provisioning should result in your microservice being successfully deployed.
After a short time, the microservice's Kubernetes services will have been assigned ephemeral IPs which you can now configure DNS for.

### Configure DNS

Now you can make your Kubernetes services' ephemeral IPs static, and configure DNS records to route domains to those static IPs.
To do this, run the `dns configure` command on each environment:
```
mscli dns configure --env=<env>
```

Configuring DNS will:
- Detect your load balancers' IPs from Kubernetes
- Change the IP from ephemeral to a reserved static IP
- Insert the load balancers' IPs into your microservice.yaml
- Configure DNS A records for the following domains:
```
https://<service-name>-api-<env>.apigateway.co          # GRPC Service (Whitelabelled)
https://<service-name>-api-<env>.vendasta-internal.com  # GRPC Service
https://<service-name>-<env>.vendasta-internal.com      # HTTPS Service
```

To ensure things are setup correctly, you can run the `dns test` command:
```
mscli dns test --env=<env>
```
This will do a DNS lookup on the above domains to ensure that they route to the IPs assigned to your load balancers.

### Deploy Cloud Endpoints
In order to hit your grpc endpoints you'll need to set up your endpoints configuration. Unfortunately most people don't have permission to add a new domain (or in this case subdomain) to our cloud project. So the initial push of endpoints for each environment needs  to be done by someone that has access to [webmaster tools](https://www.google.com/webmasters/verification/details?hl=en&domain=vendasta-internal.com)
  - Braden Bassingthwaite
  - Brent Yates
  - Shawn Gryschuk
  - Jeremy Rans
  - Jesse Redl
  - Dustin Walker

So talk to them and ask them do push endpoints for you. After they do it the first time you can do it yourself

Run `mscli app endpoints --env=[env]`(env defaults to local) and this will push up your config and return you a version number in the format of 2017-11-14v1. It will also update the version in microservice.yaml. It also creates an endpoints/env/ folder and adds some stuff there.

Add that version to your docker-compose.yaml file to run your µs locally.

A change to the endpoints version locally requires you to update docker-compose yaml file to get those changes. 
A change to endpoints (e.g. you modified request message for a RPC in your proto buffer code) requires that version number to be 
put in microservice.yaml (this should happen automatically) and a deploy for the changes to take effect.

# Using mscli with non-golang microservices

You can set a custom Dockerfile in your `microservice.yaml`, which will allow `mscli` to work with
non-golang microservices. As long as you can run the Docker images built off of your Dockerfile you
should be able to use it without modification.

Your `microservice.yaml` should begin with something like this:

``` yaml
syntax: v1
microservice:
  name: <microservice-name>
  goPackageName: github.com/vendasta/<microservice-name>
  protoPath: adwords_service/v1/adwords_service.proto
  dockerfile: Dockerfile
```

This will configure `mscli` to use the existing `Dockerfile` file in your
project repository instead of generating and using a golang-specific Dockerfile
dynamically.

For a working example, see [the Google AdWords microservice](https://github.com/vendasta/adwords-service) (which is written in Java).

# L7 Load Balancer [HTTP traffic without google auth]
There currently is no MSCLI support for this, the kubernetes configs need to be made manually (you just need to do this once). The following will need to be done for each environment (test, demo, prod)
- Create a gke folder [just one, not for each environment]
- in your-project/gke/ create a `nodeport-[env].yml` file. This will the mapping of port 80 to port 11001. Add these contents and replace [your project] and [env]
```
kind: Service
apiVersion: v1
metadata:
  name: [your-project]-node-port
  namespace: [your-project]-[env]
spec:
  type: NodePort
  selector:
    app: [your-project]
    environment: [env]
  ports:
  - protocol: TCP
    port: 80
    targetPort: 11001
```
- in ./gke create a `ingress-apigateway-co-[env].yml` file. Add these contents and replace [your-project] and [env]
```
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: [your-project]-ingress-apigateway-co
  namespace: [your-project]-[env]
spec:
  tls:
    - secretName: wildcard-apigateway-co
  backend:
    serviceName: [your-project]-node-port
    servicePort: 80
```
- create these configs in kubernetes

`kubectl create -f [Path to nodeport.yml]`

`kubectl create -f [Path to ingress.yml]`

- we are setting this up for apigateway.co so make sure you have your secondarySSlConfig set. If you've got it set up for your api traffic already [your-project]-api-[env].apigateway.co and you want to serve http traffic on [you-project]-[env].apigateway.co point [you-project]-[env].apigateway.co to your ingress ip address.
- Set up your dns entries here: https://console.cloud.google.com/net-services/dns/zones/api-gateway-co?project=repcore-prod
   - you want to point [your-project]-[env].apigateway.co to the ingress ip address. The ingress ip can be found by going to kubctlproxy -> your namespace -> discovery and load balancing -> ingresses
 
 - At this point everything should be set up but your load balancer won't be serving because more than likely the health checks it will use are wrong.
    - in the cloud console for repcore-prod go to Network Services -> Load Balancing and find your load balancer it will be called something like `k8s-um-[your-service]-[env]-[your-service]-ingress-apigateway-co--a597af20`. Open it up and you'll see [this](https://i.imgur.com/7BcQnBD.png)
    - Under `Healthy` it'll probably say 0/6 and we need it to say 6/6 so to fix that we need to fix the health check endpoint
    - Click the identifier beside `Health Check:` which takes you to the health check details page.
    - Edit the health check and change the `Request Path` to be `/healthz` as that's the actual health check endpoint

# Public Routes

## Past (<5.5.0)
To support public routes you had to specify the selectors with no requirements as well as add the routes to an AllowPublicMethods call to IAM. When you generate endpoints from this point forward you need to add the `-s` flag to skip generating the endpoints yaml files and overwriting your customizations.

## Now (>= 5.5.0)
Add your public routes to microservice.yaml file with the routes specified as you would in your browser. (The beginning slash is required)
```
microservice:
  name: snapshot
  protopath: ""
  protopaths:
  - excludeFromSDK: false
    path: snapshot/v1/api.proto
  publicroutes:
  - /snapshot.v1.SnapshotService/Get
  - /snapshot.v1.ReviewSectionService/Get
  repourl: https://github.com/vendasta/snapshot
  useinternalpackage: true
```
MSCLI will now generate your endpoints yamls with the correct selectors with no requirements so that they are public.

You will also need the line `iamAuthService.AllowPublicMethods(strings.Split(os.Getenv("PUBLIC_ROUTES"), ",")...)` in your main.go (now generated this way) so the routes are passed in to iam 


# MSCLI Updates for CI
If the updates made to mscli are intended to run in CI, you will need to update the jenkins file to reference the new mscli image.
- Publish a new mscli image with `inv mscli_dockerfile <version>` (version should match with changelog/doc.go)
- Reference the new image in the [jenkin's repo mscli template](https://github.com/vendasta/jenkins/blob/master/vars/mscliTemplate.groovy)
-- There is another template with reference to this mscli image, but it is a special case for non-mscli golang usage and typically does not need to be updated

# Updating and publishing MSCLI
- Once you have made your changes, run `inv mscli_dockerfile {{version}}` to build and publish the container to the registry
