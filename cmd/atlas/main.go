package main

import (
	"context"
	"os"
	"os/signal"

	"ariga.io/atlas/cmd/atlascmd"
	_ "ariga.io/atlas/cmd/atlascmd/docker"
	_ "ariga.io/atlas/sql/mysql"
	_ "ariga.io/atlas/sql/postgres"
	_ "ariga.io/atlas/sql/sqlite"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	atlascmd.Root.SetOut(os.Stdout)
	err := atlascmd.Root.ExecuteContext(ctx)
	// Print error from command
	if err != nil {
		atlascmd.Root.PrintErrln("Error:", err)
	}
	// Check for update.
	atlascmd.CheckForUpdate()
	// Exit code according to command success.
	if err != nil {
		os.Exit(1)
	}
}
