package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_readArguments(t *testing.T) {
	tests := []struct {
		name    string
		want    storeParams
		wantErr bool
		osArgs  []string
	}{
		{"No parameters", storeParams{}, true, []string{"cmd"}},
		{"Weird input without subcommand", storeParams{}, true, []string{"cmd", "--help"}},
		{"Subcommand with no parameters", storeParams{}, true, []string{"cmd", "store"}},
		{"Subcommand with correct parameters", storeParams{
			"test1000", "test000", "apikey_test"}, false, []string{
			"cmd", "store", "--template-id", "test1000", "--password", "test000", "--apikey", "apikey_test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args
			defer func() {
				os.Args = actualOsArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()
			os.Args = tt.osArgs
			got, err := readArguments()
			if (err != nil) != tt.wantErr {
				t.Errorf("readArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readArguments() got = %v, want %v", got, tt.want)
			}
		})
	}
}
