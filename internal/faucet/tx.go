package faucet

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (f *Faucet) Send(recipient string) error {
	backend := "file"
	_, err := f.cliexec(f.keyNodeAddr, []string{"tx", "bank", "send", f.faucetAddress, recipient,
		fmt.Sprintf("%d%s", f.creditAmount, f.denom), "--yes","--keyring-backend",backend, "--chain-id", f.chainID,"--fees","7500nhash"},
		f.keyringPassword, f.keyringPassword, f.keyringPassword)

	return err
}

func (f *Faucet) GetTotalSent(recipient string) (uint64, error) {
	args := []string{
		"query", "txs", "--events",
		fmt.Sprintf("message.sender=%s&transfer.recipient=%s", f.faucetAddress, recipient),
		"--page", "1",
		"--limit", "1000",
	}

	output, err := f.cliexec(f.keyNodeAddr,args)
	if err != nil {
		return 0, err
	}

	var result types.SearchTxsResult
	if err := f.cdc.UnmarshalJSON([]byte(output), &result); err != nil {
		return 0, err
	}

	var total uint64
	for _, v := range result.Txs {
		if len(v.GetTx().GetMsgs()) == 0 {
			return 0, fmt.Errorf("no MsgSend available in transaction")
		}

		msg := v.GetTx().GetMsgs()[0].(*bank.MsgSend)
		total += msg.Amount.AmountOf(f.denom).Uint64()
	}

	return total, nil
}
