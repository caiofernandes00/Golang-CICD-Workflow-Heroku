# Removing previous pem files
rm *.pem

# Generate Certificate simulating a CA
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=CA/emailAddress=test-ca@email.com"

# Generating a certificate for the server ap
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Server App/emailAddress=test-server@email.com"

# Using the CA private key to sign web server's CSR
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile .\server-ext.cnf

# Validate certificates
openssl verify -CAfile .\ca-cert.pem .\server-cert.pem

# Reading content from a certificate
# openssl x509 -in .\server-cert.pem -noout -text
