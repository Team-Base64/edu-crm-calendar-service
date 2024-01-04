package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	conf "main/config"
	ctrl "main/controller"
	deliv "main/delivery"
	st "main/repository"
	us "main/usecase"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"

	_ "main/docs"

	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc/keepalive"
)

var urlDB string
var urlDomain string
var tokenFile string
var credentialsFile string

func init() {
	var exist bool

	pgUser, exist := os.LookupEnv(conf.PG_USER)
	if !exist || len(pgUser) == 0 {
		log.Fatalln("could not get database host from env")
	}
	pgPwd, exist := os.LookupEnv(conf.PG_PWD)
	if !exist || len(pgPwd) == 0 {
		log.Fatalln("could not get database password from env")
	}
	pgHost, exist := os.LookupEnv(conf.PG_HOST)
	if !exist || len(pgHost) == 0 {
		log.Fatalln("could not get database host from env")
	}
	pgPort, exist := os.LookupEnv(conf.PG_PORT)
	if !exist || len(pgPort) == 0 {
		log.Fatalln("could not get database port from env")
	}
	pgDB, exist := os.LookupEnv(conf.PG_DB)
	if !exist || len(pgDB) == 0 {
		log.Fatalln("could not get database name from env")
	}

	urlDB = "postgres://" + pgUser + ":" + pgPwd + "@" + pgHost + ":" + pgPort + "/" + pgDB

	// urlDBs, exist := os.LookupEnv(conf.URL_DB)
	// if !exist || len(urlDBs) == 0 {
	// 	log.Fatalln("could not get database name from env")
	// }

	// urlDB = urlDBs

	urlDomain, exist = os.LookupEnv(conf.UrlDomain)
	if !exist || len(urlDomain) == 0 {
		log.Fatalln("could not get url domain from env")
	}

	tokenFile, exist = os.LookupEnv(conf.TokenFile)
	if !exist || len(tokenFile) == 0 {
		log.Fatalln("could not get token file path from env")
	}

	credentialsFile, exist = os.LookupEnv(conf.CredentialsFile)
	if !exist || len(credentialsFile) == 0 {
		log.Fatalln("could not get credentials file path from env")
	}

}

func main() {
	myRouter := mux.NewRouter()

	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	Store := st.NewStore(db)
	Usecase := us.NewUsecase(Store, tokenFile, credentialsFile)

	lis, err := net.Listen("tcp", conf.PortGRPCCalendar)
	if err != nil {
		log.Fatalln("cant listen grpc calendar port", err)
	}
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024),
		grpc.MaxConcurrentStreams(35),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    1 * time.Second,
			Timeout: 5 * time.Second,
		}),
	)
	ctrl.RegisterCalendarControllerServer(
		server,
		ctrl.NewCtrlService(
			Usecase,
			urlDomain,
		),
	)

	Handler := deliv.NewHandler(Usecase)

	myRouter.HandleFunc(conf.PathOAuthSetToken, Handler.SetOAUTH2Token).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathOAuthSaveToken, Handler.SaveOAUTH2TokenToFile).Methods(http.MethodGet, http.MethodOptions)

	//myRouter.HandleFunc(conf.PathWS, Handler.ServeWs).Methods(http.MethodGet, http.MethodOptions)
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)

	log.Println("starting grpc server at " + conf.PortGRPCCalendar)
	go server.Serve(lis)

	err = http.ListenAndServe(conf.PortWebCalendar, myRouter)

	if err != nil {
		log.Fatalln("cant serve", err)
	}

}
