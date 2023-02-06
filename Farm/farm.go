package Farm

type Farms struct {
	ID   int    `json:"ID" binding:"required"`
	Pond string `json:"pond" binding:"required"`
	Farm string `json:"farm" binding:"required"`
}
