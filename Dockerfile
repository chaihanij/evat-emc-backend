#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.19 as builder

# Force the go compiler to use modules
ENV GO111MODULE=on


# Update OS package and install Git
RUN apt-get update && apt-get -qy install netcat

# Set working directory
WORKDIR /go/src/evat

# Setup github credential
ADD resources/keys/id-rsa /root/.ssh/id_rsa
ADD resources/keys/id-rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan gitlab.com >> /root/.ssh/known_hosts
RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com"

# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
ADD Makefile Makefile

RUN go mod download

# Install air for local development
RUN go install github.com/cosmtrek/air@latest

# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/AlekSi/gocov-xml

# Set Docker's entry point commands
RUN go build -o /go/src/evat/evat-emc-backend app/main.go

# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------

# #App stage
FROM golang:1.19

RUN apt update && apt-get install -y curl grep sed dpkg tini tzdata && \
    apt-get clean


RUN groupadd -g 211000 appgroup && useradd -u 211000 -g 211000 -G appgroup appuser

# # Set working directory
WORKDIR /app/


# # #Get artifact from buiber stage
COPY --from=builder /go/src/evat/evat-emc-backend  /app.bin
RUN 
# # # Set Docker's entry point commands
RUN chown -R appuser:appgroup /app && chmod +x /app.bin
USER appuser

EXPOSE 8080

# # # Set Docker's entry point commands
# ENTRYPOINT ["/usr/bin/tini","--","/app/evat-emc-backend"]
# CMD ["/app/evat-emc-backend"]
CMD [ "/app.bin"]