# Start from the latest golang base image
FROM golang:latest

# Install application
RUN go get github.com/vanigabriel/Projeto-Pismo
CMD ["pwd"]
CMD ["ls"]


# Change to ROOT directory
CMD ["pwd"]
WORKDIR /app
CMD ["pwd"]

# Copy .env file
COPY .env .
RUN cp /go/bin/Projeto-Pismo .


# Command to run the executable
CMD ["./Projeto-Pismo"]
