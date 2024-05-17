FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /feeds

# CMD [ "go", "run", "cmd/main/main.go" ]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /feeds /feeds

EXPOSE 8030

USER 1001:1001

ENTRYPOINT [ "/feeds" ]