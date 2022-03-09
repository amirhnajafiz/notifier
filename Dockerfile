# build from ubuntu image
FROM golang:alpine

# switch to root
USER root

# main working directory
WORKDIR /app

# copy all file from current directory
COPY go.mod go.sum ./

# copy all of the files
COPY . .

# change work dir
WORKDIR /app/cmd

# building our file
RUN go build -o ./notifier

# change WD
WORKDIR /app

# move the exe file
RUN mv ./cmd/notifier ./notifier

# running execute file
CMD ["./notifier"]
