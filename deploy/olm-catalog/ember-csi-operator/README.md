### OLM Catalog For Ember CSI

## Deployment

```
kubectl create -f 00-pre.yaml 
kubectl create -f 01-scc.yaml # Only required for OpenShift deployments
kubectl create -f ember-csi-operator.crd.yaml
kubectl create -f 0.0.1/ember-csi-operator.v0.0.1.clusterserviceversion.yaml
```
