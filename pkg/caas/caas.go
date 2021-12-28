package caas

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zawachte/caas-rest-api/pkg/caasdb"
)

type GetClusterFindByAccountIdParams struct {
	AccountId string
}

type CaasServer struct {
	caasdbClient *caasdb.Client
}

func NewCaasServer() (*CaasServer, error) {

	dbClient, err := caasdb.NewClient()
	if err != nil {
		return nil, err
	}

	return &CaasServer{dbClient}, err
}

type Cluster struct {
	Id           string
	Text         string
	Tags         []string
	CreationDate string
	AccountId    string
	Kubeconfig   string
}

type Account struct {
	Id           string
	Username     string
	Password     string
	Email        string
	CreationDate string
	Tags         []string
}

func (cs *CaasServer) GetAccount(c *gin.Context) {
	accounts, err := cs.getAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, accounts)
	}
}

func (cs *CaasServer) PostAccount(c *gin.Context) {
	a := Account{}
	err := json.NewDecoder(c.Request.Body).Decode(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = cs.postAccount(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
func (cs *CaasServer) GetCluster(c *gin.Context) {
	clusters, err := cs.getClusters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, clusters)
	}
}
func (cs *CaasServer) PostCluster(c *gin.Context) {

	cl := Cluster{}
	err := json.NewDecoder(c.Request.Body).Decode(&cl)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = cs.postCluster(cl)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (cs *CaasServer) postCluster(cluster Cluster) error {
	var lastInsertID int

	i64, err := strconv.ParseInt(cluster.AccountId, 10, 32)
	if err != nil {
		return err
	}

	err = cs.dbPool.QueryRow(context.Background(), "INSERT INTO clusters(account_id, kubeconfig, created_on) VALUES($1, $2, $3) returning cluster_id;",
		i64,
		cluster.Kubeconfig,
		time.Now()).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CaasServer) postAccount(account Account) error {
	var lastInsertID int

	err := cs.dbPool.QueryRow(context.Background(), "INSERT INTO accounts(username, password, email, created_on) VALUES($1, $2, $3, $4) returning account_id;",
		account.Username,
		account.Password,
		account.Email,
		time.Now()).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CaasServer) GetClusterFindByAccountId(c *gin.Context, params GetClusterFindByAccountIdParams) {

	cluster, err := cs.getClusterByAccountId(c.Request.Context(), params.AccountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, cluster)
}

func (cs *CaasServer) getClusterByAccountId(ctx context.Context, accountId string) ([]Cluster, error) {

	i64, err := strconv.ParseInt(accountId, 10, 32)
	if err != nil {
		return nil, err
	}

	rows, err := cs.dbPool.Query(ctx, "select * from clusters where account_id=$1", i64)
	if err != nil {
		return nil, err
	}

	// iterate through the rows
	clusters := []Cluster{}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		cluster := Cluster{
			Id:           string(values[0].(int32)),
			AccountId:    string(values[1].(int32)),
			Kubeconfig:   values[2].(string),
			CreationDate: values[3].(time.Time).String(),
		}

		clusters = append(clusters, cluster)
	}

	return clusters, nil
}

func (cs *CaasServer) DeleteClusterId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, nil)
}

func (cs *CaasServer) GetClusterId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, nil)
}

func (cs *CaasServer) Close() {
	cs.dbPool.Close()
}

func (cs *CaasServer) getAccounts() ([]Account, error) {

	rows, err := cs.dbPool.Query(context.Background(), "select * from accounts")
	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		account := Account{
			Id:           string(values[0].(int32)),
			Username:     values[1].(string),
			Email:        values[3].(string),
			CreationDate: values[4].(time.Time).String(),
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (cs *CaasServer) getClusters() ([]Cluster, error) {

	rows, err := cs.dbPool.Query(context.Background(), "select * from clusters")
	if err != nil {
		return nil, err
	}

	clusters := []Cluster{}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		cluster := Cluster{
			Id:           string(values[0].(int32)),
			AccountId:    string(values[1].(int32)),
			Kubeconfig:   values[2].(string),
			CreationDate: values[3].(time.Time).String(),
		}

		clusters = append(clusters, cluster)
	}

	return clusters, nil
}
