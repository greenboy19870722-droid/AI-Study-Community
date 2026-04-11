 package entity

 type User struct {
      Id           uint64	`json:"id"           dc:"用户ID"` 
      Username     string	`json:"username"        dc:"用户名" v:"required|length:1,50"`
  }