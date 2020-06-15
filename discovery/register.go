package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/v3/clientv3"
	"time"
)

type EtcdServiceInfo struct {
	Info string
}

type EtcdService struct {
	Cluster string          // 集群名称
	Name    string          // 服务名称
	Info    EtcdServiceInfo // 节点信息
	stop    chan error
	leaseid clientv3.LeaseID
	client  *clientv3.Client
}

// 注册ETCD服务
func NewService(cluster, name string, info EtcdServiceInfo, hosts []string) (*EtcdService, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: 2 * time.Second,
	})

	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	// 返回服务对象
	return &EtcdService{
		Cluster: cluster,
		Name:    name,
		Info:    info,
		stop:    make(chan error),
		client:  cli,
	}, err

}

// 启动
func (s *EtcdService) Start() error {
	// 获取心跳的通道
	ch, err := s.keepLive()
	if nil != err {
		fmt.Println(err.Error())
		return err
	}
	go func() {
		// 死循环
		for {
			select {
			case <-s.stop:
				s.revoke()
				return
			case <-s.client.Ctx().Done():
				fmt.Println("[Warning]: server closed")
				return
			case /*ka*/ _, ok := <-ch:
				if !ok {
					fmt.Println("[Warning]: keep live channel closed")
					s.revoke()
					return
				} else {
					//fmt.Printf("recv reply from service:%s, ttl:%d\n", s.Name, ka.TTL)
				}
			}
		}
	}()
	return nil
}

// 保持心跳
func (s *EtcdService) keepLive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	key := s.Cluster + "/" + s.Name
	value, _ := json.Marshal(s.Info)

	// minimum lease TTL is 5-second
	resp, err := s.client.Grant(context.TODO(), 5)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}

	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("[Register]: Key:%s, Value:%s\n", key, string(value))
	s.leaseid = resp.ID

	return s.client.KeepAlive(context.TODO(), resp.ID)
}

// 设置节点信息
func (s *EtcdService) SetValue(info EtcdServiceInfo) {
	s.Info = info
	tmp, _ := json.Marshal(info)
	key := s.Cluster + "/" + s.Name
	if _, err := s.client.Put(context.TODO(), key, string(tmp), clientv3.WithLease(s.leaseid)); nil != err {
		fmt.Printf("[Error]: etcd set value failed! key:%s;value:%s", key, info)
	}

}

// 停止
func (s *EtcdService) Stop() {
	s.stop <- nil
}

// 撤销
func (s *EtcdService) revoke() error {
	_, err := s.client.Revoke(context.TODO(), s.leaseid)
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Printf("[Warning]: service:%s stop\n", s.Name)
	return nil
}
