FROM golang:1.23
WORKDIR /
COPY . .
RUN Go mod tidy
