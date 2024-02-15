package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(getBlockSpendsWithConditionsCmd)
}

var getBlockSpendsWithConditionsCmd = &cobra.Command{
    Use:   "get_block_spends_with_conditions <header_hash>",
    Args: func(cmd *cobra.Command, args []string) error {
        if err := cobra.ExactArgs(1)(cmd, args); err != nil {
            return err
        }
        if isHex(args[0]) == true {
            return nil
        }
        return fmt.Errorf("invalid hex value specified: %s", args[0])
    },
    Short: "Retrieves every coin that was spent in a block with the returned conditions",
    Long:  `Retrieves every coin that was spent in a block with the returned conditions`,
    Run: func(cmd *cobra.Command, args []string) {
        jsonData := map[string]interface{}{}
        jsonData["header_hash"] = formatHex(args[0])
        makeRequest("get_block_spends_with_conditions", jsonData)
    },
}