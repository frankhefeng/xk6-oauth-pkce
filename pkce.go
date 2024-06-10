package pkce

import (
	"crypto"
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"

	k6common "go.k6.io/k6/js/common"
	k6modules "go.k6.io/k6/js/modules"
)

type CodeChallengeMethod string

const (
	CodeChallengePlain CodeChallengeMethod = "PLAIN"
	CodeChallengeS256                      = "S256"
)

const CodeVerifierLength = 64 // 43~128 https://datatracker.ietf.org/doc/html/rfc7636#section-4.1

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// OAuthPkceModule is the entrypoint into the OAuth PKCE module.
	OAuthPkceModule struct {
		vu      k6modules.VU
		verifer CodeVerifier
	}

	// ModuleInstance represents an instance of the OAuth PKCE module.
	ModuleInstance struct {
		mod *OAuthPkceModule
	}

	CodeVerifier struct {
		Verifier        string
		ChallengeMethod CodeChallengeMethod
		Challenge       string
	}
)

var (
	_ k6modules.Module   = &RootModule{}
	_ k6modules.Instance = &ModuleInstance{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the k6modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu k6modules.VU) k6modules.Instance {
	return &ModuleInstance{
		mod: &OAuthPkceModule{
			vu: vu,
		},
	}
}

// Exports returns the exports of the OAuth PKCE module so that it can be used in test scripts.
func (mi *ModuleInstance) Exports() k6modules.Exports {
	return k6modules.Exports{Default: mi.mod}
}

func (m *OAuthPkceModule) Create(method CodeChallengeMethod) *CodeVerifier {
	v := CodeVerifier{}
	v.Verifier = m.generateCodeVerifier(CodeVerifierLength)
	v.ChallengeMethod = method

	if method == CodeChallengeS256 {
		digest := crypto.SHA256.New()
		digest.Write([]byte(v.Verifier))
		s := digest.Sum(nil)
		sEnc := base64.StdEncoding.EncodeToString(s)
		parts := strings.Split(sEnc, "=")
		sEnc = parts[0]
		r := strings.NewReplacer("+", "-", "/", "_")
		v.Challenge = r.Replace(sEnc)
	} else if method == CodeChallengePlain {
		v.Challenge = v.Verifier
	} else {
		k6common.Throw(m.vu.Runtime(),
			errors.New("Unsupported code challenge method"))
	}

	return &v
}

func (m *OAuthPkceModule) generateCodeVerifier(length int) string {
	charset := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	charsetLength := len(charset)
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[rand.Intn(charsetLength)]
	}
	return string(b)
}

func init() {
	k6modules.Register("k6/x/oauth-pkce", New())
}
