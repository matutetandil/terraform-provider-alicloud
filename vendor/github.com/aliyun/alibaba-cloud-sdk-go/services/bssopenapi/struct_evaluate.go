package bssopenapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Evaluate is a nested struct in bssopenapi response
type Evaluate struct {
	Id                 int    `json:"Id" xml:"Id"`
	GmtCreate          string `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified        string `json:"GmtModified" xml:"GmtModified"`
	UserId             int    `json:"UserId" xml:"UserId"`
	UserNick           string `json:"UserNick" xml:"UserNick"`
	OutBizId           string `json:"OutBizId" xml:"OutBizId"`
	BillId             int    `json:"BillId" xml:"BillId"`
	ItemId             int    `json:"ItemId" xml:"ItemId"`
	BillCycle          string `json:"BillCycle" xml:"BillCycle"`
	BizType            string `json:"BizType" xml:"BizType"`
	OriginalAmount     int    `json:"OriginalAmount" xml:"OriginalAmount"`
	PresentAmount      int    `json:"PresentAmount" xml:"PresentAmount"`
	CanInvoiceAmount   int    `json:"CanInvoiceAmount" xml:"CanInvoiceAmount"`
	InvoicedAmount     int    `json:"InvoicedAmount" xml:"InvoicedAmount"`
	OffsetCostAmount   int    `json:"OffsetCostAmount" xml:"OffsetCostAmount"`
	OffsetAcceptAmount int    `json:"OffsetAcceptAmount" xml:"OffsetAcceptAmount"`
	Status             int    `json:"Status" xml:"Status"`
	OpId               string `json:"OpId" xml:"OpId"`
	Name               string `json:"Name" xml:"Name"`
	BizTime            string `json:"BizTime" xml:"BizTime"`
	Type               int    `json:"Type" xml:"Type"`
}