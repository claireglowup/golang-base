package src

import (
	"database/sql"
	"golang-base/src/delivery"
	"golang-base/src/repository"
	"golang-base/src/usecase"
	"golang-base/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start()
}

type server struct {
	httpServer *gin.Engine
	db         *sql.DB
	port       string
	linkAPI    string
}

func InitServer(db *sql.DB, env *util.Config) Server {

	g := gin.Default()

	return &server{
		httpServer: g,
		db:         db,
		port:       env.HTTPServerAddress,
	}

}

func (s *server) Start() {

	repo := repository.NewStore(s.db)
	usecase := usecase.NewUsecase(repo)
	delivery := delivery.NewDelivery(usecase)

	log.Println("pake delivery bro", delivery)

	s.httpServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	err := s.httpServer.Run(s.port)
	if err != nil {
		log.Fatal(err)
	}

}
