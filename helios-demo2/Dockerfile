FROM golang:1.18-alpine AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux GOARCH=amd64 go build -v -o helios-demo2

FROM alpine:3.17.2
EXPOSE 8080
CMD ["helios-demo2"]
ENV VERSION 1.1.4
COPY --from=build /src/helios-demo2 /usr/local/bin/helios-demo2
RUN chmod +x /usr/local/bin/helios-demo2
