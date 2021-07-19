/*
This file was generated by cobra (github.com/spf13/cobra)

Licensed under the Apache License, Version 2.0 (the "License");

Command entry points are licensed under the BSD-2-Clause license
*/
package cmd

import (
	"fmt"
    "os"
    "path/filepath"

	"github.com/spf13/cobra"
    "github.com/nishakm/genv/pkg/versions"
    "github.com/nishakm/genv/pkg/workspace"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to another version of go",
	Long: `Use switch to switch to another version of go within a
genv environment.

Eg:
genv myenv
cd myenv
source bin/activate
(myenv) envtool switch go1.16.5`,
	Run: func(cmd *cobra.Command, args []string) {
        goVersion := args[0]
        // We assume that the environment has been activated
        // so we find the current GOPATH
        gopath := os.Getenv("GOPATH")
        // our environment path is relative to the gopath
        folderpath := filepath.Dir(gopath)
        // now we can get the other paths
        envpath := workspace.Envpath(folderpath)
        // get the path of the new go version
        newGoRoot := versions.GetGoRoot(goVersion)
        if newGoRoot != "" {
            // we have a new go binary to symlink to
            workspace.SetGoSym(newGoRoot, envpath)
        }
        fmt.Printf("Go version changed to %s\n. Run 'deactivate' and 'source bin/activate' again", goVersion)
        
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}