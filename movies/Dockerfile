FROM golang:1.15

RUN mkdir -p /myapp
ADD . /myapp
WORKDIR /myapp

RUN go build -o /myapp/movies.exe .
RUN chmod 777 /myapp/movies.exe

CMD ["./movies.exe"]