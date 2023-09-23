package internal

type otaServiceImpl struct{

}

func NewOTAService() OTAService{
	return &otaServiceImpl{}
}

