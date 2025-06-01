package cmd

import (
    "os"
    "fmt"

    "gochain/internal"
    "github.com/spf13/cobra"
)

var GoChain internal.Chain

var rootCli *cobra.Command = &cobra.Command{
    Use: "gochain",
    Short: "Interact with the blockchain from the terminal",
}


func init() {
}

func Execute() {
    if err := rootCli.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
