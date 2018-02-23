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

def get_immediate_subdirectories(a_dir):
    """ Returns a list of the immediate subdirectories in the given path """
    return [name for name in os.listdir(a_dir) if os.path.isdir(os.path.join(a_dir, name))]
