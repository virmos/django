package main

import (
	"github.com/fatih/color"
)

func doAuth() error {
	checkForDB()

	// migrations
	dbType := dj.DB.DataType

	tx, err := dj.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	upBytes, err := templateFS.ReadFile("templates/migrations/auth_tables." + dbType + ".sql")
	if err != nil {
		exitGracefully(err)
	}

	downBytes := []byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;")
	if err != nil {
		exitGracefully(err)
	}

	err = dj.CreatePopMigration(upBytes, downBytes, "auth", "sql")
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = dj.RunPopMigrations(tx)
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/user.go.txt", dj.RootPath + "/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/token.go.txt", dj.RootPath + "/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/remember_token.go.txt", dj.RootPath + "/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFilefromTemplate("templates/middleware/auth.go.txt", dj.RootPath + "/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth-token.go.txt", dj.RootPath + "/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/remember.go.txt", dj.RootPath + "/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/handlers/auth-handlers.go.txt", dj.RootPath + "/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/mailer/password-reset.html.tmpl", dj.RootPath + "/mail/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/mailer/password-reset.plain.tmpl", dj.RootPath + "/mail/password-reset.plain.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/login.jet", dj.RootPath + "/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/forgot.jet", dj.RootPath + "/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/reset-password.jet", dj.RootPath + "/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}


	color.Yellow("  - users, tokens, and remember_tokens migrations created and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}