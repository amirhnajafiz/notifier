# build from ubuntu image
FROM ubuntu

# switch to root
USER root

# updating apt repositories
RUN apt update

# installing golang
RUN apt install golang-go

# main working directory
WORKDIR /app

# copy all file from current directory
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# copy all of the files
COPY . ./

# building our file
RUN go build -o /notifier

# running execute file
CMD ["/notifier"]
