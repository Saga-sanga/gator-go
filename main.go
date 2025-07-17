package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/saga-sanga/gator-go/internal/config"
	"github.com/saga-sanga/gator-go/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	dbURL := conf.DBURL
	db, err := sql.Open("postgres", dbURL)

	dbQueries := database.New(db)

	gatorState := &state{
		cfg: &conf,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeed)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: clid <command> [args...]")
	}

	userCommand := command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = cmds.run(gatorState, userCommand)
	if err != nil {
		log.Fatal(err)
	}

}
