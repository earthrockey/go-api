FROM golang

WORKDIR /go/src/github.com/earthrockey
COPY . .
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]