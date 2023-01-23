#!/bin/bash
set -e
export RSA_PUBLIC_KEY=`cat resources/keys/id-rsa.pub` 
export RSA_PRIVATE_KEY=`cat resources/keys/id-rsa`
echo $RSA_PUBLIC_KEY
DEBUG=True \
MONGODB_URL='mongodb://0.0.0.0:27017' \
MONGODB_NAME=evat_db \
MONGODB_USER=evatuser \
MONGODB_PASS=asdf1234 \
MONGODB_REQUEST_TIMEOUT=60 \
ENCRYPT_KEY=b054eb59dcf46db5da45ade306d005a2473def4a51f0ee93371bd64e77ae4b20 \
JWT_TOKEN_LIFE=0 \
# air
/usr/local/bin/go run app/cmd/main.go
