# CloudSQL

## Configuration
In order to configure CloudSQL three pem files are required: client-key.pem, client-cert.pem and server-ca.pem.
These added as secrets to your kubernetes namespace. This can be done with a variation of the following command:

<pre>kubectl create secret generic <b>SECRET-NAME</b> --from-file=client-cert.pem --from-file=client-key.pem --from-file=server-ca.pem --namespace=<b>YOUR-NAMEPSACE</b></pre>

Now, in microservice.yaml **SECRET-NAME** must be added under **secrets**. Example:

<pre>secrets:
  - name: <b>SECRET-NAME</b>
    mountPath: /etc/<b>SECRET-NAME</b></pre>

Next, the newly mounted path should be set as a pod env variable. Example:

<pre>podEnv:
  - key: CLOUD_SQL_CERTS_PATH
    value: /etc/<b>SECRET-NAME</b></pre>

Finally, you must specify all the information needed to connect to the database in your podEnv:

- CLOUD_SQL_IP - The IP address of the CloudSQL instance
- CLOUD_SQL_INSTANCE_NAME - The name of the instance
- CLOUD_SQL_USERNAME - The CloudSQL username
- CLOUD_SQL_PASSWORD - The CloudSQL password

## Initializing
The client can be initialized by calling the cloudsqlclient.Initialize command.
This requires you to get the certs first.

Example (You should check for errors):
```
certs, _ := cloudsqlclient.GetCerts()
client, _ := cloudsqlclient.Initialize(
    "repcore-prod",
    cloudsqlclient.InstanceName(),
    "secondary-index-name",
    cloudsqlclient.IP(),
    cloudsqlclient.Username(),
    cloudsqlclient.Password(),
    certs.ClientCert,
    certs.ClientKey,
    certs.ServerCA,
)
```
