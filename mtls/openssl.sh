#!/bin/bash
openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -x509 -days 3650 -out ca.crt -subj /C=CN/ST=Texas/O="Denicolafamily"/CN="Denicolafamily Root"
openssl genrsa -out server.key 2048
openssl req -new -nodes -key server.key -out server.csr -subj /C=CN/ST=Texas/L=FlowerMound/O="Denicolafamily Server"/CN=localhost
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
openssl genrsa -out client.key 2048
openssl req -new -nodes -key client.key -out client.csr -subj /C=CN/ST=Texas/L=FlowerMound/O="Denicolafamily Client"/CN=brian@bjdcsa.cloud
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt
