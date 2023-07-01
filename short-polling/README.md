## Submit job

### Request
```
curl -X POST localhost:8080/submit
```

### Response
```
job:1688232899
```


## Get status

### Request
```
curl localhost:8080/status?jobId=job:1688232899
```

### Pending response
```
job pending
```

### Completed response
```
job completed
```
