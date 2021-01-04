FROM golang:1.15 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.15 as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /myapp
ADD . /myapp
WORKDIR /myapp

#Добавляем нового пользователя
RUN useradd -u 1001 myapp


#Собираем бинарный файл
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build .

FROM scratch

#Копируем зависимости нового пользователя
COPY --from=builder /etc/passwd /etc/passwd
USER myapp

COPY --from=builder /myapp /myapp
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
WORKDIR /myapp

CMD ["./movies"]
