FROM buildpack-deps:jessie-scm

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
	&& rm -rf /var/lib/apt/lists/*

ENV MONGO_USER=mongodb \
    MONGO_DATA_DIR=/var/lib/mongodb \
    MONGO_LOG_DIR=/var/log/mongodb

	RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927 \
	 && echo "deb http://repo.mongodb.org/apt/debian wheezy/mongodb-org/3.2 main" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list \
	 && apt-get update \
	 && DEBIAN_FRONTEND=noninteractive apt-get install -y mongodb-org-server mongodb-org-shell \
	 && sed 's/^bind_ip/#bind_ip/' -i /etc/mongod.conf \
	 && rm -rf /var/lib/apt/lists/*



EXPOSE 27017/tcp
RUN service mongod start

ENV GOLANG_VERSION 1.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 5470eac05d273c74ff8bac7bef5bad0b5abbd1c4052efbdbc8db45332e836b0b

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go get github.com/Cobblestone-Bridge/TurtleOverlord
RUN go get github.com/Masterminds/glide

WORKDIR /go/src/github.com/Cobblestone-Bridge/TurtleOverlord

RUN glide install

EXPOSE 8000

# Run server
