FROM golang:1.19
WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
COPY . ./
RUN go mod tidy
WORKDIR /app/apps/api
RUN go build -o /server
CMD ["/server"]