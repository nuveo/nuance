FROM golang:1.7-wheezy
MAINTAINER Avelino <thiago@avelino.xxx>

RUN apt-get update \
    && apt-get install -y \
    automake

RUN git clone https://github.com/guillermocalvo/exceptions4c.git /tmp/exceptions4c
WORKDIR /tmp/exceptions4c
RUN aclocal && autoconf && automake --add-missing && ./configure && make && make install

COPY ./ /go/src/github.com/nuveo/nuance
WORKDIR /go/src/github.com/nuveo/nuance

RUN wget https://mirror.nuveo.com.br/deb/nuance/gpg-pubkey-nuance-lnx1.asc
RUN wget https://mirror.nuveo.com.br/deb/nuance/nuance-omnipage-csdk-lib64_19.2-16357.1200_amd64.deb
RUN wget https://mirror.nuveo.com.br/deb/nuance/nuance-omnipage-csdk-devel_19.2-16357.1200_amd64.deb

RUN gpg --import gpg-pubkey-nuance-lnx1.asc
RUN dpkg -i nuance-omnipage-csdk-lib64_19.2-16357.1200_amd64.deb
RUN dpkg -i nuance-omnipage-csdk-devel_19.2-16357.1200_amd64.deb

RUN go get -u github.com/kardianos/govendor
RUN govendor sync e

ENTRYPOINT ["/go/src/github.com/nuveo/nuance/entrypoint.sh"]
CMD ["go", "run", "main.go"]
