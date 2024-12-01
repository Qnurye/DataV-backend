package models

type AdsDayPurchase struct {
	Dt          string `db:"dt" json:"dt"`
	PurchaseCnt int64  `db:"purchase_cnt" json:"purchase_cnt"`
}

type AdsProductAnalysis struct {
	ItemID          int64 `db:"item_id" json:"item_id"`
	BrowseCnt       int64 `db:"browse_cnt" json:"browse_cnt"`
	LikeCnt         int64 `db:"like_cnt" json:"like_cnt"`
	PurchaseCnt     int64 `db:"purchase_cnt" json:"purchase_cnt"`
	PopularityScore int64 `db:"popularity_score" json:"popularity_score"`
}

type AdsProductAnalysisSummary struct {
	BrowseCnt   int64 `db:"browse_cnt" json:"browse_cnt"`
	LikeCnt     int64 `db:"like_cnt" json:"like_cnt"`
	AddCartCnt  int64 `db:"add_cart_cnt" json:"add_cart_cnt"`
	PurchaseCnt int64 `db:"purchase_cnt" json:"purchase_cnt"`
}

type AdsTop10FrequentItemList struct {
	ItemList string `db:"item_list" json:"item_list"`
	Cnt      int64  `db:"cnt" json:"cnt"`
}

type ItemInformation struct {
	ID              string `db:"id" json:"id"`
	ItemID          string `db:"item_id" json:"item_id"`
	ItemCategory    string `db:"item_category" json:"item_category"`
	VisitCount      string `db:"visit_count" json:"visit_count"`
	PurchaseCount   string `db:"purchase_count" json:"purchase_count"`
	ViewCount       string `db:"view_count" json:"view_count"`
	AddToCartCount  string `db:"add_to_cart_count" json:"add_to_cart_count"`
	PopularityScore string `db:"popularity_score" json:"popularity_score"`
}

type ProvinceDistribution struct {
	ID         string `db:"id" json:"id"`
	ProvinceID string `db:"province_id" json:"province_id"`
	UserCount  string `db:"user_count" json:"user_count"`
}

type Top20PopularItemsByBehavior struct {
	ID           string `db:"id" json:"id"`
	BehaviorType string `db:"behavior_type" json:"behavior_type"`
	ItemID       string `db:"item_id" json:"item_id"`
	Popularity   string `db:"popularity" json:"popularity"`
}

type UserBehavior struct {
	ID           string `db:"id" json:"id"`
	UserID       string `db:"user_id" json:"user_id"`
	BehaviorType string `db:"behavior_type" json:"behavior_type"`
	ItemID       string `db:"item_id" json:"item_id"`
	ItemCategory string `db:"item_category" json:"item_category"`
	VisitDate    string `db:"visit_date" json:"visit_date"`
	ProvinceID   string `db:"province_id" json:"province_id"`
}

type UserOrderData struct {
	ID      string `db:"id" json:"id"`
	UID     string `db:"uid" json:"uid"`
	OrderID string `db:"order_id" json:"order_id"`
	ItemID  string `db:"item_id" json:"item_id"`
}

type ProvinceOrder struct {
	ProvinceID int64  `db:"province_id" json:"province_id"`
	Province   string `db:"province" json:"province"`
	OrdersCnt  int64  `db:"orders_cnt" json:"orders_cnt"`
}
