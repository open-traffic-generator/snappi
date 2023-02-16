import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform


BLACK_VERSION = "22.1.0"
GO_VERSION = "1.20"
PROTOC_VERSION = "3.20.3"

# this is where go and protoc shall be installed (and expected to be present)
LOCAL_PATH = os.path.join(os.path.expanduser("~"), ".local")
# path where protoc bin shall be installed or expected to be present
LOCAL_BIN_PATH = os.path.join(LOCAL_PATH, "bin")
# path where go bin shall be installed or expected to be present
GO_BIN_PATH = os.path.join(LOCAL_PATH, "go", "bin")
# path for go package source and installations
GO_HOME_PATH = os.path.join(os.path.expanduser("~"), "go")
GO_HOME_BIN_PATH = os.path.join(GO_HOME_PATH, "bin")

os.environ["GOPATH"] = GO_HOME_PATH
os.environ["PATH"] = "{}:{}:{}:{}".format(
    os.environ["PATH"], GO_BIN_PATH, GO_HOME_BIN_PATH, LOCAL_BIN_PATH
)


def arch():
    return getattr(platform.uname(), "machine", platform.uname()[-1]).lower()


def on_arm():
    return arch() in ["arm64", "aarch64"]


def on_x86():
    return arch() == "x86_64"


def on_linux():
    print("The platform is {}".format(sys.platform))
    return "linux" in sys.platform


def get_go(version=GO_VERSION, targz=None):
    if targz is None:
        if on_arm():
            targz = "go" + version + ".linux-arm64.tar.gz"
        elif on_x86():
            targz = "go" + version + ".linux-amd64.tar.gz"
        else:
            print("host architecture not supported")
            return

    print("Installing Go ...")

    if not os.path.exists(LOCAL_PATH):
        os.mkdir(LOCAL_PATH)

    cmd = "go version 2> /dev/null"
    cmd += " || (rm -rf $(dirname {})".format(GO_BIN_PATH)
    cmd += " && curl -kL -o go-installer https://dl.google.com/go/{}".format(
        targz
    )
    cmd += " && tar -C {} -xzf go-installer".format(LOCAL_PATH)
    cmd += " && rm -rf go-installer"
    cmd += " && echo 'PATH=$PATH:{}:{}' >> ~/.profile".format(
        GO_BIN_PATH, GO_HOME_BIN_PATH
    )
    cmd += " && echo 'export GOPATH={}' >> ~/.profile)".format(GO_HOME_PATH)
    run([cmd])


def get_go_deps():
    print("Getting Go libraries for grpc / protobuf ...")
    cmd = "GO111MODULE=on CGO_ENABLED=0 go install"
    run(
        [
            cmd + " -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0",
            cmd + " -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1",
            cmd + " -v golang.org/x/tools/cmd/goimports@v0.6.0",
            cmd
            + " -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1",
        ]
    )


def get_protoc(version=PROTOC_VERSION, zipfile=None):
    if zipfile is None:
        if on_arm():
            zipfile = "protoc-" + version + "-linux-aarch_64.zip"
        elif on_x86():
            zipfile = "protoc-" + version + "-linux-x86_64.zip"
        else:
            print("host architecture not supported")
            return

    print("Installing protoc ...")

    if not os.path.exists(LOCAL_PATH):
        os.mkdir(LOCAL_PATH)

    cmd = "protoc --version 2> /dev/null || (curl -kL -o ./protoc.zip "
    cmd += "https://github.com/protocolbuffers/protobuf/releases/download/v{}/{}".format(
        version, zipfile
    )
    cmd += " && unzip -o ./protoc.zip -d {}".format(LOCAL_PATH)
    cmd += " && rm -rf ./protoc.zip"
    cmd += " && echo 'PATH=$PATH:{}' >> ~/.profile)".format(LOCAL_BIN_PATH)
    run([cmd])

def setup_ext(go_version=GO_VERSION, protoc_version=PROTOC_VERSION):
    if on_linux():
        get_go(go_version)
        get_protoc(protoc_version)
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


def testpy():
    run(
        [
            py() + " -m pip install flask",
            py() + " -m pip install pytest-cov",
            py() + " -m pytest -sv --cov=snappi --cov-report term --cov-report html:cov_report",
        ]
    )
    import re

    coverage_threshold = 50
    with open("./cov_report/index.html") as fp:
        out = fp.read()
        result = re.findall(r"data-ratio.*?[>](\d+)\b", out)[0]
        if int(result) < coverage_threshold:
            raise Exception("Coverage thresold[{0}] is NOT achieved[{1}]".format(coverage_threshold, result))
        else:
            print("Coverage thresold[{0}] is achieved[{1}]".format(coverage_threshold, result))

def testgo():
    go_coverage_threshold = 0
    # TODO: not able to run the test from main directory
    os.chdir("gosnappi")
    try:
        run([
            "go test ./... -v -coverprofile coverage.txt | tee coverage.out"
        ], raise_exception=True, msg="Exception occured while running the tests")
    finally:
        os.chdir("..")

    with open("gosnappi/coverage.out") as fp:
        out = fp.read()
        result = re.findall(r"coverage:.*\s(\d+)", out)[0]
        if int(result) < go_coverage_threshold:
            raise Exception(
                "Go tests achieved {1}% which is less than Coverage thresold {0}%,".format(
                    go_coverage_threshold, result))
        else:
            print(
                "Go tests achieved {1}% ,Coverage thresold {0}%".format(
                    go_coverage_threshold, result))


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


def run(commands, raise_exception=False, msg=None):
    """
    Executes a list of commands in a native shell and raises exception upon
    failure.
    """
    try:
        for cmd in commands:
            print(cmd)
            if sys.platform != "win32":
                cmd = cmd.encode("utf-8", errors="ignore")
            subprocess.check_call(cmd, shell=True)
    except Exception as e:
        if raise_exception:
            raise Exception(msg)
        print(e)
        sys.exit(1)


def main():
    if len(sys.argv) >= 2:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    main()
