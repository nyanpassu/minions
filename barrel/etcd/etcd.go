// package etcd

// import (
// 	"context"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/projectcalico/libcalico-go/lib/apiconfig"
// 	"github.com/projecteru2/minions/types"
// 	"github.com/etcd-io/etcd/clientv3"
// )

// const (
// 	cmpVersion = "version"
// 	cmpValue   = "value"

// 	clientTimeout    = 10 * time.Second
// 	keepaliveTime    = 30 * time.Second
// 	keepaliveTimeout = 10 * time.Second
// )

// // Etcd .
// type Etcd struct {
// 	cliv3 *clientv3.Client
// }

// // NewEtcdClient .
// func NewEtcdClient(ctx context.Context, config apiconfig.CalicoAPIConfig) (*Etcd, error) {
// 	endpoints := strings.Split(config.Spec.EtcdConfig.EtcdEndpoints, ",")
// 	cliv3, err := clientv3.New(clientv3.Config{
// 		Endpoints:            endpoints,
// 		DialTimeout:          clientTimeout,
// 		DialKeepAliveTime:    keepaliveTime,
// 		DialKeepAliveTimeout: keepaliveTimeout,
// 		Context:              ctx,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Etcd{cliv3}, nil
// }

// // Get .
// func (e *Etcd) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
// 	return e.cliv3.Get(ctx, key, opts...)
// }

// // Put save a key value
// func (e *Etcd) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
// 	return e.cliv3.Put(ctx, key, val, opts...)
// }

// // Delete delete key
// func (e *Etcd) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
// 	return e.cliv3.Delete(ctx, key, opts...)
// }

// // BatchPut .
// func (e *Etcd) BatchPut(ctx context.Context, data map[string]string, limit map[string]map[string]string, opts ...clientv3.OpOption) (*clientv3.TxnResponse, error) {
// 	ops := []clientv3.Op{}
// 	failOps := []clientv3.Op{}
// 	conds := []clientv3.Cmp{}
// 	for key, val := range data {
// 		op := clientv3.OpPut(key, val, opts...)
// 		ops = append(ops, op)
// 		if v, ok := limit[key]; ok {
// 			for method, condition := range v {
// 				switch method {
// 				case cmpVersion:
// 					cond := clientv3.Compare(clientv3.Version(key), condition, 0)
// 					conds = append(conds, cond)
// 				case cmpValue:
// 					cond := clientv3.Compare(clientv3.Value(key), condition, val)
// 					failOps = append(failOps, clientv3.OpGet(key))
// 					conds = append(conds, cond)
// 				}
// 			}
// 		}
// 	}
// 	return e.doBatchOp(ctx, conds, ops, failOps)
// }

// func (e *Etcd) doBatchOp(ctx context.Context, conds []clientv3.Cmp, ops, failOps []clientv3.Op) (*clientv3.TxnResponse, error) {
// 	if len(ops) == 0 {
// 		return nil, types.ErrNoOps
// 	}

// 	const txnLimit = 125
// 	count := len(ops) / txnLimit // stupid etcd txn, default limit is 128
// 	tail := len(ops) % txnLimit
// 	length := count
// 	if tail != 0 {
// 		length++
// 	}

// 	resps := make([]*clientv3.TxnResponse, length)
// 	errs := make([]error, length)

// 	wg := sync.WaitGroup{}
// 	doOp := func(index int, ops []clientv3.Op) {
// 		defer wg.Done()
// 		txn := e.cliv3.Txn(ctx)
// 		if len(conds) != 0 {
// 			txn = txn.If(conds...)
// 		}
// 		resp, err := txn.Then(ops...).Else(failOps...).Commit()
// 		resps[index] = resp
// 		errs[index] = err
// 	}

// 	if tail != 0 {
// 		wg.Add(1)
// 		go doOp(length-1, ops[count*txnLimit:])
// 	}

// 	for i := 0; i < count; i++ {
// 		wg.Add(1)
// 		go doOp(i, ops[i*txnLimit:(i+1)*txnLimit])
// 	}
// 	wg.Wait()

// 	for _, err := range errs {
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if len(resps) == 0 {
// 		return &clientv3.TxnResponse{}, nil
// 	}

// 	resp := resps[0]
// 	for i := 1; i < len(resps); i++ {
// 		resp.Succeeded = resp.Succeeded && resps[i].Succeeded
// 		resp.Responses = append(resp.Responses, resps[i].Responses...)
// 	}
// 	return resp, nil
// }
