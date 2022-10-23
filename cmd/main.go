package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"XMProject/app/handler"
	repos "XMProject/app/repo"
	companyRepo "XMProject/app/repo/company"
	userRepo "XMProject/app/repo/user"
	companyService "XMProject/app/service/company"
	userService "XMProject/app/service/user"
)

func main() {

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")

	db, err := repos.NewPostgresDB(dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Printf("error connect db: %v\n", err)
	}

	repoCompanies := companyRepo.NewRepoCompanies(db)
	servCompany := companyService.NewCompaniesService(repoCompanies)

	repoUsers := userRepo.NewRepoUsers(db)
	servUser := userService.NewUsersService(repoUsers)

	handle := handler.NewHandler(servCompany, servUser)

	router := handle.InitRoutes()

	log.Fatalln(http.ListenAndServe(":3333", router))
}
