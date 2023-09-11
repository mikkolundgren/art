# Create ascii from picture

Upload a picture to the service, returns ascii representation. Should support all picture formats.

Azure custom handler function.

## Deploy to Azure

- build binary file. if not for windows, change the executable path in ```host.json``` accordingly. ```GOOS=windows GOARCH=amd64 go build handler.go```
- login to Azure ```az login```
- create Azure function app 
- deploy function to cloud ```func azure functionapp publish ascii-art```

## Running locally

- install Azure functions core tools
- build go binary for windows runtime ```go build handler.go```
- run azure function locally ```func start```
- upload picture file to function ```curl -F "data=@my-image.jpg"  "http://localhost:7071/api/asciiart"```

## More info
https://learn.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers

https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other
