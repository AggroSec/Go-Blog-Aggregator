package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/config"
	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.Read()
	currentState := &state{conf: &conf}
	cmds := &commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	db, err := sql.Open("postgres", conf.DB_url)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		fmt.Println("exiting with code 1 at db connection")
		os.Exit(1)
	}

	dbQueries := database.New(db)
	currentState.db = dbQueries

	fmt.Println("Welcome to the Gator CLI!")

	if len(os.Args) < 2 {
		fmt.Println("No command provided. Available commands: login, register")
		fmt.Println("exiting with code 1 at command parsing")
		os.Exit(1)
	}

	err = cmds.run(currentState, command{os.Args[1], os.Args[2:]})
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("exiting with code 1 at command execution")
		os.Exit(1)
	}
	os.Exit(0)
}
