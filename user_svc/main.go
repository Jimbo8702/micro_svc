package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"Jimbo8702/randomThoughts/diggity-dawg/util"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	httpListenAddr := flag.String("httpAdr", ":4000", "the listen address of the http server")
	flag.Parse()
	mongoURL := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	var (
		svc UserService
		store = NewMongoStore(client)
	)
	svc = NewUserService(store)
	svc = NewLogMiddleware(svc)

	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))
}

func makeHTTPTransport(listenAddr string, svc UserService) error {
	fmt.Println("user_service HTTP transport running on port:", listenAddr)
	http.HandleFunc("/create", handleCreate(svc))
	http.HandleFunc("/read", handleRead(svc))
	http.HandleFunc("/update", handleUpdate(svc))
	http.HandleFunc("/delete", handleDelete(svc))

	// http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)
}

func handleCreate(svc UserService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var params types.CreateUserParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		dbUser := &UserDBCreateParams{
			FirstName: params.FirstName,
			LastName: params.LastName,
			Email: params.Email,
			Password: params.Password,
		}
		//change these to pointers dummy
		u, err := svc.CreateUser(ctx, dbUser)
		if err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, u)
	}
}

func handleRead(svc UserService) http.HandlerFunc {
	return  func (w http.ResponseWriter, r *http.Request) {
		var params types.ReadRequestParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		if params.ReadBy == "" {
			u, err := svc.ListUser(ctx);
			if err != nil {
				util.ServerError(w, err.Error())
				return
			}
			util.WriteJSON(w, http.StatusOK, u)
		} else {
			query := &types.ReadQuery{
				By: params.ReadBy,
				Item: params.Data,
			}
			u, err := svc.ReadUser(ctx, query)
			if err != nil {
				util.ServerError(w, err.Error())
				return
			}
			util.WriteJSON(w, http.StatusOK, u)
		}
	}
}

func handleUpdate(svc UserService) http.HandlerFunc {
	return  func (w http.ResponseWriter, r *http.Request) {
		var params types.UpdateUserParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		query :=  &types.ReadQuery{
			By: params.FilterBy,
			Item: params.FilterItem,
		}
		update := &UserDBUpdateParams{
			FirstName: params.FirstName,
			LastName: params.LastName,
			Email: params.Email,
			Password: params.Password,
			PhoneNumber: params.PhoneNumber,
		}
		if err := svc.UpdateUser(ctx, query, update); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, "user updated successfuly")
	}
}

func handleDelete(svc UserService) http.HandlerFunc {
	return  func (w http.ResponseWriter, r *http.Request) {
		var params types.DeleteRequestParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		if err := svc.DeleteUser(ctx, params.ItemID); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, "user deleted successfuly")
	}
}

