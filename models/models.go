package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	ID        uint 		`gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Goods struct {
	BaseModel
	Goods_name string			//商品名称
}

type Form struct {
	User_id int					//用户id
	Goods_id int				//商品id
	Shop_id int					//商户id
	Buy_num int					//购买数量
	Buy_time string			//购买时间
}

type Shop struct {
	BaseModel
	Goods_id int				//商品id
	Number int					//商品数量
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:zhy123@/e-commerce?charset=utf8&parseTime=True")
	if err == nil {

		DB = db
		db.AutoMigrate(&Goods{}, &Form{}, &Shop{})
		//db.Model(&)
		return db, err
	}
	return nil, err
}

//goods
func (goods *Goods) Insert() error{
	return DB.Create(goods).Error
}

func (goods *Goods) Update() error{
	return DB.Save(goods).Error
}


//form
func (form *Form) Insert() error {
	return DB.Create(form).Error
}

func (form *Form) Update() error {
	return DB.Save(form).Error
}


//shop
func (shop *Shop) Insert() error {
	return DB.Create(shop).Error
}

func GetShopByID(shop_id int) (*Shop, error){
	var shop Shop
	err := DB.Where("id = ?", shop_id).Find(&shop).Error
	return &shop, err

}

func Update(goods_id int, number int, buy_num int)  error{
	rest := number-buy_num
	fmt.Println("商品剩余数量: ", rest)
	var shop Shop
	return DB.Model(&shop).Where("goods_id = ?", goods_id).Update("number",rest).Error
}
