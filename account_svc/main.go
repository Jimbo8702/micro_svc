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
	httpListenAddr := flag.String("httpAdr", ":5000", "the listen address of the http server")
	flag.Parse()
	mongoURL := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	var (
		svc AccountService
		store = NewMongoStore(client)
	)
	svc = NewAccountService(store)
	svc = NewLogMiddleware(svc)

	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))
}

func makeHTTPTransport(listenAddr string, svc AccountService) error {
	fmt.Println("account_service HTTP transport running on port:", listenAddr)
	http.HandleFunc("/create", handleCreate(svc))
	http.HandleFunc("/read", handleRead(svc))
	http.HandleFunc("/update", handleUpdate(svc))
	http.HandleFunc("/delete", handleDelete(svc))

	// http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)
}

func handleCreate(svc AccountService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var params types.CreateAccountParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		dbAcc := &AccountDBCreateParams{
			UserID: params.UserID,
			StripeCustomerID: params.StripeCustomerID,
			StripeSubscriptionID: params.StripeSubscriptionID,
			SubscriptionStatus: params.SubscriptionStatus,
			Plan: params.Plan,
		}
		a, err := svc.CreateAccount(ctx, dbAcc)
		if err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, a)
	}
}

func handleRead(svc AccountService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var params types.ReadRequestParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		if params.ReadBy == "" {
			u, err := svc.ListAccount(ctx);
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
			u, err := svc.ReadAccount(ctx, query)
			if err != nil {
				util.ServerError(w, err.Error())
				return
			}
			util.WriteJSON(w, http.StatusOK, u)
		}
	}
}

func handleUpdate(svc AccountService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var params types.UpdateAccountParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		query :=  &types.ReadQuery{
			By: params.FilterBy,
			Item: params.FilterItem,
		}
		update := &AccountDBUpdateParams{
			StripeCustomerID: params.StripeCustomerID,
			StripeSubscriptionID: params.StripeSubscriptionID,
			SubscriptionStatus: params.SubscriptionStatus,
			Plan: params.Plan,
		}
		if err := svc.UpdateAccount(ctx, query, update); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, "account updated successfully")
	}
}

func handleDelete(svc AccountService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var params types.DeleteRequestParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		if err := svc.DeleteAccount(ctx, params.ItemID); err != nil {
			util.ServerError(w, err.Error())
			return
		}
		util.WriteJSON(w, http.StatusOK, "account deleted successfuly")
	}
}