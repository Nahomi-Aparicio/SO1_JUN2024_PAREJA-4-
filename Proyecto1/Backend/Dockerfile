                                                                                   FROM golang

WORKDIR /back

COPY . ./

RUN go mod download

ENV DB_HOST="192.168.122.116"
ENV DB_PORT="27017"
ENV DB_NAME="DB"

EXPOSE 8000

CMD ["go", "run", "Main.go"]
