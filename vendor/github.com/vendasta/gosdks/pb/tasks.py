import shutil
import os
import glob

from invoke import task
import requests


@task
def docs(ctx, file=None):
    """
    Generate documentation.
    Providing --file=my/dir/file.sd will regenerate docs for that file only

    This currently generates
    -   Sequence diagrams in a .png format based off of corresponding plaintext .sd files
    """
    def handler(file_path):
        file_name = os.path.splitext(file_path)[0]
        file_dir = os.path.dirname(file_path)
        output_file_name = os.path.join(file_dir, u"{}.png".format(file_name))
        print u"Generating sequence diagram {} for file {}".format(output_file_name, file_path)
        generate_sequence_diagram(file_path, output_file_name)
    if not file:
        return handle_files(os.getcwd(), "*.sd", handler)
    return handler(file)


def handle_files(path, file_matcher, file_handler, recurse=True):
    """
    Runs file_handler on each file matched by file_matcher.
    :param path: the system path to process
    :param file_matcher: A globbable pattern like "*", "?", "*.proto", "specific.html", etc.
    :param file_handler: A function accepting a filename
    :param recurse: Whether to evaluate subdirectories as well
    :return:
    """
    files = glob.glob(os.path.join(path, file_matcher))
    for f in files:
        file_handler(f)

    if recurse:
        subdirs = [os.path.join(path, name) for name in os.listdir(path) if os.path.isdir(os.path.join(path, name))]
        return [handle_files(d, file_matcher, file_handler, recurse=recurse) for d in subdirs]


def generate_sequence_diagram(input_file, output_file):
    """
    Use www.websequencediagrams.com to generate a PNG from a .sd plaintext file
    """
    with open(input_file, 'r') as f:
        text = f.read()

    payload = {
        "message": text,
        "style": "patent",
        "apiVersion": "1",
    }

    r = requests.post('http://www.websequencediagrams.com/index.php', data=payload)
    if r.status_code != 200:
        raise RuntimeError("websequencediagrams.com did not return a 200 when creating the image: %r", r)
    try:
        image_url = r.json()["img"]
    except KeyError:
        raise RuntimeError("no image url was returned from the server %r" % r)

    r = requests.get("http://www.websequencediagrams.com/%s" % image_url, stream=True)
    if r.status_code != 200:
        raise RuntimeError("websequencediagrams.com did not return a 200 when retrieving the image: %r", r)
    with open(output_file, 'wb') as out_file:
        shutil.copyfileobj(r.raw, out_file)
    return
