package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	RegisterConsul(name, tags string) error
	DeRegister(serviceId string) error
	GetService() map[string]*api.AgentService
	FilterService(serviceName string) map[string]*api.AgentService
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

// 封装注册方法
func (r *Registry) RegisterConsul(name, tags string) error {
	registryation := api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    []string{tags},
		Port:    r.Port,
		Address: r.Host,
	}
	Consul.Agent().ServiceRegister(&registryation)
	return nil
}

// todo:注销服务
func (r *Registry) DeRegister(serviceId string) error {
	err := Consul.Agent().ServiceDeregister(serviceId)
	return err
}

// todo:获取所有服务
func (r *Registry) GetService() map[string]*api.AgentService {
	result, err := Consul.Agent().Services()
	if err != nil {
		panic(err)
	}
	return result
}

// todo:过滤服务--根据服务名称过滤
func (r *Registry) FilterService(serviceName string) map[string]*api.AgentService {
	result, err := Consul.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, serviceName))
	if err != nil {
		log.Fatalf("服务查询失败:%v", err)
	}
	return result
}
