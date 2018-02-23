""" Build Tasks """
from invoke import task, run
import os
import glob


@task
def test(ctx, verbose=False):
    """ Runs tests for the go sdks """
    args = ''
    if verbose:
        args += '-v'
    immediate = ['./{}/...'.format(f) for f in get_immediate_subdirectories('.') if
                 f not in ['.git', '.idea', 'vendor', 'dockerfiles', 'pb']]
    run("go test {} {}".format(' '.join(immediate), args))

@task
def lint(ctx):
    """ Runs lint for the go sdks """
    immediate = ['./{}/...'.format(f) for f in get_immediate_subdirectories('.') if
                 f not in ['.git', '.idea', 'vendor', 'dockerfiles', 'pb']]
    run("golint -set_exit_status {}".format(' '.join(immediate)))

@task
def vet(ctx):
    """ Runs vet for the go sdks """
    immediate = ['./{}/...'.format(f) for f in get_immediate_subdirectories('.') if
                 f not in ['.git', '.idea', 'vendor', 'dockerfiles', 'pb']]
    run("go vet {}".format(' '.join(immediate)))    


def protofiles_in_dir(ospath):
    """ Returns the list of protofiles in a path """
    return glob.glob(os.path.join(ospath, "*.proto"))


def generate_protos(ospath, firstpath):
    """ Recursively generates any proto files inside the given directory """
    files = protofiles_in_dir(ospath)
    proto_dir = ospath.replace(firstpath + "/", "")
    if proto_dir == "":
        proto_dir = "."
    if files:
        print("Generating protofiles in directory %s" % ospath)
        cmd = "docker run --rm -v {}:/src gcr.io/repcore-prod/protoc-go:latest  --go_out=plugins=grpc:. -I=. --proto_path={} --proto_path=. {}".format(firstpath, proto_dir, " ".join([proto_dir + "/" + os.path.basename(f) for f in files]))
        print("Running command: %s" % cmd)
        run(cmd)
    subdirs = [os.path.join(ospath, name) for name in os.listdir(ospath) if os.path.isdir(os.path.join(ospath, name))]
    return [generate_protos(d, firstpath) for d in subdirs]


@task
def build_dir(ctx, verbose=False, branch="master", path=None):
    """ Build dependencies, you should only need to run this if you need new or updated protos """
    # clear target
    run("rm -rf ./pb/%s" % path)

    # clear temp
    run('rm -rf ./pb_temp')

    # clone only the head of the specified branch, without fetching branch history, into temp
    run("git clone https://github.com/vendasta/vendastaapis --branch={} --depth=1 ./pb_temp".format(branch))

    # generate proto files (assumptions inside)
    protobuf_dir(ctx, path)

    # fill target
    run('cp -rf ./pb_temp/%s ./pb/%s' % (path, path))

    # clear temp
    run('rm -rf ./pb_temp')


@task
def protobuf_dir(ctx, path):
    """ Generate protobuf files in top-level path
        path: if specified, will restrict generation to only that path
     """
    generate_protos(os.path.join(os.getcwd(), "pb_temp", path), os.path.join(os.getcwd(), "pb_temp"))


@task
def protobuf(ctx):
    """ Generate protobuf files """
    generate_protos(os.path.join(os.getcwd(), "pb"), os.path.join(os.getcwd(), "pb"))


@task
def build(ctx, verbose=False, branch="master"):
    """ Build dependencies, you should only need to run this if you need new or updated protos """
    run("rm -rf ./pb")

    # clone only the head of the specified branch, without fetching branch history
    run("git clone https://github.com/vendasta/vendastaapis --branch={} --depth=1 ./pb".format(branch))

    # strip vcs metadata
    run('rm -rf ./pb/.git*')

    # generate proto files
    protobuf(ctx)


@task
def lint(ctx, project_dir="./"):
    """ Run golint on the target directory"""
    if not project_dir.startswith("./"):
        project_dir = "./" + project_dir
    run(
        "find %s -not -path '*/vendor/*' -and -name '*.go' -and -not -name '*.pb.go' -and -not -name '*_test.go' -exec golint {} \;" % project_dir)


@task
def gosdks_dockerfile(ctx, version):
    """ Builds the gosdks ci dockerfile"""
    run("docker build -t gcr.io/repcore-prod/gosdks_ci:{} -f dockerfiles/ci/Dockerfile .".format(version))
    run("docker push gcr.io/repcore-prod/gosdks_ci:{}".format(version))


def get_immediate_subdirectories(a_dir):
    """ Returns a list of the immediate subdirectories in the given path """
    return [name for name in os.listdir(a_dir) if os.path.isdir(os.path.join(a_dir, name))]
