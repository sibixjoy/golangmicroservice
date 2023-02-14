FROM golang:1.19

WORKDIR /app


COPY . ./

RUN go build product.go

EXPOSE 8002

CMD [ "./product" ]