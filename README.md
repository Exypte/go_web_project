# go_web_project

To run the application you must have docker and docker-compose ready to use.

First step : Execute the docker-compose file to setup the containers

    1. cd ../go_web_project
    2. docker-compose up --build

Second step : My very small API is ready. Follow the links below.

    Index -> http://localhost:8080/

    Acces to every data :
    Companys -> http://localhost:8080/companys

    Acces to one data :
    Company -> http://localhost:8080/company/{id}

    Add a predifine data :
    CompanyCreate -> http://localhost:8080/companyCreate