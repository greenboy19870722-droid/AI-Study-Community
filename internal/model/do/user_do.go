package do

 type UserReq struct {
      Username string `json:"username" v:"required"`
  }

  type UserResp struct {
      Id	   uint64    `json:"id"`
  }