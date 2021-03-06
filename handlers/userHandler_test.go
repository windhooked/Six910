package handlers

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	lg "github.com/Ulbora/Level_Logger"
	man "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
	"github.com/gorilla/mux"
)

func TestSix910Handler_AddUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddUserRole(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"storeAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddCustomerUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddCustomerUserReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddCustomerUserMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddCustomerUserAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	//lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUser2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	//lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.Header.Set("clientId", "123456")
	r.SetBasicAuth("tester", "tester")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUserAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUserFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = "customer"

	sdb.MockLocalAccount = &lc

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	//lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("w.code", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUserMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	//lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateUserReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	//lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUpdateUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("superAdminRole", "superAdmin")
	w := httptest.NewRecorder()

	h.AdminUpdateUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUpdateUserAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	//sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("superAdminRole", "superAdmin")
	w := httptest.NewRecorder()

	h.AdminUpdateUser(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUpdateUserMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("superAdminRole", "superAdmin")
	w := httptest.NewRecorder()

	h.AdminUpdateUser(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUpdateUserReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("superAdminRole", "superAdmin")
	w := httptest.NewRecorder()

	h.AdminUpdateUser(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUpdateUserFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	//sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"customer", "username": "tester", "password": "tester", "customerId": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("superAdminRole", "superAdmin")
	w := httptest.NewRecorder()

	h.AdminUpdateUser(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_GetUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"storeId":  "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetUser2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"storeId":  "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.Header.Set("clientId", "123456")
	r.SetBasicAuth("tester", "tester")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetUser(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"storeId":  "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetUser(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserReq1(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		//"storeId":  "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetUser(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"storeId":  "5f",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetUser(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetAdminUserList(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = storeAdmin

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetAdminUserList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetCusUserList(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCustomerUserList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetAdminUserListAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = storeAdmin

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetAdminUserList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetCusUserListAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCustomerUserList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetAdminUserListReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = storeAdmin

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetAdminUserList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCusUserListReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCustomerUserList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetAdminUserListReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = storeAdmin

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5f",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetAdminUserList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCusUserListReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	//lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"storeId": "5f",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCustomerUserList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserByCustomer(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = false
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 3
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "StoreAdmin"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	r.Header.Set("apiKey", "123456")

	r.SetBasicAuth("tester", "tester1")

	w := httptest.NewRecorder()

	h.GetUsersByCustomer(w, r)

	fmt.Println("code:", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserByCustomerAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = false
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 3
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "StoreAdmin"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	//r.Header.Set("apiKey", "123456")

	r.SetBasicAuth("tester", "tester1")

	w := httptest.NewRecorder()

	h.GetUsersByCustomer(w, r)

	fmt.Println("code:", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserByCustomerReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = false
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 3
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "StoreAdmin"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "3c",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	r.Header.Set("apiKey", "123456")

	r.SetBasicAuth("tester", "tester1")

	w := httptest.NewRecorder()

	h.GetUsersByCustomer(w, r)

	fmt.Println("code:", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetUserByCustomerReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = false
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 3
	lc.UserName = "tester"
	lc.Enabled = true
	lc.Role = customerRole

	var lst []sdbi.LocalAccount
	lst = append(lst, lc)

	sdb.MockLocalAccountList = &lst

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "StoreAdmin"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid": "3",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	r.Header.Set("apiKey", "123456")

	r.SetBasicAuth("tester", "tester1")

	w := httptest.NewRecorder()

	h.GetUsersByCustomer(w, r)

	fmt.Println("code:", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_ResetCustomerUserPassword(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.ResetCustomerUserPassword(w, r)

	fmt.Println("Code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ResetCustomerUserPasswordAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.ResetCustomerUserPassword(w, r)

	fmt.Println("Code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_ResetCustomerUserPasswordMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.ResetCustomerUserPassword(w, r)

	fmt.Println("Code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_ResetCustomerUserPasswordFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockUpdateLocalAccountSuccess = false

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.ResetCustomerUserPassword(w, r)

	fmt.Println("Code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_ResetCustomerUserPasswordReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	lc.UserName = "tester"
	lc.Enabled = true

	sdb.MockLocalAccount = &lc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockUpdateLocalAccountSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"role":"StoreAdmin", "username": "tester", "password": "tester"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.ResetCustomerUserPassword(w, r)

	fmt.Println("Code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}
