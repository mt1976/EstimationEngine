#!/bin/bash
openssl req -out sslcert.csr -newkey rsa:2048 -nodes -keyout private.key -config san.cnf
openssl req -noout -text -in sslcert.csr | grep DNS
openssl  x509  -req  -days 365  -in sslcert.csr  -signkey private.key  -out sslcert.crt

