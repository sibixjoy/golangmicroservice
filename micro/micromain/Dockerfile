FROM golang:1.19

WORKDIR /app



COPY . ./

RUN go build micromain.go

EXPOSE 8000

CMD [ "./micromain" ]