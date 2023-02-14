FROM golang:1.19

WORKDIR /app



COPY . ./

RUN go build order.go

EXPOSE 8001

CMD [ "./order" ]