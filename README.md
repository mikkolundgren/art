# Create ascii from picture

## Running locally

- install Azure functions core tools
- build go binary ```go build handler.go```
- run azure function locally ```func start```
- upload picture file to function ```curl -F "data=@teemu.jpg"  "http://localhost:7071/api/AsciiArt"```
