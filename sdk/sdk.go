package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

type AncestorDepartment struct {
	DepartmentId *string `json:"departmentId,omitempty"` // 人员所在的分支id
	ExternalId   *string `json:"externalId,omitempty"`   // 部门数据源ID
	Name         *string `json:"name,omitempty"`         // 分支name
}

func (m *AncestorDepartment) SetDepartmentId(departmentId string) *AncestorDepartment {
	m.DepartmentId = &departmentId
	return m
}

func (m *AncestorDepartment) GetDepartmentId() string {
	if m != nil {
		return *m.DepartmentId
	}
	return ""
}

func (m *AncestorDepartment) SetExternalId(externalId string) *AncestorDepartment {
	m.ExternalId = &externalId
	return m
}

func (m *AncestorDepartment) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *AncestorDepartment) SetName(name string) *AncestorDepartment {
	m.Name = &name
	return m
}

func (m *AncestorDepartment) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

type AppRoamingOrgList struct {
	OrgId   *string `json:"orgId,omitempty"`   // 组织ID
	OrgName *string `json:"orgName,omitempty"` // 组织名称
	State   *int    `json:"state,omitempty"`   // 应用对接组织状态，0-正常， 1- 已删除
}

func (m *AppRoamingOrgList) SetOrgId(orgId string) *AppRoamingOrgList {
	m.OrgId = &orgId
	return m
}

func (m *AppRoamingOrgList) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *AppRoamingOrgList) SetOrgName(orgName string) *AppRoamingOrgList {
	m.OrgName = &orgName
	return m
}

func (m *AppRoamingOrgList) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

func (m *AppRoamingOrgList) SetState(state int) *AppRoamingOrgList {
	m.State = &state
	return m
}

func (m *AppRoamingOrgList) GetState() int {
	if m != nil {
		return *m.State
	}
	return 0
}

type BaseResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *BaseResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *BaseResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *BaseResponse) ToString() string {
	return ToJsonString(m)
}

type Department struct {
	ExternalId *string  `json:"externalId,omitempty"` // 外部数据源人员唯一ID
	Id         *string  `json:"id,omitempty"`         // 所在分支ID
	Name       *string  `json:"name,omitempty"`       // 所在分支名称
	Order      *float32 `json:"order,omitempty"`      // 人员排序值
}

func (m *Department) SetExternalId(externalId string) *Department {
	m.ExternalId = &externalId
	return m
}

func (m *Department) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *Department) SetId(id string) *Department {
	m.Id = &id
	return m
}

func (m *Department) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *Department) SetName(name string) *Department {
	m.Name = &name
	return m
}

func (m *Department) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *Department) SetOrder(order float32) *Department {
	m.Order = &order
	return m
}

func (m *Department) GetOrder() float32 {
	if m != nil {
		return *m.Order
	}
	return 0
}

type DepartmentData struct {
	AncestorDepartments []*Department `json:"ancestorDepartments,omitempty"` // 部门祖先列表
	ExternalId          *string       `json:"externalId,omitempty"`
	HasChildren         *bool         `json:"hasChildren,omitempty"`  // 是否含有子分支
	Id                  *string       `json:"id,omitempty"`           // 分支ID
	MembersCount        *int          `json:"membersCount,omitempty"` // 分支人员数
	Name                *string       `json:"name,omitempty"`         // 分支名称
}

func (m *DepartmentData) SetAncestorDepartments(ancestorDepartments []*Department) *DepartmentData {
	m.AncestorDepartments = ancestorDepartments
	return m
}

func (m *DepartmentData) GetAncestorDepartments() []*Department {
	if m != nil {
		return m.AncestorDepartments
	}
	return nil
}

func (m *DepartmentData) SetExternalId(externalId string) *DepartmentData {
	m.ExternalId = &externalId
	return m
}

func (m *DepartmentData) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *DepartmentData) SetHasChildren(hasChildren bool) *DepartmentData {
	m.HasChildren = &hasChildren
	return m
}

func (m *DepartmentData) GetHasChildren() bool {
	if m != nil {
		return *m.HasChildren
	}
	return false
}

func (m *DepartmentData) SetId(id string) *DepartmentData {
	m.Id = &id
	return m
}

func (m *DepartmentData) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *DepartmentData) SetMembersCount(membersCount int) *DepartmentData {
	m.MembersCount = &membersCount
	return m
}

func (m *DepartmentData) GetMembersCount() int {
	if m != nil {
		return *m.MembersCount
	}
	return 0
}

func (m *DepartmentData) SetName(name string) *DepartmentData {
	m.Name = &name
	return m
}

func (m *DepartmentData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

type DepartmentStaffInfo struct {
	AncestorDepartments []*Department `json:"ancestorDepartments,omitempty"` // 部门祖先列表
	AvatarId            *string       `json:"avatarId,omitempty"`
	Email               *string       `json:"email,omitempty"`
	ExternalId          *string       `json:"externalId,omitempty"` // 人员外部ID
	Id                  *string       `json:"id,omitempty"`
	Mobile              *string       `json:"mobile,omitempty"`
	MobilePhone         *MobilePhone  `json:"mobilePhone,omitempty"`
	Name                *string       `json:"name,omitempty"`
	OrgId               *int          `json:"orgId,omitempty"`
	OrgName             *string       `json:"orgName,omitempty"`
	ParentId            *string       `json:"parentId,omitempty"`
	Status              *int          `json:"status,omitempty"`
}

func (m *DepartmentStaffInfo) SetAncestorDepartments(ancestorDepartments []*Department) *DepartmentStaffInfo {
	m.AncestorDepartments = ancestorDepartments
	return m
}

func (m *DepartmentStaffInfo) GetAncestorDepartments() []*Department {
	if m != nil {
		return m.AncestorDepartments
	}
	return nil
}

func (m *DepartmentStaffInfo) SetAvatarId(avatarId string) *DepartmentStaffInfo {
	m.AvatarId = &avatarId
	return m
}

func (m *DepartmentStaffInfo) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *DepartmentStaffInfo) SetEmail(email string) *DepartmentStaffInfo {
	m.Email = &email
	return m
}

func (m *DepartmentStaffInfo) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *DepartmentStaffInfo) SetExternalId(externalId string) *DepartmentStaffInfo {
	m.ExternalId = &externalId
	return m
}

func (m *DepartmentStaffInfo) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *DepartmentStaffInfo) SetId(id string) *DepartmentStaffInfo {
	m.Id = &id
	return m
}

func (m *DepartmentStaffInfo) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *DepartmentStaffInfo) SetMobile(mobile string) *DepartmentStaffInfo {
	m.Mobile = &mobile
	return m
}

func (m *DepartmentStaffInfo) GetMobile() string {
	if m != nil {
		return *m.Mobile
	}
	return ""
}

func (m *DepartmentStaffInfo) SetMobilePhone(mobilePhone MobilePhone) *DepartmentStaffInfo {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *DepartmentStaffInfo) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *DepartmentStaffInfo) SetName(name string) *DepartmentStaffInfo {
	m.Name = &name
	return m
}

func (m *DepartmentStaffInfo) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *DepartmentStaffInfo) SetOrgId(orgId int) *DepartmentStaffInfo {
	m.OrgId = &orgId
	return m
}

func (m *DepartmentStaffInfo) GetOrgId() int {
	if m != nil {
		return *m.OrgId
	}
	return 0
}

func (m *DepartmentStaffInfo) SetOrgName(orgName string) *DepartmentStaffInfo {
	m.OrgName = &orgName
	return m
}

func (m *DepartmentStaffInfo) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

func (m *DepartmentStaffInfo) SetParentId(parentId string) *DepartmentStaffInfo {
	m.ParentId = &parentId
	return m
}

func (m *DepartmentStaffInfo) GetParentId() string {
	if m != nil {
		return *m.ParentId
	}
	return ""
}

func (m *DepartmentStaffInfo) SetStatus(status int) *DepartmentStaffInfo {
	m.Status = &status
	return m
}

func (m *DepartmentStaffInfo) GetStatus() int {
	if m != nil {
		return *m.Status
	}
	return 0
}

type Event struct {
	ChannelType *int      `json:"channelType,omitempty"` // 通道类型
	DeviceType  *int      `json:"deviceType,omitempty"`  // 推送的目标设备类型：1: android, 2: ios, 4: windows, 8: mac, 16: linux 类似于channelType字段，该字段也支持或运算符( | )。如果不填或者将该字段置0， 那么开平会将该事件推送到所有的登录设备。
	EntryId     *string   `json:"entryId,omitempty"`     // 应用的入口ID（主要用于微应用）
	EventData   *string   `json:"eventData,omitempty"`   // 事件内容，系统预定义的事件-工作台角标参数参见下面数据类型与数据格式定义。对应用自定义事件，此内容由应用服务端与应用客户端沟通协商确定，蓝信开放平台不关心具体内容。
	EventType   *string   `json:"eventType,omitempty"`   // 事件类型，目前支持的系统预定义事件有工作台红点:app_changes。 对应用自定义事件数据类型拼接的建议规则 “应用type_场景type_场景Id(openid)” 应用自定义类型，应用内区分不同的事件类型，由应用自行决定具体值（接口透传该值）。
	Expires     *int      `json:"expires,omitempty"`     // 事件的过期时间，单位是秒。默认值为0是表示永不过期。如果设置指定过期时间（非0），则应用需要实现事件拉取的回调接口。
	PushData    *PushData `json:"pushData,omitempty"`
	ReceiverIds []*string `json:"receiverIds,omitempty"` // 接收者的open staffId列表
	Version     *int      `json:"version,omitempty"`     // 可选字段，数据的版本号，要求是个时间戳，精确到微秒， 例如：1605693953610320
}

func (m *Event) SetChannelType(channelType int) *Event {
	m.ChannelType = &channelType
	return m
}

func (m *Event) GetChannelType() int {
	if m != nil {
		return *m.ChannelType
	}
	return 0
}

func (m *Event) SetDeviceType(deviceType int) *Event {
	m.DeviceType = &deviceType
	return m
}

func (m *Event) GetDeviceType() int {
	if m != nil {
		return *m.DeviceType
	}
	return 0
}

func (m *Event) SetEntryId(entryId string) *Event {
	m.EntryId = &entryId
	return m
}

func (m *Event) GetEntryId() string {
	if m != nil {
		return *m.EntryId
	}
	return ""
}

func (m *Event) SetEventData(eventData string) *Event {
	m.EventData = &eventData
	return m
}

func (m *Event) GetEventData() string {
	if m != nil {
		return *m.EventData
	}
	return ""
}

func (m *Event) SetEventType(eventType string) *Event {
	m.EventType = &eventType
	return m
}

func (m *Event) GetEventType() string {
	if m != nil {
		return *m.EventType
	}
	return ""
}

func (m *Event) SetExpires(expires int) *Event {
	m.Expires = &expires
	return m
}

func (m *Event) GetExpires() int {
	if m != nil {
		return *m.Expires
	}
	return 0
}

func (m *Event) SetPushData(pushData PushData) *Event {
	m.PushData = &pushData
	return m
}

func (m *Event) GetPushData() *PushData {
	if m != nil {
		return m.PushData
	}
	return nil
}

func (m *Event) SetReceiverIds(receiverIds []*string) *Event {
	m.ReceiverIds = receiverIds
	return m
}

func (m *Event) GetReceiverIds() []*string {
	if m != nil {
		return m.ReceiverIds
	}
	return nil
}

func (m *Event) SetVersion(version int) *Event {
	m.Version = &version
	return m
}

func (m *Event) GetVersion() int {
	if m != nil {
		return *m.Version
	}
	return 0
}

type ExtraFieldId struct {
	Category *int    `json:"category,omitempty"`
	Id       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

func (m *ExtraFieldId) SetCategory(category int) *ExtraFieldId {
	m.Category = &category
	return m
}

func (m *ExtraFieldId) GetCategory() int {
	if m != nil {
		return *m.Category
	}
	return 0
}

func (m *ExtraFieldId) SetId(id string) *ExtraFieldId {
	m.Id = &id
	return m
}

func (m *ExtraFieldId) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *ExtraFieldId) SetName(name string) *ExtraFieldId {
	m.Name = &name
	return m
}

func (m *ExtraFieldId) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

type GroupInfoTag struct {
	TagId   *string `json:"tagId,omitempty"`
	TagName *string `json:"tagName,omitempty"`
}

func (m *GroupInfoTag) SetTagId(tagId string) *GroupInfoTag {
	m.TagId = &tagId
	return m
}

func (m *GroupInfoTag) GetTagId() string {
	if m != nil {
		return *m.TagId
	}
	return ""
}

func (m *GroupInfoTag) SetTagName(tagName string) *GroupInfoTag {
	m.TagName = &tagName
	return m
}

func (m *GroupInfoTag) GetTagName() string {
	if m != nil {
		return *m.TagName
	}
	return ""
}

type Introduction struct {
	Introduction *string   `json:"introduction,omitempty"`
	MediaIds     []*string `json:"mediaIds,omitempty"`
}

func (m *Introduction) SetIntroduction(introduction string) *Introduction {
	m.Introduction = &introduction
	return m
}

func (m *Introduction) GetIntroduction() string {
	if m != nil {
		return *m.Introduction
	}
	return ""
}

func (m *Introduction) SetMediaIds(mediaIds []*string) *Introduction {
	m.MediaIds = mediaIds
	return m
}

func (m *Introduction) GetMediaIds() []*string {
	if m != nil {
		return m.MediaIds
	}
	return nil
}

type MobilePhone struct {
	CountryCode *string `json:"countryCode,omitempty"` // 手机号国家码
	Number      *string `json:"number,omitempty"`      // 手机号
}

func (m *MobilePhone) SetCountryCode(countryCode string) *MobilePhone {
	m.CountryCode = &countryCode
	return m
}

func (m *MobilePhone) GetCountryCode() string {
	if m != nil {
		return *m.CountryCode
	}
	return ""
}

func (m *MobilePhone) SetNumber(number string) *MobilePhone {
	m.Number = &number
	return m
}

func (m *MobilePhone) GetNumber() string {
	if m != nil {
		return *m.Number
	}
	return ""
}

type PushData struct {
	AndroidSoundUrl *string `json:"androidSoundUrl,omitempty"` // 安卓端个性化铃声
	AppType         *string `json:"appType,omitempty"`         // 通知栏消息拉起应用时用的应用类型：“webview”， “net_meeting”， “security_mail”， “blueprint”
	Content         *string `json:"content,omitempty"`         // 通知消息内容
	IosSoundUrl     *string `json:"iosSoundUrl,omitempty"`     // ios端个性化铃声
	Title           *string `json:"title,omitempty"`           // 通知消息标题，为空时取应用名称
	Url             *string `json:"url,omitempty"`             // 应用详情页跳转地址，用于通过推送消息拉起应用时透传给应用 。如果不需要，可以为空。
}

func (m *PushData) SetAndroidSoundUrl(androidSoundUrl string) *PushData {
	m.AndroidSoundUrl = &androidSoundUrl
	return m
}

func (m *PushData) GetAndroidSoundUrl() string {
	if m != nil {
		return *m.AndroidSoundUrl
	}
	return ""
}

func (m *PushData) SetAppType(appType string) *PushData {
	m.AppType = &appType
	return m
}

func (m *PushData) GetAppType() string {
	if m != nil {
		return *m.AppType
	}
	return ""
}

func (m *PushData) SetContent(content string) *PushData {
	m.Content = &content
	return m
}

func (m *PushData) GetContent() string {
	if m != nil {
		return *m.Content
	}
	return ""
}

func (m *PushData) SetIosSoundUrl(iosSoundUrl string) *PushData {
	m.IosSoundUrl = &iosSoundUrl
	return m
}

func (m *PushData) GetIosSoundUrl() string {
	if m != nil {
		return *m.IosSoundUrl
	}
	return ""
}

func (m *PushData) SetTitle(title string) *PushData {
	m.Title = &title
	return m
}

func (m *PushData) GetTitle() string {
	if m != nil {
		return *m.Title
	}
	return ""
}

func (m *PushData) SetUrl(url string) *PushData {
	m.Url = &url
	return m
}

func (m *PushData) GetUrl() string {
	if m != nil {
		return *m.Url
	}
	return ""
}

type Resume struct {
	Description *string `json:"description,omitempty"`
	EndDate     *string `json:"endDate,omitempty"`
	StartDate   *string `json:"startDate,omitempty"`
	Unit        *string `json:"unit,omitempty"`
}

func (m *Resume) SetDescription(description string) *Resume {
	m.Description = &description
	return m
}

func (m *Resume) GetDescription() string {
	if m != nil {
		return *m.Description
	}
	return ""
}

func (m *Resume) SetEndDate(endDate string) *Resume {
	m.EndDate = &endDate
	return m
}

func (m *Resume) GetEndDate() string {
	if m != nil {
		return *m.EndDate
	}
	return ""
}

func (m *Resume) SetStartDate(startDate string) *Resume {
	m.StartDate = &startDate
	return m
}

func (m *Resume) GetStartDate() string {
	if m != nil {
		return *m.StartDate
	}
	return ""
}

func (m *Resume) SetUnit(unit string) *Resume {
	m.Unit = &unit
	return m
}

func (m *Resume) GetUnit() string {
	if m != nil {
		return *m.Unit
	}
	return ""
}

type RoleStaff struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (m *RoleStaff) SetId(id string) *RoleStaff {
	m.Id = &id
	return m
}

func (m *RoleStaff) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *RoleStaff) SetName(name string) *RoleStaff {
	m.Name = &name
	return m
}

func (m *RoleStaff) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

type SearchScope struct {
	OrgIds    []*string `json:"orgIds,omitempty"`    // 组织ID
	SectorIds []*string `json:"sectorIds,omitempty"` // 分支/部门ID
}

func (m *SearchScope) SetOrgIds(orgIds []*string) *SearchScope {
	m.OrgIds = orgIds
	return m
}

func (m *SearchScope) GetOrgIds() []*string {
	if m != nil {
		return m.OrgIds
	}
	return nil
}

func (m *SearchScope) SetSectorIds(sectorIds []*string) *SearchScope {
	m.SectorIds = sectorIds
	return m
}

func (m *SearchScope) GetSectorIds() []*string {
	if m != nil {
		return m.SectorIds
	}
	return nil
}

type SearchStaffInfo struct {
	AvatarId    *string      `json:"avatarId,omitempty"` // 头像ID
	Email       *string      `json:"email,omitempty"`    // 员工的email
	Icon        *string      `json:"icon,omitempty"`
	Mobile      *string      `json:"mobile,omitempty"` // 员工的手机号
	MobilePhone *MobilePhone `json:"mobilePhone,omitempty"`
	Name        *string      `json:"name,omitempty"`    // 员工的姓名
	StaffId     *string      `json:"staffId,omitempty"` // 员工的open id
}

func (m *SearchStaffInfo) SetAvatarId(avatarId string) *SearchStaffInfo {
	m.AvatarId = &avatarId
	return m
}

func (m *SearchStaffInfo) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *SearchStaffInfo) SetEmail(email string) *SearchStaffInfo {
	m.Email = &email
	return m
}

func (m *SearchStaffInfo) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *SearchStaffInfo) SetIcon(icon string) *SearchStaffInfo {
	m.Icon = &icon
	return m
}

func (m *SearchStaffInfo) GetIcon() string {
	if m != nil {
		return *m.Icon
	}
	return ""
}

func (m *SearchStaffInfo) SetMobile(mobile string) *SearchStaffInfo {
	m.Mobile = &mobile
	return m
}

func (m *SearchStaffInfo) GetMobile() string {
	if m != nil {
		return *m.Mobile
	}
	return ""
}

func (m *SearchStaffInfo) SetMobilePhone(mobilePhone MobilePhone) *SearchStaffInfo {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *SearchStaffInfo) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *SearchStaffInfo) SetName(name string) *SearchStaffInfo {
	m.Name = &name
	return m
}

func (m *SearchStaffInfo) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *SearchStaffInfo) SetStaffId(staffId string) *SearchStaffInfo {
	m.StaffId = &staffId
	return m
}

func (m *SearchStaffInfo) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

type StatusInfo struct {
	BaseVersion *string `json:"baseVersion,omitempty"` // 保证每次请求时，该值递增。该值由应用服务端维护
	NoDisturb   *string `json:"noDisturb,omitempty"`   // 免打扰标识。"true":开启免打扰功能，"false":关闭免打扰功能。注：noDistrub和unreadCount为互斥项，只能同时出现一个
	UnreadCount *int    `json:"unreadCount,omitempty"` // 会话未读数。注：noDistrub和unreadCount为互斥项，只能同时出现一个
	UserId      *string `json:"userId,omitempty"`      // 通知会话所有者，通知的peerId就是通知应用的appId
	Uuid        *string `json:"uuid,omitempty"`        // 一个随机字符串(uuid)
}

func (m *StatusInfo) SetBaseVersion(baseVersion string) *StatusInfo {
	m.BaseVersion = &baseVersion
	return m
}

func (m *StatusInfo) GetBaseVersion() string {
	if m != nil {
		return *m.BaseVersion
	}
	return ""
}

func (m *StatusInfo) SetNoDisturb(noDisturb string) *StatusInfo {
	m.NoDisturb = &noDisturb
	return m
}

func (m *StatusInfo) GetNoDisturb() string {
	if m != nil {
		return *m.NoDisturb
	}
	return ""
}

func (m *StatusInfo) SetUnreadCount(unreadCount int) *StatusInfo {
	m.UnreadCount = &unreadCount
	return m
}

func (m *StatusInfo) GetUnreadCount() int {
	if m != nil {
		return *m.UnreadCount
	}
	return 0
}

func (m *StatusInfo) SetUserId(userId string) *StatusInfo {
	m.UserId = &userId
	return m
}

func (m *StatusInfo) GetUserId() string {
	if m != nil {
		return *m.UserId
	}
	return ""
}

func (m *StatusInfo) SetUuid(uuid string) *StatusInfo {
	m.Uuid = &uuid
	return m
}

func (m *StatusInfo) GetUuid() string {
	if m != nil {
		return *m.Uuid
	}
	return ""
}

type SystemMsg struct {
	Content *string `json:"content,omitempty"` // 撤回消息时展示的系统消息内容（可选）
	MediaId *string `json:"mediaId,omitempty"` // 撤回消息时展示的撤回图标的id
}

func (m *SystemMsg) SetContent(content string) *SystemMsg {
	m.Content = &content
	return m
}

func (m *SystemMsg) GetContent() string {
	if m != nil {
		return *m.Content
	}
	return ""
}

func (m *SystemMsg) SetMediaId(mediaId string) *SystemMsg {
	m.MediaId = &mediaId
	return m
}

func (m *SystemMsg) GetMediaId() string {
	if m != nil {
		return *m.MediaId
	}
	return ""
}

type TagFilter struct {
	Tags []*string `json:"tags,omitempty"`
}

func (m *TagFilter) SetTags(tags []*string) *TagFilter {
	m.Tags = tags
	return m
}

func (m *TagFilter) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TagMeta struct {
	TagId   *string `json:"tagId,omitempty"`   // 标签ID
	TagName *string `json:"tagName,omitempty"` // 标签名称
}

func (m *TagMeta) SetTagId(tagId string) *TagMeta {
	m.TagId = &tagId
	return m
}

func (m *TagMeta) GetTagId() string {
	if m != nil {
		return *m.TagId
	}
	return ""
}

func (m *TagMeta) SetTagName(tagName string) *TagMeta {
	m.TagName = &tagName
	return m
}

func (m *TagMeta) GetTagName() string {
	if m != nil {
		return *m.TagName
	}
	return ""
}

type TagStaffInfo struct {
	Id          *string      `json:"id,omitempty"`
	MobilePhone *MobilePhone `json:"mobilePhone,omitempty"`
	Name        *string      `json:"name,omitempty"`
}

func (m *TagStaffInfo) SetId(id string) *TagStaffInfo {
	m.Id = &id
	return m
}

func (m *TagStaffInfo) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *TagStaffInfo) SetMobilePhone(mobilePhone MobilePhone) *TagStaffInfo {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *TagStaffInfo) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *TagStaffInfo) SetName(name string) *TagStaffInfo {
	m.Name = &name
	return m
}

func (m *TagStaffInfo) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

type UDepartment struct {
	Id          *string `json:"id,omitempty"`          // 所在分支ID
	OrderNumber *int    `json:"orderNumber,omitempty"` // 人员排序值
}

func (m *UDepartment) SetId(id string) *UDepartment {
	m.Id = &id
	return m
}

func (m *UDepartment) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *UDepartment) SetOrderNumber(orderNumber int) *UDepartment {
	m.OrderNumber = &orderNumber
	return m
}

func (m *UDepartment) GetOrderNumber() int {
	if m != nil {
		return *m.OrderNumber
	}
	return 0
}

type V1AppRoamingOrgsFetchData struct {
	HasMore       *bool                `json:"hasMore,omitempty"`       // 是否还有下一页数据，有-true，没有-false。
	LatestVersion *string              `json:"latestVersion,omitempty"` // 返回数据集所包含的数据中最新的版本号，是一个时间戳字符串，单位：微秒，请求下一页时可以此值作为base_version传入。调用者需要保存该版本号到本地，下次增量请求数据时作为base_version传递到服务端。
	OrgList       []*AppRoamingOrgList `json:"orgList,omitempty"`
}

func (m *V1AppRoamingOrgsFetchData) SetHasMore(hasMore bool) *V1AppRoamingOrgsFetchData {
	m.HasMore = &hasMore
	return m
}

func (m *V1AppRoamingOrgsFetchData) GetHasMore() bool {
	if m != nil {
		return *m.HasMore
	}
	return false
}

func (m *V1AppRoamingOrgsFetchData) SetLatestVersion(latestVersion string) *V1AppRoamingOrgsFetchData {
	m.LatestVersion = &latestVersion
	return m
}

func (m *V1AppRoamingOrgsFetchData) GetLatestVersion() string {
	if m != nil {
		return *m.LatestVersion
	}
	return ""
}

func (m *V1AppRoamingOrgsFetchData) SetOrgList(orgList []*AppRoamingOrgList) *V1AppRoamingOrgsFetchData {
	m.OrgList = orgList
	return m
}

func (m *V1AppRoamingOrgsFetchData) GetOrgList() []*AppRoamingOrgList {
	if m != nil {
		return m.OrgList
	}
	return nil
}

type V1AppRoamingOrgsFetchResponse struct {
	Data    *V1AppRoamingOrgsFetchData `json:"data,omitempty"`
	ErrCode *int                       `json:"errCode,omitempty"`
	ErrMsg  *string                    `json:"errMsg,omitempty"`
}

func (m *V1AppRoamingOrgsFetchResponse) GetData() *V1AppRoamingOrgsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1AppRoamingOrgsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1AppRoamingOrgsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1AppRoamingOrgsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1AppTokenCreateData struct {
	AppToken  *string `json:"appToken,omitempty"`  // 应用访问APP_TOKEN
	ExpiresIn *int    `json:"expiresIn,omitempty"` // TOKEN 有效期（7200秒），建议应用根据过期时间缓存appToken, 单次获取，多次使用
}

func (m *V1AppTokenCreateData) SetAppToken(appToken string) *V1AppTokenCreateData {
	m.AppToken = &appToken
	return m
}

func (m *V1AppTokenCreateData) GetAppToken() string {
	if m != nil {
		return *m.AppToken
	}
	return ""
}

func (m *V1AppTokenCreateData) SetExpiresIn(expiresIn int) *V1AppTokenCreateData {
	m.ExpiresIn = &expiresIn
	return m
}

func (m *V1AppTokenCreateData) GetExpiresIn() int {
	if m != nil {
		return *m.ExpiresIn
	}
	return 0
}

type V1AppTokenCreateResponse struct {
	Data    *V1AppTokenCreateData `json:"data,omitempty"`
	ErrCode *int                  `json:"errCode,omitempty"`
	ErrMsg  *string               `json:"errMsg,omitempty"`
}

func (m *V1AppTokenCreateResponse) GetData() *V1AppTokenCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1AppTokenCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1AppTokenCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1AppTokenCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1ChatNotificationFetchData struct {
	NoDisturb *string `json:"noDisturb,omitempty"` // 免打扰标识
}

func (m *V1ChatNotificationFetchData) SetNoDisturb(noDisturb string) *V1ChatNotificationFetchData {
	m.NoDisturb = &noDisturb
	return m
}

func (m *V1ChatNotificationFetchData) GetNoDisturb() string {
	if m != nil {
		return *m.NoDisturb
	}
	return ""
}

type V1ChatNotificationFetchResponse struct {
	Data    *V1ChatNotificationFetchData `json:"data,omitempty"`
	ErrCode *int                         `json:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty"`
}

func (m *V1ChatNotificationFetchResponse) GetData() *V1ChatNotificationFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1ChatNotificationFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1ChatNotificationFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1ChatNotificationFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsChildrenFetchData struct {
	Departments []*DepartmentData `json:"departments,omitempty"`
}

func (m *V1DeptsChildrenFetchData) SetDepartments(departments []*DepartmentData) *V1DeptsChildrenFetchData {
	m.Departments = departments
	return m
}

func (m *V1DeptsChildrenFetchData) GetDepartments() []*DepartmentData {
	if m != nil {
		return m.Departments
	}
	return nil
}

type V1DeptsChildrenFetchResponse struct {
	Data    *V1DeptsChildrenFetchData `json:"data,omitempty"`
	ErrCode *int                      `json:"errCode,omitempty"`
	ErrMsg  *string                   `json:"errMsg,omitempty"`
}

func (m *V1DeptsChildrenFetchResponse) GetData() *V1DeptsChildrenFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1DeptsChildrenFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsChildrenFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsChildrenFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsCreateData struct {
	Department *Department `json:"department,omitempty"`
}

func (m *V1DeptsCreateData) SetDepartment(department Department) *V1DeptsCreateData {
	m.Department = &department
	return m
}

func (m *V1DeptsCreateData) GetDepartment() *Department {
	if m != nil {
		return m.Department
	}
	return nil
}

type V1DeptsCreateRequestBody struct {
	ExternalId  *string   `json:"externalId,omitempty"`  // 分支外部ID，组织通讯录数据源唯一标识分支的ID。创建后不可修改，组织内必须唯一
	Name        *string   `json:"name,omitempty"`        // 分支名称
	OrderNumber *int      `json:"orderNumber,omitempty"` // 在父分支中的次序值，值越小排序越靠前
	ParentId    *string   `json:"parentId,omitempty"`    // 父节点分支Id，根分支可以用”组织ID-0“代替，例如：524288-0， 创建根分支的时候需要指定组织，创建其他子分支的时候根据父分支决定组织。
	Tags        []*string `json:"tags,omitempty"`        // 分支标签
}

func (m *V1DeptsCreateRequestBody) SetExternalId(externalId string) *V1DeptsCreateRequestBody {
	m.ExternalId = &externalId
	return m
}

func (m *V1DeptsCreateRequestBody) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1DeptsCreateRequestBody) SetName(name string) *V1DeptsCreateRequestBody {
	m.Name = &name
	return m
}

func (m *V1DeptsCreateRequestBody) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1DeptsCreateRequestBody) SetOrderNumber(orderNumber int) *V1DeptsCreateRequestBody {
	m.OrderNumber = &orderNumber
	return m
}

func (m *V1DeptsCreateRequestBody) GetOrderNumber() int {
	if m != nil {
		return *m.OrderNumber
	}
	return 0
}

func (m *V1DeptsCreateRequestBody) SetParentId(parentId string) *V1DeptsCreateRequestBody {
	m.ParentId = &parentId
	return m
}

func (m *V1DeptsCreateRequestBody) GetParentId() string {
	if m != nil {
		return *m.ParentId
	}
	return ""
}

func (m *V1DeptsCreateRequestBody) SetTags(tags []*string) *V1DeptsCreateRequestBody {
	m.Tags = tags
	return m
}

func (m *V1DeptsCreateRequestBody) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1DeptsCreateResponse struct {
	Data    *V1DeptsCreateData `json:"data,omitempty"`
	ErrCode *int               `json:"errCode,omitempty"`
	ErrMsg  *string            `json:"errMsg,omitempty"`
}

func (m *V1DeptsCreateResponse) GetData() *V1DeptsCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1DeptsCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsDeleteResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1DeptsDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsFetchData struct {
	AncestorDepartments []*Department `json:"ancestorDepartments,omitempty"` // 部门祖先列表
	DeletedMembers      *int          `json:"deletedMembers,omitempty"`      // 已删除
	ExternalId          *string       `json:"externalId,omitempty"`          // 分支外部ID
	FrozenMembers       *int          `json:"frozenMembers,omitempty"`       // 已冻结
	HasChildren         *bool         `json:"hasChildren,omitempty"`         // 是否有子分支
	Id                  *string       `json:"id,omitempty"`                  // 当前Id
	InactiveMembers     *int          `json:"inactiveMembers,omitempty"`     // 未注册
	Name                *string       `json:"name,omitempty"`                // 分支名称
	NormalMembers       *int          `json:"normalMembers,omitempty"`       // 已注册
	Order               *float32      `json:"order,omitempty"`               // 分支顺序，越小排在越前面
	ParentId            *string       `json:"parentId,omitempty"`            // 父Id
	Tags                []*string     `json:"tags,omitempty"`                // 该分支所持有的全部标签ID列表
}

func (m *V1DeptsFetchData) SetAncestorDepartments(ancestorDepartments []*Department) *V1DeptsFetchData {
	m.AncestorDepartments = ancestorDepartments
	return m
}

func (m *V1DeptsFetchData) GetAncestorDepartments() []*Department {
	if m != nil {
		return m.AncestorDepartments
	}
	return nil
}

func (m *V1DeptsFetchData) SetDeletedMembers(deletedMembers int) *V1DeptsFetchData {
	m.DeletedMembers = &deletedMembers
	return m
}

func (m *V1DeptsFetchData) GetDeletedMembers() int {
	if m != nil {
		return *m.DeletedMembers
	}
	return 0
}

func (m *V1DeptsFetchData) SetExternalId(externalId string) *V1DeptsFetchData {
	m.ExternalId = &externalId
	return m
}

func (m *V1DeptsFetchData) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1DeptsFetchData) SetFrozenMembers(frozenMembers int) *V1DeptsFetchData {
	m.FrozenMembers = &frozenMembers
	return m
}

func (m *V1DeptsFetchData) GetFrozenMembers() int {
	if m != nil {
		return *m.FrozenMembers
	}
	return 0
}

func (m *V1DeptsFetchData) SetHasChildren(hasChildren bool) *V1DeptsFetchData {
	m.HasChildren = &hasChildren
	return m
}

func (m *V1DeptsFetchData) GetHasChildren() bool {
	if m != nil {
		return *m.HasChildren
	}
	return false
}

func (m *V1DeptsFetchData) SetId(id string) *V1DeptsFetchData {
	m.Id = &id
	return m
}

func (m *V1DeptsFetchData) GetId() string {
	if m != nil {
		return *m.Id
	}
	return ""
}

func (m *V1DeptsFetchData) SetInactiveMembers(inactiveMembers int) *V1DeptsFetchData {
	m.InactiveMembers = &inactiveMembers
	return m
}

func (m *V1DeptsFetchData) GetInactiveMembers() int {
	if m != nil {
		return *m.InactiveMembers
	}
	return 0
}

func (m *V1DeptsFetchData) SetName(name string) *V1DeptsFetchData {
	m.Name = &name
	return m
}

func (m *V1DeptsFetchData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1DeptsFetchData) SetNormalMembers(normalMembers int) *V1DeptsFetchData {
	m.NormalMembers = &normalMembers
	return m
}

func (m *V1DeptsFetchData) GetNormalMembers() int {
	if m != nil {
		return *m.NormalMembers
	}
	return 0
}

func (m *V1DeptsFetchData) SetOrder(order float32) *V1DeptsFetchData {
	m.Order = &order
	return m
}

func (m *V1DeptsFetchData) GetOrder() float32 {
	if m != nil {
		return *m.Order
	}
	return 0
}

func (m *V1DeptsFetchData) SetParentId(parentId string) *V1DeptsFetchData {
	m.ParentId = &parentId
	return m
}

func (m *V1DeptsFetchData) GetParentId() string {
	if m != nil {
		return *m.ParentId
	}
	return ""
}

func (m *V1DeptsFetchData) SetTags(tags []*string) *V1DeptsFetchData {
	m.Tags = tags
	return m
}

func (m *V1DeptsFetchData) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1DeptsFetchResponse struct {
	Data    *V1DeptsFetchData `json:"data,omitempty"`
	ErrCode *int              `json:"errCode,omitempty"`
	ErrMsg  *string           `json:"errMsg,omitempty"`
}

func (m *V1DeptsFetchResponse) GetData() *V1DeptsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1DeptsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsStaffsCreateResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1DeptsStaffsCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsStaffsCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsStaffsCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsStaffsDeleteResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1DeptsStaffsDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsStaffsDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsStaffsDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsStaffsFetchData struct {
	HasMore *bool                  `json:"hasMore,omitempty"`
	Staffs  []*DepartmentStaffInfo `json:"staffs,omitempty"`
	Total   *int                   `json:"total,omitempty"`
}

func (m *V1DeptsStaffsFetchData) SetHasMore(hasMore bool) *V1DeptsStaffsFetchData {
	m.HasMore = &hasMore
	return m
}

func (m *V1DeptsStaffsFetchData) GetHasMore() bool {
	if m != nil {
		return *m.HasMore
	}
	return false
}

func (m *V1DeptsStaffsFetchData) SetStaffs(staffs []*DepartmentStaffInfo) *V1DeptsStaffsFetchData {
	m.Staffs = staffs
	return m
}

func (m *V1DeptsStaffsFetchData) GetStaffs() []*DepartmentStaffInfo {
	if m != nil {
		return m.Staffs
	}
	return nil
}

func (m *V1DeptsStaffsFetchData) SetTotal(total int) *V1DeptsStaffsFetchData {
	m.Total = &total
	return m
}

func (m *V1DeptsStaffsFetchData) GetTotal() int {
	if m != nil {
		return *m.Total
	}
	return 0
}

type V1DeptsStaffsFetchResponse struct {
	Data    *V1DeptsStaffsFetchData `json:"data,omitempty"`
	ErrCode *int                    `json:"errCode,omitempty"`
	ErrMsg  *string                 `json:"errMsg,omitempty"`
}

func (m *V1DeptsStaffsFetchResponse) GetData() *V1DeptsStaffsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1DeptsStaffsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsStaffsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsStaffsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1DeptsUpdateRequestBody struct {
	Name        *string   `json:"name,omitempty"`        // 分支名称
	OrderNumber *int      `json:"orderNumber,omitempty"` // 在父分支中的次序值，值越小排序越靠前
	ParentId    *string   `json:"parentId,omitempty"`
	Tags        []*string `json:"tags,omitempty"` // 分支标签，必须获取后整体更新，不支持增量更新
}

func (m *V1DeptsUpdateRequestBody) SetName(name string) *V1DeptsUpdateRequestBody {
	m.Name = &name
	return m
}

func (m *V1DeptsUpdateRequestBody) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1DeptsUpdateRequestBody) SetOrderNumber(orderNumber int) *V1DeptsUpdateRequestBody {
	m.OrderNumber = &orderNumber
	return m
}

func (m *V1DeptsUpdateRequestBody) GetOrderNumber() int {
	if m != nil {
		return *m.OrderNumber
	}
	return 0
}

func (m *V1DeptsUpdateRequestBody) SetParentId(parentId string) *V1DeptsUpdateRequestBody {
	m.ParentId = &parentId
	return m
}

func (m *V1DeptsUpdateRequestBody) GetParentId() string {
	if m != nil {
		return *m.ParentId
	}
	return ""
}

func (m *V1DeptsUpdateRequestBody) SetTags(tags []*string) *V1DeptsUpdateRequestBody {
	m.Tags = tags
	return m
}

func (m *V1DeptsUpdateRequestBody) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1DeptsUpdateResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1DeptsUpdateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1DeptsUpdateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1DeptsUpdateResponse) ToString() string {
	return ToJsonString(m)
}

type V1JsApiTokenCreateData struct {
	ExpiresIn  *int    `json:"expiresIn,omitempty"`  // TOKEN 有效期（7200秒），建议应用根据过期时间缓存jsApiToken
	JsApiToken *string `json:"jsApiToken,omitempty"` // JSAPI 访问TOKEN
}

func (m *V1JsApiTokenCreateData) SetExpiresIn(expiresIn int) *V1JsApiTokenCreateData {
	m.ExpiresIn = &expiresIn
	return m
}

func (m *V1JsApiTokenCreateData) GetExpiresIn() int {
	if m != nil {
		return *m.ExpiresIn
	}
	return 0
}

func (m *V1JsApiTokenCreateData) SetJsApiToken(jsApiToken string) *V1JsApiTokenCreateData {
	m.JsApiToken = &jsApiToken
	return m
}

func (m *V1JsApiTokenCreateData) GetJsApiToken() string {
	if m != nil {
		return *m.JsApiToken
	}
	return ""
}

type V1JsApiTokenCreateResponse struct {
	Data    *V1JsApiTokenCreateData `json:"data,omitempty"`
	ErrCode *int                    `json:"errCode,omitempty"`
	ErrMsg  *string                 `json:"errMsg,omitempty"`
}

func (m *V1JsApiTokenCreateResponse) GetData() *V1JsApiTokenCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1JsApiTokenCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1JsApiTokenCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1JsApiTokenCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1MediasCreateData struct {
	CreatedTime *string `json:"createdTime,omitempty"` // 文件创建时间戳，单位微秒
	MediaId     *string `json:"mediaId,omitempty"`     // 文件ID
}

func (m *V1MediasCreateData) SetCreatedTime(createdTime string) *V1MediasCreateData {
	m.CreatedTime = &createdTime
	return m
}

func (m *V1MediasCreateData) GetCreatedTime() string {
	if m != nil {
		return *m.CreatedTime
	}
	return ""
}

func (m *V1MediasCreateData) SetMediaId(mediaId string) *V1MediasCreateData {
	m.MediaId = &mediaId
	return m
}

func (m *V1MediasCreateData) GetMediaId() string {
	if m != nil {
		return *m.MediaId
	}
	return ""
}

type V1MediasCreateResponse struct {
	Data    *V1MediasCreateData `json:"data,omitempty"`
	ErrCode *int                `json:"errCode,omitempty"`
	ErrMsg  *string             `json:"errMsg,omitempty"`
}

func (m *V1MediasCreateResponse) GetData() *V1MediasCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1MediasCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1MediasCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1MediasCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1MediasPathFetchData struct {
	MediaPath *string `json:"mediaPath,omitempty"` // 文件下载路径，非永久有效，有效期1小时
	Name      *string `json:"name,omitempty"`
	Size      *int    `json:"size,omitempty"`
	Type      *string `json:"type,omitempty"`
}

func (m *V1MediasPathFetchData) SetMediaPath(mediaPath string) *V1MediasPathFetchData {
	m.MediaPath = &mediaPath
	return m
}

func (m *V1MediasPathFetchData) GetMediaPath() string {
	if m != nil {
		return *m.MediaPath
	}
	return ""
}

func (m *V1MediasPathFetchData) SetName(name string) *V1MediasPathFetchData {
	m.Name = &name
	return m
}

func (m *V1MediasPathFetchData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1MediasPathFetchData) SetSize(size int) *V1MediasPathFetchData {
	m.Size = &size
	return m
}

func (m *V1MediasPathFetchData) GetSize() int {
	if m != nil {
		return *m.Size
	}
	return 0
}

func (m *V1MediasPathFetchData) SetType(v string) *V1MediasPathFetchData {
	m.Type = &v
	return m
}

func (m *V1MediasPathFetchData) GetType() string {
	if m != nil {
		return *m.Type
	}
	return ""
}

type V1MediasPathFetchResponse struct {
	Data    *V1MediasPathFetchData `json:"data,omitempty"`
	ErrCode *int                   `json:"errCode,omitempty"`
	ErrMsg  *string                `json:"errMsg,omitempty"`
}

func (m *V1MediasPathFetchResponse) GetData() *V1MediasPathFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1MediasPathFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1MediasPathFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1MediasPathFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1MessagesCreateData struct {
	InvalidDepartment []*string `json:"invalidDepartment,omitempty"` // 请求departmentIdList列表中的分支ID 无效，无法发送
	InvalidStaff      []*string `json:"invalidStaff,omitempty"`      // 请求staffIdList 列表中的人员ID 无效，无法发送
	MsgId             *string   `json:"msgId,omitempty"`             // 消息标识，供其他接口查询进度使用。目前只有组织内应用支持返回消息ID，ISV应用不返回ID
}

func (m *V1MessagesCreateData) SetInvalidDepartment(invalidDepartment []*string) *V1MessagesCreateData {
	m.InvalidDepartment = invalidDepartment
	return m
}

func (m *V1MessagesCreateData) GetInvalidDepartment() []*string {
	if m != nil {
		return m.InvalidDepartment
	}
	return nil
}

func (m *V1MessagesCreateData) SetInvalidStaff(invalidStaff []*string) *V1MessagesCreateData {
	m.InvalidStaff = invalidStaff
	return m
}

func (m *V1MessagesCreateData) GetInvalidStaff() []*string {
	if m != nil {
		return m.InvalidStaff
	}
	return nil
}

func (m *V1MessagesCreateData) SetMsgId(msgId string) *V1MessagesCreateData {
	m.MsgId = &msgId
	return m
}

func (m *V1MessagesCreateData) GetMsgId() string {
	if m != nil {
		return *m.MsgId
	}
	return ""
}

type V1MessagesCreateRequestBody struct {
	AccountId        *string     `json:"accountId,omitempty"`        // 普通应用不需要填，仅适用于应用使用多公号消息通道的情况，例如移动会务。accountId为公号ID/entryId为应用入口ID。优先使用accountId做为目标公号。如果accountId为空，则使用entryId指定的的应用入口所关联的公号。如果应用只有一个入口可不填
	Attach           *string     `json:"attach,omitempty"`           // 普通应用不需要填，仅适用于微应用，公号消息附加数据，目前用于传递微应用链接上下文数据，内容需要做UrlEncode。
	DepartmentIdList []*string   `json:"departmentIdList,omitempty"` // 接收者分支列表（分支下的所有人），可选，与userIdList二者间必选一个，如果需要全组织广播，则填根分支Id：orgId-0，例如：524288-0, 最多支持100个, 全组织广播时，只支持1个组织
	EntryId          *string     `json:"entryId,omitempty"`          // 普通应用不需要填，仅适用于微应用
	MsgData          interface{} `json:"msgData,omitempty"`          // type 类型名对应的同名的格式化数据。每种格式都有对应的数据类型
	MsgType          *string     `json:"msgType,omitempty"`          // 发送的消息格式，支持以下几种："text"，"oacard"，"linkCard"，"appCard"
	UserIdList       []*string   `json:"userIdList,omitempty"`       // 接收者人员列表，指定消息接收者时使用，可选，与departmentIdList二者间必选一个, 最多支持1000个
}

func (m *V1MessagesCreateRequestBody) SetAccountId(accountId string) *V1MessagesCreateRequestBody {
	m.AccountId = &accountId
	return m
}

func (m *V1MessagesCreateRequestBody) GetAccountId() string {
	if m != nil {
		return *m.AccountId
	}
	return ""
}

func (m *V1MessagesCreateRequestBody) SetAttach(attach string) *V1MessagesCreateRequestBody {
	m.Attach = &attach
	return m
}

func (m *V1MessagesCreateRequestBody) GetAttach() string {
	if m != nil {
		return *m.Attach
	}
	return ""
}

func (m *V1MessagesCreateRequestBody) SetDepartmentIdList(departmentIdList []*string) *V1MessagesCreateRequestBody {
	m.DepartmentIdList = departmentIdList
	return m
}

func (m *V1MessagesCreateRequestBody) GetDepartmentIdList() []*string {
	if m != nil {
		return m.DepartmentIdList
	}
	return nil
}

func (m *V1MessagesCreateRequestBody) SetEntryId(entryId string) *V1MessagesCreateRequestBody {
	m.EntryId = &entryId
	return m
}

func (m *V1MessagesCreateRequestBody) GetEntryId() string {
	if m != nil {
		return *m.EntryId
	}
	return ""
}

func (m *V1MessagesCreateRequestBody) SetMsgData(msgData interface{}) *V1MessagesCreateRequestBody {
	m.MsgData = msgData
	return m
}

func (m *V1MessagesCreateRequestBody) GetMsgData() interface{} {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *V1MessagesCreateRequestBody) SetMsgType(msgType string) *V1MessagesCreateRequestBody {
	m.MsgType = &msgType
	return m
}

func (m *V1MessagesCreateRequestBody) GetMsgType() string {
	if m != nil {
		return *m.MsgType
	}
	return ""
}

func (m *V1MessagesCreateRequestBody) SetUserIdList(userIdList []*string) *V1MessagesCreateRequestBody {
	m.UserIdList = userIdList
	return m
}

func (m *V1MessagesCreateRequestBody) GetUserIdList() []*string {
	if m != nil {
		return m.UserIdList
	}
	return nil
}

type V1MessagesCreateResponse struct {
	Data    *V1MessagesCreateData `json:"data,omitempty"`
	ErrCode *int                  `json:"errCode,omitempty"`
	ErrMsg  *string               `json:"errMsg,omitempty"`
}

func (m *V1MessagesCreateResponse) GetData() *V1MessagesCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1MessagesCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1MessagesCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1MessagesCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1MessagesNotificationCreateData struct {
	MsgIds []*string       `json:"msgIds,omitempty"`
	Status []*BaseResponse `json:"status,omitempty"`
}

func (m *V1MessagesNotificationCreateData) SetMsgIds(msgIds []*string) *V1MessagesNotificationCreateData {
	m.MsgIds = msgIds
	return m
}

func (m *V1MessagesNotificationCreateData) GetMsgIds() []*string {
	if m != nil {
		return m.MsgIds
	}
	return nil
}

func (m *V1MessagesNotificationCreateData) SetStatus(status []*BaseResponse) *V1MessagesNotificationCreateData {
	m.Status = status
	return m
}

func (m *V1MessagesNotificationCreateData) GetStatus() []*BaseResponse {
	if m != nil {
		return m.Status
	}
	return nil
}

type V1MessagesNotificationCreateRequestBody struct {
	MsgData     interface{} `json:"msgData,omitempty"`     // 通过openApi发送私聊消息，内容根据type具体定义，消息体类型
	MsgType     *string     `json:"msgType,omitempty"`     // 消息的类型，文本等
	ReceiverIds []*string   `json:"receiverIds,omitempty"` // 消息接收者的openId列表, 最多1000个
	SenderId    *string     `json:"senderId,omitempty"`    // 如果不提供该字段，则必须要有userToken，userToken与该字段至少有一个
	Uuid        *string     `json:"uuid,omitempty"`        // 一个随机字符串(uuid)
}

func (m *V1MessagesNotificationCreateRequestBody) SetMsgData(msgData interface{}) *V1MessagesNotificationCreateRequestBody {
	m.MsgData = msgData
	return m
}

func (m *V1MessagesNotificationCreateRequestBody) GetMsgData() interface{} {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *V1MessagesNotificationCreateRequestBody) SetMsgType(msgType string) *V1MessagesNotificationCreateRequestBody {
	m.MsgType = &msgType
	return m
}

func (m *V1MessagesNotificationCreateRequestBody) GetMsgType() string {
	if m != nil {
		return *m.MsgType
	}
	return ""
}

func (m *V1MessagesNotificationCreateRequestBody) SetReceiverIds(receiverIds []*string) *V1MessagesNotificationCreateRequestBody {
	m.ReceiverIds = receiverIds
	return m
}

func (m *V1MessagesNotificationCreateRequestBody) GetReceiverIds() []*string {
	if m != nil {
		return m.ReceiverIds
	}
	return nil
}

func (m *V1MessagesNotificationCreateRequestBody) SetSenderId(senderId string) *V1MessagesNotificationCreateRequestBody {
	m.SenderId = &senderId
	return m
}

func (m *V1MessagesNotificationCreateRequestBody) GetSenderId() string {
	if m != nil {
		return *m.SenderId
	}
	return ""
}

func (m *V1MessagesNotificationCreateRequestBody) SetUuid(uuid string) *V1MessagesNotificationCreateRequestBody {
	m.Uuid = &uuid
	return m
}

func (m *V1MessagesNotificationCreateRequestBody) GetUuid() string {
	if m != nil {
		return *m.Uuid
	}
	return ""
}

type V1MessagesNotificationCreateResponse struct {
	Data    *V1MessagesNotificationCreateData `json:"data,omitempty"`
	ErrCode *int                              `json:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty"`
}

func (m *V1MessagesNotificationCreateResponse) GetData() *V1MessagesNotificationCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1MessagesNotificationCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1MessagesNotificationCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1MessagesNotificationCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1MessagesRevokeData struct {
	SubStatus []*BaseResponse `json:"subStatus,omitempty"`
}

func (m *V1MessagesRevokeData) SetSubStatus(subStatus []*BaseResponse) *V1MessagesRevokeData {
	m.SubStatus = subStatus
	return m
}

func (m *V1MessagesRevokeData) GetSubStatus() []*BaseResponse {
	if m != nil {
		return m.SubStatus
	}
	return nil
}

type V1MessagesRevokeRequestBody struct {
	ChatType   *string    `json:"chatType,omitempty"`   // 消息类型： staff, group, notification, account
	MessageIds []*string  `json:"messageIds,omitempty"` // 消息ID列表
	SenderId   *string    `json:"senderId,omitempty"`   // 私聊（staff），群聊（group）时必须要填 senderId (staffId)
	SysMsg     *SystemMsg `json:"sysMsg,omitempty"`
}

func (m *V1MessagesRevokeRequestBody) SetChatType(chatType string) *V1MessagesRevokeRequestBody {
	m.ChatType = &chatType
	return m
}

func (m *V1MessagesRevokeRequestBody) GetChatType() string {
	if m != nil {
		return *m.ChatType
	}
	return ""
}

func (m *V1MessagesRevokeRequestBody) SetMessageIds(messageIds []*string) *V1MessagesRevokeRequestBody {
	m.MessageIds = messageIds
	return m
}

func (m *V1MessagesRevokeRequestBody) GetMessageIds() []*string {
	if m != nil {
		return m.MessageIds
	}
	return nil
}

func (m *V1MessagesRevokeRequestBody) SetSenderId(senderId string) *V1MessagesRevokeRequestBody {
	m.SenderId = &senderId
	return m
}

func (m *V1MessagesRevokeRequestBody) GetSenderId() string {
	if m != nil {
		return *m.SenderId
	}
	return ""
}

func (m *V1MessagesRevokeRequestBody) SetSysMsg(sysMsg SystemMsg) *V1MessagesRevokeRequestBody {
	m.SysMsg = &sysMsg
	return m
}

func (m *V1MessagesRevokeRequestBody) GetSysMsg() *SystemMsg {
	if m != nil {
		return m.SysMsg
	}
	return nil
}

type V1MessagesRevokeResponse struct {
	Data    *V1MessagesRevokeData `json:"data,omitempty"`
	ErrCode *int                  `json:"errCode,omitempty"`
	ErrMsg  *string               `json:"errMsg,omitempty"`
}

func (m *V1MessagesRevokeResponse) GetData() *V1MessagesRevokeData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1MessagesRevokeResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1MessagesRevokeResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1MessagesRevokeResponse) ToString() string {
	return ToJsonString(m)
}

type V1OrgExtraFieldIdsFetchData struct {
	ExtraFieldIds []*ExtraFieldId `json:"extraFieldIds,omitempty"`
	HasMore       *bool           `json:"hasMore,omitempty"`
	LatestVersion *string         `json:"latestVersion,omitempty"`
	Total         *int            `json:"total,omitempty"`
}

func (m *V1OrgExtraFieldIdsFetchData) SetExtraFieldIds(extraFieldIds []*ExtraFieldId) *V1OrgExtraFieldIdsFetchData {
	m.ExtraFieldIds = extraFieldIds
	return m
}

func (m *V1OrgExtraFieldIdsFetchData) GetExtraFieldIds() []*ExtraFieldId {
	if m != nil {
		return m.ExtraFieldIds
	}
	return nil
}

func (m *V1OrgExtraFieldIdsFetchData) SetHasMore(hasMore bool) *V1OrgExtraFieldIdsFetchData {
	m.HasMore = &hasMore
	return m
}

func (m *V1OrgExtraFieldIdsFetchData) GetHasMore() bool {
	if m != nil {
		return *m.HasMore
	}
	return false
}

func (m *V1OrgExtraFieldIdsFetchData) SetLatestVersion(latestVersion string) *V1OrgExtraFieldIdsFetchData {
	m.LatestVersion = &latestVersion
	return m
}

func (m *V1OrgExtraFieldIdsFetchData) GetLatestVersion() string {
	if m != nil {
		return *m.LatestVersion
	}
	return ""
}

func (m *V1OrgExtraFieldIdsFetchData) SetTotal(total int) *V1OrgExtraFieldIdsFetchData {
	m.Total = &total
	return m
}

func (m *V1OrgExtraFieldIdsFetchData) GetTotal() int {
	if m != nil {
		return *m.Total
	}
	return 0
}

type V1OrgExtraFieldIdsFetchResponse struct {
	Data    *V1OrgExtraFieldIdsFetchData `json:"data,omitempty"`
	ErrCode *int                         `json:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty"`
}

func (m *V1OrgExtraFieldIdsFetchResponse) GetData() *V1OrgExtraFieldIdsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1OrgExtraFieldIdsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1OrgExtraFieldIdsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1OrgExtraFieldIdsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1OrgFetchData struct {
	IconUrl *string `json:"iconUrl,omitempty"` // logo资源ID
	OrgId   *string `json:"orgId,omitempty"`   // 组织id
	OrgName *string `json:"orgName,omitempty"` // 组织名称
}

func (m *V1OrgFetchData) SetIconUrl(iconUrl string) *V1OrgFetchData {
	m.IconUrl = &iconUrl
	return m
}

func (m *V1OrgFetchData) GetIconUrl() string {
	if m != nil {
		return *m.IconUrl
	}
	return ""
}

func (m *V1OrgFetchData) SetOrgId(orgId string) *V1OrgFetchData {
	m.OrgId = &orgId
	return m
}

func (m *V1OrgFetchData) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1OrgFetchData) SetOrgName(orgName string) *V1OrgFetchData {
	m.OrgName = &orgName
	return m
}

func (m *V1OrgFetchData) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

type V1OrgFetchResponse struct {
	Data    *V1OrgFetchData `json:"data,omitempty"`
	ErrCode *int            `json:"errCode,omitempty"`
	ErrMsg  *string         `json:"errMsg,omitempty"`
}

func (m *V1OrgFetchResponse) GetData() *V1OrgFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1OrgFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1OrgFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1OrgFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1OrgRoleCreateData struct {
	RoleId *string `json:"roleId,omitempty"`
}

func (m *V1OrgRoleCreateData) SetRoleId(roleId string) *V1OrgRoleCreateData {
	m.RoleId = &roleId
	return m
}

func (m *V1OrgRoleCreateData) GetRoleId() string {
	if m != nil {
		return *m.RoleId
	}
	return ""
}

type V1OrgRoleCreateRequestBody struct {
	Creator     *string `json:"creator,omitempty"`     // 角色创建者
	Description *string `json:"description,omitempty"` // 角色描述
	RoleName    *string `json:"roleName,omitempty"`    // 角色名称（名称需要组织内唯一，不可重复，建议使用：应用开发商名称+应用名称+角色名称）
}

func (m *V1OrgRoleCreateRequestBody) SetCreator(creator string) *V1OrgRoleCreateRequestBody {
	m.Creator = &creator
	return m
}

func (m *V1OrgRoleCreateRequestBody) GetCreator() string {
	if m != nil {
		return *m.Creator
	}
	return ""
}

func (m *V1OrgRoleCreateRequestBody) SetDescription(description string) *V1OrgRoleCreateRequestBody {
	m.Description = &description
	return m
}

func (m *V1OrgRoleCreateRequestBody) GetDescription() string {
	if m != nil {
		return *m.Description
	}
	return ""
}

func (m *V1OrgRoleCreateRequestBody) SetRoleName(roleName string) *V1OrgRoleCreateRequestBody {
	m.RoleName = &roleName
	return m
}

func (m *V1OrgRoleCreateRequestBody) GetRoleName() string {
	if m != nil {
		return *m.RoleName
	}
	return ""
}

type V1OrgRoleCreateResponse struct {
	Data    *V1OrgRoleCreateData `json:"data,omitempty"`
	ErrCode *int                 `json:"errCode,omitempty"`
	ErrMsg  *string              `json:"errMsg,omitempty"`
}

func (m *V1OrgRoleCreateResponse) GetData() *V1OrgRoleCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1OrgRoleCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1OrgRoleCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1OrgRoleCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1OrgRoleMembersFetchData struct {
	Staffs []*RoleStaff `json:"staffs,omitempty"`
}

func (m *V1OrgRoleMembersFetchData) SetStaffs(staffs []*RoleStaff) *V1OrgRoleMembersFetchData {
	m.Staffs = staffs
	return m
}

func (m *V1OrgRoleMembersFetchData) GetStaffs() []*RoleStaff {
	if m != nil {
		return m.Staffs
	}
	return nil
}

type V1OrgRoleMembersFetchResponse struct {
	Data    *V1OrgRoleMembersFetchData `json:"data,omitempty"`
	ErrCode *int                       `json:"errCode,omitempty"`
	ErrMsg  *string                    `json:"errMsg,omitempty"`
}

func (m *V1OrgRoleMembersFetchResponse) GetData() *V1OrgRoleMembersFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1OrgRoleMembersFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1OrgRoleMembersFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1OrgRoleMembersFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1RoleMemberCreateRequestBody struct {
	Creator  *string   `json:"creator,omitempty"`  // 创建者
	RoleId   *string   `json:"roleId,omitempty"`   // 角色ID
	StaffIds []*string `json:"staffIds,omitempty"` // 角色成员列表
}

func (m *V1RoleMemberCreateRequestBody) SetCreator(creator string) *V1RoleMemberCreateRequestBody {
	m.Creator = &creator
	return m
}

func (m *V1RoleMemberCreateRequestBody) GetCreator() string {
	if m != nil {
		return *m.Creator
	}
	return ""
}

func (m *V1RoleMemberCreateRequestBody) SetRoleId(roleId string) *V1RoleMemberCreateRequestBody {
	m.RoleId = &roleId
	return m
}

func (m *V1RoleMemberCreateRequestBody) GetRoleId() string {
	if m != nil {
		return *m.RoleId
	}
	return ""
}

func (m *V1RoleMemberCreateRequestBody) SetStaffIds(staffIds []*string) *V1RoleMemberCreateRequestBody {
	m.StaffIds = staffIds
	return m
}

func (m *V1RoleMemberCreateRequestBody) GetStaffIds() []*string {
	if m != nil {
		return m.StaffIds
	}
	return nil
}

type V1RoleMemberCreateResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1RoleMemberCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1RoleMemberCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1RoleMemberCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1RoleMemberDeleteRequestBody struct {
	Operator *string `json:"operator,omitempty"`
}

func (m *V1RoleMemberDeleteRequestBody) SetOperator(operator string) *V1RoleMemberDeleteRequestBody {
	m.Operator = &operator
	return m
}

func (m *V1RoleMemberDeleteRequestBody) GetOperator() string {
	if m != nil {
		return *m.Operator
	}
	return ""
}

type V1RoleMemberDeleteResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1RoleMemberDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1RoleMemberDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1RoleMemberDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsCreateData struct {
	StaffId *string `json:"staffId,omitempty"`
}

func (m *V1StaffsCreateData) SetStaffId(staffId string) *V1StaffsCreateData {
	m.StaffId = &staffId
	return m
}

func (m *V1StaffsCreateData) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

type V1StaffsCreateRequestBody struct {
	Address              *string           `json:"address,omitempty"`              // 人员(办公)地址
	AvatarId             *string           `json:"avatarId,omitempty"`             // 人员在蓝信里的头像
	AvatarUrl            *string           `json:"avatarUrl,omitempty"`            // 头像地址
	Birthdate            *string           `json:"birthdate,omitempty"`            // 生日
	Career               []*Resume         `json:"career,omitempty"`               // 职业
	Departments          []*UDepartment    `json:"departments,omitempty"`          // 部门
	DisablePasswordReset *bool             `json:"disablePasswordReset,omitempty"` // 账密登录方式，首次登录时是否禁用强制修改密码。 禁用 - true 不禁用 - false (默认)
	Duties               *string           `json:"duties,omitempty"`               // 职务
	Education            []*Resume         `json:"education,omitempty"`            // 教育程度
	Email                *string           `json:"email,omitempty"`                // 电子邮箱地址, loginWays指定邮箱登录时不能为空
	EmployeeNumber       *string           `json:"employeeNumber,omitempty"`       // 人员号
	ExtAttr1             *string           `json:"extAttr1,omitempty"`             // 扩展字段
	ExtAttr2             *string           `json:"extAttr2,omitempty"`             // 扩展字段
	ExternalId           *string           `json:"externalId,omitempty"`           // 人员外部ID，组织通讯录数据源唯一标识人员的ID。注意，如果是账号密码方式登录蓝信，应该使用loginName字段。如果是手机号方式登录蓝信，又想保留组织内的人员唯一ID，可以使用该externalId 字段。externalId 和 employeeNumber类似，用于组织内人员的唯一标识。创建后不可修改，组织内必须唯一。
	ExtraFieldSet        map[string]string `json:"extraFieldSet,omitempty"`        // 自定义扩展属性，k代表扩展字段ID，这个ID由蓝信管理后台预定义，应用可以通过以下接口获取获取组织预定义的人员扩展字段ID列表。
	ExtraPhones          []*MobilePhone    `json:"extraPhones,omitempty"`          // 附加联系方式
	Gender               *int              `json:"gender,omitempty"`               // 性别：0-保密， 1-男， 2-女
	IdNumber             *string           `json:"idNumber,omitempty"`             // 身份证号
	Introduction         *Introduction     `json:"introduction,omitempty"`
	LoginName            *string           `json:"loginName,omitempty"`     // 账密登录方式的登录账号，loginWays指定账密方式登录时不能为空。创建后不可修改，组织内必须唯一。
	LoginPassword        *string           `json:"loginPassword,omitempty"` // 登录方式为账密登录时，设置账户的登录密码，不能为空，为避免密码泄漏，要求原始密码使用哈希算法计算哈希值后再填充该字段，具体哈希算法需要联系组织管理员确认(算法为组织配置，例如SHA256)。
	LoginWays            []*int            `json:"loginWays,omitempty"`     // 蓝信登录方式：0-手机号， 1-邮箱， 2-账密。不填时默认手机号且手机号不能为空。目前一个人只支持一种登录方式。特别说明：如果指定账密登录方式时，需要通过蓝信超级管理员在对应的组织上创建一个组织标识
	MobilePhone          *MobilePhone      `json:"mobilePhone,omitempty"`
	Name                 *string           `json:"name,omitempty"`          // 人员姓名
	Nationality          *string           `json:"nationality,omitempty"`   // 国家
	NativePlace          *string           `json:"nativePlace,omitempty"`   // 籍贯
	OrgId                *string           `json:"orgId,omitempty"`         // 人员所在组织Id
	Parties              *string           `json:"parties,omitempty"`       // 党派
	Signature            *string           `json:"signature,omitempty"`     // 人员个人签名
	SmsInvitation        *bool             `json:"smsInvitation,omitempty"` // 是否发送验证码
	Status               *int              `json:"status,omitempty"`        // 状态
	Tags                 []*string         `json:"tags,omitempty"`          // 人员标签信息
}

func (m *V1StaffsCreateRequestBody) SetAddress(address string) *V1StaffsCreateRequestBody {
	m.Address = &address
	return m
}

func (m *V1StaffsCreateRequestBody) GetAddress() string {
	if m != nil {
		return *m.Address
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetAvatarId(avatarId string) *V1StaffsCreateRequestBody {
	m.AvatarId = &avatarId
	return m
}

func (m *V1StaffsCreateRequestBody) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetAvatarUrl(avatarUrl string) *V1StaffsCreateRequestBody {
	m.AvatarUrl = &avatarUrl
	return m
}

func (m *V1StaffsCreateRequestBody) GetAvatarUrl() string {
	if m != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetBirthdate(birthdate string) *V1StaffsCreateRequestBody {
	m.Birthdate = &birthdate
	return m
}

func (m *V1StaffsCreateRequestBody) GetBirthdate() string {
	if m != nil {
		return *m.Birthdate
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetCareer(career []*Resume) *V1StaffsCreateRequestBody {
	m.Career = career
	return m
}

func (m *V1StaffsCreateRequestBody) GetCareer() []*Resume {
	if m != nil {
		return m.Career
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetDepartments(departments []*UDepartment) *V1StaffsCreateRequestBody {
	m.Departments = departments
	return m
}

func (m *V1StaffsCreateRequestBody) GetDepartments() []*UDepartment {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetDisablePasswordReset(disablePasswordReset bool) *V1StaffsCreateRequestBody {
	m.DisablePasswordReset = &disablePasswordReset
	return m
}

func (m *V1StaffsCreateRequestBody) GetDisablePasswordReset() bool {
	if m != nil {
		return *m.DisablePasswordReset
	}
	return false
}

func (m *V1StaffsCreateRequestBody) SetDuties(duties string) *V1StaffsCreateRequestBody {
	m.Duties = &duties
	return m
}

func (m *V1StaffsCreateRequestBody) GetDuties() string {
	if m != nil {
		return *m.Duties
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetEducation(education []*Resume) *V1StaffsCreateRequestBody {
	m.Education = education
	return m
}

func (m *V1StaffsCreateRequestBody) GetEducation() []*Resume {
	if m != nil {
		return m.Education
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetEmail(email string) *V1StaffsCreateRequestBody {
	m.Email = &email
	return m
}

func (m *V1StaffsCreateRequestBody) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetEmployeeNumber(employeeNumber string) *V1StaffsCreateRequestBody {
	m.EmployeeNumber = &employeeNumber
	return m
}

func (m *V1StaffsCreateRequestBody) GetEmployeeNumber() string {
	if m != nil {
		return *m.EmployeeNumber
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetExtAttr1(extAttr1 string) *V1StaffsCreateRequestBody {
	m.ExtAttr1 = &extAttr1
	return m
}

func (m *V1StaffsCreateRequestBody) GetExtAttr1() string {
	if m != nil {
		return *m.ExtAttr1
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetExtAttr2(extAttr2 string) *V1StaffsCreateRequestBody {
	m.ExtAttr2 = &extAttr2
	return m
}

func (m *V1StaffsCreateRequestBody) GetExtAttr2() string {
	if m != nil {
		return *m.ExtAttr2
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetExternalId(externalId string) *V1StaffsCreateRequestBody {
	m.ExternalId = &externalId
	return m
}

func (m *V1StaffsCreateRequestBody) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetExtraFieldSet(extraFieldSet map[string]string) *V1StaffsCreateRequestBody {
	m.ExtraFieldSet = extraFieldSet
	return m
}

func (m *V1StaffsCreateRequestBody) GetExtraFieldSet() map[string]string {
	if m != nil {
		return m.ExtraFieldSet
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetExtraPhones(extraPhones []*MobilePhone) *V1StaffsCreateRequestBody {
	m.ExtraPhones = extraPhones
	return m
}

func (m *V1StaffsCreateRequestBody) GetExtraPhones() []*MobilePhone {
	if m != nil {
		return m.ExtraPhones
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetGender(gender int) *V1StaffsCreateRequestBody {
	m.Gender = &gender
	return m
}

func (m *V1StaffsCreateRequestBody) GetGender() int {
	if m != nil {
		return *m.Gender
	}
	return 0
}

func (m *V1StaffsCreateRequestBody) SetIdNumber(idNumber string) *V1StaffsCreateRequestBody {
	m.IdNumber = &idNumber
	return m
}

func (m *V1StaffsCreateRequestBody) GetIdNumber() string {
	if m != nil {
		return *m.IdNumber
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetIntroduction(introduction Introduction) *V1StaffsCreateRequestBody {
	m.Introduction = &introduction
	return m
}

func (m *V1StaffsCreateRequestBody) GetIntroduction() *Introduction {
	if m != nil {
		return m.Introduction
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetLoginName(loginName string) *V1StaffsCreateRequestBody {
	m.LoginName = &loginName
	return m
}

func (m *V1StaffsCreateRequestBody) GetLoginName() string {
	if m != nil {
		return *m.LoginName
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetLoginPassword(loginPassword string) *V1StaffsCreateRequestBody {
	m.LoginPassword = &loginPassword
	return m
}

func (m *V1StaffsCreateRequestBody) GetLoginPassword() string {
	if m != nil {
		return *m.LoginPassword
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetLoginWays(loginWays []*int) *V1StaffsCreateRequestBody {
	m.LoginWays = loginWays
	return m
}

func (m *V1StaffsCreateRequestBody) GetLoginWays() []*int {
	if m != nil {
		return m.LoginWays
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetMobilePhone(mobilePhone MobilePhone) *V1StaffsCreateRequestBody {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *V1StaffsCreateRequestBody) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *V1StaffsCreateRequestBody) SetName(name string) *V1StaffsCreateRequestBody {
	m.Name = &name
	return m
}

func (m *V1StaffsCreateRequestBody) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetNationality(nationality string) *V1StaffsCreateRequestBody {
	m.Nationality = &nationality
	return m
}

func (m *V1StaffsCreateRequestBody) GetNationality() string {
	if m != nil {
		return *m.Nationality
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetNativePlace(nativePlace string) *V1StaffsCreateRequestBody {
	m.NativePlace = &nativePlace
	return m
}

func (m *V1StaffsCreateRequestBody) GetNativePlace() string {
	if m != nil {
		return *m.NativePlace
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetOrgId(orgId string) *V1StaffsCreateRequestBody {
	m.OrgId = &orgId
	return m
}

func (m *V1StaffsCreateRequestBody) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetParties(parties string) *V1StaffsCreateRequestBody {
	m.Parties = &parties
	return m
}

func (m *V1StaffsCreateRequestBody) GetParties() string {
	if m != nil {
		return *m.Parties
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetSignature(signature string) *V1StaffsCreateRequestBody {
	m.Signature = &signature
	return m
}

func (m *V1StaffsCreateRequestBody) GetSignature() string {
	if m != nil {
		return *m.Signature
	}
	return ""
}

func (m *V1StaffsCreateRequestBody) SetSmsInvitation(smsInvitation bool) *V1StaffsCreateRequestBody {
	m.SmsInvitation = &smsInvitation
	return m
}

func (m *V1StaffsCreateRequestBody) GetSmsInvitation() bool {
	if m != nil {
		return *m.SmsInvitation
	}
	return false
}

func (m *V1StaffsCreateRequestBody) SetStatus(status int) *V1StaffsCreateRequestBody {
	m.Status = &status
	return m
}

func (m *V1StaffsCreateRequestBody) GetStatus() int {
	if m != nil {
		return *m.Status
	}
	return 0
}

func (m *V1StaffsCreateRequestBody) SetTags(tags []*string) *V1StaffsCreateRequestBody {
	m.Tags = tags
	return m
}

func (m *V1StaffsCreateRequestBody) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1StaffsCreateResponse struct {
	Data    *V1StaffsCreateData `json:"data,omitempty"`
	ErrCode *int                `json:"errCode,omitempty"`
	ErrMsg  *string             `json:"errMsg,omitempty"`
}

func (m *V1StaffsCreateResponse) GetData() *V1StaffsCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1StaffsCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsDeleteData interface{}

type V1StaffsDeleteResponse struct {
	Data    *V1StaffsDeleteData `json:"data,omitempty"`
	ErrCode *int                `json:"errCode,omitempty"`
	ErrMsg  *string             `json:"errMsg,omitempty"`
}

func (m *V1StaffsDeleteResponse) GetData() *V1StaffsDeleteData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1StaffsDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsDeptAncestorsFetchData struct {
	AncestorDepartments []*AncestorDepartment `json:"ancestorDepartments,omitempty"`
}

func (m *V1StaffsDeptAncestorsFetchData) SetAncestorDepartments(ancestorDepartments []*AncestorDepartment) *V1StaffsDeptAncestorsFetchData {
	m.AncestorDepartments = ancestorDepartments
	return m
}

func (m *V1StaffsDeptAncestorsFetchData) GetAncestorDepartments() []*AncestorDepartment {
	if m != nil {
		return m.AncestorDepartments
	}
	return nil
}

type V1StaffsDeptAncestorsFetchResponse struct {
	Data    []*V1StaffsDeptAncestorsFetchData `json:"data,omitempty"`
	ErrCode *int                              `json:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty"`
}

func (m *V1StaffsDeptAncestorsFetchResponse) GetData() []*V1StaffsDeptAncestorsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1StaffsDeptAncestorsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsDeptAncestorsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsDeptAncestorsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsFetchData struct {
	AvatarId    *string       `json:"avatarId,omitempty"`    // 人员图像资源ID
	AvatarUrl   *string       `json:"avatarUrl,omitempty"`   // 人员图像下载地址
	Departments []*Department `json:"departments,omitempty"` // 所在部门信息
	Gender      *int          `json:"gender,omitempty"`      // 性别 （可选值： 0 ： 保密 ， 1： 男 ， 2 ：女）
	Name        *string       `json:"name,omitempty"`        // 人员姓名
	OrgId       *string       `json:"orgId,omitempty"`       // 人员所在组织Id
	OrgName     *string       `json:"orgName,omitempty"`     // 人员组织名
	Signature   *string       `json:"signature,omitempty"`   // 签名
	StaffId     *string       `json:"staffId,omitempty"`     // 人员ID
	Status      *int          `json:"status,omitempty"`      // 成员状态 INACTIVE = 0, 未注册; NORMAL = 1, 已注册; FROZEN = 2, 已冻结; DELETED = 3, 已删除
}

func (m *V1StaffsFetchData) SetAvatarId(avatarId string) *V1StaffsFetchData {
	m.AvatarId = &avatarId
	return m
}

func (m *V1StaffsFetchData) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *V1StaffsFetchData) SetAvatarUrl(avatarUrl string) *V1StaffsFetchData {
	m.AvatarUrl = &avatarUrl
	return m
}

func (m *V1StaffsFetchData) GetAvatarUrl() string {
	if m != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *V1StaffsFetchData) SetDepartments(departments []*Department) *V1StaffsFetchData {
	m.Departments = departments
	return m
}

func (m *V1StaffsFetchData) GetDepartments() []*Department {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *V1StaffsFetchData) SetGender(gender int) *V1StaffsFetchData {
	m.Gender = &gender
	return m
}

func (m *V1StaffsFetchData) GetGender() int {
	if m != nil {
		return *m.Gender
	}
	return 0
}

func (m *V1StaffsFetchData) SetName(name string) *V1StaffsFetchData {
	m.Name = &name
	return m
}

func (m *V1StaffsFetchData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1StaffsFetchData) SetOrgId(orgId string) *V1StaffsFetchData {
	m.OrgId = &orgId
	return m
}

func (m *V1StaffsFetchData) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1StaffsFetchData) SetOrgName(orgName string) *V1StaffsFetchData {
	m.OrgName = &orgName
	return m
}

func (m *V1StaffsFetchData) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

func (m *V1StaffsFetchData) SetSignature(signature string) *V1StaffsFetchData {
	m.Signature = &signature
	return m
}

func (m *V1StaffsFetchData) GetSignature() string {
	if m != nil {
		return *m.Signature
	}
	return ""
}

func (m *V1StaffsFetchData) SetStaffId(staffId string) *V1StaffsFetchData {
	m.StaffId = &staffId
	return m
}

func (m *V1StaffsFetchData) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

func (m *V1StaffsFetchData) SetStatus(status int) *V1StaffsFetchData {
	m.Status = &status
	return m
}

func (m *V1StaffsFetchData) GetStatus() int {
	if m != nil {
		return *m.Status
	}
	return 0
}

type V1StaffsFetchResponse struct {
	Data    *V1StaffsFetchData `json:"data,omitempty"`
	ErrCode *int               `json:"errCode,omitempty"`
	ErrMsg  *string            `json:"errMsg,omitempty"`
}

func (m *V1StaffsFetchResponse) GetData() *V1StaffsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1StaffsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsInforFetchData struct {
	AvatarId       *string           `json:"avatarId,omitempty"`
	AvatarUrl      *string           `json:"avatarUrl,omitempty"`
	Birthdate      *string           `json:"birthdate,omitempty"`
	Career         []*Resume         `json:"career,omitempty"`
	Departments    []*Department     `json:"departments,omitempty"`
	Duties         *string           `json:"duties,omitempty"`
	Education      []*Resume         `json:"education,omitempty"`
	Email          *string           `json:"email,omitempty"`
	EmployeeNumber *string           `json:"employeeNumber,omitempty"`
	ExternalId     *string           `json:"externalId,omitempty"`
	ExtraFieldSet  map[string]string `json:"extraFieldSet,omitempty"`
	ExtraPhones    []*MobilePhone    `json:"extraPhones,omitempty"`
	Gender         *int              `json:"gender,omitempty"`
	IdNumber       *string           `json:"idNumber,omitempty"`
	Introduction   *Introduction     `json:"introduction,omitempty"`
	LoginName      *string           `json:"loginName,omitempty"`
	LoginPassword  *string           `json:"loginPassword,omitempty"`
	LoginWays      []*int            `json:"loginWays,omitempty"`
	MobilePhone    *MobilePhone      `json:"mobilePhone,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Nationality    *string           `json:"nationality,omitempty"`
	NativePlace    *string           `json:"nativePlace,omitempty"`
	OrgId          *string           `json:"orgId,omitempty"`
	OrgName        *string           `json:"orgName,omitempty"`
	Parties        *string           `json:"parties,omitempty"`
	Signature      *string           `json:"signature,omitempty"`
	Status         *int              `json:"status,omitempty"`
	Tags           []*string         `json:"tags,omitempty"`
}

func (m *V1StaffsInforFetchData) SetAvatarId(avatarId string) *V1StaffsInforFetchData {
	m.AvatarId = &avatarId
	return m
}

func (m *V1StaffsInforFetchData) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetAvatarUrl(avatarUrl string) *V1StaffsInforFetchData {
	m.AvatarUrl = &avatarUrl
	return m
}

func (m *V1StaffsInforFetchData) GetAvatarUrl() string {
	if m != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetBirthdate(birthdate string) *V1StaffsInforFetchData {
	m.Birthdate = &birthdate
	return m
}

func (m *V1StaffsInforFetchData) GetBirthdate() string {
	if m != nil {
		return *m.Birthdate
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetCareer(career []*Resume) *V1StaffsInforFetchData {
	m.Career = career
	return m
}

func (m *V1StaffsInforFetchData) GetCareer() []*Resume {
	if m != nil {
		return m.Career
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetDepartments(departments []*Department) *V1StaffsInforFetchData {
	m.Departments = departments
	return m
}

func (m *V1StaffsInforFetchData) GetDepartments() []*Department {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetDuties(duties string) *V1StaffsInforFetchData {
	m.Duties = &duties
	return m
}

func (m *V1StaffsInforFetchData) GetDuties() string {
	if m != nil {
		return *m.Duties
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetEducation(education []*Resume) *V1StaffsInforFetchData {
	m.Education = education
	return m
}

func (m *V1StaffsInforFetchData) GetEducation() []*Resume {
	if m != nil {
		return m.Education
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetEmail(email string) *V1StaffsInforFetchData {
	m.Email = &email
	return m
}

func (m *V1StaffsInforFetchData) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetEmployeeNumber(employeeNumber string) *V1StaffsInforFetchData {
	m.EmployeeNumber = &employeeNumber
	return m
}

func (m *V1StaffsInforFetchData) GetEmployeeNumber() string {
	if m != nil {
		return *m.EmployeeNumber
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetExternalId(externalId string) *V1StaffsInforFetchData {
	m.ExternalId = &externalId
	return m
}

func (m *V1StaffsInforFetchData) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetExtraFieldSet(extraFieldSet map[string]string) *V1StaffsInforFetchData {
	m.ExtraFieldSet = extraFieldSet
	return m
}

func (m *V1StaffsInforFetchData) GetExtraFieldSet() map[string]string {
	if m != nil {
		return m.ExtraFieldSet
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetExtraPhones(extraPhones []*MobilePhone) *V1StaffsInforFetchData {
	m.ExtraPhones = extraPhones
	return m
}

func (m *V1StaffsInforFetchData) GetExtraPhones() []*MobilePhone {
	if m != nil {
		return m.ExtraPhones
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetGender(gender int) *V1StaffsInforFetchData {
	m.Gender = &gender
	return m
}

func (m *V1StaffsInforFetchData) GetGender() int {
	if m != nil {
		return *m.Gender
	}
	return 0
}

func (m *V1StaffsInforFetchData) SetIdNumber(idNumber string) *V1StaffsInforFetchData {
	m.IdNumber = &idNumber
	return m
}

func (m *V1StaffsInforFetchData) GetIdNumber() string {
	if m != nil {
		return *m.IdNumber
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetIntroduction(introduction Introduction) *V1StaffsInforFetchData {
	m.Introduction = &introduction
	return m
}

func (m *V1StaffsInforFetchData) GetIntroduction() *Introduction {
	if m != nil {
		return m.Introduction
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetLoginName(loginName string) *V1StaffsInforFetchData {
	m.LoginName = &loginName
	return m
}

func (m *V1StaffsInforFetchData) GetLoginName() string {
	if m != nil {
		return *m.LoginName
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetLoginPassword(loginPassword string) *V1StaffsInforFetchData {
	m.LoginPassword = &loginPassword
	return m
}

func (m *V1StaffsInforFetchData) GetLoginPassword() string {
	if m != nil {
		return *m.LoginPassword
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetLoginWays(loginWays []*int) *V1StaffsInforFetchData {
	m.LoginWays = loginWays
	return m
}

func (m *V1StaffsInforFetchData) GetLoginWays() []*int {
	if m != nil {
		return m.LoginWays
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetMobilePhone(mobilePhone MobilePhone) *V1StaffsInforFetchData {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *V1StaffsInforFetchData) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *V1StaffsInforFetchData) SetName(name string) *V1StaffsInforFetchData {
	m.Name = &name
	return m
}

func (m *V1StaffsInforFetchData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetNationality(nationality string) *V1StaffsInforFetchData {
	m.Nationality = &nationality
	return m
}

func (m *V1StaffsInforFetchData) GetNationality() string {
	if m != nil {
		return *m.Nationality
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetNativePlace(nativePlace string) *V1StaffsInforFetchData {
	m.NativePlace = &nativePlace
	return m
}

func (m *V1StaffsInforFetchData) GetNativePlace() string {
	if m != nil {
		return *m.NativePlace
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetOrgId(orgId string) *V1StaffsInforFetchData {
	m.OrgId = &orgId
	return m
}

func (m *V1StaffsInforFetchData) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetOrgName(orgName string) *V1StaffsInforFetchData {
	m.OrgName = &orgName
	return m
}

func (m *V1StaffsInforFetchData) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetParties(parties string) *V1StaffsInforFetchData {
	m.Parties = &parties
	return m
}

func (m *V1StaffsInforFetchData) GetParties() string {
	if m != nil {
		return *m.Parties
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetSignature(signature string) *V1StaffsInforFetchData {
	m.Signature = &signature
	return m
}

func (m *V1StaffsInforFetchData) GetSignature() string {
	if m != nil {
		return *m.Signature
	}
	return ""
}

func (m *V1StaffsInforFetchData) SetStatus(status int) *V1StaffsInforFetchData {
	m.Status = &status
	return m
}

func (m *V1StaffsInforFetchData) GetStatus() int {
	if m != nil {
		return *m.Status
	}
	return 0
}

func (m *V1StaffsInforFetchData) SetTags(tags []*string) *V1StaffsInforFetchData {
	m.Tags = tags
	return m
}

func (m *V1StaffsInforFetchData) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1StaffsInforFetchResponse struct {
	Data    *V1StaffsInforFetchData `json:"data,omitempty"`
	ErrCode *int                    `json:"errCode,omitempty"`
	ErrMsg  *string                 `json:"errMsg,omitempty"`
}

func (m *V1StaffsInforFetchResponse) GetData() *V1StaffsInforFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1StaffsInforFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsInforFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsInforFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1StaffsUpdateRequestBody struct {
	Address        *string                `json:"address,omitempty"`        // 人员(办公)地址
	AvatarId       *string                `json:"avatarId,omitempty"`       // 人员在蓝信里的头像
	Birthdate      *string                `json:"birthdate,omitempty"`      // 出生日期。格式：yyyy-mm-dd
	Career         []*Resume              `json:"career,omitempty"`         // 职业 日期格式：yyyy-mm-dd
	Departments    []*UDepartment         `json:"departments,omitempty"`    // 人员所在分支，含分支Id和人员在分支里的排序。人员在分支下排序的序号，不确定时可填0，服务端为忽略0值。有效值从1开始
	Duties         *string                `json:"duties,omitempty"`         // 职务
	Education      []*Resume              `json:"education,omitempty"`      // 教育程度 日期格式：yyyy-mm-dd
	Email          *string                `json:"email,omitempty"`          // 电子邮箱地址
	EmployeeNumber *string                `json:"employeeNumber,omitempty"` // 人员号
	ExtAttr1       *string                `json:"extAttr1,omitempty"`       // 扩展字段
	ExtAttr2       *string                `json:"extAttr2,omitempty"`       // 扩展字段
	ExternalId     *string                `json:"externalId,omitempty"`
	ExtraFieldSet  map[string]interface{} `json:"extraFieldSet,omitempty"` // 自定义扩展属性，k代表扩展字段id。
	ExtraPhones    []*MobilePhone         `json:"extraPhones,omitempty"`   // 附加联系方式
	Gender         *int                   `json:"gender,omitempty"`        // 性别：0-保密， 1-男， 2-女
	IdNumber       *string                `json:"idNumber,omitempty"`      // 身份证号
	Introduction   *Introduction          `json:"introduction,omitempty"`
	LoginName      *string                `json:"loginName,omitempty"` // 人员使用人员名登录蓝信时的人员名，也称staffId。可由组织在创建时指定，并代表一定含义比如工号，创建后不可修改，组织内必须唯一。
	LoginWays      []*int                 `json:"loginWays,omitempty"` // 蓝信登录方式：0-手机号， 1-邮箱， 2-账密。不填时默认手机号且手机号不能为空
	MobilePhone    *MobilePhone           `json:"mobilePhone,omitempty"`
	Name           *string                `json:"name,omitempty"`        // 人员姓名
	Nationality    *string                `json:"nationality,omitempty"` // 民族
	NativePlace    *string                `json:"nativePlace,omitempty"` // 籍贯
	Parties        *string                `json:"parties,omitempty"`     // 党派
	Signature      *string                `json:"signature,omitempty"`   // 签名
	Status         *int                   `json:"status,omitempty"`      // 成员状态, 更新时允许值 ：NORMAL= 1, 已注册; FROZEN = 2, 冻结
	Tags           []*string              `json:"tags,omitempty"`        // 人员标签信息
}

func (m *V1StaffsUpdateRequestBody) SetAddress(address string) *V1StaffsUpdateRequestBody {
	m.Address = &address
	return m
}

func (m *V1StaffsUpdateRequestBody) GetAddress() string {
	if m != nil {
		return *m.Address
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetAvatarId(avatarId string) *V1StaffsUpdateRequestBody {
	m.AvatarId = &avatarId
	return m
}

func (m *V1StaffsUpdateRequestBody) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetBirthdate(birthdate string) *V1StaffsUpdateRequestBody {
	m.Birthdate = &birthdate
	return m
}

func (m *V1StaffsUpdateRequestBody) GetBirthdate() string {
	if m != nil {
		return *m.Birthdate
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetCareer(career []*Resume) *V1StaffsUpdateRequestBody {
	m.Career = career
	return m
}

func (m *V1StaffsUpdateRequestBody) GetCareer() []*Resume {
	if m != nil {
		return m.Career
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetDepartments(departments []*UDepartment) *V1StaffsUpdateRequestBody {
	m.Departments = departments
	return m
}

func (m *V1StaffsUpdateRequestBody) GetDepartments() []*UDepartment {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetDuties(duties string) *V1StaffsUpdateRequestBody {
	m.Duties = &duties
	return m
}

func (m *V1StaffsUpdateRequestBody) GetDuties() string {
	if m != nil {
		return *m.Duties
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetEducation(education []*Resume) *V1StaffsUpdateRequestBody {
	m.Education = education
	return m
}

func (m *V1StaffsUpdateRequestBody) GetEducation() []*Resume {
	if m != nil {
		return m.Education
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetEmail(email string) *V1StaffsUpdateRequestBody {
	m.Email = &email
	return m
}

func (m *V1StaffsUpdateRequestBody) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetEmployeeNumber(employeeNumber string) *V1StaffsUpdateRequestBody {
	m.EmployeeNumber = &employeeNumber
	return m
}

func (m *V1StaffsUpdateRequestBody) GetEmployeeNumber() string {
	if m != nil {
		return *m.EmployeeNumber
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetExtAttr1(extAttr1 string) *V1StaffsUpdateRequestBody {
	m.ExtAttr1 = &extAttr1
	return m
}

func (m *V1StaffsUpdateRequestBody) GetExtAttr1() string {
	if m != nil {
		return *m.ExtAttr1
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetExtAttr2(extAttr2 string) *V1StaffsUpdateRequestBody {
	m.ExtAttr2 = &extAttr2
	return m
}

func (m *V1StaffsUpdateRequestBody) GetExtAttr2() string {
	if m != nil {
		return *m.ExtAttr2
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetExternalId(externalId string) *V1StaffsUpdateRequestBody {
	m.ExternalId = &externalId
	return m
}

func (m *V1StaffsUpdateRequestBody) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetExtraFieldSet(extraFieldSet map[string]interface{}) *V1StaffsUpdateRequestBody {
	m.ExtraFieldSet = extraFieldSet
	return m
}

func (m *V1StaffsUpdateRequestBody) GetExtraFieldSet() map[string]interface{} {
	if m != nil {
		return m.ExtraFieldSet
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetExtraPhones(extraPhones []*MobilePhone) *V1StaffsUpdateRequestBody {
	m.ExtraPhones = extraPhones
	return m
}

func (m *V1StaffsUpdateRequestBody) GetExtraPhones() []*MobilePhone {
	if m != nil {
		return m.ExtraPhones
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetGender(gender int) *V1StaffsUpdateRequestBody {
	m.Gender = &gender
	return m
}

func (m *V1StaffsUpdateRequestBody) GetGender() int {
	if m != nil {
		return *m.Gender
	}
	return 0
}

func (m *V1StaffsUpdateRequestBody) SetIdNumber(idNumber string) *V1StaffsUpdateRequestBody {
	m.IdNumber = &idNumber
	return m
}

func (m *V1StaffsUpdateRequestBody) GetIdNumber() string {
	if m != nil {
		return *m.IdNumber
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetIntroduction(introduction Introduction) *V1StaffsUpdateRequestBody {
	m.Introduction = &introduction
	return m
}

func (m *V1StaffsUpdateRequestBody) GetIntroduction() *Introduction {
	if m != nil {
		return m.Introduction
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetLoginName(loginName string) *V1StaffsUpdateRequestBody {
	m.LoginName = &loginName
	return m
}

func (m *V1StaffsUpdateRequestBody) GetLoginName() string {
	if m != nil {
		return *m.LoginName
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetLoginWays(loginWays []*int) *V1StaffsUpdateRequestBody {
	m.LoginWays = loginWays
	return m
}

func (m *V1StaffsUpdateRequestBody) GetLoginWays() []*int {
	if m != nil {
		return m.LoginWays
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetMobilePhone(mobilePhone MobilePhone) *V1StaffsUpdateRequestBody {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *V1StaffsUpdateRequestBody) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *V1StaffsUpdateRequestBody) SetName(name string) *V1StaffsUpdateRequestBody {
	m.Name = &name
	return m
}

func (m *V1StaffsUpdateRequestBody) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetNationality(nationality string) *V1StaffsUpdateRequestBody {
	m.Nationality = &nationality
	return m
}

func (m *V1StaffsUpdateRequestBody) GetNationality() string {
	if m != nil {
		return *m.Nationality
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetNativePlace(nativePlace string) *V1StaffsUpdateRequestBody {
	m.NativePlace = &nativePlace
	return m
}

func (m *V1StaffsUpdateRequestBody) GetNativePlace() string {
	if m != nil {
		return *m.NativePlace
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetParties(parties string) *V1StaffsUpdateRequestBody {
	m.Parties = &parties
	return m
}

func (m *V1StaffsUpdateRequestBody) GetParties() string {
	if m != nil {
		return *m.Parties
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetSignature(signature string) *V1StaffsUpdateRequestBody {
	m.Signature = &signature
	return m
}

func (m *V1StaffsUpdateRequestBody) GetSignature() string {
	if m != nil {
		return *m.Signature
	}
	return ""
}

func (m *V1StaffsUpdateRequestBody) SetStatus(status int) *V1StaffsUpdateRequestBody {
	m.Status = &status
	return m
}

func (m *V1StaffsUpdateRequestBody) GetStatus() int {
	if m != nil {
		return *m.Status
	}
	return 0
}

func (m *V1StaffsUpdateRequestBody) SetTags(tags []*string) *V1StaffsUpdateRequestBody {
	m.Tags = tags
	return m
}

func (m *V1StaffsUpdateRequestBody) GetTags() []*string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1StaffsUpdateResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1StaffsUpdateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1StaffsUpdateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1StaffsUpdateResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagGroupsCreateData struct {
	TagGroupId *string `json:"tagGroupId,omitempty"`
}

func (m *V1TagGroupsCreateData) SetTagGroupId(tagGroupId string) *V1TagGroupsCreateData {
	m.TagGroupId = &tagGroupId
	return m
}

func (m *V1TagGroupsCreateData) GetTagGroupId() string {
	if m != nil {
		return *m.TagGroupId
	}
	return ""
}

type V1TagGroupsCreateRequestBody struct {
	OrgId            *string `json:"orgId,omitempty"`            // 组织ID
	TagGroupCategory *int    `json:"tagGroupCategory,omitempty"` // 标签类别。1：用于staff，2：用于department
	TagGroupName     *string `json:"tagGroupName,omitempty"`     // 标签分组名称（名称需要组织内唯一，不可重复，建议使用：应用开发商名称+应用名称+标签分组名称）
	TagGroupType     *string `json:"tagGroupType,omitempty"`     // 标签组类型，用于分类检索, 非必填。通过这个字段应用可以方便的管理自己所创建的标签分组，是个自定义字符串，建议使用应用的域名再加个uuid保证唯一性
}

func (m *V1TagGroupsCreateRequestBody) SetOrgId(orgId string) *V1TagGroupsCreateRequestBody {
	m.OrgId = &orgId
	return m
}

func (m *V1TagGroupsCreateRequestBody) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1TagGroupsCreateRequestBody) SetTagGroupCategory(tagGroupCategory int) *V1TagGroupsCreateRequestBody {
	m.TagGroupCategory = &tagGroupCategory
	return m
}

func (m *V1TagGroupsCreateRequestBody) GetTagGroupCategory() int {
	if m != nil {
		return *m.TagGroupCategory
	}
	return 0
}

func (m *V1TagGroupsCreateRequestBody) SetTagGroupName(tagGroupName string) *V1TagGroupsCreateRequestBody {
	m.TagGroupName = &tagGroupName
	return m
}

func (m *V1TagGroupsCreateRequestBody) GetTagGroupName() string {
	if m != nil {
		return *m.TagGroupName
	}
	return ""
}

func (m *V1TagGroupsCreateRequestBody) SetTagGroupType(tagGroupType string) *V1TagGroupsCreateRequestBody {
	m.TagGroupType = &tagGroupType
	return m
}

func (m *V1TagGroupsCreateRequestBody) GetTagGroupType() string {
	if m != nil {
		return *m.TagGroupType
	}
	return ""
}

type V1TagGroupsCreateResponse struct {
	Data    *V1TagGroupsCreateData `json:"data,omitempty"`
	ErrCode *int                   `json:"errCode,omitempty"`
	ErrMsg  *string                `json:"errMsg,omitempty"`
}

func (m *V1TagGroupsCreateResponse) GetData() *V1TagGroupsCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagGroupsCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagGroupsCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagGroupsCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagGroupsDeleteResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1TagGroupsDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagGroupsDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagGroupsDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagGroupsFetchData struct {
	TagGroupIds []*string `json:"tagGroupIds,omitempty"`
}

func (m *V1TagGroupsFetchData) SetTagGroupIds(tagGroupIds []*string) *V1TagGroupsFetchData {
	m.TagGroupIds = tagGroupIds
	return m
}

func (m *V1TagGroupsFetchData) GetTagGroupIds() []*string {
	if m != nil {
		return m.TagGroupIds
	}
	return nil
}

type V1TagGroupsFetchRequestBody struct {
	CategoryList []*int  `json:"categoryList,omitempty"` // 标签组分类。1：人员标签，2：分支标签
	OrgId        *string `json:"orgId,omitempty"`        // 组织ID
	TagGroupType *string `json:"tagGroupType,omitempty"` // 标签组类型索引
}

func (m *V1TagGroupsFetchRequestBody) SetCategoryList(categoryList []*int) *V1TagGroupsFetchRequestBody {
	m.CategoryList = categoryList
	return m
}

func (m *V1TagGroupsFetchRequestBody) GetCategoryList() []*int {
	if m != nil {
		return m.CategoryList
	}
	return nil
}

func (m *V1TagGroupsFetchRequestBody) SetOrgId(orgId string) *V1TagGroupsFetchRequestBody {
	m.OrgId = &orgId
	return m
}

func (m *V1TagGroupsFetchRequestBody) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1TagGroupsFetchRequestBody) SetTagGroupType(tagGroupType string) *V1TagGroupsFetchRequestBody {
	m.TagGroupType = &tagGroupType
	return m
}

func (m *V1TagGroupsFetchRequestBody) GetTagGroupType() string {
	if m != nil {
		return *m.TagGroupType
	}
	return ""
}

type V1TagGroupsFetchResponse struct {
	Data    *V1TagGroupsFetchData `json:"data,omitempty"`
	ErrCode *int                  `json:"errCode,omitempty"`
	ErrMsg  *string               `json:"errMsg,omitempty"`
}

func (m *V1TagGroupsFetchResponse) GetData() *V1TagGroupsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagGroupsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagGroupsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagGroupsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagGroupsInfoFetchData struct {
	TagGroupCategory *int            `json:"tagGroupCategory,omitempty"` // 标签分类。1：staff分类，2：department分类。
	TagGroupName     *string         `json:"tagGroupName,omitempty"`     // 标签组名称
	TagGroupType     *string         `json:"tagGroupType,omitempty"`     // 标签组类型（用于索引）
	Tags             []*GroupInfoTag `json:"tags,omitempty"`             // 标签列表
}

func (m *V1TagGroupsInfoFetchData) SetTagGroupCategory(tagGroupCategory int) *V1TagGroupsInfoFetchData {
	m.TagGroupCategory = &tagGroupCategory
	return m
}

func (m *V1TagGroupsInfoFetchData) GetTagGroupCategory() int {
	if m != nil {
		return *m.TagGroupCategory
	}
	return 0
}

func (m *V1TagGroupsInfoFetchData) SetTagGroupName(tagGroupName string) *V1TagGroupsInfoFetchData {
	m.TagGroupName = &tagGroupName
	return m
}

func (m *V1TagGroupsInfoFetchData) GetTagGroupName() string {
	if m != nil {
		return *m.TagGroupName
	}
	return ""
}

func (m *V1TagGroupsInfoFetchData) SetTagGroupType(tagGroupType string) *V1TagGroupsInfoFetchData {
	m.TagGroupType = &tagGroupType
	return m
}

func (m *V1TagGroupsInfoFetchData) GetTagGroupType() string {
	if m != nil {
		return *m.TagGroupType
	}
	return ""
}

func (m *V1TagGroupsInfoFetchData) SetTags(tags []*GroupInfoTag) *V1TagGroupsInfoFetchData {
	m.Tags = tags
	return m
}

func (m *V1TagGroupsInfoFetchData) GetTags() []*GroupInfoTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type V1TagGroupsInfoFetchResponse struct {
	Data    *V1TagGroupsInfoFetchData `json:"data,omitempty"`
	ErrCode *int                      `json:"errCode,omitempty"`
	ErrMsg  *string                   `json:"errMsg,omitempty"`
}

func (m *V1TagGroupsInfoFetchResponse) GetData() *V1TagGroupsInfoFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagGroupsInfoFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagGroupsInfoFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagGroupsInfoFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagGroupsUpdateRequestBody struct {
	TagGroupName *string `json:"tagGroupName,omitempty"` // 标签组的名字
}

func (m *V1TagGroupsUpdateRequestBody) SetTagGroupName(tagGroupName string) *V1TagGroupsUpdateRequestBody {
	m.TagGroupName = &tagGroupName
	return m
}

func (m *V1TagGroupsUpdateRequestBody) GetTagGroupName() string {
	if m != nil {
		return *m.TagGroupName
	}
	return ""
}

type V1TagGroupsUpdateResponse struct {
	ErrCode *int    `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}

func (m *V1TagGroupsUpdateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagGroupsUpdateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagGroupsUpdateResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagsCreateData struct {
	TagId *string `json:"tagId,omitempty"`
}

func (m *V1TagsCreateData) SetTagId(tagId string) *V1TagsCreateData {
	m.TagId = &tagId
	return m
}

func (m *V1TagsCreateData) GetTagId() string {
	if m != nil {
		return *m.TagId
	}
	return ""
}

type V1TagsCreateRequestBody struct {
	TagGroupId *string `json:"tagGroupId,omitempty"`
	TagName    *string `json:"tagName,omitempty"`
}

func (m *V1TagsCreateRequestBody) SetTagGroupId(tagGroupId string) *V1TagsCreateRequestBody {
	m.TagGroupId = &tagGroupId
	return m
}

func (m *V1TagsCreateRequestBody) GetTagGroupId() string {
	if m != nil {
		return *m.TagGroupId
	}
	return ""
}

func (m *V1TagsCreateRequestBody) SetTagName(tagName string) *V1TagsCreateRequestBody {
	m.TagName = &tagName
	return m
}

func (m *V1TagsCreateRequestBody) GetTagName() string {
	if m != nil {
		return *m.TagName
	}
	return ""
}

type V1TagsCreateResponse struct {
	Data    *V1TagsCreateData `json:"data,omitempty"`
	ErrCode *int              `json:"errCode,omitempty"`
	ErrMsg  *string           `json:"errMsg,omitempty"`
}

func (m *V1TagsCreateResponse) GetData() *V1TagsCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagsCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagsCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagsCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagsDeleteData interface{}

type V1TagsDeleteResponse struct {
	Data    *V1TagsDeleteData `json:"data,omitempty"`
	ErrCode *int              `json:"errCode,omitempty"`
	ErrMsg  *string           `json:"errMsg,omitempty"`
}

func (m *V1TagsDeleteResponse) GetData() *V1TagsDeleteData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagsDeleteResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagsDeleteResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagsDeleteResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagsFetchData struct {
	HasMore       *bool           `json:"hasMore,omitempty"`
	LatestVersion *string         `json:"latestVersion,omitempty"`
	Staffs        []*TagStaffInfo `json:"staffs,omitempty"`
	Total         *int            `json:"total,omitempty"`
}

func (m *V1TagsFetchData) SetHasMore(hasMore bool) *V1TagsFetchData {
	m.HasMore = &hasMore
	return m
}

func (m *V1TagsFetchData) GetHasMore() bool {
	if m != nil {
		return *m.HasMore
	}
	return false
}

func (m *V1TagsFetchData) SetLatestVersion(latestVersion string) *V1TagsFetchData {
	m.LatestVersion = &latestVersion
	return m
}

func (m *V1TagsFetchData) GetLatestVersion() string {
	if m != nil {
		return *m.LatestVersion
	}
	return ""
}

func (m *V1TagsFetchData) SetStaffs(staffs []*TagStaffInfo) *V1TagsFetchData {
	m.Staffs = staffs
	return m
}

func (m *V1TagsFetchData) GetStaffs() []*TagStaffInfo {
	if m != nil {
		return m.Staffs
	}
	return nil
}

func (m *V1TagsFetchData) SetTotal(total int) *V1TagsFetchData {
	m.Total = &total
	return m
}

func (m *V1TagsFetchData) GetTotal() int {
	if m != nil {
		return *m.Total
	}
	return 0
}

type V1TagsFetchRequestBody struct {
	OrgId      *string      `json:"orgId,omitempty"`
	TagFilters []*TagFilter `json:"tagFilters,omitempty"`
}

func (m *V1TagsFetchRequestBody) SetOrgId(orgId string) *V1TagsFetchRequestBody {
	m.OrgId = &orgId
	return m
}

func (m *V1TagsFetchRequestBody) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1TagsFetchRequestBody) SetTagFilters(tagFilters []*TagFilter) *V1TagsFetchRequestBody {
	m.TagFilters = tagFilters
	return m
}

func (m *V1TagsFetchRequestBody) GetTagFilters() []*TagFilter {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

type V1TagsFetchResponse struct {
	Data    *V1TagsFetchData `json:"data,omitempty"`
	ErrCode *int             `json:"errCode,omitempty"`
	ErrMsg  *string          `json:"errMsg,omitempty"`
}

func (m *V1TagsFetchResponse) GetData() *V1TagsFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagsFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagsFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagsFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagsMetaFetchData struct {
	TagMetaList []*TagMeta `json:"tagMetaList,omitempty"`
}

func (m *V1TagsMetaFetchData) SetTagMetaList(tagMetaList []*TagMeta) *V1TagsMetaFetchData {
	m.TagMetaList = tagMetaList
	return m
}

func (m *V1TagsMetaFetchData) GetTagMetaList() []*TagMeta {
	if m != nil {
		return m.TagMetaList
	}
	return nil
}

type V1TagsMetaFetchRequestBody struct {
	TagIds []*string `json:"tagIds,omitempty"`
}

func (m *V1TagsMetaFetchRequestBody) SetTagIds(tagIds []*string) *V1TagsMetaFetchRequestBody {
	m.TagIds = tagIds
	return m
}

func (m *V1TagsMetaFetchRequestBody) GetTagIds() []*string {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type V1TagsMetaFetchResponse struct {
	Data    *V1TagsMetaFetchData `json:"data,omitempty"`
	ErrCode *int                 `json:"errCode,omitempty"`
	ErrMsg  *string              `json:"errMsg,omitempty"`
}

func (m *V1TagsMetaFetchResponse) GetData() *V1TagsMetaFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagsMetaFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagsMetaFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagsMetaFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V1TagsUpdateData interface{}

type V1TagsUpdateRequestBody struct {
	TagName *string `json:"tagName,omitempty"`
}

func (m *V1TagsUpdateRequestBody) SetTagName(tagName string) *V1TagsUpdateRequestBody {
	m.TagName = &tagName
	return m
}

func (m *V1TagsUpdateRequestBody) GetTagName() string {
	if m != nil {
		return *m.TagName
	}
	return ""
}

type V1TagsUpdateResponse struct {
	Data    *V1TagsUpdateData `json:"data,omitempty"`
	ErrCode *int              `json:"errCode,omitempty"`
	ErrMsg  *string           `json:"errMsg,omitempty"`
}

func (m *V1TagsUpdateResponse) GetData() *V1TagsUpdateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1TagsUpdateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1TagsUpdateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1TagsUpdateResponse) ToString() string {
	return ToJsonString(m)
}

type V1UserTokenCreateData struct {
	ExpiresIn *int    `json:"expiresIn,omitempty"` // Token 有效期（7200秒）
	Scope     *string `json:"scope,omitempty"`     // 应用授权列表，多个授权请求需要以 scope="scope1,scope2" 形式。
	StaffId   *string `json:"staffId,omitempty"`   // 人员ID
	State     *string `json:"state,omitempty"`     // 发起请求的时候携带的随机值，和该重定向请求唯一对应。同时，也能 按 OAUTH2 协议防止 CSRF 攻击
	UserToken *string `json:"userToken,omitempty"` // 人员 Token
}

func (m *V1UserTokenCreateData) SetExpiresIn(expiresIn int) *V1UserTokenCreateData {
	m.ExpiresIn = &expiresIn
	return m
}

func (m *V1UserTokenCreateData) GetExpiresIn() int {
	if m != nil {
		return *m.ExpiresIn
	}
	return 0
}

func (m *V1UserTokenCreateData) SetScope(scope string) *V1UserTokenCreateData {
	m.Scope = &scope
	return m
}

func (m *V1UserTokenCreateData) GetScope() string {
	if m != nil {
		return *m.Scope
	}
	return ""
}

func (m *V1UserTokenCreateData) SetStaffId(staffId string) *V1UserTokenCreateData {
	m.StaffId = &staffId
	return m
}

func (m *V1UserTokenCreateData) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

func (m *V1UserTokenCreateData) SetState(state string) *V1UserTokenCreateData {
	m.State = &state
	return m
}

func (m *V1UserTokenCreateData) GetState() string {
	if m != nil {
		return *m.State
	}
	return ""
}

func (m *V1UserTokenCreateData) SetUserToken(userToken string) *V1UserTokenCreateData {
	m.UserToken = &userToken
	return m
}

func (m *V1UserTokenCreateData) GetUserToken() string {
	if m != nil {
		return *m.UserToken
	}
	return ""
}

type V1UserTokenCreateResponse struct {
	Data    *V1UserTokenCreateData `json:"data,omitempty"`
	ErrCode *int                   `json:"errCode,omitempty"`
	ErrMsg  *string                `json:"errMsg,omitempty"`
}

func (m *V1UserTokenCreateResponse) GetData() *V1UserTokenCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1UserTokenCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1UserTokenCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1UserTokenCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V1UsersFetchData struct {
	AvatarId       *string       `json:"avatarId,omitempty"`       // 头像ID
	AvatarUrl      *string       `json:"avatarUrl,omitempty"`      // 头像地址
	Department     []*Department `json:"department,omitempty"`     // 部门信息
	Email          *string       `json:"email,omitempty"`          // 邮箱地址
	EmployeeNumber *string       `json:"employeeNumber,omitempty"` // 员工号
	ExternalId     *string       `json:"externalId,omitempty"`     // 外部数据源人员唯一ID
	LoginName      *string       `json:"loginName,omitempty"`      // 登录用户名
	MobilePhone    *MobilePhone  `json:"mobilePhone,omitempty"`
	Name           *string       `json:"name,omitempty"`    // 人员名称
	OrgId          *string       `json:"orgId,omitempty"`   // 组织ID
	OrgName        *string       `json:"orgName,omitempty"` // 组织名称
	StaffId        *string       `json:"staffId,omitempty"` // 人员ID
}

func (m *V1UsersFetchData) SetAvatarId(avatarId string) *V1UsersFetchData {
	m.AvatarId = &avatarId
	return m
}

func (m *V1UsersFetchData) GetAvatarId() string {
	if m != nil {
		return *m.AvatarId
	}
	return ""
}

func (m *V1UsersFetchData) SetAvatarUrl(avatarUrl string) *V1UsersFetchData {
	m.AvatarUrl = &avatarUrl
	return m
}

func (m *V1UsersFetchData) GetAvatarUrl() string {
	if m != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *V1UsersFetchData) SetDepartment(department []*Department) *V1UsersFetchData {
	m.Department = department
	return m
}

func (m *V1UsersFetchData) GetDepartment() []*Department {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *V1UsersFetchData) SetEmail(email string) *V1UsersFetchData {
	m.Email = &email
	return m
}

func (m *V1UsersFetchData) GetEmail() string {
	if m != nil {
		return *m.Email
	}
	return ""
}

func (m *V1UsersFetchData) SetEmployeeNumber(employeeNumber string) *V1UsersFetchData {
	m.EmployeeNumber = &employeeNumber
	return m
}

func (m *V1UsersFetchData) GetEmployeeNumber() string {
	if m != nil {
		return *m.EmployeeNumber
	}
	return ""
}

func (m *V1UsersFetchData) SetExternalId(externalId string) *V1UsersFetchData {
	m.ExternalId = &externalId
	return m
}

func (m *V1UsersFetchData) GetExternalId() string {
	if m != nil {
		return *m.ExternalId
	}
	return ""
}

func (m *V1UsersFetchData) SetLoginName(loginName string) *V1UsersFetchData {
	m.LoginName = &loginName
	return m
}

func (m *V1UsersFetchData) GetLoginName() string {
	if m != nil {
		return *m.LoginName
	}
	return ""
}

func (m *V1UsersFetchData) SetMobilePhone(mobilePhone MobilePhone) *V1UsersFetchData {
	m.MobilePhone = &mobilePhone
	return m
}

func (m *V1UsersFetchData) GetMobilePhone() *MobilePhone {
	if m != nil {
		return m.MobilePhone
	}
	return nil
}

func (m *V1UsersFetchData) SetName(name string) *V1UsersFetchData {
	m.Name = &name
	return m
}

func (m *V1UsersFetchData) GetName() string {
	if m != nil {
		return *m.Name
	}
	return ""
}

func (m *V1UsersFetchData) SetOrgId(orgId string) *V1UsersFetchData {
	m.OrgId = &orgId
	return m
}

func (m *V1UsersFetchData) GetOrgId() string {
	if m != nil {
		return *m.OrgId
	}
	return ""
}

func (m *V1UsersFetchData) SetOrgName(orgName string) *V1UsersFetchData {
	m.OrgName = &orgName
	return m
}

func (m *V1UsersFetchData) GetOrgName() string {
	if m != nil {
		return *m.OrgName
	}
	return ""
}

func (m *V1UsersFetchData) SetStaffId(staffId string) *V1UsersFetchData {
	m.StaffId = &staffId
	return m
}

func (m *V1UsersFetchData) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

type V1UsersFetchResponse struct {
	Data    *V1UsersFetchData `json:"data,omitempty"`
	ErrCode *int              `json:"errCode,omitempty"`
	ErrMsg  *string           `json:"errMsg,omitempty"`
}

func (m *V1UsersFetchResponse) GetData() *V1UsersFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V1UsersFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V1UsersFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V1UsersFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V2ChatNotificationUpdateData interface{}

type V2ChatNotificationUpdateRequestBody struct {
	StatusList []*StatusInfo `json:"statusList,omitempty"`
}

func (m *V2ChatNotificationUpdateRequestBody) SetStatusList(statusList []*StatusInfo) *V2ChatNotificationUpdateRequestBody {
	m.StatusList = statusList
	return m
}

func (m *V2ChatNotificationUpdateRequestBody) GetStatusList() []*StatusInfo {
	if m != nil {
		return m.StatusList
	}
	return nil
}

type V2ChatNotificationUpdateResponse struct {
	Data    *V2ChatNotificationUpdateData `json:"data,omitempty"`
	ErrCode *int                          `json:"errCode,omitempty"`
	ErrMsg  *string                       `json:"errMsg,omitempty"`
}

func (m *V2ChatNotificationUpdateResponse) GetData() *V2ChatNotificationUpdateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V2ChatNotificationUpdateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V2ChatNotificationUpdateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V2ChatNotificationUpdateResponse) ToString() string {
	return ToJsonString(m)
}

type V2EventNotificationCreateData interface{}

type V2EventNotificationCreateRequestBody struct {
	Events []*Event `json:"events,omitempty"`
}

func (m *V2EventNotificationCreateRequestBody) SetEvents(events []*Event) *V2EventNotificationCreateRequestBody {
	m.Events = events
	return m
}

func (m *V2EventNotificationCreateRequestBody) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type V2EventNotificationCreateResponse struct {
	Data    *V2EventNotificationCreateData `json:"data,omitempty"`
	ErrCode *int                           `json:"errCode,omitempty"`
	ErrMsg  *string                        `json:"errMsg,omitempty"`
}

func (m *V2EventNotificationCreateResponse) GetData() *V2EventNotificationCreateData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V2EventNotificationCreateResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V2EventNotificationCreateResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V2EventNotificationCreateResponse) ToString() string {
	return ToJsonString(m)
}

type V2StaffsIdMappingFetchData struct {
	StaffId *string `json:"staffId,omitempty"`
}

func (m *V2StaffsIdMappingFetchData) SetStaffId(staffId string) *V2StaffsIdMappingFetchData {
	m.StaffId = &staffId
	return m
}

func (m *V2StaffsIdMappingFetchData) GetStaffId() string {
	if m != nil {
		return *m.StaffId
	}
	return ""
}

type V2StaffsIdMappingFetchResponse struct {
	Data    *V2StaffsIdMappingFetchData `json:"data,omitempty"`
	ErrCode *int                        `json:"errCode,omitempty"`
	ErrMsg  *string                     `json:"errMsg,omitempty"`
}

func (m *V2StaffsIdMappingFetchResponse) GetData() *V2StaffsIdMappingFetchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V2StaffsIdMappingFetchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V2StaffsIdMappingFetchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V2StaffsIdMappingFetchResponse) ToString() string {
	return ToJsonString(m)
}

type V2StaffsSearchData struct {
	HasMore   *bool              `json:"hasMore,omitempty"`   // 是否有更多数据
	StaffInfo []*SearchStaffInfo `json:"staffInfo,omitempty"` // 用户详情
	Total     *int               `json:"total,omitempty"`     // 总数据条数
}

func (m *V2StaffsSearchData) SetHasMore(hasMore bool) *V2StaffsSearchData {
	m.HasMore = &hasMore
	return m
}

func (m *V2StaffsSearchData) GetHasMore() bool {
	if m != nil {
		return *m.HasMore
	}
	return false
}

func (m *V2StaffsSearchData) SetStaffInfo(staffInfo []*SearchStaffInfo) *V2StaffsSearchData {
	m.StaffInfo = staffInfo
	return m
}

func (m *V2StaffsSearchData) GetStaffInfo() []*SearchStaffInfo {
	if m != nil {
		return m.StaffInfo
	}
	return nil
}

func (m *V2StaffsSearchData) SetTotal(total int) *V2StaffsSearchData {
	m.Total = &total
	return m
}

func (m *V2StaffsSearchData) GetTotal() int {
	if m != nil {
		return *m.Total
	}
	return 0
}

type V2StaffsSearchRequestBody struct {
	Keyword     *string      `json:"keyword,omitempty"`   // 关键字
	Recursive   *bool        `json:"recursive,omitempty"` // 是否递归搜索
	SearchScope *SearchScope `json:"searchScope,omitempty"`
}

func (m *V2StaffsSearchRequestBody) SetKeyword(keyword string) *V2StaffsSearchRequestBody {
	m.Keyword = &keyword
	return m
}

func (m *V2StaffsSearchRequestBody) GetKeyword() string {
	if m != nil {
		return *m.Keyword
	}
	return ""
}

func (m *V2StaffsSearchRequestBody) SetRecursive(recursive bool) *V2StaffsSearchRequestBody {
	m.Recursive = &recursive
	return m
}

func (m *V2StaffsSearchRequestBody) GetRecursive() bool {
	if m != nil {
		return *m.Recursive
	}
	return false
}

func (m *V2StaffsSearchRequestBody) SetSearchScope(searchScope SearchScope) *V2StaffsSearchRequestBody {
	m.SearchScope = &searchScope
	return m
}

func (m *V2StaffsSearchRequestBody) GetSearchScope() *SearchScope {
	if m != nil {
		return m.SearchScope
	}
	return nil
}

type V2StaffsSearchResponse struct {
	Data    *V2StaffsSearchData `json:"data,omitempty"`
	ErrCode *int                `json:"errCode,omitempty"`
	ErrMsg  *string             `json:"errMsg,omitempty"`
}

func (m *V2StaffsSearchResponse) GetData() *V2StaffsSearchData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *V2StaffsSearchResponse) GetErrCode() int {
	if m != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *V2StaffsSearchResponse) GetErrMsg() string {
	if m != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *V2StaffsSearchResponse) ToString() string {
	return ToJsonString(m)
}

func ToJsonString(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(s)
}

type V1AppRoamingOrgsFetchParams struct {
	AppToken    string  `json:"app_token"`            // app_token
	UserToken   *string `json:"user_token,omitempty"` // user_token
	PageSize    int     `json:"page_size"`            // 分页返回最大数据量，最大限制1000条，大于1000按1000条返回
	BaseVersion string  `json:"base_version"`         // 增量请求的起始时间版本号，时间戳字符串，单位：微秒
}

func (m *V1AppRoamingOrgsFetchParams) SetAppToken(appToken string) *V1AppRoamingOrgsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1AppRoamingOrgsFetchParams) SetUserToken(userToken string) *V1AppRoamingOrgsFetchParams {
	m.UserToken = &userToken
	return m
}

func (m *V1AppRoamingOrgsFetchParams) SetPageSize(pageSize int) *V1AppRoamingOrgsFetchParams {
	m.PageSize = pageSize
	return m
}

func (m *V1AppRoamingOrgsFetchParams) SetBaseVersion(baseVersion string) *V1AppRoamingOrgsFetchParams {
	m.BaseVersion = baseVersion
	return m
}

type V1AppTokenCreateParams struct {
	GrantType string `json:"grant_type"` // client_credential
	Appid     string `json:"appid"`      // 应用ID
	Secret    string `json:"secret"`     // 应用密钥
}

func (m *V1AppTokenCreateParams) SetGrantType(grantType string) *V1AppTokenCreateParams {
	m.GrantType = grantType
	return m
}

func (m *V1AppTokenCreateParams) SetAppid(appid string) *V1AppTokenCreateParams {
	m.Appid = appid
	return m
}

func (m *V1AppTokenCreateParams) SetSecret(secret string) *V1AppTokenCreateParams {
	m.Secret = secret
	return m
}

type V1ChatNotificationFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1ChatNotificationFetchParams) SetAppToken(appToken string) *V1ChatNotificationFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1ChatNotificationFetchParams) SetUserToken(userToken string) *V1ChatNotificationFetchParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsCreateParams) SetAppToken(appToken string) *V1DeptsCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsCreateParams) SetUserToken(userToken string) *V1DeptsCreateParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsChildrenFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsChildrenFetchParams) SetAppToken(appToken string) *V1DeptsChildrenFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsChildrenFetchParams) SetUserToken(userToken string) *V1DeptsChildrenFetchParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsDeleteParams) SetAppToken(appToken string) *V1DeptsDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsDeleteParams) SetUserToken(userToken string) *V1DeptsDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsFetchParams) SetAppToken(appToken string) *V1DeptsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsFetchParams) SetUserToken(userToken string) *V1DeptsFetchParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsStaffsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	Page      *int    `json:"page,omitempty"`       // 起始页码从1开始，默认值为1
	PageSize  *int    `json:"page_size,omitempty"`  // 每页显示个数，默认值是100，最大值是100
}

func (m *V1DeptsStaffsFetchParams) SetAppToken(appToken string) *V1DeptsStaffsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsStaffsFetchParams) SetUserToken(userToken string) *V1DeptsStaffsFetchParams {
	m.UserToken = &userToken
	return m
}

func (m *V1DeptsStaffsFetchParams) SetPage(page int) *V1DeptsStaffsFetchParams {
	m.Page = &page
	return m
}

func (m *V1DeptsStaffsFetchParams) SetPageSize(pageSize int) *V1DeptsStaffsFetchParams {
	m.PageSize = &pageSize
	return m
}

type V1DeptsStaffsCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsStaffsCreateParams) SetAppToken(appToken string) *V1DeptsStaffsCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsStaffsCreateParams) SetUserToken(userToken string) *V1DeptsStaffsCreateParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsStaffsDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsStaffsDeleteParams) SetAppToken(appToken string) *V1DeptsStaffsDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsStaffsDeleteParams) SetUserToken(userToken string) *V1DeptsStaffsDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1DeptsUpdateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1DeptsUpdateParams) SetAppToken(appToken string) *V1DeptsUpdateParams {
	m.AppToken = appToken
	return m
}

func (m *V1DeptsUpdateParams) SetUserToken(userToken string) *V1DeptsUpdateParams {
	m.UserToken = &userToken
	return m
}

type V1JsApiTokenCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1JsApiTokenCreateParams) SetAppToken(appToken string) *V1JsApiTokenCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1JsApiTokenCreateParams) SetUserToken(userToken string) *V1JsApiTokenCreateParams {
	m.UserToken = &userToken
	return m
}

type V1MediasCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	Type      string  `json:"type"`                 // 文件类型。取值：VIDEO = 1;IMAGE = 2;AUDIO = 3;
}

func (m *V1MediasCreateParams) SetAppToken(appToken string) *V1MediasCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1MediasCreateParams) SetUserToken(userToken string) *V1MediasCreateParams {
	m.UserToken = &userToken
	return m
}

func (m *V1MediasCreateParams) SetType(v string) *V1MediasCreateParams {
	m.Type = v
	return m
}

type V1MediasFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1MediasFetchParams) SetAppToken(appToken string) *V1MediasFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1MediasFetchParams) SetUserToken(userToken string) *V1MediasFetchParams {
	m.UserToken = &userToken
	return m
}

type V1MediasPathFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1MediasPathFetchParams) SetAppToken(appToken string) *V1MediasPathFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1MediasPathFetchParams) SetUserToken(userToken string) *V1MediasPathFetchParams {
	m.UserToken = &userToken
	return m
}

type V1MessagesCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1MessagesCreateParams) SetAppToken(appToken string) *V1MessagesCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1MessagesCreateParams) SetUserToken(userToken string) *V1MessagesCreateParams {
	m.UserToken = &userToken
	return m
}

type V1MessagesNotificationCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1MessagesNotificationCreateParams) SetAppToken(appToken string) *V1MessagesNotificationCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1MessagesNotificationCreateParams) SetUserToken(userToken string) *V1MessagesNotificationCreateParams {
	m.UserToken = &userToken
	return m
}

type V1MessagesRevokeParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1MessagesRevokeParams) SetAppToken(appToken string) *V1MessagesRevokeParams {
	m.AppToken = appToken
	return m
}

func (m *V1MessagesRevokeParams) SetUserToken(userToken string) *V1MessagesRevokeParams {
	m.UserToken = &userToken
	return m
}

type V1OrgExtraFieldIdsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	Page      *int    `json:"page,omitempty"`       // 起始页码从1开始,默认值为1
	PageSize  *int    `json:"page_size,omitempty"`  // 每页显示个数，默认值是1000，最大值是100000
}

func (m *V1OrgExtraFieldIdsFetchParams) SetAppToken(appToken string) *V1OrgExtraFieldIdsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1OrgExtraFieldIdsFetchParams) SetUserToken(userToken string) *V1OrgExtraFieldIdsFetchParams {
	m.UserToken = &userToken
	return m
}

func (m *V1OrgExtraFieldIdsFetchParams) SetPage(page int) *V1OrgExtraFieldIdsFetchParams {
	m.Page = &page
	return m
}

func (m *V1OrgExtraFieldIdsFetchParams) SetPageSize(pageSize int) *V1OrgExtraFieldIdsFetchParams {
	m.PageSize = &pageSize
	return m
}

type V1OrgFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1OrgFetchParams) SetAppToken(appToken string) *V1OrgFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1OrgFetchParams) SetUserToken(userToken string) *V1OrgFetchParams {
	m.UserToken = &userToken
	return m
}

type V1OrgRoleCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1OrgRoleCreateParams) SetAppToken(appToken string) *V1OrgRoleCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1OrgRoleCreateParams) SetUserToken(userToken string) *V1OrgRoleCreateParams {
	m.UserToken = &userToken
	return m
}

type V1OrgRoleMembersFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1OrgRoleMembersFetchParams) SetAppToken(appToken string) *V1OrgRoleMembersFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1OrgRoleMembersFetchParams) SetUserToken(userToken string) *V1OrgRoleMembersFetchParams {
	m.UserToken = &userToken
	return m
}

type V1RoleMemberCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1RoleMemberCreateParams) SetAppToken(appToken string) *V1RoleMemberCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1RoleMemberCreateParams) SetUserToken(userToken string) *V1RoleMemberCreateParams {
	m.UserToken = &userToken
	return m
}

type V1RoleMemberDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1RoleMemberDeleteParams) SetAppToken(appToken string) *V1RoleMemberDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1RoleMemberDeleteParams) SetUserToken(userToken string) *V1RoleMemberDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsCreateParams) SetAppToken(appToken string) *V1StaffsCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsCreateParams) SetUserToken(userToken string) *V1StaffsCreateParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsDeleteParams) SetAppToken(appToken string) *V1StaffsDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsDeleteParams) SetUserToken(userToken string) *V1StaffsDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsDeptAncestorsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsDeptAncestorsFetchParams) SetAppToken(appToken string) *V1StaffsDeptAncestorsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsDeptAncestorsFetchParams) SetUserToken(userToken string) *V1StaffsDeptAncestorsFetchParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsFetchParams) SetAppToken(appToken string) *V1StaffsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsFetchParams) SetUserToken(userToken string) *V1StaffsFetchParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsInforFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsInforFetchParams) SetAppToken(appToken string) *V1StaffsInforFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsInforFetchParams) SetUserToken(userToken string) *V1StaffsInforFetchParams {
	m.UserToken = &userToken
	return m
}

type V1StaffsUpdateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1StaffsUpdateParams) SetAppToken(appToken string) *V1StaffsUpdateParams {
	m.AppToken = appToken
	return m
}

func (m *V1StaffsUpdateParams) SetUserToken(userToken string) *V1StaffsUpdateParams {
	m.UserToken = &userToken
	return m
}

type V1TagGroupsCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagGroupsCreateParams) SetAppToken(appToken string) *V1TagGroupsCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagGroupsCreateParams) SetUserToken(userToken string) *V1TagGroupsCreateParams {
	m.UserToken = &userToken
	return m
}

type V1TagGroupsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagGroupsFetchParams) SetAppToken(appToken string) *V1TagGroupsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagGroupsFetchParams) SetUserToken(userToken string) *V1TagGroupsFetchParams {
	m.UserToken = &userToken
	return m
}

type V1TagGroupsDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagGroupsDeleteParams) SetAppToken(appToken string) *V1TagGroupsDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagGroupsDeleteParams) SetUserToken(userToken string) *V1TagGroupsDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1TagGroupsInfoFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagGroupsInfoFetchParams) SetAppToken(appToken string) *V1TagGroupsInfoFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagGroupsInfoFetchParams) SetUserToken(userToken string) *V1TagGroupsInfoFetchParams {
	m.UserToken = &userToken
	return m
}

type V1TagGroupsUpdateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagGroupsUpdateParams) SetAppToken(appToken string) *V1TagGroupsUpdateParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagGroupsUpdateParams) SetUserToken(userToken string) *V1TagGroupsUpdateParams {
	m.UserToken = &userToken
	return m
}

type V1TagsCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagsCreateParams) SetAppToken(appToken string) *V1TagsCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagsCreateParams) SetUserToken(userToken string) *V1TagsCreateParams {
	m.UserToken = &userToken
	return m
}

type V1TagsMetaFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagsMetaFetchParams) SetAppToken(appToken string) *V1TagsMetaFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagsMetaFetchParams) SetUserToken(userToken string) *V1TagsMetaFetchParams {
	m.UserToken = &userToken
	return m
}

type V1TagsFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	Page      *int    `json:"page,omitempty"`       // 起始页码从1开始，默认值为1
	PageSize  *int    `json:"page_size,omitempty"`  // 每页显示个数，默认值是1000，最大值是100000
}

func (m *V1TagsFetchParams) SetAppToken(appToken string) *V1TagsFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagsFetchParams) SetUserToken(userToken string) *V1TagsFetchParams {
	m.UserToken = &userToken
	return m
}

func (m *V1TagsFetchParams) SetPage(page int) *V1TagsFetchParams {
	m.Page = &page
	return m
}

func (m *V1TagsFetchParams) SetPageSize(pageSize int) *V1TagsFetchParams {
	m.PageSize = &pageSize
	return m
}

type V1TagsDeleteParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagsDeleteParams) SetAppToken(appToken string) *V1TagsDeleteParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagsDeleteParams) SetUserToken(userToken string) *V1TagsDeleteParams {
	m.UserToken = &userToken
	return m
}

type V1TagsUpdateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V1TagsUpdateParams) SetAppToken(appToken string) *V1TagsUpdateParams {
	m.AppToken = appToken
	return m
}

func (m *V1TagsUpdateParams) SetUserToken(userToken string) *V1TagsUpdateParams {
	m.UserToken = &userToken
	return m
}

type V1UsersFetchParams struct {
	AppToken  string `json:"app_token"`  // app_token
	UserToken string `json:"user_token"` // user_token
}

func (m *V1UsersFetchParams) SetAppToken(appToken string) *V1UsersFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V1UsersFetchParams) SetUserToken(userToken string) *V1UsersFetchParams {
	m.UserToken = userToken
	return m
}

type V1UserTokenCreateParams struct {
	AppToken    string  `json:"app_token"`              // app_token
	GrantType   string  `json:"grant_type"`             // 使用固定值 'authorization_code'
	Code        string  `json:"code"`                   // 人员免登录授权码
	RedirectUri *string `json:"redirect_uri,omitempty"` // redirect_uri
}

func (m *V1UserTokenCreateParams) SetAppToken(appToken string) *V1UserTokenCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V1UserTokenCreateParams) SetGrantType(grantType string) *V1UserTokenCreateParams {
	m.GrantType = grantType
	return m
}

func (m *V1UserTokenCreateParams) SetCode(code string) *V1UserTokenCreateParams {
	m.Code = code
	return m
}

func (m *V1UserTokenCreateParams) SetRedirectUri(redirectUri string) *V1UserTokenCreateParams {
	m.RedirectUri = &redirectUri
	return m
}

type V2ChatNotificationUpdateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V2ChatNotificationUpdateParams) SetAppToken(appToken string) *V2ChatNotificationUpdateParams {
	m.AppToken = appToken
	return m
}

func (m *V2ChatNotificationUpdateParams) SetUserToken(userToken string) *V2ChatNotificationUpdateParams {
	m.UserToken = &userToken
	return m
}

type V2EventNotificationCreateParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
}

func (m *V2EventNotificationCreateParams) SetAppToken(appToken string) *V2EventNotificationCreateParams {
	m.AppToken = appToken
	return m
}

func (m *V2EventNotificationCreateParams) SetUserToken(userToken string) *V2EventNotificationCreateParams {
	m.UserToken = &userToken
	return m
}

type V2StaffsIdMappingFetchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	OrgId     string  `json:"org_id"`               // 查询人员所在的组织Id
	IdType    string  `json:"id_type"`              // employ_id/mobile/mail/login/external_id
	IdValue   string  `json:"id_value"`             // id_type 对应的值：人员编号，手机号...
}

func (m *V2StaffsIdMappingFetchParams) SetAppToken(appToken string) *V2StaffsIdMappingFetchParams {
	m.AppToken = appToken
	return m
}

func (m *V2StaffsIdMappingFetchParams) SetUserToken(userToken string) *V2StaffsIdMappingFetchParams {
	m.UserToken = &userToken
	return m
}

func (m *V2StaffsIdMappingFetchParams) SetOrgId(orgId string) *V2StaffsIdMappingFetchParams {
	m.OrgId = orgId
	return m
}

func (m *V2StaffsIdMappingFetchParams) SetIdType(idType string) *V2StaffsIdMappingFetchParams {
	m.IdType = idType
	return m
}

func (m *V2StaffsIdMappingFetchParams) SetIdValue(idValue string) *V2StaffsIdMappingFetchParams {
	m.IdValue = idValue
	return m
}

type V2StaffsSearchParams struct {
	AppToken  string  `json:"app_token"`            // app_token
	UserToken *string `json:"user_token,omitempty"` // user_token
	UserId    string  `json:"user_id"`              // user_id
}

func (m *V2StaffsSearchParams) SetAppToken(appToken string) *V2StaffsSearchParams {
	m.AppToken = appToken
	return m
}

func (m *V2StaffsSearchParams) SetUserToken(userToken string) *V2StaffsSearchParams {
	m.UserToken = &userToken
	return m
}

func (m *V2StaffsSearchParams) SetUserId(userId string) *V2StaffsSearchParams {
	m.UserId = userId
	return m
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {

	// /v1/app/roaming/orgs/fetch
	V1AppRoamingOrgsFetch(ctx context.Context, params *V1AppRoamingOrgsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/apptoken/create
	V1AppTokenCreate(ctx context.Context, params *V1AppTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/chat/notification/{userid}/fetch
	V1ChatNotificationFetch(ctx context.Context, userid string, params *V1ChatNotificationFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/create
	V1DeptsCreateWithBody(ctx context.Context, params *V1DeptsCreateParams, body V1DeptsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/children/fetch
	V1DeptsChildrenFetch(ctx context.Context, departmentid string, params *V1DeptsChildrenFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/delete
	V1DeptsDelete(ctx context.Context, departmentid string, params *V1DeptsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/fetch
	V1DeptsFetch(ctx context.Context, departmentid string, params *V1DeptsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/staffs/fetch
	V1DeptsStaffsFetch(ctx context.Context, departmentid string, params *V1DeptsStaffsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/staffs/{staffid}/create
	V1DeptsStaffsCreate(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/staffs/{staffid}/delete
	V1DeptsStaffsDelete(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/departments/{departmentid}/update
	V1DeptsUpdateWithBody(ctx context.Context, departmentid string, params *V1DeptsUpdateParams, body V1DeptsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/jsapitoken/create
	V1JsApiTokenCreate(ctx context.Context, params *V1JsApiTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/medias/create
	V1MediasCreateWithBody(ctx context.Context, params *V1MediasCreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/medias/{mediaid}/fetch
	V1MediasFetch(ctx context.Context, mediaid string, params *V1MediasFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/medias/{mediaid}/path/fetch
	V1MediasPathFetch(ctx context.Context, mediaid string, params *V1MediasPathFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/messages/create
	V1MessagesCreateWithBody(ctx context.Context, params *V1MessagesCreateParams, body V1MessagesCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/messages/notification/create
	V1MessagesNotificationCreateWithBody(ctx context.Context, params *V1MessagesNotificationCreateParams, body V1MessagesNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/messages/revoke
	V1MessagesRevokeWithBody(ctx context.Context, params *V1MessagesRevokeParams, body V1MessagesRevokeRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/org/{orgid}/extrafieldids/fetch
	V1OrgExtraFieldIdsFetch(ctx context.Context, orgid string, params *V1OrgExtraFieldIdsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/org/{orgid}/fetch
	V1OrgFetch(ctx context.Context, orgid string, params *V1OrgFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/org/{orgid}/role/create
	V1OrgRoleCreateWithBody(ctx context.Context, orgid string, params *V1OrgRoleCreateParams, body V1OrgRoleCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/org/{orgid}/role/{roleid}/members/fetch
	V1OrgRoleMembersFetch(ctx context.Context, orgid string, roleid string, params *V1OrgRoleMembersFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/rolemember/create
	V1RoleMemberCreateWithBody(ctx context.Context, params *V1RoleMemberCreateParams, body V1RoleMemberCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/rolemember/{rolememberid}/delete
	V1RoleMemberDeleteWithBody(ctx context.Context, rolememberid string, params *V1RoleMemberDeleteParams, body V1RoleMemberDeleteRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/create
	V1StaffsCreateWithBody(ctx context.Context, params *V1StaffsCreateParams, body V1StaffsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/{staffid}/delete
	V1StaffsDelete(ctx context.Context, staffid string, params *V1StaffsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/{staffid}/departmentancestors/fetch
	V1StaffsDeptAncestorsFetch(ctx context.Context, staffid string, params *V1StaffsDeptAncestorsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/{staffid}/fetch
	V1StaffsFetch(ctx context.Context, staffid string, params *V1StaffsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/{staffid}/infor/fetch
	V1StaffsInforFetch(ctx context.Context, staffid string, params *V1StaffsInforFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/staffs/{staffid}/update
	V1StaffsUpdateWithBody(ctx context.Context, staffid string, params *V1StaffsUpdateParams, body V1StaffsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/taggroups/create
	V1TagGroupsCreateWithBody(ctx context.Context, params *V1TagGroupsCreateParams, body V1TagGroupsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/taggroups/fetch
	V1TagGroupsFetchWithBody(ctx context.Context, params *V1TagGroupsFetchParams, body V1TagGroupsFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/taggroups/{tag_group_id}/delete
	V1TagGroupsDelete(ctx context.Context, tagGroupId string, params *V1TagGroupsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/taggroups/{tag_group_id}/fetch
	V1TagGroupsInfoFetch(ctx context.Context, tagGroupId string, params *V1TagGroupsInfoFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/taggroups/{tag_group_id}/update
	V1TagGroupsUpdateWithBody(ctx context.Context, tagGroupId string, params *V1TagGroupsUpdateParams, body V1TagGroupsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/tags/create
	V1TagsCreateWithBody(ctx context.Context, params *V1TagsCreateParams, body V1TagsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/tags/meta/fetch
	V1TagsMetaFetchWithBody(ctx context.Context, params *V1TagsMetaFetchParams, body V1TagsMetaFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/tags/staffids/fetch
	V1TagsFetchWithBody(ctx context.Context, params *V1TagsFetchParams, body V1TagsFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/tags/{tagid}/delete
	V1TagsDelete(ctx context.Context, tagid string, params *V1TagsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/tags/{tagid}/update
	V1TagsUpdateWithBody(ctx context.Context, tagid string, params *V1TagsUpdateParams, body V1TagsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/users/fetch
	V1UsersFetch(ctx context.Context, params *V1UsersFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/usertoken/create
	V1UserTokenCreate(ctx context.Context, params *V1UserTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v2/chat/notification/update
	V2ChatNotificationUpdateWithBody(ctx context.Context, params *V2ChatNotificationUpdateParams, body V2ChatNotificationUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v2/event/notification/create
	V2EventNotificationCreateWithBody(ctx context.Context, params *V2EventNotificationCreateParams, body V2EventNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v2/staffs/id_mapping/fetch
	V2StaffsIdMappingFetch(ctx context.Context, params *V2StaffsIdMappingFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v2/staffs/search
	V2StaffsSearchWithBody(ctx context.Context, params *V2StaffsSearchParams, body V2StaffsSearchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) V1AppRoamingOrgsFetch(ctx context.Context, params *V1AppRoamingOrgsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1AppRoamingOrgsFetchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1AppTokenCreate(ctx context.Context, params *V1AppTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1AppTokenCreateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1ChatNotificationFetch(ctx context.Context, userid string, params *V1ChatNotificationFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1ChatNotificationFetchRequest(c.Server, userid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsCreateWithBody(ctx context.Context, params *V1DeptsCreateParams, body V1DeptsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsChildrenFetch(ctx context.Context, departmentid string, params *V1DeptsChildrenFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsChildrenFetchRequest(c.Server, departmentid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsDelete(ctx context.Context, departmentid string, params *V1DeptsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsDeleteRequest(c.Server, departmentid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsFetch(ctx context.Context, departmentid string, params *V1DeptsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsFetchRequest(c.Server, departmentid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsStaffsFetch(ctx context.Context, departmentid string, params *V1DeptsStaffsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsStaffsFetchRequest(c.Server, departmentid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsStaffsCreate(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsStaffsCreateRequest(c.Server, departmentid, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsStaffsDelete(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsStaffsDeleteRequest(c.Server, departmentid, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1DeptsUpdateWithBody(ctx context.Context, departmentid string, params *V1DeptsUpdateParams, body V1DeptsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1DeptsUpdateRequest(c.Server, departmentid, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1JsApiTokenCreate(ctx context.Context, params *V1JsApiTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1JsApiTokenCreateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MediasCreateWithBody(ctx context.Context, params *V1MediasCreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MediasCreateRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MediasFetch(ctx context.Context, mediaid string, params *V1MediasFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MediasFetchRequest(c.Server, mediaid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MediasPathFetch(ctx context.Context, mediaid string, params *V1MediasPathFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MediasPathFetchRequest(c.Server, mediaid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MessagesCreateWithBody(ctx context.Context, params *V1MessagesCreateParams, body V1MessagesCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MessagesCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MessagesNotificationCreateWithBody(ctx context.Context, params *V1MessagesNotificationCreateParams, body V1MessagesNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MessagesNotificationCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1MessagesRevokeWithBody(ctx context.Context, params *V1MessagesRevokeParams, body V1MessagesRevokeRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1MessagesRevokeRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1OrgExtraFieldIdsFetch(ctx context.Context, orgid string, params *V1OrgExtraFieldIdsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1OrgExtraFieldIdsFetchRequest(c.Server, orgid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1OrgFetch(ctx context.Context, orgid string, params *V1OrgFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1OrgFetchRequest(c.Server, orgid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1OrgRoleCreateWithBody(ctx context.Context, orgid string, params *V1OrgRoleCreateParams, body V1OrgRoleCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1OrgRoleCreateRequest(c.Server, orgid, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1OrgRoleMembersFetch(ctx context.Context, orgid string, roleid string, params *V1OrgRoleMembersFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1OrgRoleMembersFetchRequest(c.Server, orgid, roleid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1RoleMemberCreateWithBody(ctx context.Context, params *V1RoleMemberCreateParams, body V1RoleMemberCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1RoleMemberCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1RoleMemberDeleteWithBody(ctx context.Context, rolememberid string, params *V1RoleMemberDeleteParams, body V1RoleMemberDeleteRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1RoleMemberDeleteRequest(c.Server, rolememberid, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsCreateWithBody(ctx context.Context, params *V1StaffsCreateParams, body V1StaffsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsDelete(ctx context.Context, staffid string, params *V1StaffsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsDeleteRequest(c.Server, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsDeptAncestorsFetch(ctx context.Context, staffid string, params *V1StaffsDeptAncestorsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsDeptAncestorsFetchRequest(c.Server, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsFetch(ctx context.Context, staffid string, params *V1StaffsFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsFetchRequest(c.Server, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsInforFetch(ctx context.Context, staffid string, params *V1StaffsInforFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsInforFetchRequest(c.Server, staffid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1StaffsUpdateWithBody(ctx context.Context, staffid string, params *V1StaffsUpdateParams, body V1StaffsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1StaffsUpdateRequest(c.Server, staffid, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagGroupsCreateWithBody(ctx context.Context, params *V1TagGroupsCreateParams, body V1TagGroupsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagGroupsCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagGroupsFetchWithBody(ctx context.Context, params *V1TagGroupsFetchParams, body V1TagGroupsFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagGroupsFetchRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagGroupsDelete(ctx context.Context, tagGroupId string, params *V1TagGroupsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagGroupsDeleteRequest(c.Server, tagGroupId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagGroupsInfoFetch(ctx context.Context, tagGroupId string, params *V1TagGroupsInfoFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagGroupsInfoFetchRequest(c.Server, tagGroupId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagGroupsUpdateWithBody(ctx context.Context, tagGroupId string, params *V1TagGroupsUpdateParams, body V1TagGroupsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagGroupsUpdateRequest(c.Server, tagGroupId, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagsCreateWithBody(ctx context.Context, params *V1TagsCreateParams, body V1TagsCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagsCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagsMetaFetchWithBody(ctx context.Context, params *V1TagsMetaFetchParams, body V1TagsMetaFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagsMetaFetchRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagsFetchWithBody(ctx context.Context, params *V1TagsFetchParams, body V1TagsFetchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagsFetchRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagsDelete(ctx context.Context, tagid string, params *V1TagsDeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagsDeleteRequest(c.Server, tagid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1TagsUpdateWithBody(ctx context.Context, tagid string, params *V1TagsUpdateParams, body V1TagsUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1TagsUpdateRequest(c.Server, tagid, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1UsersFetch(ctx context.Context, params *V1UsersFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1UsersFetchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1UserTokenCreate(ctx context.Context, params *V1UserTokenCreateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1UserTokenCreateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V2ChatNotificationUpdateWithBody(ctx context.Context, params *V2ChatNotificationUpdateParams, body V2ChatNotificationUpdateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV2ChatNotificationUpdateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V2EventNotificationCreateWithBody(ctx context.Context, params *V2EventNotificationCreateParams, body V2EventNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV2EventNotificationCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V2StaffsIdMappingFetch(ctx context.Context, params *V2StaffsIdMappingFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV2StaffsIdMappingFetchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V2StaffsSearchWithBody(ctx context.Context, params *V2StaffsSearchParams, body V2StaffsSearchRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV2StaffsSearchRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func NewV1AppRoamingOrgsFetchRequest(server string, params *V1AppRoamingOrgsFetchParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/app/roaming/orgs/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page_size", runtime.ParamLocationQuery, params.PageSize); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "base_version", runtime.ParamLocationQuery, params.BaseVersion); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1AppTokenCreateRequest(server string, params *V1AppTokenCreateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/apptoken/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "grant_type", runtime.ParamLocationQuery, params.GrantType); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "appid", runtime.ParamLocationQuery, params.Appid); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "secret", runtime.ParamLocationQuery, params.Secret); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1ChatNotificationFetchRequest(server string, userid string, params *V1ChatNotificationFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userid", runtime.ParamLocationPath, userid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/chat/notification/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsCreateRequest(server string, params *V1DeptsCreateParams, body V1DeptsCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1DeptsCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1DeptsCreateRequestWithBody(server string, params *V1DeptsCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1DeptsChildrenFetchRequest(server string, departmentid string, params *V1DeptsChildrenFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/children/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsDeleteRequest(server string, departmentid string, params *V1DeptsDeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsFetchRequest(server string, departmentid string, params *V1DeptsFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsStaffsFetchRequest(server string, departmentid string, params *V1DeptsStaffsFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/staffs/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page_size", runtime.ParamLocationQuery, *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsStaffsCreateRequest(server string, departmentid string, staffid string, params *V1DeptsStaffsCreateParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/staffs/%s/create", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsStaffsDeleteRequest(server string, departmentid string, staffid string, params *V1DeptsStaffsDeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/staffs/%s/delete", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1DeptsUpdateRequest(server string, departmentid string, params *V1DeptsUpdateParams, body V1DeptsUpdateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1DeptsUpdateRequestWithBody(server, departmentid, params, "application/json", bodyReader)
}

func NewV1DeptsUpdateRequestWithBody(server string, departmentid string, params *V1DeptsUpdateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "departmentid", runtime.ParamLocationPath, departmentid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/departments/%s/update", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1JsApiTokenCreateRequest(server string, params *V1JsApiTokenCreateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/jsapitoken/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1MediasCreateRequestWithBody(server string, params *V1MediasCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/medias/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, params.Type); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1MediasFetchRequest(server string, mediaid string, params *V1MediasFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "mediaid", runtime.ParamLocationPath, mediaid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/medias/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1MediasPathFetchRequest(server string, mediaid string, params *V1MediasPathFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "mediaid", runtime.ParamLocationPath, mediaid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/medias/%s/path/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1MessagesCreateRequest(server string, params *V1MessagesCreateParams, body V1MessagesCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1MessagesCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1MessagesCreateRequestWithBody(server string, params *V1MessagesCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/messages/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1MessagesNotificationCreateRequest(server string, params *V1MessagesNotificationCreateParams, body V1MessagesNotificationCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1MessagesNotificationCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1MessagesNotificationCreateRequestWithBody(server string, params *V1MessagesNotificationCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/messages/notification/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1MessagesRevokeRequest(server string, params *V1MessagesRevokeParams, body V1MessagesRevokeRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1MessagesRevokeRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1MessagesRevokeRequestWithBody(server string, params *V1MessagesRevokeParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/messages/revoke")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1OrgExtraFieldIdsFetchRequest(server string, orgid string, params *V1OrgExtraFieldIdsFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orgid", runtime.ParamLocationPath, orgid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/org/%s/extrafieldids/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page_size", runtime.ParamLocationQuery, *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1OrgFetchRequest(server string, orgid string, params *V1OrgFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orgid", runtime.ParamLocationPath, orgid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/org/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1OrgRoleCreateRequest(server string, orgid string, params *V1OrgRoleCreateParams, body V1OrgRoleCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1OrgRoleCreateRequestWithBody(server, orgid, params, "application/json", bodyReader)
}

func NewV1OrgRoleCreateRequestWithBody(server string, orgid string, params *V1OrgRoleCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orgid", runtime.ParamLocationPath, orgid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/org/%s/role/create", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1OrgRoleMembersFetchRequest(server string, orgid string, roleid string, params *V1OrgRoleMembersFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orgid", runtime.ParamLocationPath, orgid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "roleid", runtime.ParamLocationPath, roleid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/org/%s/role/%s/members/fetch", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1RoleMemberCreateRequest(server string, params *V1RoleMemberCreateParams, body V1RoleMemberCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1RoleMemberCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1RoleMemberCreateRequestWithBody(server string, params *V1RoleMemberCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/rolemember/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1RoleMemberDeleteRequest(server string, rolememberid string, params *V1RoleMemberDeleteParams, body V1RoleMemberDeleteRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1RoleMemberDeleteRequestWithBody(server, rolememberid, params, "application/json", bodyReader)
}

func NewV1RoleMemberDeleteRequestWithBody(server string, rolememberid string, params *V1RoleMemberDeleteParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "rolememberid", runtime.ParamLocationPath, rolememberid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/rolemember/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1StaffsCreateRequest(server string, params *V1StaffsCreateParams, body V1StaffsCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1StaffsCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1StaffsCreateRequestWithBody(server string, params *V1StaffsCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1StaffsDeleteRequest(server string, staffid string, params *V1StaffsDeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1StaffsDeptAncestorsFetchRequest(server string, staffid string, params *V1StaffsDeptAncestorsFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/%s/departmentancestors/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1StaffsFetchRequest(server string, staffid string, params *V1StaffsFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1StaffsInforFetchRequest(server string, staffid string, params *V1StaffsInforFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/%s/infor/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1StaffsUpdateRequest(server string, staffid string, params *V1StaffsUpdateParams, body V1StaffsUpdateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1StaffsUpdateRequestWithBody(server, staffid, params, "application/json", bodyReader)
}

func NewV1StaffsUpdateRequestWithBody(server string, staffid string, params *V1StaffsUpdateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "staffid", runtime.ParamLocationPath, staffid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/staffs/%s/update", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagGroupsCreateRequest(server string, params *V1TagGroupsCreateParams, body V1TagGroupsCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagGroupsCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1TagGroupsCreateRequestWithBody(server string, params *V1TagGroupsCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/taggroups/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagGroupsFetchRequest(server string, params *V1TagGroupsFetchParams, body V1TagGroupsFetchRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagGroupsFetchRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1TagGroupsFetchRequestWithBody(server string, params *V1TagGroupsFetchParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/taggroups/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagGroupsDeleteRequest(server string, tagGroupId string, params *V1TagGroupsDeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tag_group_id", runtime.ParamLocationPath, tagGroupId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/taggroups/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1TagGroupsInfoFetchRequest(server string, tagGroupId string, params *V1TagGroupsInfoFetchParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tag_group_id", runtime.ParamLocationPath, tagGroupId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/taggroups/%s/fetch", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1TagGroupsUpdateRequest(server string, tagGroupId string, params *V1TagGroupsUpdateParams, body V1TagGroupsUpdateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagGroupsUpdateRequestWithBody(server, tagGroupId, params, "application/json", bodyReader)
}

func NewV1TagGroupsUpdateRequestWithBody(server string, tagGroupId string, params *V1TagGroupsUpdateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tag_group_id", runtime.ParamLocationPath, tagGroupId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/taggroups/%s/update", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagsCreateRequest(server string, params *V1TagsCreateParams, body V1TagsCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagsCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1TagsCreateRequestWithBody(server string, params *V1TagsCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/tags/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagsMetaFetchRequest(server string, params *V1TagsMetaFetchParams, body V1TagsMetaFetchRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagsMetaFetchRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1TagsMetaFetchRequestWithBody(server string, params *V1TagsMetaFetchParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/tags/meta/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagsFetchRequest(server string, params *V1TagsFetchParams, body V1TagsFetchRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagsFetchRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV1TagsFetchRequestWithBody(server string, params *V1TagsFetchParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/tags/staffids/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page_size", runtime.ParamLocationQuery, *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1TagsDeleteRequest(server string, tagid string, params *V1TagsDeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tagid", runtime.ParamLocationPath, tagid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/tags/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1TagsUpdateRequest(server string, tagid string, params *V1TagsUpdateParams, body V1TagsUpdateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1TagsUpdateRequestWithBody(server, tagid, params, "application/json", bodyReader)
}

func NewV1TagsUpdateRequestWithBody(server string, tagid string, params *V1TagsUpdateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tagid", runtime.ParamLocationPath, tagid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/tags/%s/update", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV1UsersFetchRequest(server string, params *V1UsersFetchParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/users/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, params.UserToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV1UserTokenCreateRequest(server string, params *V1UserTokenCreateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/usertoken/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "grant_type", runtime.ParamLocationQuery, params.GrantType); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "code", runtime.ParamLocationQuery, params.Code); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.RedirectUri != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "redirect_uri", runtime.ParamLocationQuery, *params.RedirectUri); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV2ChatNotificationUpdateRequest(server string, params *V2ChatNotificationUpdateParams, body V2ChatNotificationUpdateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV2ChatNotificationUpdateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV2ChatNotificationUpdateRequestWithBody(server string, params *V2ChatNotificationUpdateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/chat/notification/update")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV2EventNotificationCreateRequest(server string, params *V2EventNotificationCreateParams, body V2EventNotificationCreateRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV2EventNotificationCreateRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV2EventNotificationCreateRequestWithBody(server string, params *V2EventNotificationCreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/event/notification/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func NewV2StaffsIdMappingFetchRequest(server string, params *V2StaffsIdMappingFetchParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/staffs/id_mapping/fetch")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "org_id", runtime.ParamLocationQuery, params.OrgId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_type", runtime.ParamLocationQuery, params.IdType); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_value", runtime.ParamLocationQuery, params.IdValue); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewV2StaffsSearchRequest(server string, params *V2StaffsSearchParams, body V2StaffsSearchRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV2StaffsSearchRequestWithBody(server, params, "application/json", bodyReader)
}

func NewV2StaffsSearchRequestWithBody(server string, params *V2StaffsSearchParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/staffs/search")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_token", runtime.ParamLocationQuery, params.AppToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.UserToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_token", runtime.ParamLocationQuery, *params.UserToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_id", runtime.ParamLocationQuery, params.UserId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {

	// /v1/app/roaming/orgs/fetch
	V1AppRoamingOrgsFetchWithResponse(ctx context.Context, params *V1AppRoamingOrgsFetchParams, reqEditors ...RequestEditorFn) (*V1AppRoamingOrgsFetchResponse, error)

	// /v1/apptoken/create
	V1AppTokenCreateWithResponse(ctx context.Context, params *V1AppTokenCreateParams, reqEditors ...RequestEditorFn) (*V1AppTokenCreateResponse, error)

	// /v1/chat/notification/{userid}/fetch
	V1ChatNotificationFetchWithResponse(ctx context.Context, userid string, params *V1ChatNotificationFetchParams, reqEditors ...RequestEditorFn) (*V1ChatNotificationFetchResponse, error)

	// /v1/departments/create
	V1DeptsCreateWithBodyWithResponse(ctx context.Context, params *V1DeptsCreateParams, body V1DeptsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1DeptsCreateResponse, error)

	// /v1/departments/{departmentid}/children/fetch
	V1DeptsChildrenFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsChildrenFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsChildrenFetchResponse, error)

	// /v1/departments/{departmentid}/delete
	V1DeptsDeleteWithResponse(ctx context.Context, departmentid string, params *V1DeptsDeleteParams, reqEditors ...RequestEditorFn) (*V1DeptsDeleteResponse, error)

	// /v1/departments/{departmentid}/fetch
	V1DeptsFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsFetchResponse, error)

	// /v1/departments/{departmentid}/staffs/fetch
	V1DeptsStaffsFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsStaffsFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsFetchResponse, error)

	// /v1/departments/{departmentid}/staffs/{staffid}/create
	V1DeptsStaffsCreateWithResponse(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsCreateParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsCreateResponse, error)

	// /v1/departments/{departmentid}/staffs/{staffid}/delete
	V1DeptsStaffsDeleteWithResponse(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsDeleteParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsDeleteResponse, error)

	// /v1/departments/{departmentid}/update
	V1DeptsUpdateWithBodyWithResponse(ctx context.Context, departmentid string, params *V1DeptsUpdateParams, body V1DeptsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1DeptsUpdateResponse, error)

	// /v1/jsapitoken/create
	V1JsApiTokenCreateWithResponse(ctx context.Context, params *V1JsApiTokenCreateParams, reqEditors ...RequestEditorFn) (*V1JsApiTokenCreateResponse, error)

	// /v1/medias/create
	V1MediasCreateWithBodyWithResponse(ctx context.Context, params *V1MediasCreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*V1MediasCreateResponse, error)

	// /v1/medias/{mediaid}/fetch
	V1MediasFetchWithResponse(ctx context.Context, mediaid string, params *V1MediasFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// /v1/medias/{mediaid}/path/fetch
	V1MediasPathFetchWithResponse(ctx context.Context, mediaid string, params *V1MediasPathFetchParams, reqEditors ...RequestEditorFn) (*V1MediasPathFetchResponse, error)

	// /v1/messages/create
	V1MessagesCreateWithBodyWithResponse(ctx context.Context, params *V1MessagesCreateParams, body V1MessagesCreateRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesCreateResponse, error)

	// /v1/messages/notification/create
	V1MessagesNotificationCreateWithBodyWithResponse(ctx context.Context, params *V1MessagesNotificationCreateParams, body V1MessagesNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesNotificationCreateResponse, error)

	// /v1/messages/revoke
	V1MessagesRevokeWithBodyWithResponse(ctx context.Context, params *V1MessagesRevokeParams, body V1MessagesRevokeRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesRevokeResponse, error)

	// /v1/org/{orgid}/extrafieldids/fetch
	V1OrgExtraFieldIdsFetchWithResponse(ctx context.Context, orgid string, params *V1OrgExtraFieldIdsFetchParams, reqEditors ...RequestEditorFn) (*V1OrgExtraFieldIdsFetchResponse, error)

	// /v1/org/{orgid}/fetch
	V1OrgFetchWithResponse(ctx context.Context, orgid string, params *V1OrgFetchParams, reqEditors ...RequestEditorFn) (*V1OrgFetchResponse, error)

	// /v1/org/{orgid}/role/create
	V1OrgRoleCreateWithBodyWithResponse(ctx context.Context, orgid string, params *V1OrgRoleCreateParams, body V1OrgRoleCreateRequestBody, reqEditors ...RequestEditorFn) (*V1OrgRoleCreateResponse, error)

	// /v1/org/{orgid}/role/{roleid}/members/fetch
	V1OrgRoleMembersFetchWithResponse(ctx context.Context, orgid string, roleid string, params *V1OrgRoleMembersFetchParams, reqEditors ...RequestEditorFn) (*V1OrgRoleMembersFetchResponse, error)

	// /v1/rolemember/create
	V1RoleMemberCreateWithBodyWithResponse(ctx context.Context, params *V1RoleMemberCreateParams, body V1RoleMemberCreateRequestBody, reqEditors ...RequestEditorFn) (*V1RoleMemberCreateResponse, error)

	// /v1/rolemember/{rolememberid}/delete
	V1RoleMemberDeleteWithBodyWithResponse(ctx context.Context, rolememberid string, params *V1RoleMemberDeleteParams, body V1RoleMemberDeleteRequestBody, reqEditors ...RequestEditorFn) (*V1RoleMemberDeleteResponse, error)

	// /v1/staffs/create
	V1StaffsCreateWithBodyWithResponse(ctx context.Context, params *V1StaffsCreateParams, body V1StaffsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1StaffsCreateResponse, error)

	// /v1/staffs/{staffid}/delete
	V1StaffsDeleteWithResponse(ctx context.Context, staffid string, params *V1StaffsDeleteParams, reqEditors ...RequestEditorFn) (*V1StaffsDeleteResponse, error)

	// /v1/staffs/{staffid}/departmentancestors/fetch
	V1StaffsDeptAncestorsFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsDeptAncestorsFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsDeptAncestorsFetchResponse, error)

	// /v1/staffs/{staffid}/fetch
	V1StaffsFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsFetchResponse, error)

	// /v1/staffs/{staffid}/infor/fetch
	V1StaffsInforFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsInforFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsInforFetchResponse, error)

	// /v1/staffs/{staffid}/update
	V1StaffsUpdateWithBodyWithResponse(ctx context.Context, staffid string, params *V1StaffsUpdateParams, body V1StaffsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1StaffsUpdateResponse, error)

	// /v1/taggroups/create
	V1TagGroupsCreateWithBodyWithResponse(ctx context.Context, params *V1TagGroupsCreateParams, body V1TagGroupsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsCreateResponse, error)

	// /v1/taggroups/fetch
	V1TagGroupsFetchWithBodyWithResponse(ctx context.Context, params *V1TagGroupsFetchParams, body V1TagGroupsFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsFetchResponse, error)

	// /v1/taggroups/{tag_group_id}/delete
	V1TagGroupsDeleteWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsDeleteParams, reqEditors ...RequestEditorFn) (*V1TagGroupsDeleteResponse, error)

	// /v1/taggroups/{tag_group_id}/fetch
	V1TagGroupsInfoFetchWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsInfoFetchParams, reqEditors ...RequestEditorFn) (*V1TagGroupsInfoFetchResponse, error)

	// /v1/taggroups/{tag_group_id}/update
	V1TagGroupsUpdateWithBodyWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsUpdateParams, body V1TagGroupsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsUpdateResponse, error)

	// /v1/tags/create
	V1TagsCreateWithBodyWithResponse(ctx context.Context, params *V1TagsCreateParams, body V1TagsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1TagsCreateResponse, error)

	// /v1/tags/meta/fetch
	V1TagsMetaFetchWithBodyWithResponse(ctx context.Context, params *V1TagsMetaFetchParams, body V1TagsMetaFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagsMetaFetchResponse, error)

	// /v1/tags/staffids/fetch
	V1TagsFetchWithBodyWithResponse(ctx context.Context, params *V1TagsFetchParams, body V1TagsFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagsFetchResponse, error)

	// /v1/tags/{tagid}/delete
	V1TagsDeleteWithResponse(ctx context.Context, tagid string, params *V1TagsDeleteParams, reqEditors ...RequestEditorFn) (*V1TagsDeleteResponse, error)

	// /v1/tags/{tagid}/update
	V1TagsUpdateWithBodyWithResponse(ctx context.Context, tagid string, params *V1TagsUpdateParams, body V1TagsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1TagsUpdateResponse, error)

	// /v1/users/fetch
	V1UsersFetchWithResponse(ctx context.Context, params *V1UsersFetchParams, reqEditors ...RequestEditorFn) (*V1UsersFetchResponse, error)

	// /v1/usertoken/create
	V1UserTokenCreateWithResponse(ctx context.Context, params *V1UserTokenCreateParams, reqEditors ...RequestEditorFn) (*V1UserTokenCreateResponse, error)

	// /v2/chat/notification/update
	V2ChatNotificationUpdateWithBodyWithResponse(ctx context.Context, params *V2ChatNotificationUpdateParams, body V2ChatNotificationUpdateRequestBody, reqEditors ...RequestEditorFn) (*V2ChatNotificationUpdateResponse, error)

	// /v2/event/notification/create
	V2EventNotificationCreateWithBodyWithResponse(ctx context.Context, params *V2EventNotificationCreateParams, body V2EventNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*V2EventNotificationCreateResponse, error)

	// /v2/staffs/id_mapping/fetch
	V2StaffsIdMappingFetchWithResponse(ctx context.Context, params *V2StaffsIdMappingFetchParams, reqEditors ...RequestEditorFn) (*V2StaffsIdMappingFetchResponse, error)

	// /v2/staffs/search
	V2StaffsSearchWithBodyWithResponse(ctx context.Context, params *V2StaffsSearchParams, body V2StaffsSearchRequestBody, reqEditors ...RequestEditorFn) (*V2StaffsSearchResponse, error)
}

func (c *ClientWithResponses) V1AppRoamingOrgsFetchWithResponse(ctx context.Context, params *V1AppRoamingOrgsFetchParams, reqEditors ...RequestEditorFn) (*V1AppRoamingOrgsFetchResponse, error) {
	rsp, err := c.V1AppRoamingOrgsFetch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1AppRoamingOrgsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1AppTokenCreateWithResponse(ctx context.Context, params *V1AppTokenCreateParams, reqEditors ...RequestEditorFn) (*V1AppTokenCreateResponse, error) {
	rsp, err := c.V1AppTokenCreate(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1AppTokenCreateResponse(rsp)
}

func (c *ClientWithResponses) V1ChatNotificationFetchWithResponse(ctx context.Context, userid string, params *V1ChatNotificationFetchParams, reqEditors ...RequestEditorFn) (*V1ChatNotificationFetchResponse, error) {
	rsp, err := c.V1ChatNotificationFetch(ctx, userid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1ChatNotificationFetchResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsCreateWithBodyWithResponse(ctx context.Context, params *V1DeptsCreateParams, body V1DeptsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1DeptsCreateResponse, error) {
	rsp, err := c.V1DeptsCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsCreateResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsChildrenFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsChildrenFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsChildrenFetchResponse, error) {
	rsp, err := c.V1DeptsChildrenFetch(ctx, departmentid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsChildrenFetchResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsDeleteWithResponse(ctx context.Context, departmentid string, params *V1DeptsDeleteParams, reqEditors ...RequestEditorFn) (*V1DeptsDeleteResponse, error) {
	rsp, err := c.V1DeptsDelete(ctx, departmentid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsFetchResponse, error) {
	rsp, err := c.V1DeptsFetch(ctx, departmentid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsStaffsFetchWithResponse(ctx context.Context, departmentid string, params *V1DeptsStaffsFetchParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsFetchResponse, error) {
	rsp, err := c.V1DeptsStaffsFetch(ctx, departmentid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsStaffsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsStaffsCreateWithResponse(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsCreateParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsCreateResponse, error) {
	rsp, err := c.V1DeptsStaffsCreate(ctx, departmentid, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsStaffsCreateResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsStaffsDeleteWithResponse(ctx context.Context, departmentid string, staffid string, params *V1DeptsStaffsDeleteParams, reqEditors ...RequestEditorFn) (*V1DeptsStaffsDeleteResponse, error) {
	rsp, err := c.V1DeptsStaffsDelete(ctx, departmentid, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsStaffsDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1DeptsUpdateWithBodyWithResponse(ctx context.Context, departmentid string, params *V1DeptsUpdateParams, body V1DeptsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1DeptsUpdateResponse, error) {
	rsp, err := c.V1DeptsUpdateWithBody(ctx, departmentid, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1DeptsUpdateResponse(rsp)
}

func (c *ClientWithResponses) V1JsApiTokenCreateWithResponse(ctx context.Context, params *V1JsApiTokenCreateParams, reqEditors ...RequestEditorFn) (*V1JsApiTokenCreateResponse, error) {
	rsp, err := c.V1JsApiTokenCreate(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1JsApiTokenCreateResponse(rsp)
}

func (c *ClientWithResponses) V1MediasCreateWithBodyWithResponse(ctx context.Context, params *V1MediasCreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*V1MediasCreateResponse, error) {
	rsp, err := c.V1MediasCreateWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1MediasCreateResponse(rsp)
}

func (c *ClientWithResponses) V1MediasFetchWithResponse(ctx context.Context, mediaid string, params *V1MediasFetchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	rsp, err := c.V1MediasFetch(ctx, mediaid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *ClientWithResponses) V1MediasPathFetchWithResponse(ctx context.Context, mediaid string, params *V1MediasPathFetchParams, reqEditors ...RequestEditorFn) (*V1MediasPathFetchResponse, error) {
	rsp, err := c.V1MediasPathFetch(ctx, mediaid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1MediasPathFetchResponse(rsp)
}

func (c *ClientWithResponses) V1MessagesCreateWithBodyWithResponse(ctx context.Context, params *V1MessagesCreateParams, body V1MessagesCreateRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesCreateResponse, error) {
	rsp, err := c.V1MessagesCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1MessagesCreateResponse(rsp)
}

func (c *ClientWithResponses) V1MessagesNotificationCreateWithBodyWithResponse(ctx context.Context, params *V1MessagesNotificationCreateParams, body V1MessagesNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesNotificationCreateResponse, error) {
	rsp, err := c.V1MessagesNotificationCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1MessagesNotificationCreateResponse(rsp)
}

func (c *ClientWithResponses) V1MessagesRevokeWithBodyWithResponse(ctx context.Context, params *V1MessagesRevokeParams, body V1MessagesRevokeRequestBody, reqEditors ...RequestEditorFn) (*V1MessagesRevokeResponse, error) {
	rsp, err := c.V1MessagesRevokeWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1MessagesRevokeResponse(rsp)
}

func (c *ClientWithResponses) V1OrgExtraFieldIdsFetchWithResponse(ctx context.Context, orgid string, params *V1OrgExtraFieldIdsFetchParams, reqEditors ...RequestEditorFn) (*V1OrgExtraFieldIdsFetchResponse, error) {
	rsp, err := c.V1OrgExtraFieldIdsFetch(ctx, orgid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1OrgExtraFieldIdsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1OrgFetchWithResponse(ctx context.Context, orgid string, params *V1OrgFetchParams, reqEditors ...RequestEditorFn) (*V1OrgFetchResponse, error) {
	rsp, err := c.V1OrgFetch(ctx, orgid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1OrgFetchResponse(rsp)
}

func (c *ClientWithResponses) V1OrgRoleCreateWithBodyWithResponse(ctx context.Context, orgid string, params *V1OrgRoleCreateParams, body V1OrgRoleCreateRequestBody, reqEditors ...RequestEditorFn) (*V1OrgRoleCreateResponse, error) {
	rsp, err := c.V1OrgRoleCreateWithBody(ctx, orgid, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1OrgRoleCreateResponse(rsp)
}

func (c *ClientWithResponses) V1OrgRoleMembersFetchWithResponse(ctx context.Context, orgid string, roleid string, params *V1OrgRoleMembersFetchParams, reqEditors ...RequestEditorFn) (*V1OrgRoleMembersFetchResponse, error) {
	rsp, err := c.V1OrgRoleMembersFetch(ctx, orgid, roleid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1OrgRoleMembersFetchResponse(rsp)
}

func (c *ClientWithResponses) V1RoleMemberCreateWithBodyWithResponse(ctx context.Context, params *V1RoleMemberCreateParams, body V1RoleMemberCreateRequestBody, reqEditors ...RequestEditorFn) (*V1RoleMemberCreateResponse, error) {
	rsp, err := c.V1RoleMemberCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1RoleMemberCreateResponse(rsp)
}

func (c *ClientWithResponses) V1RoleMemberDeleteWithBodyWithResponse(ctx context.Context, rolememberid string, params *V1RoleMemberDeleteParams, body V1RoleMemberDeleteRequestBody, reqEditors ...RequestEditorFn) (*V1RoleMemberDeleteResponse, error) {
	rsp, err := c.V1RoleMemberDeleteWithBody(ctx, rolememberid, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1RoleMemberDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsCreateWithBodyWithResponse(ctx context.Context, params *V1StaffsCreateParams, body V1StaffsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1StaffsCreateResponse, error) {
	rsp, err := c.V1StaffsCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsCreateResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsDeleteWithResponse(ctx context.Context, staffid string, params *V1StaffsDeleteParams, reqEditors ...RequestEditorFn) (*V1StaffsDeleteResponse, error) {
	rsp, err := c.V1StaffsDelete(ctx, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsDeptAncestorsFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsDeptAncestorsFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsDeptAncestorsFetchResponse, error) {
	rsp, err := c.V1StaffsDeptAncestorsFetch(ctx, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsDeptAncestorsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsFetchResponse, error) {
	rsp, err := c.V1StaffsFetch(ctx, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsInforFetchWithResponse(ctx context.Context, staffid string, params *V1StaffsInforFetchParams, reqEditors ...RequestEditorFn) (*V1StaffsInforFetchResponse, error) {
	rsp, err := c.V1StaffsInforFetch(ctx, staffid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsInforFetchResponse(rsp)
}

func (c *ClientWithResponses) V1StaffsUpdateWithBodyWithResponse(ctx context.Context, staffid string, params *V1StaffsUpdateParams, body V1StaffsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1StaffsUpdateResponse, error) {
	rsp, err := c.V1StaffsUpdateWithBody(ctx, staffid, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1StaffsUpdateResponse(rsp)
}

func (c *ClientWithResponses) V1TagGroupsCreateWithBodyWithResponse(ctx context.Context, params *V1TagGroupsCreateParams, body V1TagGroupsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsCreateResponse, error) {
	rsp, err := c.V1TagGroupsCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagGroupsCreateResponse(rsp)
}

func (c *ClientWithResponses) V1TagGroupsFetchWithBodyWithResponse(ctx context.Context, params *V1TagGroupsFetchParams, body V1TagGroupsFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsFetchResponse, error) {
	rsp, err := c.V1TagGroupsFetchWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagGroupsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1TagGroupsDeleteWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsDeleteParams, reqEditors ...RequestEditorFn) (*V1TagGroupsDeleteResponse, error) {
	rsp, err := c.V1TagGroupsDelete(ctx, tagGroupId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagGroupsDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1TagGroupsInfoFetchWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsInfoFetchParams, reqEditors ...RequestEditorFn) (*V1TagGroupsInfoFetchResponse, error) {
	rsp, err := c.V1TagGroupsInfoFetch(ctx, tagGroupId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagGroupsInfoFetchResponse(rsp)
}

func (c *ClientWithResponses) V1TagGroupsUpdateWithBodyWithResponse(ctx context.Context, tagGroupId string, params *V1TagGroupsUpdateParams, body V1TagGroupsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1TagGroupsUpdateResponse, error) {
	rsp, err := c.V1TagGroupsUpdateWithBody(ctx, tagGroupId, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagGroupsUpdateResponse(rsp)
}

func (c *ClientWithResponses) V1TagsCreateWithBodyWithResponse(ctx context.Context, params *V1TagsCreateParams, body V1TagsCreateRequestBody, reqEditors ...RequestEditorFn) (*V1TagsCreateResponse, error) {
	rsp, err := c.V1TagsCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagsCreateResponse(rsp)
}

func (c *ClientWithResponses) V1TagsMetaFetchWithBodyWithResponse(ctx context.Context, params *V1TagsMetaFetchParams, body V1TagsMetaFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagsMetaFetchResponse, error) {
	rsp, err := c.V1TagsMetaFetchWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagsMetaFetchResponse(rsp)
}

func (c *ClientWithResponses) V1TagsFetchWithBodyWithResponse(ctx context.Context, params *V1TagsFetchParams, body V1TagsFetchRequestBody, reqEditors ...RequestEditorFn) (*V1TagsFetchResponse, error) {
	rsp, err := c.V1TagsFetchWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagsFetchResponse(rsp)
}

func (c *ClientWithResponses) V1TagsDeleteWithResponse(ctx context.Context, tagid string, params *V1TagsDeleteParams, reqEditors ...RequestEditorFn) (*V1TagsDeleteResponse, error) {
	rsp, err := c.V1TagsDelete(ctx, tagid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagsDeleteResponse(rsp)
}

func (c *ClientWithResponses) V1TagsUpdateWithBodyWithResponse(ctx context.Context, tagid string, params *V1TagsUpdateParams, body V1TagsUpdateRequestBody, reqEditors ...RequestEditorFn) (*V1TagsUpdateResponse, error) {
	rsp, err := c.V1TagsUpdateWithBody(ctx, tagid, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1TagsUpdateResponse(rsp)
}

func (c *ClientWithResponses) V1UsersFetchWithResponse(ctx context.Context, params *V1UsersFetchParams, reqEditors ...RequestEditorFn) (*V1UsersFetchResponse, error) {
	rsp, err := c.V1UsersFetch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1UsersFetchResponse(rsp)
}

func (c *ClientWithResponses) V1UserTokenCreateWithResponse(ctx context.Context, params *V1UserTokenCreateParams, reqEditors ...RequestEditorFn) (*V1UserTokenCreateResponse, error) {
	rsp, err := c.V1UserTokenCreate(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1UserTokenCreateResponse(rsp)
}

func (c *ClientWithResponses) V2ChatNotificationUpdateWithBodyWithResponse(ctx context.Context, params *V2ChatNotificationUpdateParams, body V2ChatNotificationUpdateRequestBody, reqEditors ...RequestEditorFn) (*V2ChatNotificationUpdateResponse, error) {
	rsp, err := c.V2ChatNotificationUpdateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV2ChatNotificationUpdateResponse(rsp)
}

func (c *ClientWithResponses) V2EventNotificationCreateWithBodyWithResponse(ctx context.Context, params *V2EventNotificationCreateParams, body V2EventNotificationCreateRequestBody, reqEditors ...RequestEditorFn) (*V2EventNotificationCreateResponse, error) {
	rsp, err := c.V2EventNotificationCreateWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV2EventNotificationCreateResponse(rsp)
}

func (c *ClientWithResponses) V2StaffsIdMappingFetchWithResponse(ctx context.Context, params *V2StaffsIdMappingFetchParams, reqEditors ...RequestEditorFn) (*V2StaffsIdMappingFetchResponse, error) {
	rsp, err := c.V2StaffsIdMappingFetch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV2StaffsIdMappingFetchResponse(rsp)
}

func (c *ClientWithResponses) V2StaffsSearchWithBodyWithResponse(ctx context.Context, params *V2StaffsSearchParams, body V2StaffsSearchRequestBody, reqEditors ...RequestEditorFn) (*V2StaffsSearchResponse, error) {
	rsp, err := c.V2StaffsSearchWithBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV2StaffsSearchResponse(rsp)
}

func ParseV1AppRoamingOrgsFetchResponse(rsp *http.Response) (*V1AppRoamingOrgsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1AppRoamingOrgsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1AppTokenCreateResponse(rsp *http.Response) (*V1AppTokenCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1AppTokenCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1ChatNotificationFetchResponse(rsp *http.Response) (*V1ChatNotificationFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1ChatNotificationFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsCreateResponse(rsp *http.Response) (*V1DeptsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsChildrenFetchResponse(rsp *http.Response) (*V1DeptsChildrenFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsChildrenFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsDeleteResponse(rsp *http.Response) (*V1DeptsDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsFetchResponse(rsp *http.Response) (*V1DeptsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsStaffsFetchResponse(rsp *http.Response) (*V1DeptsStaffsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsStaffsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsStaffsCreateResponse(rsp *http.Response) (*V1DeptsStaffsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsStaffsCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsStaffsDeleteResponse(rsp *http.Response) (*V1DeptsStaffsDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsStaffsDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1DeptsUpdateResponse(rsp *http.Response) (*V1DeptsUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1DeptsUpdateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1JsApiTokenCreateResponse(rsp *http.Response) (*V1JsApiTokenCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1JsApiTokenCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1MediasCreateResponse(rsp *http.Response) (*V1MediasCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1MediasCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1MediasFetchResponse(rsp *http.Response) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func ParseV1MediasPathFetchResponse(rsp *http.Response) (*V1MediasPathFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1MediasPathFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1MessagesCreateResponse(rsp *http.Response) (*V1MessagesCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1MessagesCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1MessagesNotificationCreateResponse(rsp *http.Response) (*V1MessagesNotificationCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1MessagesNotificationCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1MessagesRevokeResponse(rsp *http.Response) (*V1MessagesRevokeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1MessagesRevokeResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1OrgExtraFieldIdsFetchResponse(rsp *http.Response) (*V1OrgExtraFieldIdsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1OrgExtraFieldIdsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1OrgFetchResponse(rsp *http.Response) (*V1OrgFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1OrgFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1OrgRoleCreateResponse(rsp *http.Response) (*V1OrgRoleCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1OrgRoleCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1OrgRoleMembersFetchResponse(rsp *http.Response) (*V1OrgRoleMembersFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1OrgRoleMembersFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1RoleMemberCreateResponse(rsp *http.Response) (*V1RoleMemberCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1RoleMemberCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1RoleMemberDeleteResponse(rsp *http.Response) (*V1RoleMemberDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1RoleMemberDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsCreateResponse(rsp *http.Response) (*V1StaffsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsDeleteResponse(rsp *http.Response) (*V1StaffsDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsDeptAncestorsFetchResponse(rsp *http.Response) (*V1StaffsDeptAncestorsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsDeptAncestorsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsFetchResponse(rsp *http.Response) (*V1StaffsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsInforFetchResponse(rsp *http.Response) (*V1StaffsInforFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsInforFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1StaffsUpdateResponse(rsp *http.Response) (*V1StaffsUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1StaffsUpdateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagGroupsCreateResponse(rsp *http.Response) (*V1TagGroupsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagGroupsCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagGroupsFetchResponse(rsp *http.Response) (*V1TagGroupsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagGroupsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagGroupsDeleteResponse(rsp *http.Response) (*V1TagGroupsDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagGroupsDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagGroupsInfoFetchResponse(rsp *http.Response) (*V1TagGroupsInfoFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagGroupsInfoFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagGroupsUpdateResponse(rsp *http.Response) (*V1TagGroupsUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagGroupsUpdateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagsCreateResponse(rsp *http.Response) (*V1TagsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagsCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagsMetaFetchResponse(rsp *http.Response) (*V1TagsMetaFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagsMetaFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagsFetchResponse(rsp *http.Response) (*V1TagsFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagsFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagsDeleteResponse(rsp *http.Response) (*V1TagsDeleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagsDeleteResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1TagsUpdateResponse(rsp *http.Response) (*V1TagsUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1TagsUpdateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1UsersFetchResponse(rsp *http.Response) (*V1UsersFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1UsersFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV1UserTokenCreateResponse(rsp *http.Response) (*V1UserTokenCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V1UserTokenCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV2ChatNotificationUpdateResponse(rsp *http.Response) (*V2ChatNotificationUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V2ChatNotificationUpdateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV2EventNotificationCreateResponse(rsp *http.Response) (*V2EventNotificationCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V2EventNotificationCreateResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV2StaffsIdMappingFetchResponse(rsp *http.Response) (*V2StaffsIdMappingFetchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V2StaffsIdMappingFetchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func ParseV2StaffsSearchResponse(rsp *http.Response) (*V2StaffsSearchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(V2StaffsSearchResponse)
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

