# testtools-container
Application for testing environments from a container.

Its possible to run TCP connect tests, dns lookup test and http GET request tests.

To deploy a container running tests and a complete monitoring solution with grafana and prometheus see helm chart in `chart` folder.

Tests are defined in the `chart/testtools-container/values.yaml`.

Example config:
```
connectTests:
  google:
    hostname: google.com
    port: 443
    timeout: 30s
    interval: 5s
  github:
    hostname: github.com
    port: 443
    timeout: 30s
    interval: 5s
dnsTests:
  google:
    hostname: google.com
    interval: 1s
  github:
    hostname: github.com
    interval: 1s
apiTests:
  google:
    uri: https://www.google.com
    timeout: 30s
    interval: 5s
  github:
    uri: https://api.github.com/users/octocat/orgs
    timeout: 30s
    interval: 5s
```
