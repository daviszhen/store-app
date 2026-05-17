package business

import "store-app/server/internal/model"

func init() {
	Register(&GroceryLoader{})
}

type GroceryLoader struct{}

func (g *GroceryLoader) Name() string { return "grocery" }

func (g *GroceryLoader) Seed() *SeedData {
	return &SeedData{
		Store: model.Store{
			ID:    1,
			Name:  "瑞信商店",
			Theme: "#e4393c",
			Notice: "新店开业，全场9折！",
		},
		Categories: []model.Category{
			{ID: 1, Name: "时令水果", Icon: "🍎", Sort: 1},
			{ID: 2, Name: "新鲜蔬菜", Icon: "🥬", Sort: 2},
			{ID: 3, Name: "精选肉禽", Icon: "🥩", Sort: 3},
		},
		Products: []model.Product{
			{ID: 1, Name: "红富士苹果", Price: 12.80, CategoryID: 1, Description: "产地直采，新鲜脆甜", Stock: 999, Status: 1},
			{ID: 2, Name: "进口香蕉", Price: 8.80, CategoryID: 1, Description: "自然成熟，香甜软糯", Stock: 999, Status: 1},
			{ID: 3, Name: "有机西红柿", Price: 6.80, CategoryID: 2, Description: "自然生长，酸甜多汁", Stock: 999, Status: 1},
			{ID: 4, Name: "新鲜黄瓜", Price: 4.50, CategoryID: 2, Description: "清脆爽口，现摘现卖", Stock: 999, Status: 1},
			{ID: 5, Name: "精品五花肉", Price: 32.00, CategoryID: 3, Description: "当日鲜切，肥瘦相间", Stock: 999, Status: 1},
			{ID: 6, Name: "土鸡蛋10枚", Price: 15.80, CategoryID: 3, Description: "散养土鸡，营养丰富", Stock: 999, Status: 1},
		},
	}
}
