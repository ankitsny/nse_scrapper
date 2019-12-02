
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/nse_scrapper
COPY . .
RUN apk add --no-cache git
RUN apk add --update nodejs npm

# Build the view
WORKDIR /go/src/nse_scrapper/nse_view
RUN ls -al
RUN npm install --production
RUN npm run build:prod

WORKDIR /go/src/nse_scrapper

ENV GOBIN=/usr/local/go/bin

# RUN exit 1
RUN go get -d -v ./...; exit 0
RUN go install -v ./...; exit 0




#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/local/go/bin/nse_scrapper /nse_scrapper

# TODO: Mount the env file
COPY --from=builder /go/src/nse_scrapper/env /env

# Copy static files
COPY --from=builder /go/src/nse_scrapper/nse_view/build /nse_view/build


ENTRYPOINT ./nse_scrapper
LABEL Name=nse_scrapper Version=0.0.1
EXPOSE 3000
