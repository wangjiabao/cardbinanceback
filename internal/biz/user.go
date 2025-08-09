package biz

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID            uint64
	Address       string
	Card          string
	CardNumber    string
	CardOrderId   string
	CardAmount    float64
	Amount        float64
	AmountTwo     uint64
	MyTotalAmount uint64
	IsDelete      uint64
	Vip           uint64
	FirstName     string
	LastName      string
	BirthDate     string
	Email         string
	CountryCode   string
	Phone         string
	City          string
	Country       string
	Street        string
	PostalCode    string
	CardUserId    string
	ProductId     string
	MaxCardQuota  uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserRecommend struct {
	ID            uint64
	UserId        uint64
	RecommendCode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Config struct {
	ID      uint64
	KeyName string
	Name    string
	Value   string
}

type Withdraw struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	RelAmount float64
	Status    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reward struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	Reason    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   string
	One       uint64
}

type EthUserRecord struct {
	ID        int64
	UserId    int64
	Hash      string
	Amount    string
	AmountTwo uint64
	Last      int64
	CreatedAt time.Time
}

type UserRepo interface {
	SetNonceByAddress(ctx context.Context, wallet string) (int64, error)
	GetAndDeleteWalletTimestamp(ctx context.Context, wallet string) (string, error)
	GetConfigByKeys(keys ...string) ([]*Config, error)
	GetUserByAddress(address string) (*User, error)
	GetUserById(userId uint64) (*User, error)
	GetUserRecommendByUserId(userId uint64) (*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateUserRecommend(ctx context.Context, userId uint64, recommendUser *UserRecommend) (*UserRecommend, error)
	GetUserRecommendByCode(code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(code string) ([]*UserRecommend, error)
	GetUserByUserIds(userIds ...uint64) (map[uint64]*User, error)
	CreateCard(ctx context.Context, userId uint64, user *User) error
	GetAllUsers() ([]*User, error)
	UpdateCard(ctx context.Context, userId uint64, cardOrderId, card string) error
	UpdateCardNo(ctx context.Context, userId uint64) error
	UpdateCardSucces(ctx context.Context, userId uint64, cardNum string) error
	CreateCardRecommend(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error
	GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error)
	AmountTo(ctx context.Context, userId, toUserId uint64, toAddress string, amount float64) error
	Withdraw(ctx context.Context, userId uint64, amount, amountRel float64, address string) error
	GetUserRewardByUserIdPage(ctx context.Context, b *Pagination, userId uint64, reason uint64) ([]*Reward, error, int64)
	SetVip(ctx context.Context, userId uint64, vip uint64) error
	GetUsersOpenCard() ([]*User, error)
	GetUsersOpenCardStatusDoing() ([]*User, error)
	GetEthUserRecordLast() (int64, error)
	GetUserByAddresses(Addresses ...string) (map[string]*User, error)
	GetUserRecommends() ([]*UserRecommend, error)
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
	UpdateUserMyTotalAmountAdd(ctx context.Context, userId uint64, amount uint64) error
	UpdateWithdraw(ctx context.Context, id uint64, status string) (*Withdraw, error)
}

type UserUseCase struct {
	repo UserRepo
	tx   Transaction
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		tx:   tx,
		log:  log.NewHelper(logger),
	}
}

type Pagination struct {
	PageNum  int
	PageSize int
}

// 后台

func (uuc *UserUseCase) GetEthUserRecordLast() (int64, error) {
	return uuc.repo.GetEthUserRecordLast()
}
func (uuc *UserUseCase) GetUserByAddress(Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(Addresses...)
}

func (uuc *UserUseCase) DepositNew(ctx context.Context, userId uint64, amount uint64, eth *EthUserRecord, system bool) error {
	// 推荐人
	var (
		err error
	)

	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		// 充值记录
		if !system {
			_, err = uuc.repo.CreateEthUserRecordListByHash(ctx, &EthUserRecord{
				Hash:      eth.Hash,
				UserId:    eth.UserId,
				Amount:    eth.Amount,
				AmountTwo: amount,
				Last:      eth.Last,
			})
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", userId, amount)
		return err
	}

	// 推荐人
	var (
		userRecommend       *UserRecommend
		tmpRecommendUserIds []string
	)
	userRecommend, err = uuc.repo.GetUserRecommendByUserId(userId)
	if nil != err {
		return err
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {
		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		// 增加业绩
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateUserMyTotalAmountAdd(ctx, tmpUserId, amount)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, eth)
			continue
		}
	}

	return nil
}

var lockHandle sync.Mutex

func (uuc *UserUseCase) OpenCardHandle(ctx context.Context) error {
	lockHandle.Lock()
	defer lockHandle.Unlock()

	var (
		userOpenCard []*User
		err          error
	)

	userOpenCard, err = uuc.repo.GetUsersOpenCard()
	if nil != err {
		return err
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	//var (
	//	products          *CardProductListResponse
	//	productIdUse      string
	//	productIdUseInt64 uint64
	//	maxCardQuota      int
	//)
	//products, err = GetCardProducts()
	//if nil == products || nil != err {
	//	fmt.Println("产品信息错误1")
	//	return nil
	//}
	//
	//for _, v := range products.Rows {
	//	if 0 < len(v.ProductId) && "ENABLED" == v.ProductStatus {
	//		productIdUse = v.ProductId
	//		maxCardQuota = v.MaxCardQuota
	//		productIdUseInt64, err = strconv.ParseUint(productIdUse, 10, 64)
	//		if nil != err {
	//			fmt.Println("产品信息错误2")
	//			return nil
	//		}
	//		fmt.Println("当前选择产品信息", productIdUse, maxCardQuota, v)
	//		break
	//	}
	//}
	//
	//if 0 >= maxCardQuota {
	//	fmt.Println("产品信息错误3")
	//	return nil
	//}
	//
	//if 0 >= productIdUseInt64 {
	//	fmt.Println("产品信息错误4")
	//	return nil
	//}

	for _, user := range userOpenCard {
		//var (
		//	resCreatCardholder *CreateCardholderResponse
		//)
		//resCreatCardholder, err = CreateCardholderRequest(productIdUseInt64, user)
		//if nil == resCreatCardholder || 200 != resCreatCardholder.Code || err != nil {
		//	fmt.Println("持卡人订单创建失败", user, resCreatCardholder, err)
		//	continue
		//}
		//if 0 > len(resCreatCardholder.Data.HolderID) {
		//	fmt.Println("持卡人订单信息错误", user, resCreatCardholder, err)
		//	continue
		//}
		//fmt.Println("持卡人信息", user, resCreatCardholder)
		//

		var (
			holderId          uint64
			productIdUseInt64 uint64
			resCreatCard      *CreateCardResponse
			openRes           = true
		)
		if 5 > len(user.CardUserId) {
			fmt.Println("持卡人id空", user)
			openRes = false
		}
		holderId, err = strconv.ParseUint(user.CardUserId, 10, 64)
		if nil != err {
			fmt.Println("持卡人错误2")
			openRes = false
		}
		if 0 >= holderId {
			fmt.Println("持卡人错误3")
			openRes = false
		}
		if 5 > len(user.CardUserId) {
			fmt.Println("持卡人id空", user)
			openRes = false
		}

		if 0 >= user.MaxCardQuota {
			fmt.Println("最大额度错误", user)
			openRes = false
		}

		if 5 > len(user.ProductId) {
			fmt.Println("productid空", user)
			openRes = false
		}
		productIdUseInt64, err = strconv.ParseUint(user.ProductId, 10, 64)
		if nil != err {
			fmt.Println("产品信息错误1")
			openRes = false
		}
		if 0 >= productIdUseInt64 {
			fmt.Println("产品信息错误2")
			openRes = false
		}

		if !openRes {
			fmt.Println("回滚了用户", user)
			err = uuc.backCard(ctx, user.ID)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}

			continue
		}

		resCreatCard, err = CreateCardRequestWithSign(0, holderId, productIdUseInt64)
		if nil == resCreatCard || 200 != resCreatCard.Code || err != nil {
			fmt.Println("开卡订单创建失败", user, resCreatCard, err)
			err = uuc.backCard(ctx, user.ID)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}
		fmt.Println("开卡信息：", user, resCreatCard)

		if 0 >= len(resCreatCard.Data.CardID) || 0 >= len(resCreatCard.Data.CardOrderID) {
			fmt.Println("开卡订单信息错误", resCreatCard, err)
			err = uuc.backCard(ctx, user.ID)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateCard(ctx, user.ID, resCreatCard.Data.CardOrderID, resCreatCard.Data.CardID)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "开卡后，写入mysql错误", err, user, resCreatCard)
			return nil
		}
	}

	return nil
}

var cardStatusLockHandle sync.Mutex

func (uuc *UserUseCase) CardStatusHandle(ctx context.Context) error {
	cardStatusLockHandle.Lock()
	defer cardStatusLockHandle.Unlock()

	var (
		userOpenCard []*User
		err          error
	)

	userOpenCard, err = uuc.repo.GetUsersOpenCardStatusDoing()
	if nil != err {
		return err
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = uuc.repo.GetAllUsers()
	if nil == users {
		fmt.Println("用户无")
		return nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	for _, user := range userOpenCard {
		//
		var (
			resHolder   *QueryCardHolderResponse
			tmpHolderId uint64
		)
		tmpHolderId, err = strconv.ParseUint(user.CardUserId, 10, 64)
		if err != nil || 0 >= tmpHolderId {
			fmt.Println(user, err, "解析持卡人id错误")
			continue
		}

		resHolder, err = QueryCardHolderWithSign(tmpHolderId)
		if nil == resHolder || err != nil || 200 != resHolder.Code {
			fmt.Println(user, err, "持卡人信息请求错误", resHolder)
			continue
		}

		if "active" == resHolder.Data.Status {

		} else if "pending" == resHolder.Data.Status {
			continue
		} else {
			fmt.Println(user, err, "持卡人创建失败", user, resHolder)
			err = uuc.backCard(ctx, user.ID)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		// 查询状态。成功分红
		var (
			resCard *CardInfoResponse
		)
		if 2 >= len(user.Card) {
			continue
		}

		resCard, err = GetCardInfoRequestWithSign(user.Card)
		if nil == resCard || 200 != resCard.Code || err != nil {
			fmt.Println(resCard, err)
			continue
		}

		if "ACTIVE" == resCard.Data.CardStatus {
			fmt.Println("开卡状态，激活：", resCard, user.ID)
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.UpdateCardSucces(ctx, user.ID, resCard.Data.Pan)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err，开卡成功", err, user.ID)
				continue
			}
		} else if "PENDING" == resCard.Data.CardStatus || "PROGRESS" == resCard.Data.CardStatus {
			fmt.Println("开卡状态，待处理：", resCard, user.ID)
			continue
		} else {
			fmt.Println("开卡状态，失败：", resCard, user.ID)
			err = uuc.backCard(ctx, user.ID)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		// 分红
		var (
			userRecommend *UserRecommend
		)
		tmpRecommendUserIds := make([]string, 0)
		// 推荐
		userRecommend, err = uuc.repo.GetUserRecommendByUserId(user.ID)
		if nil == userRecommend {
			fmt.Println(err, "信息错误", err, user)
			return nil
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		totalTmp := len(tmpRecommendUserIds) - 1
		lastVip := uint64(0)
		for i := totalTmp; i >= 0; i-- {
			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("开卡遍历，信息缺失：", tmpUserId)
				continue
			}

			if 0 == usersMap[tmpUserId].Vip && 0 == lastVip {
				continue
			}

			if lastVip >= usersMap[tmpUserId].Vip {
				fmt.Println("开卡遍历，vip信息设置错误：", usersMap[tmpUserId], lastVip)
				break
			}

			tmpRate := usersMap[tmpUserId].Vip - lastVip
			lastVip = usersMap[tmpUserId].Vip
			tmpAmount := tmpRate

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.CreateCardRecommend(ctx, tmpUserId, float64(tmpAmount), usersMap[tmpUserId].Vip, user.Address)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward", err, user, usersMap[tmpUserId])
			}
		}
	}

	return nil
}

func (uuc *UserUseCase) backCard(ctx context.Context, userId uint64) error {
	var (
		err error
	)
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateCardNo(ctx, userId)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("err")
		return err
	}

	return nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error) {
	return uuc.repo.GetWithdrawPassOrRewardedFirst(ctx)
}

func (uuc *UserUseCase) GetUserByUserIds(userIds ...uint64) (map[uint64]*User, error) {
	return uuc.repo.GetUserByUserIds(userIds...)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id uint64) (*Withdraw, error) {
	return uuc.repo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id uint64) (*Withdraw, error) {
	return uuc.repo.UpdateWithdraw(ctx, id, "success")
}

func GenerateSign(params map[string]interface{}, signKey string) string {
	// 1. 排除 sign 字段
	var keys []string
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接 key + value 字符串
	var sb strings.Builder
	sb.WriteString(signKey)

	for _, k := range keys {
		sb.WriteString(k)
		value := params[k]

		var strValue string
		switch v := value.(type) {
		case string:
			strValue = v
		case float64, int, int64, bool:
			strValue = fmt.Sprintf("%v", v)
		default:
			// map、slice 等复杂类型用 JSON 编码
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				strValue = ""
			} else {
				strValue = string(jsonBytes)
			}
		}
		sb.WriteString(strValue)
	}

	signString := sb.String()
	//fmt.Println("md5前字符串", signString)

	// 3. 进行 MD5 加密
	hash := md5.Sum([]byte(signString))
	return hex.EncodeToString(hash[:])
}

type CreateCardResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CardID      string `json:"cardId"`
		CardOrderID string `json:"cardOrderId"`
		CreateTime  string `json:"createTime"`
		CardStatus  string `json:"cardStatus"`
		OrderStatus string `json:"orderStatus"`
	} `json:"data"`
}

func CreateCardRequestWithSign(cardAmount uint64, cardholderId uint64, cardProductId uint64) (*CreateCardResponse, error) {
	//url := "https://test-api.ispay.com/dev-api/vcc/api/v1/cards/create"
	//url := "https://www.ispay.com/prod-api/vcc/api/v1/cards/create"
	baseUrl := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/create"

	reqBody := map[string]interface{}{
		"merchantId":    "322338",
		"cardCurrency":  "USD",
		"cardAmount":    cardAmount,
		"cardholderId":  cardholderId,
		"cardProductId": cardProductId,
		"cardSpendRule": map[string]interface{}{
			"dailyLimit":   250000,
			"monthlyLimit": 1000000,
		},
		"cardRiskControl": map[string]interface{}{
			"allowedMerchants": []string{"ONLINE"},
			"blockedCountries": []string{},
		},
	}

	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	// 请求体（包括嵌套结构）
	reqBody["sign"] = sign

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	//fmt.Println("请求报文:", string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	fmt.Println("响应报文:", string(body)) // ← 打印响应内容

	var result CreateCardResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("开卡，JSON 解析失败:", err)
		return nil, err
	}

	return &result, nil
}

type CardProductListResponse struct {
	Total int           `json:"total"`
	Rows  []CardProduct `json:"rows"`
	Code  int           `json:"code"`
	Msg   string        `json:"msg"`
}

type CardProduct struct {
	ProductId          string       `json:"productId"` // ← 改成 string
	ProductName        string       `json:"productName"`
	ModeType           string       `json:"modeType"`
	CardBin            string       `json:"cardBin"`
	CardForm           []string     `json:"cardForm"`
	MaxCardQuota       int          `json:"maxCardQuota"`
	CardScheme         string       `json:"cardScheme"`
	NoPinPaymentAmount []AmountItem `json:"noPinPaymentAmount"`
	CardCurrency       []string     `json:"cardCurrency"`
	CreateTime         string       `json:"createTime"`
	UpdateTime         string       `json:"updateTime"`
	ProductStatus      string       `json:"productStatus"`
}

type AmountItem struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func GetCardProducts() (*CardProductListResponse, error) {
	baseURL := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/products/all"

	reqBody := map[string]interface{}{
		"merchantId": "322338",
	}

	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")

	params := url.Values{}
	params.Set("merchantId", "322338")
	params.Set("sign", sign)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Language", "zh_CN")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println("响应报文:", string(body))

	var result CardProductListResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return nil, err
	}

	//fmt.Println(result)

	return &result, nil
}

type CreateCardholderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		HolderID    string `json:"holderId"`
		Email       string `json:"email"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		BirthDate   string `json:"birthDate"`
		CountryCode string `json:"countryCode"`
		PhoneNumber string `json:"phoneNumber"`

		DeliveryAddress DeliveryAddress `json:"deliveryAddress"`
		ProofFile       ProofFile       `json:"proofFile"`
	} `json:"data"`
}

type DeliveryAddress struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Street  string `json:"street"`
}

type ProofFile struct {
	FileBase64 string `json:"fileBase64"`
	FileType   string `json:"fileType"`
}

func CreateCardholderRequest(productId uint64, user *User) (*CreateCardholderResponse, error) {
	//baseURL := "https://www.ispay.com/prod-api/vcc/api/v1/cards/holders/create"
	baseURL := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/holders/create"

	reqBody := map[string]interface{}{
		"productId":   productId,
		"merchantId":  "322338",
		"email":       user.Email,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"birthDate":   user.BirthDate,
		"countryCode": user.CountryCode,
		"phoneNumber": user.Phone,
		"deliveryAddress": map[string]interface{}{
			"city":       user.City,
			"country":    user.Country,
			"street":     user.Street,
			"postalCode": user.PostalCode,
		},
	}

	// 生成签名
	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9") // 用你的密钥替换
	reqBody["sign"] = sign

	// 构造请求
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %v", err)
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("new request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http do error: %v", err)
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("响应报文:", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status not ok: %v", resp.StatusCode)
	}

	var result CreateCardholderResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &result, nil
}

type CardInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CardID     string `json:"cardId"`
		Pan        string `json:"pan"`
		CardStatus string `json:"cardStatus"`
		Holder     struct {
			HolderID string `json:"holderId"`
		} `json:"holder"`
	} `json:"data"`
}

func GetCardInfoRequestWithSign(cardId string) (*CardInfoResponse, error) {
	baseUrl := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/info"
	//baseUrl := "https://www.ispay.com/prod-api/vcc/api/v1/cards/info"

	reqBody := map[string]interface{}{
		"merchantId": "322338",
		"cardId":     cardId, // 如果需要传 cardId，根据实际接口文档添加
	}

	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	reqBody["sign"] = sign

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", string(body))
	}

	//fmt.Println("响应报文:", string(body))

	var result CardInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("卡信息 JSON 解析失败:", err)
		return nil, err
	}

	return &result, nil
}

type CardHolderData struct {
	HolderId        int64           `json:"holderId"`
	Email           string          `json:"email"`
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	Gender          string          `json:"gender"`
	BirthDate       string          `json:"birthDate"`
	CountryCode     string          `json:"countryCode"`
	PhoneNumber     string          `json:"phoneNumber"`
	Status          string          `json:"status"`
	DeliveryAddress DeliveryAddress `json:"deliveryAddress"`
	ProofFile       string          `json:"proofFile"`
}

type QueryCardHolderResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data CardHolderData `json:"data"`
}

func QueryCardHolderWithSign(holderId uint64) (*QueryCardHolderResponse, error) {
	baseUrl := "https://www.ispay.com/prod-api/vcc/api/v1/cards/holders/query"

	// 请求体
	reqBody := map[string]interface{}{
		"holderId":   holderId,
		"merchantId": "322338",
	}

	// 生成签名
	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	reqBody["sign"] = sign

	// 转 JSON
	jsonData, _ := json.Marshal(reqBody)

	// 打印调试
	//fmt.Println("签名:", sign)
	//fmt.Println("请求报文:", string(jsonData))

	// 创建 HTTP 请求
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	//fmt.Println("响应报文:", string(body))

	// 解析结果
	var result QueryCardHolderResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
