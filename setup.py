"""
To build distribution: python setup.py sdist bdist_wheel --universal
"""
import os
import setuptools
from setuptools.command.install import install
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

grpc_requires = []
with open(os.path.join(base_dir, "grpc-requirements.txt"), "r+") as fd:
    grpc_requires = fd.readlines()
    grpc_requires = grpc_requires[1:]


class InstallCommand(install):
    user_options = install.user_options + [
        ('noGrpc', None, '<description for this custom option>'),
    ]

    def initialize_options(self):
        install.initialize_options(self)
        self.noGrpc = False

    def finalize_options(self):
        global install_requires
        skip_requirement= ["grpc", "protobuf"]
        min_reqs =[]
        if self.noGrpc:
            for skip in skip_requirement:
                for requirement in install_requires:
                    if requirement.startswith(skip):
                        continue
                    else:
                        min_reqs.append(requirement)
        install_requires=min_reqs
        install.finalize_options(self)

    def run(self):
        print(self.noGrpc)
        install.run(self)

print("after")

setuptools.setup(
    name=pkg_name,
    version=version,
    cmdclass={"install": InstallCommand},
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
    extras_require={
        "ixnetwork": ["snappi_ixnetwork==0.9.1"],
        "trex": ["snappi_trex"],
        "convergence": ["snappi_convergence==0.4.1"],
        "testing": ["pytest", "flask"],
    },
)
