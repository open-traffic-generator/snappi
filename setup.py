"""
To build distribution: python setup.py sdist bdist_wheel --universal
"""
import os
import setuptools
from version import Version

pkg_name = Version.package_name
version = Version.version
models_version = Version.models_version

# read long description from readme.md
base_dir = os.path.abspath(os.path.dirname(__file__))
with open(os.path.join(base_dir, "readme.md")) as fd:
    long_description = fd.read()

with open("models-release", "w") as out:
    out.write("v" + models_version)

install_requires = []
with open(os.path.join(base_dir, "requirements.txt"), "r+") as fd:
    install_requires = fd.readlines()
    install_requires = install_requires[1:]

setuptools.setup(
    name=pkg_name,
    version=version,
    description="The Snappi Open Traffic Generator Python Package",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/open-traffic-generator/snappi",
    author="Keysight Technologies",
    author_email="andy.balogh@keysight.com",
    license="MIT",
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Testing :: Traffic Generation",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",

    ],
    keywords="snappi testing open traffic generator automation",
    include_package_data=True,
    packages=[pkg_name],
    python_requires=">=3.7, <4",
    install_requires=install_requires,
    extras_require={
        "ixnetwork": ["snappi_ixnetwork==1.34.0"],
        "trex": ["snappi_trex"],
        "testing": ["pytest", "flask"],
        "telemetry": ["opentelemetry-api==1.17.0 ; python_version >= '3.7'",
                      "opentelemetry-sdk ; python_version >= '3.7'",
                      "opentelemetry-exporter-otlp-proto-http ; python_version >= '3.7'",
                      "opentelemetry-exporter-otlp-proto-grpc ; python_version >= '3.7'",
                      "opentelemetry-instrumentation-grpc ; python_version >= '3.7'",
                      "opentelemetry-instrumentation-requests ; python_version >= '3.7'"]
    },
)
