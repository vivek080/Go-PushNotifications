
# Go-PushNotifications

This Project is for sending push Notifications to user devices using Golang


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file and add the serviceAccountKey file in the directory 


`GOOGLE_APPLICATION_CREDENTIALS = "serviceAccountKey.json"`


## Deployment

To deploy this project run

```bash
  go run main.go
```


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