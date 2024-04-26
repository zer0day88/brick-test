
## README

### Mock API Doc : [mock api documentation](https://apidog.com/apidoc/shared-bb3ed2b0-3782-457e-a566-a7c39534962f)

### List mock endpoint :
1. Inquiry Account - https://mock.apidog.com/m1/520155-0-default/snap/v1.0/account-inquiry
2. Transfer - https://mock.apidog.com/m1/520155-0-default/snap/v1.0/transfer-bank

### Service endpoint
1. POST /v1/pg/inquiry-account\
   This API used to inquiry bank account info.\
   sample request
   ```json
   {
    "beneficiaryBankCode": "111",
    "beneficiaryAccountNumber": "111"
   }
   ```
2. POST /v1/pg/transfer\
   This Api used to transfer amount\
   sample request
   ```json
   {
    "partnerReferenceNo": "f8ebb65f-7db7-46c8-9a0d-a204665528be",
    "amount": {
        "value": "200000.00",
        "currency": "IDR"
    },
    "beneficiaryAccountNumber": "444401000155501",
    "beneficiaryAccountName": "Jaden Malone",
    "beneficiaryBankCode": "001",
    "beneficiaryBankName": "Bank ABC",
    "sourceAccountNumber": "111101000155505"
    
   }
   ```
3. POST /v1/pg/notify-transfer\
   This API is used to notify transfer status from bank to internal system.\
   sample request
   ```json
   {
    "referenceNo": "33b18d06-bf74-4c12-8fd7-3f75515482eb",
    "transactionStatus": "SUCCESS"
   }
   ```

### How to run the projects using docker-compose
```shell
docker-compose build
docker-compose up
```
This command will build the docker image then run the container\
and make the service exposed to port 3000 and database exposed to port 5432

### Elaboration
First of all, This service have a purpose like gateway between client and bank.\
This is just a simple example and flow for the use case.\
I think in real production environment this service should be more advance.\
we can add more security using signature utilizing private and public key,\
then the other improvement maybe such as using event driven, cache and fault handling,etc.\

Okay, back to the core topic, i will elaborate the first endpoint to inquiry account\
so the purpose of the inquiry account is validating the receiver bank account.\
what i do is just make the proxy between bank and client. \
then the second one is transfer, this endpoint should be more complex than the previous one\
we need to add idempotency checking, validating more the data before send or save the transaction\
also we need to optimize the concurrency, btw we can also use event driven here and handle the error\
that we could retry independently.\
And then the last one is the endpoint that waiting for the status from the bank service.\
This endpoint will update the status transfer by ref no.\
So many scenario that i not yet implement here like when account invalid, bank service timeout,etc.


