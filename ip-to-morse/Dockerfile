FROM golang:alpine as build
WORKDIR /go/src/github.com/ip-to-morse
COPY morse.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix cgo .

FROM alpine
COPY --from=build /go/src/github.com/ip-to-morse/ip-to-morse .
EXPOSE 5555
CMD ["./ip-to-morse"] 
