package discovery

import (
	"Michael-Min/octopus/config"
)

type CrawlerDiscovery struct {
	Discovery EtcdDis
	Hosts     []string
}

func NewCrawlerDiscover() *CrawlerDiscovery {
	dis := EtcdDis{Cluster: "crawler"}
	hosts := []string{config.EtcdHost1, config.EtcdHost2, config.EtcdHost3}

	return &CrawlerDiscovery{
		Discovery: dis,
		Hosts:     hosts,
	}
}

func (d *CrawlerDiscovery) GetList(service string) ([]string, bool) {
	var list []string
	if nodes, ok := d.Discovery.GetServiceInfoAllNode(service); ok {
		for _, node := range nodes {
			if _, key, ok := SplitServiceNameKey(node.Key); ok {
				list = append(list, key)
			}
		}
		return list, true
	}
	return nil, false
}

func (d *CrawlerDiscovery) GetRandomIterm(service string) (string, bool) {
	if node, ok := d.Discovery.GetServiceInfoRandom(service); ok {
		if _, key, ok := SplitServiceNameKey(node.Key); ok {
			return key, true
		}
	}
	return "", false
}
