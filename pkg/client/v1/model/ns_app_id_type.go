// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAppIdType Enum.

type NsAppIdType string

const (
	cNsAppIdTypeExchangeDag NsAppIdType = "exchange_dag"
	cNsAppIdTypeSql2012     NsAppIdType = "sql2012"
	cNsAppIdTypeSql2014     NsAppIdType = "sql2014"
	cNsAppIdTypeInval       NsAppIdType = "inval"
	cNsAppIdTypeSql2005     NsAppIdType = "sql2005"
	cNsAppIdTypeSql2016     NsAppIdType = "sql2016"
	cNsAppIdTypeExchange    NsAppIdType = "exchange"
	cNsAppIdTypeSql2017     NsAppIdType = "sql2017"
	cNsAppIdTypeSql2008     NsAppIdType = "sql2008"
	cNsAppIdTypeHyperv      NsAppIdType = "hyperv"
)

var pNsAppIdTypeExchangeDag NsAppIdType
var pNsAppIdTypeSql2012 NsAppIdType
var pNsAppIdTypeSql2014 NsAppIdType
var pNsAppIdTypeInval NsAppIdType
var pNsAppIdTypeSql2005 NsAppIdType
var pNsAppIdTypeSql2016 NsAppIdType
var pNsAppIdTypeExchange NsAppIdType
var pNsAppIdTypeSql2017 NsAppIdType
var pNsAppIdTypeSql2008 NsAppIdType
var pNsAppIdTypeHyperv NsAppIdType

// Export
var NsAppIdTypeExchangeDag *NsAppIdType
var NsAppIdTypeSql2012 *NsAppIdType
var NsAppIdTypeSql2014 *NsAppIdType
var NsAppIdTypeInval *NsAppIdType
var NsAppIdTypeSql2005 *NsAppIdType
var NsAppIdTypeSql2016 *NsAppIdType
var NsAppIdTypeExchange *NsAppIdType
var NsAppIdTypeSql2017 *NsAppIdType
var NsAppIdTypeSql2008 *NsAppIdType
var NsAppIdTypeHyperv *NsAppIdType

func init() {
	pNsAppIdTypeExchangeDag = cNsAppIdTypeExchangeDag
	NsAppIdTypeExchangeDag = &pNsAppIdTypeExchangeDag

	pNsAppIdTypeSql2012 = cNsAppIdTypeSql2012
	NsAppIdTypeSql2012 = &pNsAppIdTypeSql2012

	pNsAppIdTypeSql2014 = cNsAppIdTypeSql2014
	NsAppIdTypeSql2014 = &pNsAppIdTypeSql2014

	pNsAppIdTypeInval = cNsAppIdTypeInval
	NsAppIdTypeInval = &pNsAppIdTypeInval

	pNsAppIdTypeSql2005 = cNsAppIdTypeSql2005
	NsAppIdTypeSql2005 = &pNsAppIdTypeSql2005

	pNsAppIdTypeSql2016 = cNsAppIdTypeSql2016
	NsAppIdTypeSql2016 = &pNsAppIdTypeSql2016

	pNsAppIdTypeExchange = cNsAppIdTypeExchange
	NsAppIdTypeExchange = &pNsAppIdTypeExchange

	pNsAppIdTypeSql2017 = cNsAppIdTypeSql2017
	NsAppIdTypeSql2017 = &pNsAppIdTypeSql2017

	pNsAppIdTypeSql2008 = cNsAppIdTypeSql2008
	NsAppIdTypeSql2008 = &pNsAppIdTypeSql2008

	pNsAppIdTypeHyperv = cNsAppIdTypeHyperv
	NsAppIdTypeHyperv = &pNsAppIdTypeHyperv

}
