
# Go-PushNotifications

This Project is for sending push Notifications to user devices using Golang


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file and add the serviceAccountKey file in the directory 


`GOOGLE_APPLICATION_CREDENTIALS = "serviceAccountKey.json"`


## Local Deployment

Install Docker before runing this project from `https://docs.docker.com/engine/install/ubuntu/`

To run this project run

```bash
  docker-compose up
```

enter `localhost:5000/pushNotification` endpoint in postman to use the REST API.


## Documentation

[Documentation](https://linktodocumentation)

POST Method

{
    "data": {
        "type": "notification",
        "attributes": {
            "title": "",
            "body": "",
            "tokens": [
                "device Token"
            ]
        }
    }
}