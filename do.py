import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform


os.environ["GOPATH"] = os.path.join(
    os.path.dirname(os.path.abspath(__file__)), ".local"
)
os.environ["PATH"] = os.environ["PATH"] + ":{0}/go/bin:{0}/bin".format(os.environ["GOPATH"])


def on_arm():
    arch = platform.uname().machine.lower()
    return arch == "arm64" or arch == "aarch64"


def on_x86():
    arch = platform.uname().machine.lower()
    return arch == "x86_64"


def on_linux():
    return sys.platform == "linux"


def get_go():
    version = "1.17"
    targz = None

    if on_arm():
        targz = "go" + version + ".linux-arm64.tar.gz"
    elif on_x86():
        targz = "go" + version + ".linux-amd64.tar.gz"
    else:
        print("host architecture not supported")
        return

    if not os.path.exists(os.environ["GOPATH"]):
        os.mkdir(os.environ["GOPATH"])

    print("Installing Go ...")
    cmd = "go version 2> /dev/null || curl -kL https://dl.google.com/go/" + targz
    cmd += " | tar -C " + os.environ["GOPATH"] + " -xzf -"
    run([cmd])


def get_go_deps():
    print("Getting Go libraries for grpc / protobuf ...")
    cmd = "GO111MODULE=on go get -v"
    run(
        [
            cmd + " google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0",
            cmd + " google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0",
            cmd + " golang.org/x/tools/cmd/goimports"
        ]
    )


def get_protoc():
    version = "3.17.3"
    zipfile = None

    if on_arm():
        zipfile = "protoc-" + version + "-linux-aarch_64.zip"
    elif on_x86():
        zipfile = "protoc-" + version + "-linux-x86_64.zip"
    else:
        print("host architecture not supported")
        return

    print("Installing protoc ...")
    cmd = "protoc --version 2> /dev/null || ( curl -kL -o ./protoc.zip "
    cmd += "https://github.com/protocolbuffers/protobuf/releases/download/v"
    cmd += version + "/" + zipfile
    cmd += ' && unzip ./protoc.zip -d ' + os.environ["GOPATH"]
    cmd += ' && rm -rf ./protoc.zip )'
    run([cmd])


def setup_ext():
    if on_linux():
        get_go()
        get_protoc()
        get_go_deps()
    else:
        print("Skipping go and protoc installation on non-linux platform ...")


def setup():
    run(
        [
            py() + " -m pip install --upgrade pip",
            py() + " -m pip install --upgrade virtualenv",
            py() + " -m virtualenv .env",
        ]
    )


def init():
    run(
        [
            py() + " -m pip install -r requirements.txt",
        ]
    )


def lint():
    paths = [pkg()[0], "tests", "setup.py", "do.py"]

    run(
        [
            py() + " -m black " + " ".join(paths),
            py() + " -m flake8 " + " ".join(paths),
            py() + " -m pytype " + " ".join(paths),
        ]
    )


def test():
    run(
        [
            py() + " -m pytest -sv",
        ]
    )


def dist():
    clean()
    run(
        [
            py() + " setup.py sdist bdist_wheel --universal",
        ]
    )
    print(os.listdir("dist"))


def install():
    wheel = "{}-{}-py2.py3-none-any.whl".format(*pkg())
    run(
        [
            "{} -m pip install --upgrade --force-reinstall {}[testing]".format(
                py(), os.path.join("dist", wheel)
            ),
        ]
    )


def release():
    run(
        [
            py() + " -m pip install --upgrade twine",
            "{} -m twine upload -u {} -p {} dist/*".format(
                py(),
                os.environ["PYPI_USERNAME"],
                os.environ["PYPI_PASSWORD"],
            ),
        ]
    )


def clean():
    """
    Removes filenames or dirnames matching provided patterns.
    """
    pwd_patterns = [
        ".pytype",
        "dist",
        "build",
        "*.egg-info",
    ]
    recursive_patterns = [
        ".pytest_cache",
        "__pycache__",
        "*.pyc",
        "*.log",
    ]

    for pattern in pwd_patterns:
        for path in pattern_find(".", pattern, recursive=False):
            rm_path(path)

    for pattern in recursive_patterns:
        for path in pattern_find(".", pattern, recursive=True):
            rm_path(path)


def version():
    print(pkg()[-1])


def pkg():
    """
    Returns name of python package in current directory and its version.
    """
    try:
        return pkg.pkg
    except AttributeError:
        with open("setup.py") as f:
            out = f.read()
            name = re.findall(r"pkg_name = \"(.+)\"", out)[0]
            version = re.findall(r"version = \"(.+)\"", out)[0]

            pkg.pkg = (name, version)
        return pkg.pkg


def rm_path(path):
    """
    Removes a path if it exists.
    """
    if os.path.exists(path):
        if os.path.isdir(path):
            shutil.rmtree(path)
        else:
            os.remove(path)


def pattern_find(src, pattern, recursive=True):
    """
    Recursively searches for a dirname or filename matching given pattern and
    returns all the matches.
    """
    matches = []

    if not recursive:
        for name in os.listdir(src):
            if fnmatch.fnmatch(name, pattern):
                matches.append(os.path.join(src, name))
        return matches

    for dirpath, dirnames, filenames in os.walk(src):
        for names in [dirnames, filenames]:
            for name in names:
                if fnmatch.fnmatch(name, pattern):
                    matches.append(os.path.join(dirpath, name))

    return matches


def py():
    """
    Returns path to python executable to be used.
    """
    try:
        return py.path
    except AttributeError:
        py.path = os.path.join(".env", "bin", "python")
        if not os.path.exists(py.path):
            py.path = sys.executable

        # since some paths may contain spaces
        py.path = '"' + py.path + '"'
        return py.path


def run(commands):
    """
    Executes a list of commands in a native shell and raises exception upon
    failure.
    """
    try:
        for cmd in commands:
            print(cmd)
            subprocess.check_call(
                cmd.encode("utf-8", errors="ignore"), shell=True
            )
    except Exception:
        sys.exit(1)


def main():
    if len(sys.argv) >= 2:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    main()
