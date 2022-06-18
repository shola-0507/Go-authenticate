FROM golang:1.18.3 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -o /go-auth-api

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /go-auth-api /go-auth-api

EXPOSE 8080

ENTRYPOINT [ "/go-auth-api" ]