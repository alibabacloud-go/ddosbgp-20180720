// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddIpRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetInstanceId(v string) *AddIpRequest {
	s.InstanceId = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

func (s *AddIpRequest) SetRegionId(v string) *AddIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpRequest) SetResourceGroupId(v string) *AddIpRequest {
	s.ResourceGroupId = &v
	return s
}

type AddIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpResponseBody) SetRequestId(v string) *AddIpResponseBody {
	s.RequestId = &v
	return s
}

type AddIpResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetHeaders(v map[string]*string) *AddIpResponse {
	s.Headers = v
	return s
}

func (s *AddIpResponse) SetStatusCode(v int32) *AddIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

type CheckAccessLogAuthRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckAccessLogAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthRequest) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthRequest) SetRegionId(v string) *CheckAccessLogAuthRequest {
	s.RegionId = &v
	return s
}

func (s *CheckAccessLogAuthRequest) SetResourceGroupId(v string) *CheckAccessLogAuthRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckAccessLogAuthResponseBody struct {
	AccessLogAuth *bool   `json:"AccessLogAuth,omitempty" xml:"AccessLogAuth,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccessLogAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponseBody) SetAccessLogAuth(v bool) *CheckAccessLogAuthResponseBody {
	s.AccessLogAuth = &v
	return s
}

func (s *CheckAccessLogAuthResponseBody) SetRequestId(v string) *CheckAccessLogAuthResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccessLogAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckAccessLogAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAccessLogAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponse) SetHeaders(v map[string]*string) *CheckAccessLogAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckAccessLogAuthResponse) SetStatusCode(v int32) *CheckAccessLogAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccessLogAuthResponse) SetBody(v *CheckAccessLogAuthResponseBody) *CheckAccessLogAuthResponse {
	s.Body = v
	return s
}

type CheckGrantRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckGrantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

type CheckGrantResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetHeaders(v map[string]*string) *CheckGrantResponse {
	s.Headers = v
	return s
}

func (s *CheckGrantResponse) SetStatusCode(v int32) *CheckGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

type ConfigSchedruleOnDemandRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ConfigSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandRequest) SetInstanceId(v string) *ConfigSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRegionId(v string) *ConfigSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleAction(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleName(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleSwitch(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoMode(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetTimeZone(v string) *ConfigSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type ConfigSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponseBody) SetRequestId(v string) *ConfigSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type ConfigSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *ConfigSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetStatusCode(v int32) *ConfigSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetBody(v *ConfigSchedruleOnDemandResponseBody) *ConfigSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type CreateSchedruleOnDemandRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandRequest) SetInstanceId(v string) *CreateSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRegionId(v string) *CreateSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleAction(v string) *CreateSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleName(v string) *CreateSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleSwitch(v string) *CreateSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoMode(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetTimeZone(v string) *CreateSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type CreateSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponseBody) SetRequestId(v string) *CreateSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type CreateSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *CreateSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetStatusCode(v int32) *CreateSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetBody(v *CreateSchedruleOnDemandResponseBody) *CreateSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DeleteBlackholeRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetInstanceId(v string) *DeleteBlackholeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetRegionId(v string) *DeleteBlackholeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetResourceGroupId(v string) *DeleteBlackholeRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteBlackholeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBlackholeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponseBody) SetRequestId(v string) *DeleteBlackholeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBlackholeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetHeaders(v map[string]*string) *DeleteBlackholeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackholeResponse) SetStatusCode(v int32) *DeleteBlackholeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

type DeleteIpRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetInstanceId(v string) *DeleteIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetRegionId(v string) *DeleteIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpRequest) SetResourceGroupId(v string) *DeleteIpRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpResponseBody) SetRequestId(v string) *DeleteIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetHeaders(v map[string]*string) *DeleteIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpResponse) SetStatusCode(v int32) *DeleteIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpResponse) SetBody(v *DeleteIpResponseBody) *DeleteIpResponse {
	s.Body = v
	return s
}

type DeleteSchedruleOnDemandRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandRequest) SetInstanceId(v string) *DeleteSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRegionId(v string) *DeleteSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRuleName(v string) *DeleteSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

type DeleteSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponseBody) SetRequestId(v string) *DeleteSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *DeleteSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetStatusCode(v int32) *DeleteSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetBody(v *DeleteSchedruleOnDemandResponseBody) *DeleteSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DescribeDdosEventRequest struct {
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetInstanceId(v string) *DescribeDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetIp(v string) *DescribeDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetRegionId(v string) *DescribeDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetResourceGroupId(v string) *DescribeDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeDdosEventResponseBody struct {
	Events    []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventResponseBodyEvents struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mbps      *int32  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeDdosEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetHeaders(v map[string]*string) *DescribeDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventResponse) SetStatusCode(v int32) *DescribeDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

type DescribeExcpetionCountRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeExcpetionCountResponseBody struct {
	ExceptionIpCount *int32  `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty"`
	ExpireTimeCount  *int32  `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcpetionCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponseBody) SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExpireTimeCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetRequestId(v string) *DescribeExcpetionCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExcpetionCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExcpetionCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExcpetionCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponse) SetHeaders(v map[string]*string) *DescribeExcpetionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcpetionCountResponse) SetStatusCode(v int32) *DescribeExcpetionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse {
	s.Body = v
	return s
}

type DescribeInstanceListRequest struct {
	InstanceIdList  *string                           `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	InstanceType    *string                           `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Ip              *string                           `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpVersion       *string                           `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Orderby         *string                           `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	Orderdire       *string                           `json:"Orderdire,omitempty" xml:"Orderdire,omitempty"`
	PageNo          *int32                            `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark          *string                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceGroupId *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceType(v string) *DescribeInstanceListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIp(v string) *DescribeInstanceListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIpVersion(v string) *DescribeInstanceListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderby(v string) *DescribeInstanceListRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderdire(v string) *DescribeInstanceListRequest {
	s.Orderdire = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRegionId(v string) *DescribeInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRemark(v string) *DescribeInstanceListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListRequest) SetResourceGroupId(v string) *DescribeInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest {
	s.Tag = v
	return s
}

type DescribeInstanceListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequestTag) SetKey(v string) *DescribeInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceListRequestTag) SetValue(v string) *DescribeInstanceListRequestTag {
	s.Value = &v
	return s
}

type DescribeInstanceListResponseBody struct {
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceListResponseBodyInstanceList struct {
	AutoRenewal       *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IpType            *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.BlackholdingCount = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetIpType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetHeaders(v map[string]*string) *DescribeInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceListResponse) SetStatusCode(v int32) *DescribeInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	InstanceIdList  *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIdList(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetRegionId(v string) *DescribeInstanceSpecsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetResourceGroupId(v string) *DescribeInstanceSpecsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	AvailableDefenseTimes         *int32                                                    `json:"AvailableDefenseTimes,omitempty" xml:"AvailableDefenseTimes,omitempty"`
	AvailableDeleteBlackholeCount *string                                                   `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	InstanceId                    *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsFullDefenseMode             *int32                                                    `json:"IsFullDefenseMode,omitempty" xml:"IsFullDefenseMode,omitempty"`
	PackConfig                    *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	Region                        *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	TotalDefenseTimes             *int32                                                    `json:"TotalDefenseTimes,omitempty" xml:"TotalDefenseTimes,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetIsFullDefenseMode(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.IsFullDefenseMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PackConfig = v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetTotalDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.TotalDefenseTimes = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig struct {
	Bandwidth       *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BindIpCount     *int32 `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty"`
	IpAdvanceThre   *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	IpBasicThre     *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	IpSpec          *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	NormalBandwidth *int32 `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	PackAdvThre     *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	PackBasicThre   *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBandwidth(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBindIpCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpAdvanceThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpSpec(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetNormalBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.NormalBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackAdvThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetStatusCode(v int32) *DescribeInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeOnDemandDdosEventRequest struct {
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOnDemandDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventRequest) SetEndTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetInstanceId(v string) *DescribeOnDemandDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetIp(v string) *DescribeOnDemandDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageNo(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageSize(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetRegionId(v string) *DescribeOnDemandDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetResourceGroupId(v string) *DescribeOnDemandDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetStartTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeOnDemandDdosEventResponseBody struct {
	Events    []*DescribeOnDemandDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBody) SetEvents(v []*DescribeOnDemandDdosEventResponseBodyEvents) *DescribeOnDemandDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetRequestId(v string) *DescribeOnDemandDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetTotal(v int64) *DescribeOnDemandDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeOnDemandDdosEventResponseBodyEvents struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mbps      *int32  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetIp(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetPps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStatus(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeOnDemandDdosEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOnDemandDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOnDemandDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponse) SetHeaders(v map[string]*string) *DescribeOnDemandDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetStatusCode(v int32) *DescribeOnDemandDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetBody(v *DescribeOnDemandDdosEventResponseBody) *DescribeOnDemandDdosEventResponse {
	s.Body = v
	return s
}

type DescribeOnDemandInstanceStatusRequest struct {
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusRequest) SetInstanceIdList(v []*string) *DescribeOnDemandInstanceStatusRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeOnDemandInstanceStatusRequest) SetRegionId(v string) *DescribeOnDemandInstanceStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponseBody struct {
	Instances []*DescribeOnDemandInstanceStatusResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetInstances(v []*DescribeOnDemandInstanceStatusResponseBodyInstances) *DescribeOnDemandInstanceStatusResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetRequestId(v string) *DescribeOnDemandInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponseBodyInstances struct {
	Declared   *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Net        *string `json:"Net,omitempty" xml:"Net,omitempty"`
	RegistedAs *string `json:"RegistedAs,omitempty" xml:"RegistedAs,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDeclared(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Declared = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDesc(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Desc = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetInstanceId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetMode(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Mode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetNet(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Net = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetRegistedAs(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.RegistedAs = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetUserId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.UserId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOnDemandInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOnDemandInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeOnDemandInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetStatusCode(v int32) *DescribeOnDemandInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetBody(v *DescribeOnDemandInstanceStatusResponseBody) *DescribeOnDemandInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderBy         *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderDir        *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetInstanceId(v string) *DescribeOpEntitiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderBy(v string) *DescribeOpEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderDir(v string) *DescribeOpEntitiesRequest {
	s.OrderDir = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetRegionId(v string) *DescribeOpEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	EntityType   *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	OpAction     *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePackIpListRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductName     *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackIpListRequest) SetInstanceId(v string) *DescribePackIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackIpListRequest) SetIp(v string) *DescribePackIpListRequest {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageNo(v int32) *DescribePackIpListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageSize(v int32) *DescribePackIpListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackIpListRequest) SetProductName(v string) *DescribePackIpListRequest {
	s.ProductName = &v
	return s
}

func (s *DescribePackIpListRequest) SetRegionId(v string) *DescribePackIpListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackIpListRequest) SetResourceGroupId(v string) *DescribePackIpListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePackIpListResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	IpList    []*DescribePackIpListResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePackIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBody) SetCode(v string) *DescribePackIpListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribePackIpListResponseBody) SetRequestId(v string) *DescribePackIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetSuccess(v bool) *DescribePackIpListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetTotal(v int32) *DescribePackIpListResponseBody {
	s.Total = &v
	return s
}

type DescribePackIpListResponseBodyIpList struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePackIpListResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBodyIpList) SetIp(v string) *DescribePackIpListResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetProduct(v string) *DescribePackIpListResponseBodyIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRemark(v string) *DescribePackIpListResponseBodyIpList {
	s.Remark = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetStatus(v string) *DescribePackIpListResponseBodyIpList {
	s.Status = &v
	return s
}

type DescribePackIpListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponse) SetHeaders(v map[string]*string) *DescribePackIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribePackIpListResponse) SetStatusCode(v int32) *DescribePackIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackIpListResponse) SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTrafficRequest struct {
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowType        *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Ipnet           *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetFlowType(v string) *DescribeTrafficRequest {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficRequest) SetInstanceId(v string) *DescribeTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetIp(v string) *DescribeTrafficRequest {
	s.Ip = &v
	return s
}

func (s *DescribeTrafficRequest) SetIpnet(v string) *DescribeTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTrafficRequest) SetRegionId(v string) *DescribeTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTrafficRequest) SetResourceGroupId(v string) *DescribeTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeTrafficResponseBody struct {
	FlowList  []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBody) SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrafficResponseBody) SetRequestId(v string) *DescribeTrafficResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrafficResponseBodyFlowList struct {
	AttackBps *int64  `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	AttackPps *int64  `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	FlowType  *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Kbps      *int32  `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	Time      *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackBps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackPps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetName(v string) *DescribeTrafficResponseBodyFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetPps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

type DescribeTrafficResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetHeaders(v map[string]*string) *DescribeTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficResponse) SetStatusCode(v int32) *DescribeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

type GetSlsOpenStatusRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusRequest) SetRegionId(v string) *GetSlsOpenStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetSlsOpenStatusRequest) SetResourceGroupId(v string) *GetSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type GetSlsOpenStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsOpenStatus *bool   `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s GetSlsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponseBody) SetRequestId(v string) *GetSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

type GetSlsOpenStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponse) SetHeaders(v map[string]*string) *GetSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSlsOpenStatusResponse) SetStatusCode(v int32) *GetSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSlsOpenStatusResponse) SetBody(v *GetSlsOpenStatusResponseBody) *GetSlsOpenStatusResponse {
	s.Body = v
	return s
}

type ListOpenedAccessLogInstancesRequest struct {
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListOpenedAccessLogInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageNumber(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageSize(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBody struct {
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsConfigStatus []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
	TotalCount      *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetRequestId(v string) *ListOpenedAccessLogInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) *ListOpenedAccessLogInstancesResponseBody {
	s.SlsConfigStatus = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetTotalCount(v int32) *ListOpenedAccessLogInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBodySlsConfigStatus struct {
	Enable     *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetEnable(v bool) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetInstanceId(v string) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.InstanceId = &v
	return s
}

type ListOpenedAccessLogInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOpenedAccessLogInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenedAccessLogInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponse) SetHeaders(v map[string]*string) *ListOpenedAccessLogInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetStatusCode(v int32) *ListOpenedAccessLogInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetBody(v *ListOpenedAccessLogInstancesResponseBody) *ListOpenedAccessLogInstancesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken       *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType    *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag             []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyRemarkRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRemarkRequest) SetInstanceId(v string) *ModifyRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRegionId(v string) *ModifyRemarkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRemark(v string) *ModifyRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyRemarkRequest) SetResourceGroupId(v string) *ModifyRemarkRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponseBody) SetRequestId(v string) *ModifyRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRemarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponse) SetHeaders(v map[string]*string) *ModifyRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyRemarkResponse) SetStatusCode(v int32) *ModifyRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRemarkResponse) SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse {
	s.Body = v
	return s
}

type QuerySchedruleOnDemandRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandRequest) SetInstanceId(v string) *QuerySchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandRequest) SetRegionId(v string) *QuerySchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type QuerySchedruleOnDemandResponseBody struct {
	InstanceId *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleConfig []*QuerySchedruleOnDemandResponseBodyRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Repeated"`
	RuleStatus []*QuerySchedruleOnDemandResponseBodyRuleStatus `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty" type:"Repeated"`
	UserId     *string                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBody) SetInstanceId(v string) *QuerySchedruleOnDemandResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRequestId(v string) *QuerySchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleConfig(v []*QuerySchedruleOnDemandResponseBodyRuleConfig) *QuerySchedruleOnDemandResponseBody {
	s.RuleConfig = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleStatus(v []*QuerySchedruleOnDemandResponseBodyRuleStatus) *QuerySchedruleOnDemandResponseBody {
	s.RuleStatus = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetUserId(v string) *QuerySchedruleOnDemandResponseBody {
	s.UserId = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleConfig struct {
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleAction(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleAction = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionCnt(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionCnt = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionKpps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionKpps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionMbps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionMbps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleName(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleName = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleSwitch(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleSwitch = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoBeginTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoEndTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoEndTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoMode(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoMode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetTimeZone(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.TimeZone = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleStatus struct {
	Net             *string `json:"Net,omitempty" xml:"Net,omitempty"`
	RuleSchedStatus *string `json:"RuleSchedStatus,omitempty" xml:"RuleSchedStatus,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetNet(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.Net = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetRuleSchedStatus(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.RuleSchedStatus = &v
	return s
}

type QuerySchedruleOnDemandResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponse) SetHeaders(v map[string]*string) *QuerySchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetStatusCode(v int32) *QuerySchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetBody(v *QuerySchedruleOnDemandResponseBody) *QuerySchedruleOnDemandResponse {
	s.Body = v
	return s
}

type SetInstanceModeOnDemandRequest struct {
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetInstanceModeOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandRequest) SetInstanceIdList(v []*string) *SetInstanceModeOnDemandRequest {
	s.InstanceIdList = v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetMode(v string) *SetInstanceModeOnDemandRequest {
	s.Mode = &v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetRegionId(v string) *SetInstanceModeOnDemandRequest {
	s.RegionId = &v
	return s
}

type SetInstanceModeOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceModeOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponseBody) SetRequestId(v string) *SetInstanceModeOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceModeOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetInstanceModeOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstanceModeOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponse) SetHeaders(v map[string]*string) *SetInstanceModeOnDemandResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetStatusCode(v int32) *SetInstanceModeOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetBody(v *SetInstanceModeOnDemandResponseBody) *SetInstanceModeOnDemandResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType    *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag             []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":          tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAccessLogAuthWithOptions(request *CheckAccessLogAuthRequest, runtime *util.RuntimeOptions) (_result *CheckAccessLogAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccessLogAuth"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAccessLogAuth(request *CheckAccessLogAuthRequest) (_result *CheckAccessLogAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CheckAccessLogAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckGrant"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigSchedruleOnDemandWithOptions(request *ConfigSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigSchedruleOnDemand(request *ConfigSchedruleOnDemandRequest) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.ConfigSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSchedruleOnDemandWithOptions(request *CreateSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *CreateSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedruleOnDemand(request *CreateSchedruleOnDemandRequest) (_result *CreateSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CreateSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBlackhole"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSchedruleOnDemandWithOptions(request *DeleteSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchedruleOnDemand(request *DeleteSchedruleOnDemandRequest) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.DeleteSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *util.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExcpetionCount"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExcpetionCount(request *DescribeExcpetionCountRequest) (_result *DescribeExcpetionCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DescribeExcpetionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Orderby)) {
		query["Orderby"] = request.Orderby
	}

	if !tea.BoolValue(util.IsUnset(request.Orderdire)) {
		query["Orderdire"] = request.Orderdire
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecs"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOnDemandDdosEventWithOptions(request *DescribeOnDemandDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOnDemandDdosEvent(request *DescribeOnDemandDdosEventRequest) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.DescribeOnDemandDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOnDemandInstanceStatusWithOptions(request *DescribeOnDemandInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandInstanceStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOnDemandInstanceStatus(request *DescribeOnDemandInstanceStatusRequest) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.DescribeOnDemandInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDir)) {
		query["OrderDir"] = request.OrderDir
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpEntities"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *util.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackIpList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackIpList(request *DescribePackIpListRequest) (_result *DescribePackIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DescribePackIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FlowType)) {
		query["FlowType"] = request.FlowType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Ipnet)) {
		query["Ipnet"] = request.Ipnet
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraffic"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSlsOpenStatusWithOptions(request *GetSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *GetSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSlsOpenStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSlsOpenStatus(request *GetSlsOpenStatusRequest) (_result *GetSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.GetSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenedAccessLogInstancesWithOptions(request *ListOpenedAccessLogInstancesRequest, runtime *util.RuntimeOptions) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenedAccessLogInstances"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenedAccessLogInstances(request *ListOpenedAccessLogInstancesRequest) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.ListOpenedAccessLogInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRemark"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRemark(request *ModifyRemarkRequest) (_result *ModifyRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.ModifyRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySchedruleOnDemandWithOptions(request *QuerySchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *QuerySchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySchedruleOnDemand(request *QuerySchedruleOnDemandRequest) (_result *QuerySchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.QuerySchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstanceModeOnDemandWithOptions(request *SetInstanceModeOnDemandRequest, runtime *util.RuntimeOptions) (_result *SetInstanceModeOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstanceModeOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstanceModeOnDemand(request *SetInstanceModeOnDemandRequest) (_result *SetInstanceModeOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.SetInstanceModeOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
