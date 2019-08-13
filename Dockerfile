FROM pangpanglabs/golang:builder AS builder
WORKDIR /go/src/area-china-api
COPY . .
# disable cgo
ENV CGO_ENABLED=0

RUN git -C $GOPATH/src/github.com/pangpanglabs/echoswagger pull
# build steps
RUN echo ">>> 1: go version" && go version
RUN echo ">>> 2: go get" && go get -v -d
RUN echo ">>> 3: go install" && go install
 
# make application docker image use alpine
FROM pangpanglabs/alpine-ssl
WORKDIR /go/bin/
# copy config files to image
COPY --from=builder /go/src/area-china-api/*.yml ./
# copy execute file to image
COPY --from=builder /go/bin/area-china-api ./
EXPOSE 8080
CMD ["./area-china-api"]

