package common

const (
	Get     = "GET"

	APIContextPattern = "/{context:(?:%s|)}%s"

	APIHealthUrl        = "health"
	APIHealthStatusUrl = APIHealthUrl + "/emperor"

	ContextPath = "/"
)