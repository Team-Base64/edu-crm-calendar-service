package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"time"

	conf "main/config"
	googleApi "main/delivery/calendar"
	grpcCalendar "main/delivery/grpc/calendar"
	protoCalendar "main/delivery/grpc/calendar/proto"
	pgStore "main/repository/pg"
	calendarUc "main/usecase/calendar"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var urlDB string
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

	tokenFile, exist = os.LookupEnv(conf.TOKEN_FILE)
	if !exist || len(tokenFile) == 0 {
		log.Fatalln("could not get token file path from env")
	}

	credentialsFile, exist = os.LookupEnv(conf.CREDENTILS_FILE)
	if !exist || len(credentialsFile) == 0 {
		log.Fatalln("could not get credentials file path from env")
	}

}

func main() {
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	store := pgStore.NewStore(db)
	calendar := googleApi.NewGoogleCalendar(tokenFile, credentialsFile)
	usecase := calendarUc.NewCalendarUsecase(store, calendar)

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
	protoCalendar.RegisterCalendarServer(
		server,
		grpcCalendar.NewCalendarGrpcHandler(usecase),
	)

	log.Println("starting grpc server at " + conf.PortGRPCCalendar)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("cant serve", err)
	}

}
