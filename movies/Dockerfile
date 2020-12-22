FROM golang:1.15 as builder

RUN mkdir -p /myapp
ADD . /myapp
WORKDIR /myapp


#Собираем бинарный файл
RUN GOOS=windows GOARCH=amd64 CGO_ENABLE=0 \
    go build -o /myapp/movies.exe .

#Предоставляем права все всем
RUN chmod 777 /myapp/movies.exe

FROM scratch

COPY --from=builder /myapp /myapp

CMD ["./movies.exe"]