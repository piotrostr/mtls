#!/bin/bash

openssl req -new -x509 -sha256 -keyout ca.key -out ca.crt -days 3650

