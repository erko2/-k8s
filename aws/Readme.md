#kubernetes deployment samples on aws kubernetes cluster. 

helm install app-ingress ingress-nginx/ingress-nginx \
     --namespace ingress-main \
     --set controller.replicaCount=2 \
     --set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
     --set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux