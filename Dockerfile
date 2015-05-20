FROM google/golang
MAINTAINER Wind Zhu "jiangmail@126.com"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/windzhu/test
RUN go install github.com/windzhu/test

EXPOSE 80
CMD ["/gopath/app/bin/test"]