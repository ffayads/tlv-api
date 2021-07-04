FROM golang
LABEL maintainer="felixfayad@paymentez.com"
# Set up app
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/api/tlv-api/
# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
ENV GO111MODULE=auto
COPY . .
RUN go build -o tlv-api
EXPOSE 80
CMD ["./tlv-api"]
