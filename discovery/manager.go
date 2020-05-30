package discovery

import (
	"fmt"
	"strings"
)

type EtcdDis struct {
	Cluster string
	// 注册；key:服务名称
	MapRegister map[string]*EtcdService
	// 监听相关的服务
	MapWatch map[string]*EtcdMaster
}

//----------------------------------------------------------------------------------
// 注册相关的函数
//----------------------------------------------------------------------------------
// 注册
func (d *EtcdDis) Register(host []string, service, key string, value EtcdServiceInfo) error{
	name := service + "/" + key
	var s *EtcdService
	var e error
	if s, e = NewService(d.Cluster, name, value, host); nil != e {
		fmt.Printf("[Manager]: Register service:%s error:%s\n", service, e.Error())
		return e
	}

	if nil == d.MapRegister {
		d.MapRegister = make(map[string]*EtcdService)
	}

	if _, ok := d.MapRegister[name]; ok {
		fmt.Printf("[Manager]: Service:%s Have Registered!\n", name)
		return nil
	}

	d.MapRegister[name] = s
	// w维持心跳
	e = s.Start()
	return e
}

// 更新
func (d *EtcdDis) UpdateInfo(service, key string, info EtcdServiceInfo) {
	name := service + "/" + key
	if sri, ok := d.MapRegister[name]; ok {
		if nil != sri {
			sri.SetValue(info)
		}
	}
}

// 停止
func (d *EtcdDis) Stop(service, key string) {
	name := service + "/" + key
	if sri, ok := d.MapRegister[name]; ok {
		sri.Stop()
	}
}


//----------------------------------------------------------------------------------
// 监听相关的函数
//----------------------------------------------------------------------------------
func (d *EtcdDis) Watch(host []string, service string) {
	var w *EtcdMaster
	var e error
	if w, e = NewMaster(host, d.Cluster, service); nil != e {
		fmt.Printf("[Manager]: Watch Service:%s Failed!Error:%s\n", service, e.Error())
		return
	}

	if nil == d.MapWatch {
		d.MapWatch = make(map[string]*EtcdMaster)
	}

	if _, ok := d.MapWatch[service]; ok {
		fmt.Printf("[Manager]: Service:%s Have Watch!\n", service)
		return
	}

	d.MapWatch[service] = w
}

// 获取服务的节点信息-随机获取
func (d *EtcdDis) GetServiceInfoRandom(service string) (EtcdNode, bool) {
	if nil == d.MapWatch {
		fmt.Println("[Manager]: MapWatch is nil")
		return EtcdNode{}, false
	}

	if v, ok := d.MapWatch[service]; ok {
		if nil != v {
			if n, ok1 := v.GetNodeRandom(); ok1 {
				return n, true
			}
		}
	} else {
		fmt.Printf("[Manager]: Service:%s Not Be Watched!\n", service)
	}

	return EtcdNode{}, false
}

// 获取服务的节点信息-全部获取
func (d *EtcdDis) GetServiceInfoAllNode(service string) ([]EtcdNode, bool) {
	if nil == d.MapWatch {
		fmt.Println("MapWatch is nil")
		return []EtcdNode{}, false
	}

	if v, ok := d.MapWatch[service]; ok {
		if nil != v {
			return v.GetAllNodes(), true
		}
	} else {
		fmt.Printf("[Manager]: Service:%s Not Be Watched!\n", service)
	}

	return []EtcdNode{}, false
}

//----------------------------------------------------------------------------------
// 工具相关函数
//----------------------------------------------------------------------------------
// 拆分service name、key；返回bool true表示成功;false表示失败
func SplitServiceNameKey(dir string) (string, string, bool) {
	if idx := strings.Index(dir, "/"); -1 != idx {
		name := dir[:idx]
		key := dir[idx+1:]
		return name, key, true
	}
	return "", "", false
}
