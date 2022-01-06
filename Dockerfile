####################################
# STEP 1 build executable binary
####################################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk add ca-certificates
WORKDIR $GOPATH/src/quipper/quiz_master

#copy all the content to container
COPY .. .

##Fix go mod cant download without using proxy
ENV GOPROXY="https://goproxy.cn,direct"

# Build the binary
RUN export CGO_ENABLED=0 && go build -o quiz_master

#change the permission on binary
RUN chmod +x quiz_master

##############################################
# STEP 2 build a small image using alpine:3.14
##############################################
FROM alpine:3.14

# Copy our static executable.
COPY --from=builder /quiz_master ./scheduler

# Run the entrypoints.
ENTRYPOINT [ "./quiz_master" ]