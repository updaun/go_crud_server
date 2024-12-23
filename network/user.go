package network

import (
	"fmt"
	"sync"

	"example.com/m/types"
	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다")

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다")
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		User:        nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다")

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다")

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}
