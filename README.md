# xk6-oauth-pkce
`xk6-oauth-pkce` is a [k6](https://go.k6.io/k6) [extension](https://github.com/grafana/xk6) provides capability to generate [OAuth PKCE](https://datatracker.ietf.org/doc/html/rfc7636) code verifier and code challenge.

## Build
To build a `k6` binary with this extension, install [xk6](https://github.com/grafana/xk6) and build your custom k6 binary with the this extension:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/frankhefeng/xk6-oauth-pkce@latest
  ```

# example

```javascript
import pkce from 'k6/x/oauth-pkce'

export default function () {
    const verifier = pkce.create("S256")
    console.log(verifier)
    console.log(verifier.verifier)
    console.log(verifier.challenge)
}
```
