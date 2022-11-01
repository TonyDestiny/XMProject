FROM golang:1.19-alpine

COPY . /XMProject
WORKDIR /XMProject
RUN go mod download

RUN go build -o /XMProject/build/docker-gs-ping /XMProject/cmd

EXPOSE 8080

CMD [ "/XMProject/build/docker-gs-ping" ]