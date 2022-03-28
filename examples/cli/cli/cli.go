// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/circl-dev/go-swagger/examples/cli/client"

	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/circl-dev/runtime"
	httptransport "github.com/circl-dev/runtime/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.AToDoListApplication, error) {
	hostname := viper.GetString("hostname")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, client.DefaultBasePath, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONConsumer()

	r.Producers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONProducer()

	auth, err := makeAuthInfoWriter(cmd)
	if err != nil {
		return nil, err
	}
	r.DefaultAuthentication = auth

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	if err := registerAuthInoWriterFlags(rootCmd); err != nil {
		return nil, err
	}
	// add all operation groups
	operationGroupTodosCmd, err := makeOperationGroupTodosCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTodosCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

// registerAuthInoWriterFlags registers all flags needed to perform authentication
func registerAuthInoWriterFlags(cmd *cobra.Command) error {
	/*x-todolist-token */
	cmd.PersistentFlags().String("x-todolist-token", "", ``)
	viper.BindPFlag("x-todolist-token", cmd.PersistentFlags().Lookup("x-todolist-token"))
	return nil
}

// makeAuthInfoWriter retrieves cmd flags and construct an auth info writer
func makeAuthInfoWriter(cmd *cobra.Command) (runtime.ClientAuthInfoWriter, error) {
	auths := []runtime.ClientAuthInfoWriter{}
	/*x-todolist-token */
	if viper.IsSet("x-todolist-token") {
		XTodolistTokenKey := viper.GetString("x-todolist-token")
		auths = append(auths, httptransport.APIKeyAuth("x-todolist-token", "header", XTodolistTokenKey))
	}
	if len(auths) == 0 {
		logDebugf("Warning: No auth params detected.")
		return nil, nil
	}
	// compose all auths together
	return httptransport.Compose(auths...), nil
}

func makeOperationGroupTodosCmd() (*cobra.Command, error) {
	operationGroupTodosCmd := &cobra.Command{
		Use:  "todos",
		Long: ``,
	}

	operationAddOneCmd, err := makeOperationTodosAddOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationAddOneCmd)

	operationDestroyOneCmd, err := makeOperationTodosDestroyOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationDestroyOneCmd)

	operationFindTodosCmd, err := makeOperationTodosFindTodosCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationFindTodosCmd)

	operationUpdateOneCmd, err := makeOperationTodosUpdateOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationUpdateOneCmd)

	return operationGroupTodosCmd, nil
}
