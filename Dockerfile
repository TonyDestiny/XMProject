FROM golang:1.19-alpine

COPY . /XMProject
WORKDIR /XMProject
RUN go mod download

RUN go build -o /docker-gs-ping /XMProject/cmd

EXPOSE 8080

CMD [ "/docker-gs-ping" ]