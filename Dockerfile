FROM golang:1.19 as build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o playlist

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/playlist /
CMD ["/playlist"]