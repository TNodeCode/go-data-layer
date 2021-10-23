package httpdao

var httpDaos map[string]*HttpDao = make(map[string]*HttpDao)

func NewHttpDao(name string, url string, defaultHeaders map[string]string) *HttpDao {
	if dao, ok := httpDaos[name]; ok {
		return dao
	}

	httpDaos[name] = &HttpDao{
		Url:            url,
		DefaultHeaders: defaultHeaders,
	}

	return httpDaos[name]
}

func GetHttpDao(name string) *HttpDao {
	if dao, ok := httpDaos[name]; ok {
		return dao
	}

	return nil
}

func getDefaultHeaders() map[string]string {
	defaultHeaders := make(map[string]string)
	defaultHeaders["Content-Type"] = "application/json"
	return defaultHeaders
}
