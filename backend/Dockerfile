FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /decidr-app-api ./cmd/api

FROM alpine:latest 

WORKDIR /

COPY --from=build /decidr-app-api /decidr-app-alpine

EXPOSE 8080

USER daemon:daemon

ARG MONGO_URL
ENV MONGO_URL=$MONGO_URL

ENTRYPOINT [ "/decidr-app-alpine" ]

