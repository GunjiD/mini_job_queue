FROM golang:1.18-alpine

WORKDIR /usr/src/routes

COPY ./routes/go.mod ./routes/go.sum ./
RUN go mod download && go mod verify

COPY ./routes .
RUN go build -v -o /usr/local/bin/routes ./

#RUN go mod download

EXPOSE 8080

CMD ["routes"]
