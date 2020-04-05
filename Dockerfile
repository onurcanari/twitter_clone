FROM golang
LABEL maintainer="Onurcan Ari <arionurcan@gmail.com>"
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
WORKDIR /dist
COPY ./pkg/db/DB.db /dist/pkg/db/DB.db
RUN cp /build/main .
EXPOSE 8888
CMD ["/dist/main"]

