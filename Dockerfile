FROM golang:1.13 as build


ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn


WORKDIR /go/cache

ADD go.mod .
#ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go


FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/app /

CMD ["/app"]



#FROM golang:1.13.0-stretch
#
#COPY . /bazinga-eye
#WORKDIR /bazinga-eye
#
#ENV GO111MODULE=on
#ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o bazinga-eye
#
#CMD ["./bazinga-eye"]

