package business

import "store-app/server/internal/model"

// SeedData 业务类型的种子数据结构
type SeedData struct {
	Store      model.Store
	Categories []model.Category
	Products   []model.Product
}

// Loader 业务类型种子数据加载器接口
type Loader interface {
	Name() string
	Seed() *SeedData
}

var registry = map[string]Loader{}

// Register 注册业务类型
func Register(l Loader) {
	registry[l.Name()] = l
}

// Get 获取指定业务类型的种子数据，不存在返回 nil
func Get(businessType string) *SeedData {
	if l, ok := registry[businessType]; ok {
		return l.Seed()
	}
	return nil
}

// List 列出所有已注册的业务类型
func List() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}
