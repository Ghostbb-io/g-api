# ginx
包裝gin常用函數

## GetToken(c *gin.Context) string
> Get jwt token from header (Authorization: Bearer xxx)
 
## ParseJSON(c *gin.Context, obj interface{}) error
> Parse body json data to struct

## ParseParamID(c *gin.Context, key string) uint64
> Param returns the value of the URL param

## ParseParam(c *gin.Context, key string) string
> Param returns the value of the URL param

## ParseQuery(c *gin.Context, obj interface{}) error
> Parse query parameter to struct

## ParseForm(c *gin.Context, obj interface{}) error
> Parse body form data to struct

## SetUserID(c *gin.Context, userID uint)
> 將UserID寫入上下文

## GetUserID(c *gin.Context) uint
> 獲取UserID

## SetUserUUID(c *gin.Context, uuid uuid.UUID)
> 將User UUID寫入上下文

## GetUserUUID(c *gin.Context) uuid.UUID
> 獲取User UUID

## SetUserName(c *gin.Context, username string)
> 將user name寫入上下文

## GetUserName(c *gin.Context) string
> 獲取user name

## SetRole(c *gin.Context, roles []string)
> 將role寫入上下文

## GetRole(c *gin.Context) []string
> 獲取role