# The base go-image
FROM golang:1.19
 
# Create a directory for the app
RUN mkdir /build
# Copy all files from the current directory to the app directory
COPY . /build
WORKDIR /build

# Copy all files from the current directory to the app directory
COPY . /build/main

# RUN go env -w GO111MODULE=auto
RUN go mod init simpleserver
RUN go mod tidy

# Run command as described:
# go build will build an executable file named server in the current directory
RUN cd main && go build -o simpleserver . 

EXPOSE 5050
 
# Run the server executable
CMD [ "/build/simpleserver" ]