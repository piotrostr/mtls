#!/bin/bash

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650

