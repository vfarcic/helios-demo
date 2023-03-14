FROM golang:1.18-alpine AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux GOARCH=amd64 go build -v -o helios-demo 

FROM alpine:3.17.2
EXPOSE 8080
CMD ["helios-demo"]
ENV VERSION 1.1.4
COPY --from=build /src/helios-demo /usr/local/bin/helios-demo
RUN chmod +x /usr/local/bin/helios-demo
