package types

type GrantType int

const (
	ClientCredentialsGrant GrantType = iota
	DeviceCodeGrant
)

func (gt GrantType) String() string {
	switch gt {
	case ClientCredentialsGrant:
		return "client_credentials"
	case DeviceCodeGrant:
		return "device_code"
	}

	return "unknown"
}