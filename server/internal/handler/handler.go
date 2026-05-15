package handler

import (
"fmt"
"net/http"
"strconv"
"time"

"store-app/server/internal/database"
"store-app/server/internal/model"

"github.com/gin-gonic/gin"
)

func json(c *gin.Context, code int, msg string, data interface{}) {
c.JSON(code, gin.H{"code": code, "msg": msg, "data": data})
}

// ---------- Store ----------

func GetStore(c *gin.Context) {
var store model.Store
if err := database.DB.First(&store, 1).Error; err != nil {
json(c, 404, "店铺不存在", nil)
return
}
json(c, 200, "ok", store)
}

func UpdateStore(c *gin.Context) {
var store model.Store
if err := database.DB.First(&store, 1).Error; err != nil {
json(c, 404, "店铺不存在", nil)
return
}
if err := c.ShouldBindJSON(&store); err != nil {
json(c, 400, "参数错误", nil)
return
}
database.DB.Model(&store).Updates(store)
json(c, 200, "ok", store)
}

// ---------- Category ----------

func GetCategories(c *gin.Context) {
var categories []model.Category
database.DB.Order("sort").Find(&categories)
json(c, 200, "ok", categories)
}

func CreateCategory(c *gin.Context) {
var cat model.Category
if err := c.ShouldBindJSON(&cat); err != nil {
json(c, 400, "参数错误", nil)
return
}
database.DB.Create(&cat)
json(c, 200, "ok", cat)
}

func UpdateCategory(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
var cat model.Category
if err := database.DB.First(&cat, id).Error; err != nil {
json(c, 404, "分类不存在", nil)
return
}
if err := c.ShouldBindJSON(&cat); err != nil {
json(c, 400, "参数错误", nil)
return
}
database.DB.Save(&cat)
json(c, 200, "ok", cat)
}

func DeleteCategory(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
database.DB.Delete(&model.Category{}, id)
json(c, 200, "ok", nil)
}

// ---------- Product ----------

func GetProducts(c *gin.Context) {
var products []model.Product
query := database.DB.Preload("Category")
if cid := c.Query("category_id"); cid != "" {
query = query.Where("category_id = ?", cid)
}
query.Order("created_at desc").Find(&products)
json(c, 200, "ok", products)
}

func GetProduct(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
var product model.Product
if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
json(c, 404, "商品不存在", nil)
return
}
json(c, 200, "ok", product)
}

func CreateProduct(c *gin.Context) {
var p model.Product
if err := c.ShouldBindJSON(&p); err != nil {
json(c, http.StatusBadRequest, "参数错误", nil)
return
}
database.DB.Create(&p)
json(c, 200, "ok", p)
}

func UpdateProduct(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
var p model.Product
if err := database.DB.First(&p, id).Error; err != nil {
json(c, 404, "商品不存在", nil)
return
}
if err := c.ShouldBindJSON(&p); err != nil {
json(c, 400, "参数错误", nil)
return
}
database.DB.Save(&p)
json(c, 200, "ok", p)
}

func DeleteProduct(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
database.DB.Delete(&model.Product{}, id)
json(c, 200, "ok", nil)
}

// ---------- Cart ----------

func GetCart(c *gin.Context) {
userID := c.Query("user_id")
if userID == "" {
userID = "default"
}
var items []model.Cart
database.DB.Preload("Product").Where("user_id = ?", userID).Find(&items)
json(c, 200, "ok", items)
}

func AddToCart(c *gin.Context) {
var req struct {
UserID    string `json:"user_id"`
ProductID int    `json:"product_id"`
Quantity  int    `json:"quantity"`
}
if err := c.ShouldBindJSON(&req); err != nil {
json(c, 400, "参数错误", nil)
return
}
if req.UserID == "" {
req.UserID = "default"
}
if req.Quantity <= 0 {
req.Quantity = 1
}

var existing model.Cart
err := database.DB.Where("user_id = ? AND product_id = ?", req.UserID, req.ProductID).First(&existing).Error
if err == nil {
existing.Quantity += req.Quantity
database.DB.Save(&existing)
json(c, 200, "ok", existing)
return
}

cart := model.Cart{
UserID:    req.UserID,
ProductID: req.ProductID,
Quantity:  req.Quantity,
}
database.DB.Create(&cart)
json(c, 200, "ok", cart)
}

func UpdateCartItem(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
var req struct {
Quantity int `json:"quantity"`
}
if err := c.ShouldBindJSON(&req); err != nil {
json(c, 400, "参数错误", nil)
return
}
var item model.Cart
if err := database.DB.First(&item, id).Error; err != nil {
json(c, 404, "购物车项不存在", nil)
return
}
if req.Quantity <= 0 {
database.DB.Delete(&item)
json(c, 200, "ok", nil)
return
}
item.Quantity = req.Quantity
database.DB.Save(&item)
json(c, 200, "ok", item)
}

func DeleteCartItem(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
database.DB.Delete(&model.Cart{}, id)
json(c, 200, "ok", nil)
}

// ---------- Order ----------

func CreateOrder(c *gin.Context) {
var req struct {
UserID       string `json:"user_id"`
ContactName  string `json:"contact_name"`
ContactPhone string `json:"contact_phone"`
ContactAddr  string `json:"contact_addr"`
Remark       string `json:"remark"`
}
if err := c.ShouldBindJSON(&req); err != nil {
json(c, 400, "参数错误", nil)
return
}
if req.UserID == "" {
req.UserID = "default"
}

var cartItems []model.Cart
database.DB.Preload("Product").Where("user_id = ?", req.UserID).Find(&cartItems)
if len(cartItems) == 0 {
json(c, 400, "购物车为空", nil)
return
}

var total float64
for _, item := range cartItems {
total += item.Product.Price * float64(item.Quantity)
}

orderNo := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
time.Now().Year(), time.Now().Month(), time.Now().Day(),
time.Now().Hour(), time.Now().Minute(), time.Now().Second())

order := model.Order{
OrderNo:      orderNo,
UserID:       req.UserID,
TotalAmount:  total,
Status:       1,
ContactName:  req.ContactName,
ContactPhone: req.ContactPhone,
ContactAddr:  req.ContactAddr,
Remark:       req.Remark,
}

tx := database.DB.Begin()
if err := tx.Create(&order).Error; err != nil {
tx.Rollback()
json(c, 500, "创建订单失败", nil)
return
}

for _, item := range cartItems {
oi := model.OrderItem{
OrderID:     order.ID,
ProductID:   item.ProductID,
ProductName: item.Product.Name,
Price:       item.Product.Price,
Quantity:    item.Quantity,
}
if err := tx.Create(&oi).Error; err != nil {
tx.Rollback()
json(c, 500, "创建订单明细失败", nil)
return
}
}

tx.Where("user_id = ?", req.UserID).Delete(&model.Cart{})
tx.Commit()

database.DB.Preload("Items").First(&order, order.ID)
json(c, 200, "ok", order)
}

func GetOrders(c *gin.Context) {
userID := c.Query("user_id")
if userID == "" {
userID = "default"
}
var orders []model.Order
database.DB.Preload("Items").Where("user_id = ?", userID).Order("created_at desc").Find(&orders)
json(c, 200, "ok", orders)
}

func GetOrder(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
var order model.Order
if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
json(c, 404, "订单不存在", nil)
return
}
json(c, 200, "ok", order)
}
