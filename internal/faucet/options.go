package faucet

const (
	DefaultAppCli        = "provenanced"
	DefaultKeyName       = "faucet"
	DefaultDenom         = "nhash"
	DefaultNodeAddr      = "127.0.0.1"
	DefaultCreditAmount  = 10
	DefaultMaximumCredit = 10000
)

func defaultOptions() *Options {
	return &Options{
		AppCli:       DefaultAppCli,
		KeyName:      DefaultKeyName,
		Denom:        DefaultDenom,
		KeyNodeAddr:  DefaultNodeAddr,
		CreditAmount: DefaultCreditAmount,
		MaxCredit:    DefaultMaximumCredit,
	}
}

type Options struct {
	AppCli          string
	KeyringPassword string
	KeyName         string
	KeyMnemonic     string
	KeyNodeAddr     string
	Denom           string
	CreditAmount    uint64
	MaxCredit       uint64
}

type Option func(*Options)

func CliName(s string) Option {
	return func(opts *Options) {
		opts.AppCli = s
	}
}

func KeyringPassword(s string) Option {
	return func(opts *Options) {
		opts.KeyringPassword = s
	}
}

func KeyName(s string) Option {
	return func(opts *Options) {
		opts.KeyName = s
	}
}

func WithMnemonic(s string) Option {
	return func(opts *Options) {
		opts.KeyMnemonic = s
	}
}

func WithNodeAddr(s string) Option {
	return func(opts *Options) {
		opts.KeyNodeAddr = s
	}
}

func Denom(s string) Option {
	return func(opts *Options) {
		opts.Denom = s
	}
}

func CreditAmount(v uint64) Option {
	return func(opts *Options) {
		opts.CreditAmount = v
	}
}

func MaxCredit(v uint64) Option {
	return func(opts *Options) {
		opts.MaxCredit = v
	}
}
