package main

import (
	"fmt"

	"github.com/centrifuge/go-centrifuge/crypto"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

func init() {

	//specific param
	var curveTypeParam string
	var messageParam string
	var signatureParam string
	var publicKeyFileParam string
	var ethereumSignFlag bool

	var verifyMsgCmd = &cobra.Command{
		Use:   "verify",
		Short: "verify a signature",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			signatureBytes, err := hexutil.Decode(signatureParam)
			if err != nil {
				log.Fatal(err)
			}

			publicKey, err := utils.ReadKeyFromPemFile(publicKeyFileParam, utils.PublicKey)

			if err != nil {
				log.Fatal(err)
			}
			correct := crypto.VerifyMessage(publicKey, []byte(messageParam), signatureBytes, curveTypeParam, ethereumSignFlag)
			fmt.Println(correct)
		},
	}

	rootCmd.AddCommand(verifyMsgCmd)
	verifyMsgCmd.Flags().StringVarP(&messageParam, "message", "m", "", "message to verify")
	verifyMsgCmd.Flags().StringVarP(&publicKeyFileParam, "public", "q", "", "public key path")
	verifyMsgCmd.Flags().StringVarP(&curveTypeParam, "type", "t", "", "type of the curve (supported:'secp256k1')")
	verifyMsgCmd.Flags().StringVarP(&signatureParam, "signature", "s", "", "signature")
	verifyMsgCmd.Flags().BoolVarP(&ethereumSignFlag, "ethereum", "e", false, "verify message which was signed with Ethereum")
}
