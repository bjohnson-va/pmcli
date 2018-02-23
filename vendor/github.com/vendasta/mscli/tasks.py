""" Build Tasks """
from invoke import task, run

@task
def sdkgen_dockerfile(ctx, version):
    """ Builds the sdkgen dockerfile"""
    run("docker build -t gcr.io/repcore-prod/sdkgen:{} -f dockerfiles/sdkgen/Dockerfile .".format(version))
    run("gcloud docker -- push gcr.io/repcore-prod/sdkgen:{}".format(version))

@task
def sdkgen_test(ctx, language):
    """ Builds a local sdkgen dockerfile and runs it for the specified language on all test protos """
    run("rm -rf ./tools/sdkgen/tests/output/*")
    run("docker build -t sdkgen:latest -f dockerfiles/sdkgen/Dockerfile .")
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/executive_report/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/executive_report --sdkgen_out=\"{}&executive-report:/output\" /protos/executive_report/api.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/account_group/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/account_group --sdkgen_out=\"{}&account-group:/output\" /protos/account_group/api.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/iam/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/iam --sdkgen_out=\"{}&iam:/output\" /protos/iam/api.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/sales_opportunities/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/sales_opportunities --sdkgen_out=\"{}&sales-opportunities:/output\" /protos/sales_opportunities/test.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/google_my_business/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/google_my_business --sdkgen_out=\"{}&google_my_business:/output\" /protos/google_my_business/api.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/marketplace_packages/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/marketplace_packages --sdkgen_out=\"{}&marketplace_packages:/output\" /protos/marketplace_packages/v1/api.proto /protos/marketplace_packages/v1/addon.proto /protos/marketplace_packages/v1/package.proto /protos/marketplace_packages/v1/product.proto".format(language))
    run("docker run -t -v `pwd`/tools/sdkgen/tests/test_protos/:/protos/ -v `pwd`/tools/sdkgen/tests/output/billing/:/output/ sdkgen:latest -I/protos/ --proto_path=/protos/billing --sdkgen_out=\"{}&billing:/output\" /protos/billing/api.proto".format(language))


@task
def mscli_dockerfile(ctx, version):
    """ Builds the mscli dockerfile"""
    run("docker build -t gcr.io/repcore-prod/mscli:{} -f dockerfiles/mscli/Dockerfile .".format(version))
    run("gcloud docker -- push gcr.io/repcore-prod/mscli:{}".format(version))

@task
def build(ctx):
    """ Runs go build """
    run("go build ./...")

@task
def test(ctx):
    """ Runs tests """
    run("go test ./...")

@task
def vet(ctx):
    """ Runs go vet """
    run("go vet ./...")

@task
def lint(ctx):
    """ Runs golint """
    run("golint ./cmd/... ./pkg/... ./tools/...")

@task(pre=[build, test, lint, vet])
def check(ctx):
    """ Runs all CI tasks """
    pass
