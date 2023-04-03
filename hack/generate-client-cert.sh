#!/bin/bash

openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr
openssl x509 -req -sha256 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 3650

