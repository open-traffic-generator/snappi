FROM ubuntu:20.04
ENV SRC_ROOT=/home/otg/snappi
ENV GOPATH=${SRC_ROOT}/.local
ENV PATH="${PATH}:${GOPATH}/go/bin"
RUN mkdir -p ${SRC_ROOT}
# Get project source, install dependencies and build it
COPY . ${SRC_ROOT}/
RUN apt-get update \
	&& apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    && apt-get -y install curl git vim unzip python-is-python3 python3-pip
WORKDIR ${SRC_ROOT}
CMD ["/bin/bash"]
