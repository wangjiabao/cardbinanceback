package biz

import (
	"bytes"
	pb "cardbinance/api/user/v1"
	"cardbinance/internal/pkg/middleware/auth"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	jwt2 "github.com/golang-jwt/jwt/v5"
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

type Admin struct {
	ID       int64
	Password string
	Account  string
	Type     string
}

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
	VipTwo        uint64
	VipThree      uint64
	CardTwo       uint64
	CanVip        uint64
	UserCount     uint64
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
	GetUserByCard(card string) (*User, error)
	GetUserByCardUserId(cardUserId string) (*User, error)
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
	UpdateCardNo(ctx context.Context, userId uint64, amount float64) error
	UpdateCardSucces(ctx context.Context, userId uint64, cardNum string) error
	CreateCardRecommend(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error
	CreateCardRecommendTwo(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error
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
	InsertCardRecord(ctx context.Context, userId, recordType uint64, remark string, code string, opt string) error
	UpdateCardTwo(ctx context.Context, id uint64) error
	GetUserCardTwo() ([]*Reward, error)
	GetUsers(b *Pagination, address string) ([]*User, error, int64)
	GetAdminByAccount(ctx context.Context, account string, password string) (*Admin, error)
	SetCanVip(ctx context.Context, userId uint64, lock uint64) (bool, error)
	SetVipThree(ctx context.Context, userId uint64, vipThree uint64) (bool, error)
	SetUserCount(ctx context.Context, userId uint64) (bool, error)
	GetConfigs() ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
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
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}

			continue
		}

		//
		var (
			resHolder *QueryCardHolderResponse
		)

		resHolder, err = QueryCardHolderWithSign(holderId, productIdUseInt64)
		if nil == resHolder || err != nil || 200 != resHolder.Code {
			fmt.Println(user, err, "持卡人信息请求错误", resHolder)
			continue
		}

		if "active" == resHolder.Data.Status {

		} else if "pending" == resHolder.Data.Status {
			continue
		} else {
			fmt.Println(user, err, "持卡人创建失败", resHolder)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		resCreatCard, err = CreateCardRequestWithSign(0, holderId, productIdUseInt64)
		if nil == resCreatCard || 200 != resCreatCard.Code || err != nil {
			fmt.Println("开卡订单创建失败", user, resCreatCard, err)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}
		fmt.Println("开卡信息：", user, resCreatCard)

		if 0 >= len(resCreatCard.Data.CardID) || 0 >= len(resCreatCard.Data.CardOrderID) {
			fmt.Println("开卡订单信息错误", resCreatCard, err)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
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
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
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

		tmpTopVip := uint64(10)
		if 30 == user.VipTwo {
			tmpTopVip = 30
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

			if usersMap[tmpUserId].VipTwo != user.VipTwo {
				fmt.Println("开卡遍历，信息缺失，不是一个vip区域：", usersMap[tmpUserId], user)
				continue
			}

			if tmpTopVip < usersMap[tmpUserId].Vip {
				fmt.Println("开卡遍历，vip信息设置错误：", usersMap[tmpUserId], lastVip)
				break
			}

			// 小于等于上一个级别，跳过
			if usersMap[tmpUserId].Vip <= lastVip {
				continue
			}

			tmpAmount := usersMap[tmpUserId].Vip - lastVip // 极差
			lastVip = usersMap[tmpUserId].Vip

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

var cardTwoStatusLockHandle sync.Mutex

func (uuc *UserUseCase) CardTwoStatusHandle(ctx context.Context) error {
	cardTwoStatusLockHandle.Lock()
	defer cardTwoStatusLockHandle.Unlock()

	var (
		userOpenCard []*Reward
		err          error
	)

	var (
		configs       []*Config
		vipThreeThree uint64
		vipThreeTwo   uint64
		vipThreeOne   uint64
	)

	// 配置
	configs, err = uuc.repo.GetConfigByKeys("vip_three_three", "vip_three_two", "vip_three_one")
	if nil != configs {
		for _, vConfig := range configs {
			if "vip_three_three" == vConfig.KeyName {
				vipThreeThree, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "vip_three_two" == vConfig.KeyName {
				vipThreeTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "vip_three_one" == vConfig.KeyName {
				vipThreeOne, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
		}
	}

	userOpenCard, err = uuc.repo.GetUserCardTwo()
	if nil != err {
		return err
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = uuc.repo.GetAllUsers()
	if nil == users {
		fmt.Println("开卡2，用户无")
		return nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	for _, userCard := range userOpenCard {
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateCardTwo(ctx, userCard.ID)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward 2", err, userCard)
			continue
		}

		if _, ok := usersMap[userCard.UserId]; !ok {
			fmt.Println("开卡2，信息缺失：", userCard)
			continue
		}
		user := usersMap[userCard.UserId]

		// 分红
		var (
			userRecommend *UserRecommend
		)
		tmpRecommendUserIds := make([]string, 0)
		// 推荐
		userRecommend, err = uuc.repo.GetUserRecommendByUserId(user.ID)
		if nil == userRecommend {
			fmt.Println(err, "开卡2，信息错误", err, user)
			return nil
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		tmpTopVip := uint64(3)
		totalTmp := len(tmpRecommendUserIds) - 1
		lastVip := uint64(0)
		lastAmount := uint64(0)
		for i := totalTmp; i >= 0; i-- {
			if vipThreeThree <= lastAmount {
				break
			}

			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("开卡2遍历，信息缺失：", tmpUserId)
				continue
			}

			if tmpTopVip < usersMap[tmpUserId].VipThree {
				fmt.Println("开卡2遍历，vip信息设置错误：", usersMap[tmpUserId], lastVip)
				break
			}

			// 小于等于上一个级别，跳过
			if usersMap[tmpUserId].VipThree <= lastVip {
				continue
			}
			lastVip = usersMap[tmpUserId].VipThree

			// 奖励
			tmpAmount := uint64(0)
			if 1 == usersMap[tmpUserId].VipThree {
				if vipThreeOne <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误1：", usersMap[tmpUserId], lastVip, vipThreeOne, lastAmount)
					continue
				}

				tmpAmount = vipThreeOne - lastAmount
				lastAmount = vipThreeOne
			} else if 2 == usersMap[tmpUserId].VipThree {
				if vipThreeTwo <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误2：", usersMap[tmpUserId], lastVip, vipThreeTwo, lastAmount)
					continue
				}

				tmpAmount = vipThreeTwo - lastAmount
				lastAmount = vipThreeTwo
			} else if 3 == usersMap[tmpUserId].VipThree {
				if vipThreeThree <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误3：", usersMap[tmpUserId], lastVip, vipThreeThree, lastAmount)
					continue
				}

				tmpAmount = vipThreeThree - lastAmount
				lastAmount = vipThreeThree
			} else {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.CreateCardRecommendTwo(ctx, tmpUserId, float64(tmpAmount), usersMap[tmpUserId].Vip, user.Address)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward 2", err, user, usersMap[tmpUserId])
			}
		}
	}

	return nil
}

func (uuc *UserUseCase) backCard(ctx context.Context, userId uint64, amount float64) error {
	var (
		err error
	)
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateCardNo(ctx, userId, amount)
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

func (uuc *UserUseCase) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest, ca string) (*pb.AdminLoginReply, error) {
	var (
		admin *Admin
		err   error
	)

	res := &pb.AdminLoginReply{}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	fmt.Println(password)
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != err {
		return res, err
	}

	claims := auth.CustomClaims{
		UserId:   uint64(admin.ID),
		UserType: "admin",
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2天过期
			Issuer:    "game",
		},
	}

	token, err := auth.CreateToken(claims, ca)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return res, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *pb.AdminRewardListRequest) (*pb.AdminRewardListReply, error) {
	var (
		userSearch  *User
		userId      uint64 = 0
		userRewards []*Reward
		users       map[uint64]*User
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		err         error
		count       int64
	)
	res := &pb.AdminRewardListReply{
		Rewards: make([]*pb.AdminRewardListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(req.Address)
		if nil != err || nil == userSearch {
			return res, nil
		}

		userId = userSearch.ID
	}

	userRewards, err, count = uuc.repo.GetUserRewardByUserIdPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.Reason)
	if nil != err {
		return res, nil
	}
	res.Count = uint64(count)

	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		res.Rewards = append(res.Rewards, &pb.AdminRewardListReply_List{
			CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:     fmt.Sprintf("%.2f", vUserReward.Amount),
			Address:    tmpUser,
			Reason:     vUserReward.Reason,
			AddressTwo: vUserReward.Address,
			One:        vUserReward.One,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListReply, error) {
	var (
		users   []*User
		userIds []uint64
		count   int64
		err     error
	)

	res := &pb.AdminUserListReply{
		Users: make([]*pb.AdminUserListReply_UserList, 0),
	}

	users, err, count = uuc.repo.GetUsers(&Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address)
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		myLowUser         map[uint64][]*UserRecommend
		userRecommendsMap map[uint64]*UserRecommend
	)

	myLowUser = make(map[uint64][]*UserRecommend, 0)
	userRecommendsMap = make(map[uint64]*UserRecommend, 0)

	userRecommends, err = uuc.repo.GetUserRecommends()
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId uint64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		usersAll []*User
		usersMap map[uint64]*User
	)
	usersAll, err = uuc.repo.GetAllUsers()
	if nil == usersAll {
		return nil, nil
	}
	usersMap = make(map[uint64]*User, 0)

	for _, vUsers := range usersAll {
		usersMap[vUsers.ID] = vUsers
	}

	for _, vUsers := range users {
		// 推荐人
		var (
			userRecommend *UserRecommend
		)

		addressMyRecommend := ""
		if _, ok := userRecommendsMap[vUsers.ID]; ok {
			userRecommend = userRecommendsMap[vUsers.ID]

			if nil != userRecommend && "" != userRecommend.RecommendCode {
				var (
					tmpRecommendUserIds   []string
					myUserRecommendUserId uint64
				)
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
				if 2 <= len(tmpRecommendUserIds) {
					myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
				}

				if 0 < myUserRecommendUserId {
					if _, ok2 := usersMap[myUserRecommendUserId]; ok2 {
						addressMyRecommend = usersMap[myUserRecommendUserId].Address
					}
				}
			}
		}

		lenUsers := uint64(0)
		if _, ok := myLowUser[vUsers.ID]; ok {
			lenUsers = uint64(len(myLowUser[vUsers.ID]))
		}

		res.Users = append(res.Users, &pb.AdminUserListReply_UserList{
			UserId:             vUsers.ID,
			CreatedAt:          vUsers.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:            vUsers.Address,
			Amount:             fmt.Sprintf("%.2f", vUsers.Amount),
			Vip:                vUsers.Vip,
			CanVip:             vUsers.CanVip,
			VipThree:           vUsers.VipThree,
			MyRecommendAddress: addressMyRecommend,
			HistoryRecommend:   lenUsers,
			MyTotalAmount:      vUsers.MyTotalAmount,
			Card:               vUsers.Card,
			CardNumber:         vUsers.CardNumber,
			CardOrderId:        vUsers.CardOrderId,
			UserCount:          vUsers.UserCount,
			VipTwo:             vUsers.VipTwo,
			CardTwo:            vUsers.CardTwo,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) UpdateCanVip(ctx context.Context, req *pb.UpdateCanVipRequest) (*pb.UpdateCanVipReply, error) {
	var (
		err  error
		lock uint64
	)

	res := &pb.UpdateCanVipReply{}

	if 1 == req.SendBody.CanVip {
		lock = 1
	} else {
		lock = 0
	}

	_, err = uuc.repo.SetCanVip(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) SetVipThree(ctx context.Context, req *pb.SetVipThreeRequest) (*pb.SetVipThreeReply, error) {
	var (
		err  error
		lock uint64
	)

	res := &pb.SetVipThreeReply{}

	if 1 == req.SendBody.VipThree {
		lock = 1
	} else if 2 == req.SendBody.VipThree {
		lock = 2
	} else if 3 == req.SendBody.VipThree {
		lock = 3
	} else {
		lock = 0
	}

	_, err = uuc.repo.SetVipThree(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *pb.AdminConfigUpdateRequest) (*pb.AdminConfigUpdateReply, error) {
	var (
		err error
	)

	res := &pb.AdminConfigUpdateReply{}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = uuc.repo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *pb.AdminConfigRequest) (*pb.AdminConfigReply, error) {
	var (
		configs []*Config
	)

	res := &pb.AdminConfigReply{
		Config: make([]*pb.AdminConfigReply_List, 0),
	}

	configs, _ = uuc.repo.GetConfigs()
	if nil == configs {
		return res, nil
	}

	for _, v := range configs {
		res.Config = append(res.Config, &pb.AdminConfigReply_List{
			Id:    int64(v.ID),
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) SetUserCount(ctx context.Context, req *pb.SetUserCountRequest) (*pb.SetUserCountReply, error) {
	var (
		err error
	)

	res := &pb.SetUserCountReply{}

	_, err = uuc.repo.SetUserCount(ctx, req.SendBody.UserId)
	if nil != err {
		return res, err
	}

	return res, nil
}

type CardUserHandle struct {
	MerchantId string `json:"merchantId"`
	HolderId   string `json:"holderId"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

type CardCreateData struct {
	MerchantId string `json:"merchantId"`
	//ReferenceCode string `json:"referenceCode"`
	Remark     string `json:"remark"`
	CardId     string `json:"cardId"`
	CardNumber string `json:"cardNumber"`
	//Opt string `json:"opt"`
}

type RechargeData struct {
	MerchantId string `json:"merchantId"`
	//ReferenceCode string `json:"referenceCode"`
	//Opt string `json:"opt"`
	Remark     string `json:"remark"`
	CardId     string `json:"cardId"`
	CardNumber string `json:"cardNumber"`
}

func (uuc *UserUseCase) CallBackHandleOne(ctx context.Context, r *CardUserHandle) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCardUserId(r.HolderId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 1, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
}

func (uuc *UserUseCase) CallBackHandleTwo(ctx context.Context, r *CardCreateData) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCard(r.CardId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 2, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
}

func (uuc *UserUseCase) CallBackHandleThree(ctx context.Context, r *RechargeData) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCard(r.CardId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 3, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
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
	HolderId    string `json:"holderId"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birthDate"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	Status      string `json:"status"`
}

type QueryCardHolderResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data CardHolderData `json:"data"`
}

func QueryCardHolderWithSign(holderId uint64, productId uint64) (*QueryCardHolderResponse, error) {
	baseUrl := "https://www.ispay.com/prod-api/vcc/api/v1/cards/holders/query"

	// 请求体
	reqBody := map[string]interface{}{
		"holderId":   holderId,
		"merchantId": "322338",
		"productId":  productId,
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
