## Hello Logs!

Basic golang application printing logs!

### Sample Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-logs
spec:
  replicas: 10
  selector:
    matchLabels:
      app: hello-logs
  template:
    metadata:
      labels:
        app: hello-logs
    spec:
      containers:
        - name: hello-logs
          image: alessandroargentieri/hello-logs:latest
          env:
            - name: MSG
              value: "Hello Talos!"
            - name: RANDOM
              value: "false"
```
