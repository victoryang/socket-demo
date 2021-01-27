# Traefik

## Traefik Install
```bash
#!/bin/bash

helm install traefik traefik/traefik

kubectl patch service traefik -p '{"spec":{"type":"NodePort"}}'
```

## Backend

**build image**
docker build -t test/test_traefik:v1 .

```bash
FROM alpine

ENV PATH /usr/bin:${PATH}

EXPOSE 9000

COPY app /usr/bin/

CMD ["/usr/bin/app"]
```

```yaml
# Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
   name: test-traefik
   namespace: default
spec:
   replicas: 1
   selector:
      matchLabels:
          app: myapp
   template:
      metadata:
          labels:
             app: myapp
      spec:
          containers:
          - name: test-traefik
            image: test/test_traefik:v1
            ports:
            - name: http
              containerPort: 9000

# Service
apiVersion: v1
kind: Service
metadata:
    name: test-traefik
    labels:
        app: myapp
spec:
    type: ClusterIP
    ports:
        - port: 9000
        targetPort: 9000
    selector:
        app: myapp

# Ingress
apiVersion: "networking.k8s.io/v1beta1"
kind: "Ingress"
metadata:
    name: "test-ingress"
    annotations:
          kubernetes.io/ingress.class: traefik
spec:
    rules:
        - http:
             paths:
             - path: /
               backend:
                    serviceName: test-traefik
                    servicePort: 9000
```