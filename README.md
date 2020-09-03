# Curso Microserviços - Desafio k8s Golang hpa

## Parte 1

- Código Go pasta src
- Arquivos k8s pasta k8s
- Imagem docker hub: https://hub.docker.com/r/gmeneguz/go-hpa

## Parte 2

- Arquivo k8s/hpa.yaml

# Parte 3

```
kubectl run -it --rm load-generator --image=busybox /bin/sh

while true; do wget -q -O- http://go-hpa-service.default.svc.cluster.local:8000; done;
```

# Desafio Deploy Contínuo

Arquivo Cloudbuild para deploy: cloudbuild.prod.yaml
