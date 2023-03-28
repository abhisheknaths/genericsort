FROM golang:1.20.2-alpine3.17
ENV BENCH_COMMAND=
WORKDIR /genericsort
COPY . .
RUN go get github.com/stretchr/testify/assert
CMD ${BENCH_COMMAND}