FROM golang:1.19

WORKDIR /app


COPY . . 

RUN go build user.go

EXPOSE 8004

CMD [ "./user" ]