# FROM golang:1.16-alpine

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o ./main

# EXPOSE 8300

# CMD [ "./main" ]

FROM golang:1.19

WORKDIR /app

COPY . .

#RUN go get

RUN go build main.go

EXPOSE 8300

CMD [ "./main" ]