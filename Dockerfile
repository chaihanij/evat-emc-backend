#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.19 as builder

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV XDG_RUNTIME_DIR=/tmp


# Update OS package and install Git
RUN apt-get update && apt-get -qy install locales tzdata wkhtmltopdf \
    && sed -i 's/# th_/th_/' /etc/locale.gen \
    && locale-gen \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime


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
ADD resources/template/index.html /var/app/template/index.html
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

RUN apt update && apt-get install -y curl grep sed dpkg tini tzdata wkhtmltopdf locales tzdata \
    && sed -i 's/# th_/th_/' /etc/locale.gen \
    && locale-gen \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && apt-get clean

ENV XDG_RUNTIME_DIR=/tmp

WORKDIR /app
COPY --from=builder /var/app/template/index.html   /var/app/template/index.html
COPY --from=builder /go/src/evat/evat-emc-backend  /app.bin


EXPOSE 8080

# # # Set Docker's entry point commands
ENTRYPOINT ["/usr/bin/tini","--","/app.bin"]
# CMD ["/app/evat-emc-backend"]
# CMD [ "/app.bin"]
