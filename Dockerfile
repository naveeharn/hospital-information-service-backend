FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -tags netgo -ldflags '-s -w' -o hospital_information_service_backend

EXPOSE 4011

CMD [ "./hospital_information_service_backend" ]