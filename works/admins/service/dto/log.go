package dto

type LogQueryDTO struct {
	UserName    string `form:"userName"`
	OperationIP string `form:"operationIp"`
	Status      *int   `form:"status"`
	StartTime   string `form:"startTime"`
	EndTime     string `form:"endTime"`
}

type OperationLogCreateDTO struct {
	UserID            int    `json:"userId"`
	UserName          string `json:"userName"`
	OperationIP       string `json:"operationIp"`
	OperationLocation string `json:"operationLocation"`
	OperationBrowser  string `json:"operationBrowser"`
	OperationOS       string `json:"operationOs"`
	OperationMethod   string `json:"operationMethod"`
	OperationPath     string `json:"operationPath"`
	OperationModule   string `json:"operationModule"`
	OperationContent  string `json:"operationContent"`
	OperationStatus   int    `json:"operationStatus"`
	OperationLatency  int    `json:"operationLatency"`
	OperationRequest  string `json:"operationRequest"`
	OperationResponse string `json:"operationResponse"`
}

type OperationLogInfo struct {
	LogID             int    `json:"logId"`
	UserID            int    `json:"userId"`
	UserName          string `json:"userName"`
	OperationIP       string `json:"operationIp"`
	OperationLocation string `json:"operationLocation"`
	OperationBrowser  string `json:"operationBrowser"`
	OperationOS       string `json:"operationOs"`
	OperationMethod   string `json:"operationMethod"`
	OperationPath     string `json:"operationPath"`
	OperationModule   string `json:"operationModule"`
	OperationContent  string `json:"operationContent"`
	OperationStatus   int    `json:"operationStatus"`
	OperationLatency  int    `json:"operationLatency"`
	OperationRequest  string `json:"operationRequest"`
	OperationResponse string `json:"operationResponse"`
	CreatedAt         string `json:"createdAt"`
}

type LoginLogCreateDTO struct {
	UserID        int    `json:"userId"`
	UserName      string `json:"userName"`
	LoginIP       string `json:"loginIp"`
	LoginLocation string `json:"loginLocation"`
	LoginBrowser  string `json:"loginBrowser"`
	LoginOS       string `json:"loginOs"`
	LoginStatus   int8   `json:"loginStatus"`
	LoginMessage  string `json:"loginMessage"`
}

type LoginLogInfo struct {
	LogID         int    `json:"logId"`
	UserID        int    `json:"userId"`
	UserName      string `json:"userName"`
	LoginIP       string `json:"loginIp"`
	LoginLocation string `json:"loginLocation"`
	LoginBrowser  string `json:"loginBrowser"`
	LoginOS       string `json:"loginOs"`
	LoginStatus   int8   `json:"loginStatus"`
	LoginMessage  string `json:"loginMessage"`
	CreatedAt     string `json:"createdAt"`
}
