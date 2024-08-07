##### Stage 1 #####

FROM golang:1.22-alpine as builder

RUN apk --no-cache add tzdata
# CERT PACKAGES
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true 
        
### Create new directly and set it as working directory
RUN mkdir -p /project
WORKDIR /project

### Copy Go application dependency files
COPY go.mod .
COPY go.sum .

### Download Go application module dependencies
RUN go mod download

### Copy actual source code for building the application
COPY . .

### CGO has to be disabled cross platform builds
### Otherwise the application won't be able to start
ENV CGO_ENABLED=0

### Build the Go app for a linux OS
### 'scratch' and 'alpine' both are Linux distributions

RUN GOOS=linux go build -o app src/cmd/main.go

##### Stage 2 #####

### Define the running image
FROM scratch

### Set working directory
WORKDIR /dist

### Copy built binary application from 'builder' image
COPY --from=builder /project/app .
COPY --from=builder /project/deploy/conf/production.env ./production.env
COPY --from=builder /project/deploy/mail_template ./mail_template

# copy the ca-certificate.crt from the build stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENV TZ=Asia/Ho_Chi_Minh

### Run the binary application
CMD ["./app", "-conf=production.env"]