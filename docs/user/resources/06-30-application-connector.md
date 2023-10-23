# CompassConnection

The `applicationconnectors.operator.kyma-project.io` CustomResourceDefinition (CRD)
is a detailed description of the kind of data and the format used to preserve
the configuration and status of the Application Connector Module .
The `ApplicationConnector` custom resource (CR) contains the state and statuses of the module installed on the cluster.

To get the up-to-date CRD and show the output in the `yaml` format, run this command:

```bash
kubectl get crd applicationconnectors.operator.kyma-project.io -o yaml
```
## Sample custom resource

```
apiVersion: operator.kyma-project.io/v1alpha1
kind: ApplicationConnector
metadata:
  namespace: kyma-system
  labels:
    app.kubernetes.io/name: applicationconnector
    app.kubernetes.io/instance: applicationconnector-sample
  name: applicationconnector-sample
spec:
  domainName: "my.domain"
  appConnValidator:
    logLevel: "info"
    logFormat: "json"
  appGateway:
    proxyTimeout: "10s"
    requestTimeout: "10s"
    logLevel: "info"
  runtimeAgent:
    controllerSyncPeriod: "180s"
    minimalConfigSyncTime: "15s"
    certValidityRenewalThreshold: "0.3"
```
