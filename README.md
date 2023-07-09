# xk6-oauth-pkce

`xk6-oauth-pkce` is a [k6](https://go.k6.io/k6) [extension](https://github.com/grafana/xk6) provides the capability to generate [OAuth PKCE](https://datatracker.ietf.org/doc/html/rfc7636) code verifier and code challenge.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download [xk6](https://github.com/grafana/xk6):

```bash
$ go install go.k6.io/xk6/cmd/xk6@latest
```

2. [Build the k6 binary](https://github.com/grafana/xk6#command-usage):

```bash
$ xk6 build --with github.com/frankhefeng/xk6-oauth-pkce@latest
```

# Example

```javascript
import pkce from "k6/x/oauth-pkce";

export default function () {
  const verifier = pkce.create("S256");
  console.log(verifier);
  console.log(verifier.verifier);
  console.log(verifier.challenge);
}
```

Result output:

```shell
./k6 run plain.js

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: plain.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] d4pvC50QvNMo9kpvxnWJWIBzst3LZlZVhA4bBhd1wGBwc7Qdvl4YErYx8XzOFHOt  source=console
INFO[0000] "PLAIN"                                       source=console
INFO[0000] d4pvC50QvNMo9kpvxnWJWIBzst3LZlZVhA4bBhd1wGBwc7Qdvl4YErYx8XzOFHOt  source=console

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=171.26µs min=171.26µs med=171.26µs max=171.26µs p(90)=171.26µs p(95)=171.26µs
     iterations...........: 1   3703.703704/s


running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU
```

```shell
./k6 run s256.js

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: s256.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] {"verifier":"35UoJSej40X87vZLreSjUck3QfkjpvAw07aKOgwQY3FMjT2OfRG3G9aMdzkZhMtU","challenge_method":"S256","challenge":"HjKS897Yv8kjD5AX-s0mBsug4EzbjgiVSmUubkHNz58"}  source=console
INFO[0000] 35UoJSej40X87vZLreSjUck3QfkjpvAw07aKOgwQY3FMjT2OfRG3G9aMdzkZhMtU  source=console
INFO[0000] HjKS897Yv8kjD5AX-s0mBsug4EzbjgiVSmUubkHNz58   source=console

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=250.45µs min=250.45µs med=250.45µs max=250.45µs p(90)=250.45µs p(95)=250.45µs
     iterations...........: 1   3003.003003/s


running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU
```
