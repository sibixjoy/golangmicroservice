FROM golang:1.19

WORKDIR /app



COPY . ./

RUN go build inventory.go

EXPOSE 8003

CMD [ "./inventory" ]