package modeladapter

type UserValuesRequest struct {
	UsrId string `json:"usrId"`
}

type CacOrderTerminalsRequest struct {
	Dealer   *string   `json:"dealer"`
	Entidad  *string   `json:"entidad"`
	ID       *string   `json:"id"`
	Pin      *string   `json:"pin"`
	ProID    *string   `json:"proId"`
	Reserva  *string   `json:"reserva"`
	Status   *[]string `json:"status"`
	UsrPlace *string   `json:"usrPlace"`
}

type UserTaskRequest struct {
	UsrIds []string `json:"usrIds"`
}

type UserTaskPriorityRequest struct {
	UsrIds []string `json:"usrIds"`
}
