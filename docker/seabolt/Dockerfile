# start from golang image based on alpine-3.8
FROM golang:1.10-alpine3.8 AS dev-build
# add our cgo dependencies
RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev git curl pkgconfig
# clone seabolt-1.7.0 source code
RUN git clone -b v1.7.0 https://github.com/neo4j-drivers/seabolt.git /seabolt 
# invoke cmake build and install artifacts - default location is /usr/local
WORKDIR /seabolt/build
# CMAKE_INSTALL_LIBDIR=lib is a hack where we override default lib64 to lib to workaround a defect
# in our generated pkg-config file 
RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target install

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# get the sample source code from gist
WORKDIR /go/src/blog
ADD https://gist.githubusercontent.com/ali-ince/baefaf990f7baedadece133520f31310/raw/bc2087227f600d547d4cab80afd2d872f462f1eb/completed.go main.go
# install dependencies
RUN dep init && dep ensure -add github.com/neo4j/neo4j-go-driver/neo4j
# build statically linked executable
RUN go build -o sample -tags seabolt_static main.go

# base the image to neo4j
FROM neo4j:3.4
# copy the statically linked executable
COPY --from=dev-build /go/src/blog/sample /blog/sample
