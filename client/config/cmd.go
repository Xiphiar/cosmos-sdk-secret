package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// Cmd returns a CLI command to interactively create an application CLI
// config file.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <key> [value]",
		Short: "Create or query an application CLI configuration file",

		// This is here to prevent config validation before changing it.
		// It overrides the rootCmd `PersistentPreRunE` to allow to correct a faulty configuration.
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Create context with flags to get the right home dir
			clientCtx, err := client.ReadPersistentCommandFlags(client.GetClientContextFromCmd(cmd).WithViper(""), cmd.Flags())
			if err != nil {
				return err
			}

			// Create config.toml if it doesn't exist yet
			_, err = GetConfigOrDefault(clientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			return nil
		},
		RunE: runConfigCmd,
		Args: cobra.RangeArgs(0, 2),
	}
	return cmd
}

func runConfigCmd(cmd *cobra.Command, args []string) error {
	clientCtx := client.GetClientContextFromCmd(cmd)
	configPath := filepath.Join(clientCtx.HomeDir, "config")

	conf, err := getClientConfig(configPath, clientCtx.Viper)
	if err != nil {
		return fmt.Errorf("couldn't get client config: %v", err)
	}

	switch len(args) {
	case 0:
		// print all client config fields to stdout
		s, err := json.MarshalIndent(conf, "", "\t")
		if err != nil {
			return err
		}
		cmd.Println(string(s))

	case 1:
		// it's a get
		key := args[0]

		switch key {
		case flags.FlagChainID:
			cmd.Println(conf.ChainID)
		case flags.FlagKeyringBackend:
			cmd.Println(conf.KeyringBackend)
		case tmcli.OutputFlag:
			cmd.Println(conf.Output)
		case flags.FlagNode:
			cmd.Println(conf.Node)
		case flags.FlagBroadcastMode:
			cmd.Println(conf.BroadcastMode)
		default:
			err := errUnknownConfigKey(key)
			return fmt.Errorf("couldn't get the value for the key: %v, error:  %v", key, err)
		}

	case 2:
		// it's set
		key, value := args[0], args[1]

		switch key {
		case flags.FlagChainID:
			conf.SetChainID(value)
		case flags.FlagKeyringBackend:
			conf.SetKeyringBackend(value)
		case tmcli.OutputFlag:
			conf.SetOutput(value)
		case flags.FlagNode:
			conf.SetNode(value)
		case flags.FlagBroadcastMode:
			conf.SetBroadcastMode(value)
		default:
			return errUnknownConfigKey(key)
		}

		confFile := filepath.Join(configPath, "client.toml")
		if err := writeConfigToFile(confFile, conf); err != nil {
			return fmt.Errorf("could not write client config to the file: %v", err)
		}

	default:
		panic("cound not execute config command")
	}

	return nil
}

func errUnknownConfigKey(key string) error {
	return fmt.Errorf("unknown configuration key: %q", key)
}
