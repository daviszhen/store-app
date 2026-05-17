package business

import "store-app/server/internal/model"

func init() {
	Register(&CoffeeLoader{})
}

type CoffeeLoader struct{}

func (c *CoffeeLoader) Name() string { return "coffee" }

func (c *CoffeeLoader) Seed() *SeedData {
	return &SeedData{
		Store: model.Store{
			ID:           1,
			Name:         "慢时光咖啡",
			Theme:        "#6F4E37",
			Notice:       "慢下来，享受一杯好咖啡 ☕",
			BusinessType: "coffee",
		},
		Categories: []model.Category{
			{ID: 1, Name: "咖啡系列", Icon: "☕", Sort: 1},
			{ID: 2, Name: "茶饮系列", Icon: "🍵", Sort: 2},
			{ID: 3, Name: "甜品轻食", Icon: "🍰", Sort: 3},
			{ID: 4, Name: "咖啡豆&周边", Icon: "🫘", Sort: 4},
		},
		Products: []model.Product{
			// 咖啡系列
			{ID: 1, Name: "美式咖啡", Price: 18.00, CategoryID: 1, Description: "经典美式，醇厚回甘", Stock: 999, Status: 1},
			{ID: 2, Name: "拿铁", Price: 25.00, CategoryID: 1, Description: "意式浓缩与蒸汽牛奶的完美融合", Stock: 999, Status: 1},
			{ID: 3, Name: "卡布奇诺", Price: 26.00, CategoryID: 1, Description: "浓缩咖啡搭配绵密奶泡", Stock: 999, Status: 1},
			{ID: 4, Name: "摩卡", Price: 28.00, CategoryID: 1, Description: "巧克力与咖啡的甜蜜邂逅", Stock: 999, Status: 1},
			{ID: 5, Name: "澳白", Price: 24.00, CategoryID: 1, Description: "丝滑细腻，咖啡风味更突出", Stock: 999, Status: 1},
			{ID: 6, Name: "浓缩咖啡", Price: 16.00, CategoryID: 1, Description: "意式经典，浓郁强劲", Stock: 999, Status: 1},
			// 茶饮系列
			{ID: 7, Name: "抹茶拿铁", Price: 26.00, CategoryID: 2, Description: "日式抹茶融入香醇牛奶", Stock: 999, Status: 1},
			{ID: 8, Name: "伯爵红茶", Price: 20.00, CategoryID: 2, Description: "经典英式红茶，佛手柑清香", Stock: 999, Status: 1},
			{ID: 9, Name: "茉莉花茶", Price: 18.00, CategoryID: 2, Description: "茉莉花香，清新淡雅", Stock: 999, Status: 1},
			{ID: 10, Name: "柠檬茶", Price: 16.00, CategoryID: 2, Description: "鲜切柠檬，酸甜清爽", Stock: 999, Status: 1},
			{ID: 11, Name: "热巧克力", Price: 22.00, CategoryID: 2, Description: "浓郁可可，温暖甜蜜", Stock: 999, Status: 1},
			// 甜品轻食
			{ID: 12, Name: "提拉米苏", Price: 32.00, CategoryID: 3, Description: "经典意式甜品，入口即化", Stock: 999, Status: 1},
			{ID: 13, Name: "芝士蛋糕", Price: 28.00, CategoryID: 3, Description: "纽约风格，绵密醇厚", Stock: 999, Status: 1},
			{ID: 14, Name: "牛角包", Price: 15.00, CategoryID: 3, Description: "法式酥脆，黄油飘香", Stock: 999, Status: 1},
			{ID: 15, Name: "火腿三明治", Price: 22.00, CategoryID: 3, Description: "新鲜现做，营养搭配", Stock: 999, Status: 1},
			{ID: 16, Name: "马卡龙", Price: 18.00, CategoryID: 3, Description: "法式小圆饼，缤纷口味", Stock: 999, Status: 1},
			// 咖啡豆&周边
			{ID: 17, Name: "拼配咖啡豆 250g", Price: 68.00, CategoryID: 4, Description: "门店同款拼配，均衡醇厚", Stock: 999, Status: 1},
			{ID: 18, Name: "单品咖啡豆 200g", Price: 88.00, CategoryID: 4, Description: "埃塞俄比亚单一产地，花果香调", Stock: 999, Status: 1},
			{ID: 19, Name: "品牌随行杯", Price: 59.00, CategoryID: 4, Description: "双层不锈钢，保温6小时", Stock: 999, Status: 1},
			{ID: 20, Name: "手冲壶", Price: 128.00, CategoryID: 4, Description: "细口设计，精准控流", Stock: 999, Status: 1},
		},
	}
}
