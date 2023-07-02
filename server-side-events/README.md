## Client code
```
let sse = new EventSource("http://localhost:8080/stream");
sse.onmessage = console.log
```