"""
To build distribution: python setup.py sdist bdist_wheel --universal
"""
import os
import setuptools
import openapiart
import requests
import shutil

pkg_name = "snappi"
go_pkg_name = "gosnappi"
model_protobuf_name = "otg"
version = "0.8.6"
models_version = "0.8.5"

# read long description from readme.md
base_dir = os.path.abspath(os.path.dirname(__file__))
with open(os.path.join(base_dir, "readme.md")) as fd:
    long_description = fd.read()

# download openapi.yaml
openapi_url = "https://github.com/open-traffic-generator/models/releases/download/v{version}/openapi.yaml".format(
    version=models_version
)
response = requests.request("GET", openapi_url, allow_redirects=True)
assert response.status_code == 200
with open(os.path.join("openapi.yaml"), "wb") as fp:
    fp.write(response.content)

openapiart.OpenApiArt(
    api_files=["openapi.yaml"],
    protobuf_name=model_protobuf_name,
    artifact_dir="artifacts",
    extension_prefix=pkg_name,
).GeneratePythonSdk(package_name=pkg_name).GenerateGoSdk(
    package_dir="github.com/open-traffic-generator/snappi/%s" % go_pkg_name, package_name=go_pkg_name
).GenerateGoServer(
    module_path="github.com/open-traffic-generator/snappi/%s" % go_pkg_name,
    models_prefix=go_pkg_name,
    models_path="github.com/open-traffic-generator/snappi/%s" % go_pkg_name
).GoTidy(relative_package_dir=go_pkg_name)

if os.path.exists(pkg_name):
    shutil.rmtree(pkg_name, ignore_errors=True)

# remove unwanted files
shutil.copytree(os.path.join("artifacts", pkg_name), pkg_name)
shutil.copyfile(
    os.path.join("artifacts", "requirements.txt"),
    os.path.join(base_dir, "pkg_requires.txt")
)
shutil.copyfile(
    os.path.join(base_dir, "artifacts", model_protobuf_name + ".proto"),
    os.path.join(base_dir, model_protobuf_name + ".proto")
)
shutil.rmtree("artifacts", ignore_errors=True)
for name in os.listdir(pkg_name):
    path = os.path.join(pkg_name, name)
    print(path + " will be published")

doc_dir = os.path.join(pkg_name, "docs")
os.mkdir(doc_dir)
shutil.move("openapi.yaml", doc_dir)

with open("models-release", "w") as out:
    out.write("v" + models_version)
install_requires = []
with open(os.path.join(base_dir, "pkg_requires.txt"), "r+") as fd:
    install_requires = fd.readlines()
    install_requires = install_requires[1:]

setuptools.setup(
    name=pkg_name,
    version=version,
    description="The Snappi Open Traffic Generator Python Package",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/open-traffic-generator/snappi",
    author="ajbalogh",
    author_email="andy.balogh@keysight.com",
    license="MIT",
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Testing :: Traffic Generation",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3",
    ],
    keywords="snappi testing open traffic generator automation",
    include_package_data=True,
    packages=[pkg_name],
    python_requires=">=2.7, <4",
    install_requires=install_requires,
    # install_requires=[
    #     "requests",
    #     "pyyaml",
    #     "jsonpath-ng",
    #     "typing",
    #     "typing-extensions",
    #     "grpcio==1.38.0 ; python_version > '2.7'", # Warning - changing versions need thorough investigation
    #     "grpcio-tools==1.38.0 ; python_version > '2.7'", # Warning - changing versions need thorough investigation
    #     "grpcio==1.35.0 ; python_version == '2.7'", # Warning - changing versions need thorough investigation
    #     "grpcio-tools==1.35.0 ; python_version == '2.7'", # Warning - changing versions need thorough investigation
    #     "protobuf==3.15.0" # Warning - changing versions need thorough investigation
    # ],
    extras_require={
        "ixnetwork": ["snappi_ixnetwork==0.8.2"],
        "trex": ["snappi_trex"],
        "convergence": ["snappi_convergence==0.3.4"],
        "testing": ["pytest", "flask"],
    },
)
