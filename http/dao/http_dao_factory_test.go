package httpdao

import "testing"

func TestNewHttpDao(t *testing.T) {
	var dao *HttpDao
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao("dao1", "https://host1.com", defaultHeaders)
	t.Logf(dao.Url)

	if dao.Url != "https://host1.com" {
		t.Fatalf("Url does not match the url that was passed to the factory method")
	}
}

func TestGetHttpDao(t *testing.T) {
	var dao1, dao2, dao3, dao4, dao5 *HttpDao
	defaultHeaders := make(map[string]string)

	dao1 = NewHttpDao("dao1", "https://host1.com", defaultHeaders)
	dao2 = NewHttpDao("dao2", "https://host2.com", defaultHeaders)

	dao3 = GetHttpDao("dao1")
	dao4 = GetHttpDao("dao2")
	dao5 = GetHttpDao("dao3")

	t.Logf(dao1.Url)
	t.Logf(dao2.Url)
	t.Logf(dao3.Url)
	t.Logf(dao4.Url)

	if dao1.Url != "https://host1.com" {
		t.Fatalf("Url 1 does not match the url that was passed to the factory method")
	}

	if dao1.Url != "https://host1.com" {
		t.Fatalf("Url 2 does not match the url that was passed to the factory method")
	}

	if dao1 != dao3 {
		t.Fatalf("dao1 and dao3 should be the same object")
	}

	if dao2 != dao4 {
		t.Fatalf("dao2 and dao4 should be the same object")
	}

	if dao5 != nil {
		t.Fatalf("dao5 should be nil because no dao object was registered under the name 'dao3'")
	}
}
