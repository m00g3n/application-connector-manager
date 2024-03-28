# Security

## Client Certificates

To provide maximum security, in the [Compass mode](README.md), Application Connector uses the mutual TLS protocol with Client Authentication enabled. As a result, anyone attempting to connect to Application Connector must present a valid client certificate, which is dedicated to a specific Application. In this way, the traffic is fully encrypted, and the client has a valid identity.

### TLS Certificate Verification for External Systems

By default, the TLS certificate verification is enabled when sending data and requests to every application.
You can [disable the TLS certificate verification](tutorials/01-50-disable-tls-certificate-verification.md) in the communication between Kyma and an application to allow Kyma to send requests and data to an unsecured application. Disabling the certificate verification can be useful in certain testing scenarios.
