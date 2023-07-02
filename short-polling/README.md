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
curl localhost:8080/status?jobId=job:1688232899 -w  "%{time_starttransfer}\n
```

### Pending response
```
0.001634
job pending
```

### Completed response
```
0.001590
job completed
```
