package dto

type LogQueryDTO struct {
	AdminName   string `form:"adminName"`
	OperationIP string `form:"operationIp"`
	Status      *int   `form:"status"`
	StartTime   string `form:"startTime"`
	EndTime     string `form:"endTime"`
}

type OperationLogCreateDTO struct {
	AdminID           int    `json:"adminId"`
	AdminName         string `json:"adminName"`
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
	AdminID           int    `json:"adminId"`
	AdminName         string `json:"adminName"`
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
