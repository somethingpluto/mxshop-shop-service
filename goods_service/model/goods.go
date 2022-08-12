package model

// Category
// @Description: 目录表结构
//
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

// TableName
// @Description: 自定义表名
// @receiver Category
// @return string
//
func (Category) TableName() string {
	return "category"
}

// Brands
// @Description: 品牌表结构
//
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

// GoodsCategoryBrand
// @Description: 商品目录表结构
//
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

// TableName
// @Description: 自定义表名
// @receiver GoodsCategoryBrand
// @return string
//
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// Banner
// @Description: 横幅表结构
//
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

// TableName
// @Description: 自定义表名
// @receiver Banner
// @return string
//
func (Banner) TableName() string {
	return "banner"
}

// Goods
// @Description: 商品表结构
//
type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name            string   `gorm:"type:varchar(100);not null"`
	GoodsSn         string   `gorm:"type:varchar(50);not null"`
	ClickNum        int32    `gorm:"type:int;default:0;not null"`
	SoldNum         int32    `gorm:"type:int;default:0;not null"`
	FavNum          int32    `gorm:"type:int;default:0;not null"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`
	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`
}
