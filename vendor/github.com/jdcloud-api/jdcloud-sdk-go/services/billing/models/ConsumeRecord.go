// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type ConsumeRecord struct {

    /* 消费记录数据库唯一id (Optional) */
    Id int `json:"id"`

    /* 消费记录id (Optional) */
    BillingRecordId int `json:"billingRecordId"`

    /* appCode (Optional) */
    AppCode string `json:"appCode"`

    /* serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 区域 (Optional) */
    Region string `json:"region"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 规格 (Optional) */
    Formula string `json:"formula"`

    /* 计费类型 (Optional) */
    BillingType int `json:"billingType"`

    /* 价格快照 (Optional) */
    PriceSnapShot string `json:"priceSnapShot"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 创建日期 (Optional) */
    CreateTime string `json:"createTime"`

    /* 账单金额 (Optional) */
    BillFee int `json:"billFee"`

    /* 账单金额保留小数点后2位 (Optional) */
    BillFee2 int `json:"billFee2"`

    /* 折扣金额 (Optional) */
    DiscountFee int `json:"discountFee"`

    /* 优惠券id (Optional) */
    CouponId string `json:"couponId"`

    /* 优惠金额 (Optional) */
    CouponFee int `json:"couponFee"`

    /* 交易单号 (Optional) */
    TransactionNo string `json:"transactionNo"`

    /* null (Optional) */
    IsBillGenerated int `json:"isBillGenerated"`

    /* 子账单id (Optional) */
    SubBillId int `json:"subBillId"`

    /* 退款单号 (Optional) */
    RefundNo string `json:"refundNo"`
}