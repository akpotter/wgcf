package wireguard

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

var profileTemplate = `[Interface]
PrivateKey = {{ .PrivateKey }}
DNS = 1.1.1.1
Address = {{ .Address1 }}
Address = {{ .Address2 }}
[Peer]
PublicKey = {{ .PublicKey }}
AllowedIPs = 0.0.0.0/0
AllowedIPs = ::/0
Endpoint = {{ .Endpoint }}`

type Profile struct {
	profileString string
}

type ProfileData struct {
	PrivateKey string
	Address1   string
	Address2   string
	PublicKey  string
	Endpoint   string
}

func NewProfile(data *ProfileData) (*Profile, error) {
	profileString, err := generateProfile(data)
	if err != nil {
		return nil, err
	}
	return &Profile{profileString: profileString}, nil
}

func generateProfile(data *ProfileData) (string, error) {
	t, err := template.New("").Parse(profileTemplate)
	if err != nil {
		return "", err
	}
	var result bytes.Buffer
	if err := t.Execute(&result, data); err != nil {
		return "", err
	}
	return result.String(), nil
}

func (p *Profile) Save(profileFile string) error {
	return ioutil.WriteFile(profileFile, []byte(p.profileString), 0600)
}
