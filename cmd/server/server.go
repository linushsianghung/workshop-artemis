package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/linushsianghung/internal/app/persistent"
	"github.com/linushsianghung/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	kind     string
	port     string
	operator persistent.NoSqlOperator
}

func NewHTTPServer(port string, operator persistent.NoSqlOperator) *Server {
	return &Server{
		kind:     "Restful",
		port:     port,
		operator: operator,
	}
}

func InitServer(s *Server) {
	http.HandleFunc("/", handleHelloWorld)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatalf("Failed to start server on port: %s", s.port)
		return
	}
	log.Printf("Server is running on port: %s", s.port)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application-json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method Not Allowed!!!")); err != nil {
			log.Println("Failed to write response!!!")
			json.NewEncoder(w).Encode(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Hello World")
	if err != nil {
		log.Println("Failed to encode response!!!")
		json.NewEncoder(w).Encode(err)
		return
	}

}

// FetchConnectionString fetch connection information of Database
// TODO: Improvement for security by using Environment Variables or Command Flags
func FetchConnectionString(isLocal bool) string {
	if isLocal {
		return configs.GetConfigStr("connection.nosql.localhost")
	}

	host := configs.GetConfigStr("connection.nosql.host")
	port := configs.GetConfigStr("connection.nosql.port")
	username := configs.GetConfigStr("connection.nosql.username")
	password := configs.GetConfigStr("connection.nosql.password")
	db := configs.GetConfigStr("connection.nosql.database")
	// rs := configs.GetConfigStr("connection.nosql.replicaset")
	return fmt.Sprintf("mongodb+srv://%s:%s@%s%s/%s?authSource=admin&retryWrites=true&w=majority", username, password, host, port, db)
}
