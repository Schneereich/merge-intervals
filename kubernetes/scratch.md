```
kubectl create deployment merge-intervals --image=schneereich/merge-intervals
kubectl create deployment merge-intervals --image=schneereich/merge-intervals --port=80
kubectl get deployments
kubectl get pods
kubectl get events
kubectl config view
kubectl expose deployment merge-intervals --port=80 --target-port=80 --name=merge-intervals 
kubectl expose deployment merge-intervals --type=NodePort --name=merge-intervals 
kubectl get services

kubectl delete deployment merge-intervals
kubectl delete svc merge-intervals
```