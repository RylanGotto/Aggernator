FROM golang:1.23-bookworm

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o ./go_scrap2 .

EXPOSE 8000

CMD [ "./go_scrap2" ]