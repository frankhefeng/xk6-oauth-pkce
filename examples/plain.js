import pkce from 'k6/x/oauth-pkce'

export default function () {
    const verifier = pkce.create("PLAIN")
    console.log(verifier.verifier)
    console.log(verifier.challenge_method)
    console.log(verifier.challenge)
}
