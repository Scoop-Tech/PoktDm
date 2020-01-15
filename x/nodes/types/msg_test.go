package types

import (
	sdk "github.com/pokt-network/posmint/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"math/rand"
	"reflect"
	"testing"
)

func TestMsgBeginUnstake_GetSignBytes(t *testing.T) {
	type fields struct {
		Address sdk.Address
	}
	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	mesg := MsgBeginUnstake{
		Address: va,
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"Test GetSignBytes", fields{va}, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(mesg))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgBeginUnstake{
				Address: tt.fields.Address,
			}
			if got := msg.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestMsgBeginUnstake_GetSigners(t *testing.T) {
	type fields struct {
		Address sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	mesg := MsgBeginUnstake{
		Address: va,
	}

	tests := []struct {
		name   string
		fields fields
		want   []sdk.Address
	}{
		{"Test GetSigners", fields{va}, []sdk.Address{sdk.Address(mesg.Address)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgBeginUnstake{
				Address: tt.fields.Address,
			}
			if got := msg.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgBeginUnstake_Route(t *testing.T) {
	type fields struct {
		Address sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Route", fields{va}, ModuleName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgBeginUnstake{
				Address: tt.fields.Address,
			}
			if got := msg.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgBeginUnstake_Type(t *testing.T) {
	type fields struct {
		Address sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Type", fields{va}, "begin_unstaking_validator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgBeginUnstake{
				Address: tt.fields.Address,
			}
			if got := msg.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgBeginUnstake_ValidateBasic(t *testing.T) {
	type fields struct {
		Address sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   sdk.Error
	}{
		{"Test Validate Basic error", fields{nil}, sdk.NewError(codespace, CodeInvalidInput, "validator address is nil")},
		{"Test Validate Basic pass", fields{va}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgBeginUnstake{
				Address: tt.fields.Address,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgSend_GetSignBytes(t *testing.T) {
	type fields struct {
		FromAddress sdk.Address
		ToAddress   sdk.Address
		Amount      sdk.Int
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	rand.Read(pub[:])
	va2 := sdk.Address(pub.Address())

	mesg := MsgSend{
		FromAddress: va,
		ToAddress:   va2,
		Amount:      sdk.OneInt(),
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"Test GetSignBytes", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(mesg))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgSend{
				FromAddress: tt.fields.FromAddress,
				ToAddress:   tt.fields.ToAddress,
				Amount:      tt.fields.Amount,
			}
			if got := msg.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgSend_GetSigners(t *testing.T) {
	type fields struct {
		FromAddress sdk.Address
		ToAddress   sdk.Address
		Amount      sdk.Int
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	rand.Read(pub[:])
	va2 := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   []sdk.Address
	}{
		{"Test GetSigners", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, []sdk.Address{sdk.Address(va)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgSend{
				FromAddress: tt.fields.FromAddress,
				ToAddress:   tt.fields.ToAddress,
				Amount:      tt.fields.Amount,
			}
			if got := msg.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgSend_Route(t *testing.T) {
	type fields struct {
		FromAddress sdk.Address
		ToAddress   sdk.Address
		Amount      sdk.Int
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	rand.Read(pub[:])
	va2 := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Route", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, ModuleName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgSend{
				FromAddress: tt.fields.FromAddress,
				ToAddress:   tt.fields.ToAddress,
				Amount:      tt.fields.Amount,
			}
			if got := msg.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgSend_Type(t *testing.T) {
	type fields struct {
		FromAddress sdk.Address
		ToAddress   sdk.Address
		Amount      sdk.Int
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	rand.Read(pub[:])
	va2 := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Type", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, "send"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgSend{
				FromAddress: tt.fields.FromAddress,
				ToAddress:   tt.fields.ToAddress,
				Amount:      tt.fields.Amount,
			}
			if got := msg.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgSend_ValidateBasic(t *testing.T) {
	type fields struct {
		FromAddress sdk.Address
		ToAddress   sdk.Address
		Amount      sdk.Int
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	rand.Read(pub[:])
	va2 := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   sdk.Error
	}{
		{"Test ValidateBasic ok", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, nil},
		{"Test ValidateBasic empty from", fields{
			FromAddress: nil,
			ToAddress:   va2,
			Amount:      sdk.OneInt(),
		}, ErrBadValidatorAddr(DefaultCodespace)},
		{"Test ValidateBasic empty to", fields{
			FromAddress: va,
			ToAddress:   nil,
			Amount:      sdk.OneInt(),
		}, ErrBadValidatorAddr(DefaultCodespace)},
		{"Test ValidateBasic bad amount", fields{
			FromAddress: va,
			ToAddress:   va2,
			Amount:      sdk.NewInt(-1),
		}, ErrBadSendAmount(DefaultCodespace)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgSend{
				FromAddress: tt.fields.FromAddress,
				ToAddress:   tt.fields.ToAddress,
				Amount:      tt.fields.Amount,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgStake_GetSignBytes(t *testing.T) {
	type fields struct {
		Address    sdk.Address
		PubKey     crypto.PubKey
		Chains     []string
		Value      sdk.Int
		ServiceURL string
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	chains := []string{"b60d7bdd334cd3768d43f14a05c7fe7e886ba5bcb77e1064530052fed1a3f145"}
	value := sdk.OneInt()
	surl := "www.pokt.network"

	mesg := MsgStake{
		Address:    va,
		PubKey:     pub,
		Chains:     chains,
		Value:      value,
		ServiceURL: surl,
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"Test SignBytes", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(mesg))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgStake{
				Address:    tt.fields.Address,
				PubKey:     tt.fields.PubKey,
				Chains:     tt.fields.Chains,
				Value:      tt.fields.Value,
				ServiceURL: tt.fields.ServiceURL,
			}
			if got := msg.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgStake_GetSigners(t *testing.T) {
	type fields struct {
		Address    sdk.Address
		PubKey     crypto.PubKey
		Chains     []string
		Value      sdk.Int
		ServiceURL string
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	chains := []string{"b60d7bdd334cd3768d43f14a05c7fe7e886ba5bcb77e1064530052fed1a3f145"}
	value := sdk.OneInt()
	surl := "www.pokt.network"

	tests := []struct {
		name   string
		fields fields
		want   []sdk.Address
	}{
		{"Test GetSigners", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, []sdk.Address{sdk.Address(va)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgStake{
				Address:    tt.fields.Address,
				PubKey:     tt.fields.PubKey,
				Chains:     tt.fields.Chains,
				Value:      tt.fields.Value,
				ServiceURL: tt.fields.ServiceURL,
			}
			if got := msg.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgStake_Route(t *testing.T) {
	type fields struct {
		Address    sdk.Address
		PubKey     crypto.PubKey
		Chains     []string
		Value      sdk.Int
		ServiceURL string
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	chains := []string{"b60d7bdd334cd3768d43f14a05c7fe7e886ba5bcb77e1064530052fed1a3f145"}
	value := sdk.OneInt()
	surl := "www.pokt.network"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Route", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, RouterKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgStake{
				Address:    tt.fields.Address,
				PubKey:     tt.fields.PubKey,
				Chains:     tt.fields.Chains,
				Value:      tt.fields.Value,
				ServiceURL: tt.fields.ServiceURL,
			}
			if got := msg.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgStake_Type(t *testing.T) {
	type fields struct {
		Address    sdk.Address
		PubKey     crypto.PubKey
		Chains     []string
		Value      sdk.Int
		ServiceURL string
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	chains := []string{"b60d7bdd334cd3768d43f14a05c7fe7e886ba5bcb77e1064530052fed1a3f145"}
	value := sdk.OneInt()
	surl := "www.pokt.network"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Type", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, "stake_validator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgStake{
				Address:    tt.fields.Address,
				PubKey:     tt.fields.PubKey,
				Chains:     tt.fields.Chains,
				Value:      tt.fields.Value,
				ServiceURL: tt.fields.ServiceURL,
			}
			if got := msg.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgStake_ValidateBasic(t *testing.T) {
	type fields struct {
		Address    sdk.Address
		PubKey     crypto.PubKey
		Chains     []string
		Value      sdk.Int
		ServiceURL string
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())
	chains := []string{"b60d7bdd334cd3768d43f14a05c7fe7e886ba5bcb77e1064530052fed1a3f145"}
	value := sdk.OneInt()
	surl := "www.pokt.network"

	tests := []struct {
		name   string
		fields fields
		want   sdk.Error
	}{
		{"Test Validate Basic ok", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, nil},
		{"Test Validate Basic bad address", fields{
			Address:    nil,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: surl,
		}, ErrNilValidatorAddr(DefaultCodespace)},
		{"Test Validate Basic bad value", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      sdk.NewInt(-1),
			ServiceURL: surl,
		}, ErrBadDelegationAmount(DefaultCodespace)},
		{"Test Validate Basic bad Chains", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     []string{},
			Value:      value,
			ServiceURL: surl,
		}, ErrNoChains(DefaultCodespace)},
		{"Test Validate Basic bad chain in Chains", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     []string{""},
			Value:      value,
			ServiceURL: surl,
		}, ErrNoChains(DefaultCodespace)},
		{"Test Validate Basic bad serviceURL", fields{
			Address:    va,
			PubKey:     pub,
			Chains:     chains,
			Value:      value,
			ServiceURL: "",
		}, ErrNoServiceURL(DefaultCodespace)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgStake{
				Address:    tt.fields.Address,
				PubKey:     tt.fields.PubKey,
				Chains:     tt.fields.Chains,
				Value:      tt.fields.Value,
				ServiceURL: tt.fields.ServiceURL,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUnjail_GetSignBytes(t *testing.T) {
	type fields struct {
		ValidatorAddr sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	mesg := MsgUnjail{
		ValidatorAddr: va,
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"Test GetSignBytes", fields{va}, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(mesg))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUnjail{
				ValidatorAddr: tt.fields.ValidatorAddr,
			}
			if got := msg.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUnjail_GetSigners(t *testing.T) {
	type fields struct {
		ValidatorAddr sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   []sdk.Address
	}{
		{"Test GetSigners", fields{va}, []sdk.Address{sdk.Address(va)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUnjail{
				ValidatorAddr: tt.fields.ValidatorAddr,
			}
			if got := msg.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUnjail_Route(t *testing.T) {
	type fields struct {
		ValidatorAddr sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Route", fields{va}, ModuleName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUnjail{
				ValidatorAddr: tt.fields.ValidatorAddr,
			}
			if got := msg.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUnjail_Type(t *testing.T) {
	type fields struct {
		ValidatorAddr sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Test Type", fields{va}, "unjail"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUnjail{
				ValidatorAddr: tt.fields.ValidatorAddr,
			}
			if got := msg.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUnjail_ValidateBasic(t *testing.T) {
	type fields struct {
		ValidatorAddr sdk.Address
	}

	var pub ed25519.PubKeyEd25519
	rand.Read(pub[:])
	va := sdk.Address(pub.Address())

	tests := []struct {
		name   string
		fields fields
		want   sdk.Error
	}{
		{"Test ValidateBasic OK", fields{va}, nil},
		{"Test ValidateBasic bad address", fields{nil}, ErrBadValidatorAddr(DefaultCodespace)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUnjail{
				ValidatorAddr: tt.fields.ValidatorAddr,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
