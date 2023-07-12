import pkce from 'k6/x/oauth-pkce'

export default function () {
    const verifier = pkce.create("S256")
    console.log(verifier)
    console.log(verifier.verifier)
    console.log(verifier.challenge)
}
