package model

type UserRequest struct {
	UsrId string `json:"usrId" validate:"required"`
}

type Status struct {
	Status      int
	Description string
}

type User struct {
	AccountingItemsInfo struct {
		PendingIntermidiateMaterials int `json:"pendingIntermidiateMaterials"`
		PendingMaterials             int `json:"pendingMaterials"`
		PendingProducts              int `json:"pendingProducts"`
	} `json:"accountingItemsInfo"`
	EntitiesInfo struct {
		Canal        string `json:"canal"`
		Dealer       string `json:"dealer"`
		DealerDni    string `json:"dealerDni"`
		Entidad      string `json:"entidad"`
		EntidadPadre string `json:"entidadPadre"`
		EntidadTipo  string `json:"entidadTipo"`
		Estructura   string `json:"estructura"`
	} `json:"entitiesInfo"`
	UserInfo struct {
		DescripcionGrupo string   `json:"descripcionGrupo"`
		Grupo            string   `json:"grupo"`
		TaskPriority     string   `json:"taskPriority"`
		Tasks            []string `json:"tasks"`
		Usuario          string   `json:"usuario"`
	} `json:"userInfo"`
	WarehouseInfo struct {
		Almacen               string `json:"almacen"`
		Centro                string `json:"centro"`
		UserPlace             string `json:"userPlace"`
		UserPlaceCentralizing string `json:"userPlaceCentralizing"`
		UserPlaceDescription  string `json:"userPlaceDescription"`
		UserPlaceParent       string `json:"userPlaceParent"`
	} `json:"warehouseInfo"`
}
