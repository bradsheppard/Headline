FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api api
COPY model model
COPY proto proto
COPY main.go ./

RUN go build -o /backend

EXPOSE 50051

CMD [ "/backend" ]

