# go_web_project

To run the application you must have docker and docker-compose ready to use.

First step : Execute the docker-compose file to setup the containers

    1. cd ../go_web_project
    2. docker-compose up --build

Second step : My very small API is ready. Follow the links below.

    Index -> http://localhost:8080/

    Acces to every data :
    Companys -> http://localhost:8080/companys

    Acces to one data by id :
    Company -> http://localhost:8080/company/{id}

    Acces to one data by name :
    Company -> http://localhost:8080/company/name/{name}

    Add a predifine data :
    CompanyCreate -> http://localhost:8080/companyCreate

Third step : I added a client/server in my application. The server is run with the docker-compose but the server should be execute localy.

    cd ../client
    go run main.go

The goal of this third part is to simulate a authentification service. I used the RSA authentification to secure the connection beetwen the client and the server.
When you run the client he's going to ask you to write a name. This name is encrypt with the server public key. After that the server ask my api to know if the name exist in the database. Finally the server encrypt the answer for the client.

I work on a small go module to learn more about them : https://github.com/Exypte/go_encryption.git