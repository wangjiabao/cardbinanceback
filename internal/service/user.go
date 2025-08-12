package service

import (
	pb "cardbinance/api/user/v1"
	"cardbinance/internal/biz"
	"cardbinance/internal/conf"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServer

	uuc *biz.UserUseCase
	log *log.Helper
	ca  *conf.Auth
}

func NewUserService(uuc *biz.UserUseCase, logger log.Logger, ca *conf.Auth) *UserService {
	return &UserService{uuc: uuc, log: log.NewHelper(logger), ca: ca}
}

func (u *UserService) OpenCardHandle(ctx context.Context, req *pb.OpenCardHandleRequest) (*pb.OpenCardHandleReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)

	var (
		err error
	)
	for i := 1; i <= 10; i++ {
		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		err = u.uuc.OpenCardHandle(ctx)
		if nil != err {
			fmt.Println(err)
		}
		time.Sleep(5 * time.Second)
	}

	return nil, nil
}

func (u *UserService) CardStatusHandle(ctx context.Context, req *pb.CardStatusHandleRequest) (*pb.CardStatusHandleReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)

	var (
		err error
	)
	for i := 1; i <= 10; i++ {
		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		err = u.uuc.CardStatusHandle(ctx)
		if nil != err {
			fmt.Println(err)
		}
		time.Sleep(5 * time.Second)
	}

	return nil, nil
}

func (u *UserService) RewardCardTwo(ctx context.Context, req *pb.RewardCardTwoRequest) (*pb.RewardCardTwoReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)

	var (
		err error
	)
	for i := 1; i <= 10; i++ {
		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		err = u.uuc.CardTwoStatusHandle(ctx)
		if nil != err {
			fmt.Println(err)
		}
		time.Sleep(5 * time.Second)
	}

	return nil, nil
}

func (u *UserService) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)

	for i := 1; i <= 10; i++ {
		var (
			depositUsdtResult []*userDeposit
			depositUsers      map[string]*biz.User
			fromAccount       []string
			userLength        int64
			last              int64
			err               error
		)

		last, err = u.uuc.GetEthUserRecordLast()
		if nil != err {
			fmt.Println(err)
			continue
		}

		if -1 == last {
			fmt.Println(err)
			continue
		}

		// 0x0299e92df88c034F6425e78b6f6A367e84160B45 test
		// 0x5d4bAA2A7a73dEF7685d036AAE993662B0Ef2f8F rel
		userLength, err = getUserLength("0x0876D2b69D53Bf6e5710Aa6b46ea3a739596F864")
		if nil != err {
			fmt.Println(err)
		}

		if -1 == userLength {
			continue
		}

		if 0 == userLength {
			break
		}

		if last >= userLength {
			break
		}

		// 0x0299e92df88c034F6425e78b6f6A367e84160B454 test
		// 0x5d4bAA2A7a73dEF7685d036AAE993662B0Ef2f8F rel
		depositUsdtResult, err = getUserInfo(last, userLength-1, "0x0876D2b69D53Bf6e5710Aa6b46ea3a739596F864")
		if nil != err {
			break
		}

		now := time.Now().UTC()
		//fmt.Println(now, end)
		if end.Before(now) {
			break
		}

		if 0 >= len(depositUsdtResult) {
			break
		}

		for _, vUser := range depositUsdtResult {
			fromAccount = append(fromAccount, vUser.Address)
		}

		depositUsers, err = u.uuc.GetUserByAddress(fromAccount...)
		if nil != depositUsers {
			// 统计开始
			for _, vUser := range depositUsdtResult { // 主查usdt
				if _, ok := depositUsers[vUser.Address]; !ok { // 用户不存在
					continue
				}

				var (
					tmpValue int64
				)

				if 10 <= vUser.Amount {
					tmpValue = vUser.Amount
				} else {
					return &pb.DepositReply{}, nil
				}

				// 充值
				err = u.uuc.DepositNew(ctx, depositUsers[vUser.Address].ID, uint64(tmpValue), &biz.EthUserRecord{ // 两种币的记录
					UserId:    int64(depositUsers[vUser.Address].ID),
					Amount:    strconv.FormatInt(tmpValue, 10) + "00000000000000000000",
					AmountTwo: uint64(vUser.Amount),
					Last:      userLength,
				}, false)
				if nil != err {
					fmt.Println(err)
				}
			}
		}

		time.Sleep(5 * time.Second)
	}

	return nil, nil
}

func FloatTo18DecimalsString(f float64) string {
	// 最多保留 18 位小数（足够精度）
	str := strconv.FormatFloat(f, 'f', -1, 64)

	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""

	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	// 计算还需要补多少个0
	padZeros := 18 - len(decimalPart)
	if padZeros < 0 {
		// 多余则截断
		decimalPart = decimalPart[:18]
	} else {
		// 不足则补0
		decimalPart += strings.Repeat("0", padZeros)
	}

	return integerPart + decimalPart
}

func (u *UserService) AdminWithdrawEth(ctx context.Context, req *pb.AdminWithdrawEthRequest) (*pb.AdminWithdrawEthReply, error) {
	var (
		withdraw     *biz.Withdraw
		userIds      []uint64
		userIdsMap   map[uint64]uint64
		users        map[uint64]*biz.User
		tokenAddress string
		err          error
	)
	end := time.Now().UTC().Add(50 * time.Second)

	for j := 1; j <= 10; j++ {
		now := time.Now().UTC()
		//fmt.Println(now, end)
		if end.Before(now) {
			break
		}

		withdraw, err = u.uuc.GetWithdrawPassOrRewardedFirst(ctx)
		if nil == withdraw {
			break
		}

		userIdsMap = make(map[uint64]uint64, 0)
		userIdsMap[withdraw.UserId] = withdraw.UserId
		for _, v := range userIdsMap {
			userIds = append(userIds, v)
		}

		users, err = u.uuc.GetUserByUserIds(userIds...)
		if nil != err {
			return nil, err
		}

		if _, ok := users[withdraw.UserId]; !ok {
			continue
		}

		tokenAddress = "0x55d398326f99059fF775485246999027B3197955"
		_, err = u.uuc.UpdateWithdrawDoing(ctx, withdraw.ID)
		if nil != err {
			continue
		}

		tmpUrl1 := "https://bsc-dataseed4.binance.org/"
		withDrawAmount := FloatTo18DecimalsString(withdraw.RelAmount)
		if len(withDrawAmount) <= 15 {
			fmt.Println(withDrawAmount, withdraw)
			_, err = u.uuc.UpdateWithdrawSuccess(ctx, withdraw.ID)
			continue
		}

		for i := 0; i <= 5; i++ {
			//fmt.Println(11111, user.ToAddress, v.Amount, balanceInt)
			_, err = toToken("", users[withdraw.UserId].Address, withDrawAmount, tokenAddress, tmpUrl1)
			if err == nil {
				_, err = u.uuc.UpdateWithdrawSuccess(ctx, withdraw.ID)
				//time.Sleep(3 * time.Second)
				break
			} else {
				fmt.Println(err)
				if 0 == i {
					tmpUrl1 = "https://bsc-dataseed1.binance.org"
				} else if 1 == i {
					tmpUrl1 = "https://bsc-dataseed3.binance.org"
				} else if 2 == i {
					tmpUrl1 = "https://bsc-dataseed2.binance.org"
				} else if 3 == i {
					tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com/"
				} else if 4 == i {
					tmpUrl1 = "https://bsc-dataseed.binance.org"
				}
				fmt.Println(33331, err, users[withdraw.UserId].Address, withdraw.Address, withDrawAmount, tokenAddress)
				time.Sleep(3 * time.Second)
			}
		}

		time.Sleep(5 * time.Second)
	}

	return &pb.AdminWithdrawEthReply{}, nil
}

func (u *UserService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	return u.uuc.AdminLogin(ctx, req, u.ca.JwtKey)
}

func (u *UserService) AdminRewardList(ctx context.Context, req *pb.AdminRewardListRequest) (*pb.AdminRewardListReply, error) {
	return u.uuc.AdminRewardList(ctx, req)
}

func (u *UserService) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListReply, error) {
	return u.uuc.AdminUserList(ctx, req)
}

func (u *UserService) UpdateCanVip(ctx context.Context, req *pb.UpdateCanVipRequest) (*pb.UpdateCanVipReply, error) {
	return u.uuc.UpdateCanVip(ctx, req)
}

func (u *UserService) SetVipThree(ctx context.Context, req *pb.SetVipThreeRequest) (*pb.SetVipThreeReply, error) {
	return u.uuc.SetVipThree(ctx, req)
}

func (u *UserService) SetUserCount(ctx context.Context, req *pb.SetUserCountRequest) (*pb.SetUserCountReply, error) {
	return u.uuc.SetUserCount(ctx, req)
}

func (u *UserService) AdminConfig(ctx context.Context, req *pb.AdminConfigRequest) (*pb.AdminConfigReply, error) {
	return u.uuc.AdminConfig(ctx, req)
}

func (u *UserService) AdminConfigUpdate(ctx context.Context, req *pb.AdminConfigUpdateRequest) (*pb.AdminConfigUpdateReply, error) {
	return u.uuc.AdminConfigUpdate(ctx, req)
}

type CallbackRequest struct {
	Version   string          `json:"version"`
	EventName string          `json:"eventName"`
	EventType string          `json:"eventType"`
	EventId   string          `json:"eventId"`
	SourceId  string          `json:"sourceId"`
	Data      json.RawMessage `json:"data"` // 用 RawMessage 接收动态结构
}

func (u *UserService) CallBack(w http.ResponseWriter, r *http.Request) {
	// 从 http.Request 获取 context.Context
	ctx := r.Context()

	//body, _ := io.ReadAll(r.Body)
	//fmt.Println("Raw Body:", string(body))

	var req CallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 只要 Version 和 EventType，例如
	//fmt.Println("Version:", req.Version)
	//fmt.Println("EventType:", req.EventType)
	//fmt.Println("auth:", r.Header.Get("Authorization"))
	//for k, v := range r.Header {
	//	fmt.Println("Header:", k, v)
	//}

	eventType := req.EventType

	switch {
	case strings.HasPrefix(eventType, "vcc.card.recharge.fai"):
		var rechargeData *biz.RechargeData
		if err := json.Unmarshal(req.Data, &rechargeData); err != nil {
			fmt.Println("Parse recharge data failed:", err, string(req.Data))
		}
		_ = u.uuc.CallBackHandleThree(ctx, rechargeData)

	case strings.HasPrefix(eventType, "vcc.cardholder.create.fail"):
		var cardholderData *biz.CardUserHandle
		if err := json.Unmarshal(req.Data, &cardholderData); err != nil {
			fmt.Println("Parse cardholder data failed:", err, string(req.Data))
		}
		_ = u.uuc.CallBackHandleOne(ctx, cardholderData)

	case strings.HasPrefix(eventType, "vcc.card.create.fai"):
		var createData *biz.CardCreateData
		if err := json.Unmarshal(req.Data, &createData); err != nil {
			fmt.Println("Parse create data failed:", err, string(req.Data))
		}
		_ = u.uuc.CallBackHandleTwo(ctx, createData)

	default:
		fmt.Println("Unhandled event type:", eventType, string(req.Data))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func getUserLength(address string) (int64, error) {
	url1 := "https://bsc-dataseed4.binance.org/"

	var balInt int64
	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals, err := instance.GetUserLength(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		balInt = bals.Int64()
		break
	}

	return balInt, nil
}

type userDeposit struct {
	Address string
	Amount  int64
}

func getUserInfo(start int64, end int64, address string) ([]*userDeposit, error) {
	url1 := "https://bsc-dataseed4.binance.org/"

	var (
		bals  []common.Address
		bals2 []*big.Int
	)
	users := make([]*userDeposit, 0)

	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals, err = instance.GetUsersByIndex(&bind.CallOpts{}, new(big.Int).SetInt64(start), new(big.Int).SetInt64(end))
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		break
	}

	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals2, err = instance.GetUsersAmountByIndex(&bind.CallOpts{}, new(big.Int).SetInt64(start), new(big.Int).SetInt64(end))
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		break
	}

	if len(bals) != len(bals2) {
		fmt.Println("数量不一致，错误")
		return users, nil
	}

	for k, v := range bals {
		users = append(users, &userDeposit{
			Address: v.String(),
			Amount:  bals2[k].Int64(),
		})
	}

	return users, nil
}

func toToken(userPrivateKey string, toAccount string, withdrawAmount string, withdrawTokenAddress string, url1 string) (string, error) {
	client, err := ethclient.Dial(url1)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(withdrawTokenAddress)
	instance, err := NewDfil(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tmpWithdrawAmount, _ := new(big.Int).SetString(withdrawAmount, 10)
	_, err = instance.Transfer(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	}, common.HexToAddress(toAccount), tmpWithdrawAmount)
	if err != nil {
		return "", err
	}

	return "", nil
}
