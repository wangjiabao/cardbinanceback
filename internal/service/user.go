package service

import (
	pb "cardbinance/api/user/v1"
	"cardbinance/internal/biz"
	"cardbinance/internal/conf"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"strconv"
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
