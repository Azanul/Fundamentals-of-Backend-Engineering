## Request
```
curl -X POST -F "myFile=@Test.png" http://localhost:8000/upload
```

## Response
```
Successfully Uploaded File
```

## Server log
```
File Upload Endpoint Hit
Uploaded File: Test.png
File Size: 40404
MIME Header: map[Content-Disposition:[form-data; name="myFile"; filename="Test.png"] Content-Type:[image/png]]
```