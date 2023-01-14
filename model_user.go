/*
VRChat API Documentation


API version: 1.10.1
Contact: me@ariesclark.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
)

// User struct for User
type User struct {
	AllowAvatarCopying bool `json:"allowAvatarCopying"`
	Bio string `json:"bio"`
	BioLinks []string `json:"bioLinks"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string `json:"currentAvatarImageUrl"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string `json:"currentAvatarThumbnailImageUrl"`
	DateJoined string `json:"date_joined"`
	DeveloperType DeveloperType `json:"developerType"`
	// A users visual display name. This is what shows up in-game, and can different from their `username`. Changing display name is restricted to a cooldown period.
	DisplayName string `json:"displayName"`
	FriendKey string `json:"friendKey"`
	FriendRequestStatus *string `json:"friendRequestStatus,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id string `json:"id"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	InstanceId *string `json:"instanceId,omitempty"`
	// Either their `friendKey`, or empty string if you are not friends. Unknown usage.
	IsFriend bool `json:"isFriend"`
	// Either a date-time or empty string.
	LastActivity string `json:"last_activity"`
	// Either a date-time or empty string.
	LastLogin string `json:"last_login"`
	// This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform string `json:"last_platform"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	Location *string `json:"location,omitempty"`
	Note *string `json:"note,omitempty"`
	ProfilePicOverride string `json:"profilePicOverride"`
	State UserState `json:"state"`
	Status UserStatus `json:"status"`
	StatusDescription string `json:"statusDescription"`
	//  
	Tags []string `json:"tags"`
	TravelingToInstance *string `json:"travelingToInstance,omitempty"`
	TravelingToLocation *string `json:"travelingToLocation,omitempty"`
	TravelingToWorld *string `json:"travelingToWorld,omitempty"`
	UserIcon string `json:"userIcon"`
	// -| A users unique name, used during login. This is different from `displayName` which is what shows up in-game. A users `username` can never be changed.' **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	// Deprecated
	Username *string `json:"username,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId *string `json:"worldId,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(allowAvatarCopying bool, bio string, bioLinks []string, currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, friendKey string, id string, isFriend bool, lastActivity string, lastLogin string, lastPlatform string, profilePicOverride string, state UserState, status UserStatus, statusDescription string, tags []string, userIcon string) *User {
	this := User{}
	this.AllowAvatarCopying = allowAvatarCopying
	this.Bio = bio
	this.BioLinks = bioLinks
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DateJoined = dateJoined
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.FriendKey = friendKey
	this.Id = id
	this.IsFriend = isFriend
	this.LastActivity = lastActivity
	this.LastLogin = lastLogin
	this.LastPlatform = lastPlatform
	this.ProfilePicOverride = profilePicOverride
	this.State = state
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	this.UserIcon = userIcon
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var allowAvatarCopying bool = true
	this.AllowAvatarCopying = allowAvatarCopying
	var developerType DeveloperType = DEVELOPERTYPE_NONE
	this.DeveloperType = developerType
	var state UserState = USERSTATE_OFFLINE
	this.State = state
	var status UserStatus = USERSTATUS_OFFLINE
	this.Status = status
	return &this
}

// GetAllowAvatarCopying returns the AllowAvatarCopying field value
func (o *User) GetAllowAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAvatarCopying
}

// GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *User) GetAllowAvatarCopyingOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AllowAvatarCopying, true
}

// SetAllowAvatarCopying sets field value
func (o *User) SetAllowAvatarCopying(v bool) {
	o.AllowAvatarCopying = v
}

// GetBio returns the Bio field value
func (o *User) GetBio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *User) GetBioOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *User) SetBio(v string) {
	o.Bio = v
}

// GetBioLinks returns the BioLinks field value
func (o *User) GetBioLinks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value
// and a boolean to check if the value has been set.
func (o *User) GetBioLinksOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BioLinks, true
}

// SetBioLinks sets field value
func (o *User) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *User) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *User) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *User) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *User) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDateJoined returns the DateJoined field value
func (o *User) GetDateJoined() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateJoined
}

// GetDateJoinedOk returns a tuple with the DateJoined field value
// and a boolean to check if the value has been set.
func (o *User) GetDateJoinedOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DateJoined, true
}

// SetDateJoined sets field value
func (o *User) SetDateJoined(v string) {
	o.DateJoined = v
}

// GetDeveloperType returns the DeveloperType field value
func (o *User) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *User) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *User) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *User) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *User) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFriendKey returns the FriendKey field value
func (o *User) GetFriendKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value
// and a boolean to check if the value has been set.
func (o *User) GetFriendKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FriendKey, true
}

// SetFriendKey sets field value
func (o *User) SetFriendKey(v string) {
	o.FriendKey = v
}

// GetFriendRequestStatus returns the FriendRequestStatus field value if set, zero value otherwise.
func (o *User) GetFriendRequestStatus() string {
	if o == nil || isNil(o.FriendRequestStatus) {
		var ret string
		return ret
	}
	return *o.FriendRequestStatus
}

// GetFriendRequestStatusOk returns a tuple with the FriendRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFriendRequestStatusOk() (*string, bool) {
	if o == nil || isNil(o.FriendRequestStatus) {
    return nil, false
	}
	return o.FriendRequestStatus, true
}

// HasFriendRequestStatus returns a boolean if a field has been set.
func (o *User) HasFriendRequestStatus() bool {
	if o != nil && !isNil(o.FriendRequestStatus) {
		return true
	}

	return false
}

// SetFriendRequestStatus gets a reference to the given string and assigns it to the FriendRequestStatus field.
func (o *User) SetFriendRequestStatus(v string) {
	o.FriendRequestStatus = &v
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *User) GetInstanceId() string {
	if o == nil || isNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.InstanceId) {
    return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *User) HasInstanceId() bool {
	if o != nil && !isNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *User) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetIsFriend returns the IsFriend field value
func (o *User) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *User) GetIsFriendOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *User) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastActivity returns the LastActivity field value
func (o *User) GetLastActivity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value
// and a boolean to check if the value has been set.
func (o *User) GetLastActivityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LastActivity, true
}

// SetLastActivity sets field value
func (o *User) SetLastActivity(v string) {
	o.LastActivity = v
}

// GetLastLogin returns the LastLogin field value
func (o *User) GetLastLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LastLogin, true
}

// SetLastLogin sets field value
func (o *User) SetLastLogin(v string) {
	o.LastLogin = v
}

// GetLastPlatform returns the LastPlatform field value
func (o *User) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *User) GetLastPlatformOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *User) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *User) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *User) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *User) SetLocation(v string) {
	o.Location = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *User) GetNote() string {
	if o == nil || isNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNoteOk() (*string, bool) {
	if o == nil || isNil(o.Note) {
    return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *User) HasNote() bool {
	if o != nil && !isNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *User) SetNote(v string) {
	o.Note = &v
}

// GetProfilePicOverride returns the ProfilePicOverride field value
func (o *User) GetProfilePicOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value
// and a boolean to check if the value has been set.
func (o *User) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProfilePicOverride, true
}

// SetProfilePicOverride sets field value
func (o *User) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = v
}

// GetState returns the State field value
func (o *User) GetState() UserState {
	if o == nil {
		var ret UserState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *User) GetStateOk() (*UserState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *User) SetState(v UserState) {
	o.State = v
}

// GetStatus returns the Status field value
func (o *User) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *User) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *User) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *User) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *User) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *User) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *User) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *User) SetTags(v []string) {
	o.Tags = v
}

// GetTravelingToInstance returns the TravelingToInstance field value if set, zero value otherwise.
func (o *User) GetTravelingToInstance() string {
	if o == nil || isNil(o.TravelingToInstance) {
		var ret string
		return ret
	}
	return *o.TravelingToInstance
}

// GetTravelingToInstanceOk returns a tuple with the TravelingToInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToInstanceOk() (*string, bool) {
	if o == nil || isNil(o.TravelingToInstance) {
    return nil, false
	}
	return o.TravelingToInstance, true
}

// HasTravelingToInstance returns a boolean if a field has been set.
func (o *User) HasTravelingToInstance() bool {
	if o != nil && !isNil(o.TravelingToInstance) {
		return true
	}

	return false
}

// SetTravelingToInstance gets a reference to the given string and assigns it to the TravelingToInstance field.
func (o *User) SetTravelingToInstance(v string) {
	o.TravelingToInstance = &v
}

// GetTravelingToLocation returns the TravelingToLocation field value if set, zero value otherwise.
func (o *User) GetTravelingToLocation() string {
	if o == nil || isNil(o.TravelingToLocation) {
		var ret string
		return ret
	}
	return *o.TravelingToLocation
}

// GetTravelingToLocationOk returns a tuple with the TravelingToLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToLocationOk() (*string, bool) {
	if o == nil || isNil(o.TravelingToLocation) {
    return nil, false
	}
	return o.TravelingToLocation, true
}

// HasTravelingToLocation returns a boolean if a field has been set.
func (o *User) HasTravelingToLocation() bool {
	if o != nil && !isNil(o.TravelingToLocation) {
		return true
	}

	return false
}

// SetTravelingToLocation gets a reference to the given string and assigns it to the TravelingToLocation field.
func (o *User) SetTravelingToLocation(v string) {
	o.TravelingToLocation = &v
}

// GetTravelingToWorld returns the TravelingToWorld field value if set, zero value otherwise.
func (o *User) GetTravelingToWorld() string {
	if o == nil || isNil(o.TravelingToWorld) {
		var ret string
		return ret
	}
	return *o.TravelingToWorld
}

// GetTravelingToWorldOk returns a tuple with the TravelingToWorld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToWorldOk() (*string, bool) {
	if o == nil || isNil(o.TravelingToWorld) {
    return nil, false
	}
	return o.TravelingToWorld, true
}

// HasTravelingToWorld returns a boolean if a field has been set.
func (o *User) HasTravelingToWorld() bool {
	if o != nil && !isNil(o.TravelingToWorld) {
		return true
	}

	return false
}

// SetTravelingToWorld gets a reference to the given string and assigns it to the TravelingToWorld field.
func (o *User) SetTravelingToWorld(v string) {
	o.TravelingToWorld = &v
}

// GetUserIcon returns the UserIcon field value
func (o *User) GetUserIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value
// and a boolean to check if the value has been set.
func (o *User) GetUserIconOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UserIcon, true
}

// SetUserIcon sets field value
func (o *User) SetUserIcon(v string) {
	o.UserIcon = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
// Deprecated
func (o *User) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
// Deprecated
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *User) GetWorldId() string {
	if o == nil || isNil(o.WorldId) {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetWorldIdOk() (*string, bool) {
	if o == nil || isNil(o.WorldId) {
    return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *User) HasWorldId() bool {
	if o != nil && !isNil(o.WorldId) {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *User) SetWorldId(v string) {
	o.WorldId = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowAvatarCopying"] = o.AllowAvatarCopying
	}
	if true {
		toSerialize["bio"] = o.Bio
	}
	if true {
		toSerialize["bioLinks"] = o.BioLinks
	}
	if true {
		toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	}
	if true {
		toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	}
	if true {
		toSerialize["date_joined"] = o.DateJoined
	}
	if true {
		toSerialize["developerType"] = o.DeveloperType
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["friendKey"] = o.FriendKey
	}
	if !isNil(o.FriendRequestStatus) {
		toSerialize["friendRequestStatus"] = o.FriendRequestStatus
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if true {
		toSerialize["isFriend"] = o.IsFriend
	}
	if true {
		toSerialize["last_activity"] = o.LastActivity
	}
	if true {
		toSerialize["last_login"] = o.LastLogin
	}
	if true {
		toSerialize["last_platform"] = o.LastPlatform
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if true {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.TravelingToInstance) {
		toSerialize["travelingToInstance"] = o.TravelingToInstance
	}
	if !isNil(o.TravelingToLocation) {
		toSerialize["travelingToLocation"] = o.TravelingToLocation
	}
	if !isNil(o.TravelingToWorld) {
		toSerialize["travelingToWorld"] = o.TravelingToWorld
	}
	if true {
		toSerialize["userIcon"] = o.UserIcon
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.WorldId) {
		toSerialize["worldId"] = o.WorldId
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


