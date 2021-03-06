package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

/*
 Six910 is a shopping cart and E-commerce system.

 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// AddOrder godoc
// @Summary Add a new order
// @Description Adds a new order to a store
// @Tags Order
// @Accept  json
// @Produce  json
// @Param order body six910-database-interface.Order true "order"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/order/add [post]
func (h *Six910Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var addorURL = "/six910/rs/order/add"
	var aorc jv.Claim
	aorc.Role = customerRole
	aorc.URL = addorURL
	aorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	clientID := r.Header.Get("clientId")
	var auth bool
	if clientID != "" {
		auth = h.processSecurity(r, &aorc)
	} else {
		auth = h.processBasicSecurity(r, &aorc)
	}
	h.Log.Debug("order add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aor sdbi.Order
			aorsuc, aorerr := h.ProcessBody(r, &aor)
			h.Log.Debug("aorsuc: ", aorsuc)
			h.Log.Debug("aor: ", aor)
			h.Log.Debug("aorerr: ", aorerr)
			if !aorsuc && aorerr != nil {
				http.Error(w, aorerr.Error(), http.StatusBadRequest)
			} else {
				aorres := h.Manager.AddOrder(&aor)
				h.Log.Debug("aorres: ", *aorres)
				if aorres.Success && aorres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aorres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aorf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aorf)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateOrder godoc
// @Summary Update a order
// @Description Update order data
// @Tags Order
// @Accept  json
// @Produce  json
// @Param order body six910-database-interface.Order true "order"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/order/update [put]
func (h *Six910Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var uorURL = "/six910/rs/order/update"
	var uorc jv.Claim
	uorc.Role = customerRole
	uorc.URL = uorURL
	uorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uorc)
	h.Log.Debug("order update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uor sdbi.Order
			uorsuc, uorerr := h.ProcessBody(r, &uor)
			h.Log.Debug("uorsuc: ", uorsuc)
			h.Log.Debug("uor: ", uor)
			h.Log.Debug("uorerr: ", uorerr)
			if !uorsuc && uorerr != nil {
				http.Error(w, uorerr.Error(), http.StatusBadRequest)
			} else {
				uorres := h.Manager.UpdateOrder(&uor)
				h.Log.Debug("uorres: ", *uorres)
				if uorres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uorres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uorf m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uorf)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetOrder godoc
// @Summary Get details of a order by id
// @Description Get details of a order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "order id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Order
// @Router /rs/order/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var gorURL = "/six910/rs/order/get"
	var gorc jv.Claim
	gorc.Role = customerRole
	gorc.URL = gorURL
	gorc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	clientID := r.Header.Get("clientId")
	var auth bool
	if clientID != "" {
		auth = h.processSecurity(r, &gorc)
	} else {
		auth = h.processBasicSecurity(r, &gorc)
	}
	h.Log.Debug("order get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var goridStr = vars["id"]
			var gorstoreIDStr = vars["storeId"]
			id, goriderr := strconv.ParseInt(goridStr, 10, 64)
			storeID, gorsiderr := strconv.ParseInt(gorstoreIDStr, 10, 64)
			var gorres *sdbi.Order
			if goriderr == nil && gorsiderr == nil {
				gorres = h.Manager.GetOrder(id, storeID)
				h.Log.Debug("gorres: ", gorres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Order
				gorres = &nc
			}
			resJSON, _ := json.Marshal(gorres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetOrderList godoc
// @Summary Get list of order
// @Description Get list of order for a store
// @Tags Order
// @Accept  json
// @Produce  json
// @Param cid path string true "customer Id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Order
// @Router /rs/order/get/list/{cid}/{storeId} [get]
func (h *Six910Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {
	var gorlURL = "/six910/rs/order/list"
	var gorcl jv.Claim
	gorcl.Role = customerRole
	gorcl.URL = gorlURL
	gorcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	clientID := r.Header.Get("clientId")
	var auth bool
	if clientID != "" {
		auth = h.processSecurity(r, &gorcl)
	} else {
		auth = h.processBasicSecurity(r, &gorcl)
	}
	h.Log.Debug("order get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var orlcidStr = vars["cid"]
			var orlstoreIDStr = vars["storeId"]
			cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, sorlserr := strconv.ParseInt(orlstoreIDStr, 10, 64)
			var gorlres *[]sdbi.Order
			if sorlciderr == nil && sorlserr == nil {
				gorlres = h.Manager.GetOrderList(cID, storeID)
				h.Log.Debug("get order list: ", gorlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Order{}
				gorlres = &nc
			}
			resJSON, _ := json.Marshal(gorlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetStoreOrderList godoc
// @Summary Get list of store orders
// @Description Get list of order for a store
// @Tags Order
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Order
// @Router /rs/order/get/store/list/{storeId} [get]
func (h *Six910Handler) GetStoreOrderList(w http.ResponseWriter, r *http.Request) {
	var gorlURL = "/six910/rs/store/order/list"
	var gsorcl jv.Claim
	gsorcl.Role = customerRole
	gsorcl.URL = gorlURL
	gsorcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gsorcl)
	h.Log.Debug("order get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var sorlstoreIDStr = vars["storeId"]
			storeID, sorlserr := strconv.ParseInt(sorlstoreIDStr, 10, 64)
			var sgorlres *[]sdbi.Order
			if sorlserr == nil {
				sgorlres = h.Manager.GetStoreOrderList(storeID)
				h.Log.Debug("get store order list: ", sgorlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Order{}
				sgorlres = &nc
			}
			resJSON, _ := json.Marshal(sgorlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetStoreOrderListByStatus godoc
// @Summary Get list of store orders by status
// @Description Get list of order for a store
// @Tags Order
// @Accept  json
// @Produce  json
// @Param status path string true "order status"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Order
// @Router /rs/order/get/store/list/status/{status}/{storeId} [get]
func (h *Six910Handler) GetStoreOrderListByStatus(w http.ResponseWriter, r *http.Request) {
	var gorlURL = "/six910/rs/store/order/list/status"
	var gsosrcl jv.Claim
	gsosrcl.Role = customerRole
	gsosrcl.URL = gorlURL
	gsosrcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gsosrcl)
	h.Log.Debug("order get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var statusStr = vars["status"]
			var sosrlstoreIDStr = vars["storeId"]
			//cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, sorlserr := strconv.ParseInt(sosrlstoreIDStr, 10, 64)
			var sgsorlres *[]sdbi.Order
			if sorlserr == nil {
				sgsorlres = h.Manager.GetStoreOrderListByStatus(statusStr, storeID)
				h.Log.Debug("get store order list by status: ", sgsorlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Order{}
				sgsorlres = &nc
			}
			resJSON, _ := json.Marshal(sgsorlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetOrderCountData GetOrderCountData
func (h *Six910Handler) GetOrderCountData(w http.ResponseWriter, r *http.Request) {
	var gorcdURL = "/six910/rs/store/order/count/data"
	var goccl jv.Claim
	goccl.Role = customerRole
	goccl.URL = gorcdURL
	goccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &goccl)
	h.Log.Debug("order get count data authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var oscdstoreIDStr = vars["storeId"]
			//cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, orcdserr := strconv.ParseInt(oscdstoreIDStr, 10, 64)
			var gorcdres *[]sdbi.OrderCountData
			if orcdserr == nil {
				gorcdres = h.Manager.GetOrderCountData(storeID)
				h.Log.Debug("get order count data list: ", gorcdres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.OrderCountData{}
				gorcdres = &nc
			}
			resJSON, _ := json.Marshal(gorcdres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetOrderSalesData GetOrderSalesData
func (h *Six910Handler) GetOrderSalesData(w http.ResponseWriter, r *http.Request) {
	var gorsdURL = "/six910/rs/store/order/sales/data"
	var goscl jv.Claim
	goscl.Role = customerRole
	goscl.URL = gorsdURL
	goscl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &goscl)
	h.Log.Debug("order get sales data authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var ossdstoreIDStr = vars["storeId"]
			//cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, orsdserr := strconv.ParseInt(ossdstoreIDStr, 10, 64)
			var gorsdres *[]sdbi.OrderSalesData
			if orsdserr == nil {
				gorsdres = h.Manager.GetOrderSalesData(storeID)
				h.Log.Debug("get order sales data list: ", gorsdres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.OrderSalesData{}
				gorsdres = &nc
			}
			resJSON, _ := json.Marshal(gorsdres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteOrder godoc
// @Summary Delete a order
// @Description Delete a order from the store
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "order id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/order/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var dorURL = "/six910/rs/order/delete"
	var dorc jv.Claim
	dorc.Role = customerRole
	dorc.URL = dorURL
	dorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dorc)
	h.Log.Debug("order delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var doridStr = vars["id"]
			var dorstoreIDStr = vars["storeId"]
			id, doriderr := strconv.ParseInt(doridStr, 10, 64)
			storeID, doridserr := strconv.ParseInt(dorstoreIDStr, 10, 64)
			var dorres *m.Response
			if doriderr == nil && doridserr == nil {
				dorres = h.Manager.DeleteOrder(id, storeID)
				h.Log.Debug("dorres: ", *dorres)
				if dorres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dorres = &nc
			}
			resJSON, _ := json.Marshal(dorres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
