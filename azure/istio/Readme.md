## To get services on the istio namespace:
```
   kubectl get svc -n istio-system
```
## To port forward the grafana/kiali dashboard:
   <!-- kubectl port-forward svc/grafana -n istio-system 3000 -->
   <!-- kubectl port-forward svc/kiali -n istio-system 20001 -->
```
    kubectl port-forward svc/grafana -n istio-system [port] 
```
```
   - kubectl port-forward svc/kiali -n istio-system [port] 
```