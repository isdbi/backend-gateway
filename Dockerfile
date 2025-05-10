# Specifies a parent image
FROM golang:1.24.1
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod tidy
 


# Builds your app with optional configuration
RUN go build ./cmd/api
 
# Tells Docker which network port your container listens on
EXPOSE 8080



RUN chmod a+x ./api
# Specifies the executable command that runs when the container starts
CMD [ "./api" ]
