
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/nse_scrapper
COPY . .
RUN apk add --no-cache git

ENV GOBIN=/usr/local/go/bin

# RUN exit 1
RUN go get -d -v ./...; exit 0
RUN go install -v ./...; exit 0




#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/local/go/bin/nse_scrapper /nse_scrapper
COPY --from=builder /go/src/nse_scrapper/env /env
ENTRYPOINT ./nse_scrapper
LABEL Name=nse_scrapper Version=0.0.1
EXPOSE 3000
