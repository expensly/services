FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/expensly/services

ENV USER ewan
ENV HTTP_ADDR 8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET fuANhlJY9YU9beLo

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://expensly@192.168.59.103:5001/services?sslmode=disable&password=expensly

WORKDIR /go/src/github.com/expensly/services

RUN godep go build
RUN ./services

EXPOSE 8888