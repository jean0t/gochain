package cmd

import (
    "os"
    "fmt"

    "gochain/internal"
    "github.com/spf13/cobra"
)

var (
    GoChain internal.Chain
    sender string
    receiver string
    valueSent float64
)


var rootCli *cobra.Command = &cobra.Command{
    Use: "gochain",
    Short: "Interact with the blockchain from the terminal",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if _, err := os.Stat("GoChain.json"); err != nil {
            internal.CreateChain()
        }
            GoChain.Load()
    },
}

var addCmd *cobra.Command = &cobra.Command{
    Use: "add",
    Short: "Add a transaction to the blockchain",
    Run: func( cmd *cobra.Command, args []string) {
        if sender == "" || receiver == "" || valueSent <= 0 {
            fmt.Println("[--from | --to | --value] are all required.")
            return
        }

        var data internal.Data = internal.CreateData(sender, receiver, valueSent)
        var success bool = GoChain.AddNode(data)
        if !success {
            fmt.Println("Something went wrong")
            return
        }
        if err := GoChain.Save(); err != nil {
            fmt.Println("Something went wrong saving to the database")
            return
        }
        fmt.Println("Success adding transaction")
    },
}

var listCmd *cobra.Command = &cobra.Command{
    Use: "list",
    Short: "Print the blockchain as JSON to output",
    Run: func(cmd *cobra.Command, args []string) {
        GoChain.PrintChain()
    },
}

func init() {
    addCmd.Flags().StringVar(&sender, "from", "", "Sender address (required)")
    addCmd.Flags().StringVar(&receiver, "to", "", "Receiver address (required)")
    addCmd.Flags().Float64Var(&valueSent, "value", 0, "Value to transfer (required)")
    addCmd.MarkFlagRequired("from")
    addCmd.MarkFlagRequired("to")
    addCmd.MarkFlagRequired("value")

    rootCli.AddCommand(addCmd, listCmd)
}

func Execute() {
    if err := rootCli.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
