# Evve: Event Easier With EVVE - Events Ticketing API

## About
Evve is a web-based application that operates in the field of event ticketing and online event organizing. This application is aimed at buyers who want to buy digital tickets for an event and is also aimed at partners, especially event organizers who want to market their event tickets online.

This application is integrated with the 3rd party payment gateway, namely Midtrans

## Tech Stack
- Go
- Echo Framework
- MySQL
- GORM
- Docker
- GCP

## Entity Relationship Diagram
<img src="assets\images\erd capstone project.drawio.png" width= 600>

## Installation
1. Clone:

```
git clone https://github.com/Capstone-Project-Team-2/Backend-Immersive.git
```
2. Go to the Backend-Immersive directory
```
cd Backend-Immersive
```

3. Jalankan perintah berikut
- Enter the package name you want to <b>package-name</b>
```
go mod init package-name
go mod tidy
```

4. Create a file in .env format (local.env for local development)

5. Write as follows in the .env file. Adjust to your needs
```
export JWT_KEY = 'your-jwt-key'
export DBUSER = 'your-db-username'
export DBPASS = 'your-db-password'
export DBHOST = 'your-db-host'
export DBPORT = 'your-db-port'
export DBNAME = 'your-db-name'
export KEY_SERVER_MIDTRANS = 'your-midtrans-server-key'
export KEY_CLIENT_MIDTRANS = 'your-midtrans-client-key'
```
Midtrans Server and Client keys are found on your Midtrans Dashboard. For reference, please check [Midtrans](https://support.midtrans.com/hc/id)

6. For file upload purposes, create and save <i>Google Application Credentials</i> in a file with the name keys.json. For references to <i>Google Application Credentials</i>, please check [reference](https://adityarama1210.medium.com/simple-golang-api-uploader-using-google-cloud-storage-3d5e45df74a5) or [reference](https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go)


## API Documentation
This API documentation can be viewed on [SwaggerHub](https://app.swaggerhub.com/apis-docs/Capstone-BE18/Capstone-Project-Tickets/1.0.0).

## Collaborator
- Farihatul Ilmiyya - [Github](https://github.com/Farihatul-ilmiyya)
- Mohammad Hadi Hamdam - [Github](https://github.com/Hadi1Hamdam)
