# Go API client for vrchatapi


# Welcome to the VRChat API

Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**.
This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat.
The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.

The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance
with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities.
Malicious usage or spamming the API may result in account termination.
Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.

![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)

Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported.
VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**.
Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support.
We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing.
If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.

# Getting Started

The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more.
The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website.
This documentation focuses only on the Web API.

The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects.
Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes.
Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.

<div class=\"callout callout-error\">
  <strong>🛑 Warning! Do not touch Photon!</strong><br>
  Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination.
</div>

<div class=\"callout callout-info\">
  <strong>ℹ️ API Key and Authentication</strong><br>
  The API Key has always been the same and is currently <code>JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26</code>.
  Read <a href=\"#tag--authentication\">Authentication</a> for how to log in.
</div>

# Using the API

For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source
API client that's great for sending requests to the API in an orderly fashion.
Insomnia allows you to send data in the format that's required for VRChat's API.
It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to
[vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.

For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs.
This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls
rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification,
sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon
as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:

* [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat)
* [Dart](https://pub.dev/packages/vrchat_dart)
* [Rust](https://crates.io/crates/vrchatapi)
* [C#](https://github.com/vrchatapi/vrchatapi-csharp)
* [Python](https://github.com/vrchatapi/vrchatapi-python)

# Pagination

Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br>
Using both the limit and offset parameters allows you to easily paginate through a large number of objects.

| Query Parameter | Type | Description |
| ----------|--|------- |
| `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.|
| `offset` | integer  | A zero-based offset from the default object sorting.|

If a request returns fewer objects than the `limit` parameter, there are no more items available to return.

# Contribution

Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries?
This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given
commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.

[![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.10.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/VRChatAPI](https://github.com/VRChatAPI)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import vrchatapi "github.com/vrchatapi/vrchatapi-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), vrchatapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), vrchatapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), vrchatapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), vrchatapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.vrchat.cloud/api/1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationApi* | [**CheckUserExists**](docs/AuthenticationApi.md#checkuserexists) | **Get** /auth/exists | Check User Exists
*AuthenticationApi* | [**DeleteUser**](docs/AuthenticationApi.md#deleteuser) | **Put** /users/{userId}/delete | Delete User
*AuthenticationApi* | [**GetCurrentUser**](docs/AuthenticationApi.md#getcurrentuser) | **Get** /auth/user | Login and/or Get Current User Info
*AuthenticationApi* | [**Logout**](docs/AuthenticationApi.md#logout) | **Put** /logout | Logout
*AuthenticationApi* | [**Verify2FA**](docs/AuthenticationApi.md#verify2fa) | **Post** /auth/twofactorauth/totp/verify | Verify 2FA code
*AuthenticationApi* | [**Verify2FAEmailCode**](docs/AuthenticationApi.md#verify2faemailcode) | **Post** /auth/twofactorauth/emailotp/verify | Verify 2FA email code
*AuthenticationApi* | [**VerifyAuthToken**](docs/AuthenticationApi.md#verifyauthtoken) | **Get** /auth | Verify Auth Token
*AuthenticationApi* | [**VerifyRecoveryCode**](docs/AuthenticationApi.md#verifyrecoverycode) | **Post** /auth/twofactorauth/otp/verify | Verify 2FA code with Recovery code
*AvatarsApi* | [**CreateAvatar**](docs/AvatarsApi.md#createavatar) | **Post** /avatars | Create Avatar
*AvatarsApi* | [**DeleteAvatar**](docs/AvatarsApi.md#deleteavatar) | **Delete** /avatars/{avatarId} | Delete Avatar
*AvatarsApi* | [**GetAvatar**](docs/AvatarsApi.md#getavatar) | **Get** /avatars/{avatarId} | Get Avatar
*AvatarsApi* | [**GetFavoritedAvatars**](docs/AvatarsApi.md#getfavoritedavatars) | **Get** /avatars/favorites | List Favorited Avatars
*AvatarsApi* | [**GetOwnAvatar**](docs/AvatarsApi.md#getownavatar) | **Get** /users/{userId}/avatar | Get Own Avatar
*AvatarsApi* | [**SearchAvatars**](docs/AvatarsApi.md#searchavatars) | **Get** /avatars | Search Avatars
*AvatarsApi* | [**SelectAvatar**](docs/AvatarsApi.md#selectavatar) | **Put** /avatars/{avatarId}/select | Select Avatar
*AvatarsApi* | [**SelectFallbackAvatar**](docs/AvatarsApi.md#selectfallbackavatar) | **Put** /avatars/{avatarId}/selectFallback | Select Fallback Avatar
*AvatarsApi* | [**UpdateAvatar**](docs/AvatarsApi.md#updateavatar) | **Put** /avatars/{avatarId} | Update Avatar
*EconomyApi* | [**GetCurrentSubscriptions**](docs/EconomyApi.md#getcurrentsubscriptions) | **Get** /auth/user/subscription | Get Current Subscriptions
*EconomyApi* | [**GetLicenseGroup**](docs/EconomyApi.md#getlicensegroup) | **Get** /licenseGroups/{licenseGroupId} | Get License Group
*EconomyApi* | [**GetSteamTransaction**](docs/EconomyApi.md#getsteamtransaction) | **Get** /Steam/transactions/{transactionId} | Get Steam Transaction
*EconomyApi* | [**GetSteamTransactions**](docs/EconomyApi.md#getsteamtransactions) | **Get** /Steam/transactions | List Steam Transactions
*EconomyApi* | [**GetSubscriptions**](docs/EconomyApi.md#getsubscriptions) | **Get** /subscriptions | List Subscriptions
*FavoritesApi* | [**AddFavorite**](docs/FavoritesApi.md#addfavorite) | **Post** /favorites | Add Favorite
*FavoritesApi* | [**ClearFavoriteGroup**](docs/FavoritesApi.md#clearfavoritegroup) | **Delete** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Clear Favorite Group
*FavoritesApi* | [**GetFavorite**](docs/FavoritesApi.md#getfavorite) | **Get** /favorites/{favoriteId} | Show Favorite
*FavoritesApi* | [**GetFavoriteGroup**](docs/FavoritesApi.md#getfavoritegroup) | **Get** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Show Favorite Group
*FavoritesApi* | [**GetFavoriteGroups**](docs/FavoritesApi.md#getfavoritegroups) | **Get** /favorite/groups | List Favorite Groups
*FavoritesApi* | [**GetFavorites**](docs/FavoritesApi.md#getfavorites) | **Get** /favorites | List Favorites
*FavoritesApi* | [**RemoveFavorite**](docs/FavoritesApi.md#removefavorite) | **Delete** /favorites/{favoriteId} | Remove Favorite
*FavoritesApi* | [**UpdateFavoriteGroup**](docs/FavoritesApi.md#updatefavoritegroup) | **Put** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Update Favorite Group
*FilesApi* | [**CreateFile**](docs/FilesApi.md#createfile) | **Post** /file | Create File
*FilesApi* | [**CreateFileVersion**](docs/FilesApi.md#createfileversion) | **Post** /file/{fileId} | Create File Version
*FilesApi* | [**DeleteFile**](docs/FilesApi.md#deletefile) | **Delete** /file/{fileId} | Delete File
*FilesApi* | [**DeleteFileVersion**](docs/FilesApi.md#deletefileversion) | **Delete** /file/{fileId}/{versionId} | Delete File Version
*FilesApi* | [**DownloadFileVersion**](docs/FilesApi.md#downloadfileversion) | **Get** /file/{fileId}/{versionId} | Download File Version
*FilesApi* | [**FinishFileDataUpload**](docs/FilesApi.md#finishfiledataupload) | **Put** /file/{fileId}/{versionId}/{fileType}/finish | Finish FileData Upload
*FilesApi* | [**GetFile**](docs/FilesApi.md#getfile) | **Get** /file/{fileId} | Show File
*FilesApi* | [**GetFileDataUploadStatus**](docs/FilesApi.md#getfiledatauploadstatus) | **Get** /file/{fileId}/{versionId}/{fileType}/status | Check FileData Upload Status
*FilesApi* | [**GetFiles**](docs/FilesApi.md#getfiles) | **Get** /files | List Files
*FilesApi* | [**StartFileDataUpload**](docs/FilesApi.md#startfiledataupload) | **Put** /file/{fileId}/{versionId}/{fileType}/start | Start FileData Upload
*FriendsApi* | [**DeleteFriendRequest**](docs/FriendsApi.md#deletefriendrequest) | **Delete** /user/{userId}/friendRequest | Delete Friend Request
*FriendsApi* | [**Friend**](docs/FriendsApi.md#friend) | **Post** /user/{userId}/friendRequest | Send Friend Request
*FriendsApi* | [**GetFriendStatus**](docs/FriendsApi.md#getfriendstatus) | **Get** /user/{userId}/friendStatus | Check Friend Status
*FriendsApi* | [**GetFriends**](docs/FriendsApi.md#getfriends) | **Get** /auth/user/friends | List Friends
*FriendsApi* | [**Unfriend**](docs/FriendsApi.md#unfriend) | **Delete** /auth/user/friends/{userId} | Unfriend
*GroupsApi* | [**AddGroupGalleryImage**](docs/GroupsApi.md#addgroupgalleryimage) | **Post** /groups/{groupId}/galleries/{groupGalleryId}/images | Add Group Gallery Image
*GroupsApi* | [**AddGroupMemberRole**](docs/GroupsApi.md#addgroupmemberrole) | **Put** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Add Role to GroupMember
*GroupsApi* | [**BanGroupMember**](docs/GroupsApi.md#bangroupmember) | **Post** /groups/{groupId}/bans | Ban Group Member
*GroupsApi* | [**CancelGroupRequest**](docs/GroupsApi.md#cancelgrouprequest) | **Delete** /groups/{groupId}/requests | Cancel Group Join Request
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /groups | Create Group
*GroupsApi* | [**CreateGroupAnnouncement**](docs/GroupsApi.md#creategroupannouncement) | **Post** /groups/{groupId}/announcement | Create Group Announcement
*GroupsApi* | [**CreateGroupGallery**](docs/GroupsApi.md#creategroupgallery) | **Post** /groups/{groupId}/galleries | Create Group Gallery
*GroupsApi* | [**CreateGroupInvite**](docs/GroupsApi.md#creategroupinvite) | **Post** /groups/{groupId}/invites | Invite User to Group
*GroupsApi* | [**CreateGroupRole**](docs/GroupsApi.md#creategrouprole) | **Post** /groups/{groupId}/roles | Create GroupRole
*GroupsApi* | [**DeleteGroup**](docs/GroupsApi.md#deletegroup) | **Delete** /groups/{groupId} | Delete Group
*GroupsApi* | [**DeleteGroupAnnouncement**](docs/GroupsApi.md#deletegroupannouncement) | **Delete** /groups/{groupId}/announcement | Delete Group Announcement
*GroupsApi* | [**DeleteGroupGallery**](docs/GroupsApi.md#deletegroupgallery) | **Delete** /groups/{groupId}/galleries/{groupGalleryId} | Delete Group Gallery
*GroupsApi* | [**DeleteGroupGalleryImage**](docs/GroupsApi.md#deletegroupgalleryimage) | **Delete** /groups/{groupId}/galleries/{groupGalleryId}/images/{groupGalleryImageId} | Delete Group Gallery Image
*GroupsApi* | [**DeleteGroupInvite**](docs/GroupsApi.md#deletegroupinvite) | **Delete** /groups/{groupId}/invites/{userId} | Delete User Invite
*GroupsApi* | [**DeleteGroupRole**](docs/GroupsApi.md#deletegrouprole) | **Delete** /groups/{groupId}/roles/{groupRoleId} | Delete Group Role
*GroupsApi* | [**GetGroup**](docs/GroupsApi.md#getgroup) | **Get** /groups/{groupId} | Get Group by ID
*GroupsApi* | [**GetGroupAnnouncements**](docs/GroupsApi.md#getgroupannouncements) | **Get** /groups/{groupId}/announcement | Get Group Announcement
*GroupsApi* | [**GetGroupAuditLogs**](docs/GroupsApi.md#getgroupauditlogs) | **Get** /groups/{groupId}/auditLogs | Get Group Audit Logs
*GroupsApi* | [**GetGroupBans**](docs/GroupsApi.md#getgroupbans) | **Get** /groups/{groupId}/bans | Get Group Bans
*GroupsApi* | [**GetGroupGalleryImages**](docs/GroupsApi.md#getgroupgalleryimages) | **Get** /groups/{groupId}/galleries/{groupGalleryId} | Get Group Gallery Images
*GroupsApi* | [**GetGroupInvites**](docs/GroupsApi.md#getgroupinvites) | **Get** /groups/{groupId}/invites | Get Group Invites Sent
*GroupsApi* | [**GetGroupMember**](docs/GroupsApi.md#getgroupmember) | **Get** /groups/{groupId}/members/{userId} | Get Group Member
*GroupsApi* | [**GetGroupMembers**](docs/GroupsApi.md#getgroupmembers) | **Get** /groups/{groupId}/members | List Group Members
*GroupsApi* | [**GetGroupPermissions**](docs/GroupsApi.md#getgrouppermissions) | **Get** /groups/{groupId}/permissions | List Group Permissions
*GroupsApi* | [**GetGroupRequests**](docs/GroupsApi.md#getgrouprequests) | **Get** /groups/{groupId}/requests | Get Group Join Requests
*GroupsApi* | [**GetGroupRoles**](docs/GroupsApi.md#getgrouproles) | **Get** /groups/{groupId}/roles | Get Group Roles
*GroupsApi* | [**JoinGroup**](docs/GroupsApi.md#joingroup) | **Post** /groups/{groupId}/join | Join Group
*GroupsApi* | [**KickGroupMember**](docs/GroupsApi.md#kickgroupmember) | **Delete** /groups/{groupId}/members/{userId} | Kick Group Member
*GroupsApi* | [**LeaveGroup**](docs/GroupsApi.md#leavegroup) | **Post** /groups/{groupId}/leave | Leave Group
*GroupsApi* | [**RemoveGroupMemberRole**](docs/GroupsApi.md#removegroupmemberrole) | **Delete** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Remove Role from GroupMember
*GroupsApi* | [**RespondGroupJoinRequest**](docs/GroupsApi.md#respondgroupjoinrequest) | **Put** /groups/{groupId}/requests/{userId} | Respond Group Join request
*GroupsApi* | [**UnbanGroupMember**](docs/GroupsApi.md#unbangroupmember) | **Delete** /groups/{groupId}/bans/{userId} | Unban Group Member
*GroupsApi* | [**UpdateGroup**](docs/GroupsApi.md#updategroup) | **Put** /groups/{groupId} | Update Group
*GroupsApi* | [**UpdateGroupGallery**](docs/GroupsApi.md#updategroupgallery) | **Put** /groups/{groupId}/galleries/{groupGalleryId} | Update Group Gallery
*GroupsApi* | [**UpdateGroupMember**](docs/GroupsApi.md#updategroupmember) | **Put** /groups/{groupId}/members/{userId} | Update Group Member
*GroupsApi* | [**UpdateGroupRole**](docs/GroupsApi.md#updategrouprole) | **Put** /groups/{groupId}/roles/{groupRoleId} | Update Group Role
*InstancesApi* | [**GetInstance**](docs/InstancesApi.md#getinstance) | **Get** /instances/{worldId}:{instanceId} | Get Instance
*InstancesApi* | [**GetInstanceByShortName**](docs/InstancesApi.md#getinstancebyshortname) | **Get** /instances/s/{shortName} | Get Instance By Short Name
*InstancesApi* | [**GetShortName**](docs/InstancesApi.md#getshortname) | **Get** /instances/{worldId}:{instanceId}/shortName | Get Instance Short Name
*InstancesApi* | [**SendSelfInvite**](docs/InstancesApi.md#sendselfinvite) | **Post** /instances/{worldId}:{instanceId}/invite | Send Self Invite
*InviteApi* | [**GetInviteMessage**](docs/InviteApi.md#getinvitemessage) | **Get** /message/{userId}/{messageType}/{slot} | Get Invite Message
*InviteApi* | [**GetInviteMessages**](docs/InviteApi.md#getinvitemessages) | **Get** /message/{userId}/{messageType} | List Invite Messages
*InviteApi* | [**InviteMyselfTo**](docs/InviteApi.md#invitemyselfto) | **Post** /invite/myself/to/{worldId}:{instanceId} | Invite Myself To Instance
*InviteApi* | [**InviteUser**](docs/InviteApi.md#inviteuser) | **Post** /invite/{userId} | Invite User
*InviteApi* | [**RequestInvite**](docs/InviteApi.md#requestinvite) | **Post** /requestInvite/{userId} | Request Invite
*InviteApi* | [**ResetInviteMessage**](docs/InviteApi.md#resetinvitemessage) | **Delete** /message/{userId}/{messageType}/{slot} | Reset Invite Message
*InviteApi* | [**RespondInvite**](docs/InviteApi.md#respondinvite) | **Post** /invite/{notificationId}/response | Respond Invite
*InviteApi* | [**UpdateInviteMessage**](docs/InviteApi.md#updateinvitemessage) | **Put** /message/{userId}/{messageType}/{slot} | Update Invite Message
*NotificationsApi* | [**AcceptFriendRequest**](docs/NotificationsApi.md#acceptfriendrequest) | **Put** /auth/user/notifications/{notificationId}/accept | Accept Friend Request
*NotificationsApi* | [**ClearNotifications**](docs/NotificationsApi.md#clearnotifications) | **Put** /auth/user/notifications/clear | Clear All Notifications
*NotificationsApi* | [**DeleteNotification**](docs/NotificationsApi.md#deletenotification) | **Put** /auth/user/notifications/{notificationId}/hide | Delete Notification
*NotificationsApi* | [**GetNotifications**](docs/NotificationsApi.md#getnotifications) | **Get** /auth/user/notifications | List Notifications
*NotificationsApi* | [**MarkNotificationAsRead**](docs/NotificationsApi.md#marknotificationasread) | **Put** /auth/user/notifications/{notificationId}/see | Mark Notification As Read
*PermissionsApi* | [**GetAssignedPermissions**](docs/PermissionsApi.md#getassignedpermissions) | **Get** /auth/permissions | Get Assigned Permissions
*PermissionsApi* | [**GetPermission**](docs/PermissionsApi.md#getpermission) | **Get** /permissions/{permissionId} | Get Permission
*PlayermoderationApi* | [**ClearAllPlayerModerations**](docs/PlayermoderationApi.md#clearallplayermoderations) | **Delete** /auth/user/playermoderations | Clear All Player Moderations
*PlayermoderationApi* | [**DeletePlayerModeration**](docs/PlayermoderationApi.md#deleteplayermoderation) | **Delete** /auth/user/playermoderations/{playerModerationId} | Delete Player Moderation
*PlayermoderationApi* | [**GetPlayerModeration**](docs/PlayermoderationApi.md#getplayermoderation) | **Get** /auth/user/playermoderations/{playerModerationId} | Get Player Moderation
*PlayermoderationApi* | [**GetPlayerModerations**](docs/PlayermoderationApi.md#getplayermoderations) | **Get** /auth/user/playermoderations | Search Player Moderations
*PlayermoderationApi* | [**ModerateUser**](docs/PlayermoderationApi.md#moderateuser) | **Post** /auth/user/playermoderations | Moderate User
*PlayermoderationApi* | [**UnmoderateUser**](docs/PlayermoderationApi.md#unmoderateuser) | **Put** /auth/user/unplayermoderate | Unmoderate User
*SystemApi* | [**GetCSS**](docs/SystemApi.md#getcss) | **Get** /css/app.css | Download CSS
*SystemApi* | [**GetConfig**](docs/SystemApi.md#getconfig) | **Get** /config | Fetch API Config
*SystemApi* | [**GetCurrentOnlineUsers**](docs/SystemApi.md#getcurrentonlineusers) | **Get** /visits | Current Online Users
*SystemApi* | [**GetHealth**](docs/SystemApi.md#gethealth) | **Get** /health | Check API Health
*SystemApi* | [**GetInfoPush**](docs/SystemApi.md#getinfopush) | **Get** /infoPush | Show Information Notices
*SystemApi* | [**GetJavaScript**](docs/SystemApi.md#getjavascript) | **Get** /js/app.js | Download JavaScript
*SystemApi* | [**GetSystemTime**](docs/SystemApi.md#getsystemtime) | **Get** /time | Current System Time
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{userId} | Get User by ID
*UsersApi* | [**GetUserByName**](docs/UsersApi.md#getuserbyname) | **Get** /users/{username}/name | Get User by Username
*UsersApi* | [**GetUserGroupRequests**](docs/UsersApi.md#getusergrouprequests) | **Get** /users/{userId}/groups/requested | Get User Group Requests
*UsersApi* | [**GetUserGroups**](docs/UsersApi.md#getusergroups) | **Get** /users/{userId}/groups | Get User Groups
*UsersApi* | [**SearchUsers**](docs/UsersApi.md#searchusers) | **Get** /users | Search All Users
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /users/{userId} | Update User Info
*WorldsApi* | [**CreateWorld**](docs/WorldsApi.md#createworld) | **Post** /worlds | Create World
*WorldsApi* | [**DeleteWorld**](docs/WorldsApi.md#deleteworld) | **Delete** /worlds/{worldId} | Delete World
*WorldsApi* | [**GetActiveWorlds**](docs/WorldsApi.md#getactiveworlds) | **Get** /worlds/active | List Active Worlds
*WorldsApi* | [**GetFavoritedWorlds**](docs/WorldsApi.md#getfavoritedworlds) | **Get** /worlds/favorites | List Favorited Worlds
*WorldsApi* | [**GetRecentWorlds**](docs/WorldsApi.md#getrecentworlds) | **Get** /worlds/recent | List Recent Worlds
*WorldsApi* | [**GetWorld**](docs/WorldsApi.md#getworld) | **Get** /worlds/{worldId} | Get World by ID
*WorldsApi* | [**GetWorldInstance**](docs/WorldsApi.md#getworldinstance) | **Get** /worlds/{worldId}/{instanceId} | Get World Instance
*WorldsApi* | [**GetWorldMetadata**](docs/WorldsApi.md#getworldmetadata) | **Get** /worlds/{worldId}/metadata | Get World Metadata
*WorldsApi* | [**GetWorldPublishStatus**](docs/WorldsApi.md#getworldpublishstatus) | **Get** /worlds/{worldId}/publish | Get World Publish Status
*WorldsApi* | [**PublishWorld**](docs/WorldsApi.md#publishworld) | **Put** /worlds/{worldId}/publish | Publish World
*WorldsApi* | [**SearchWorlds**](docs/WorldsApi.md#searchworlds) | **Get** /worlds | Search All Worlds
*WorldsApi* | [**UnpublishWorld**](docs/WorldsApi.md#unpublishworld) | **Delete** /worlds/{worldId}/publish | Unpublish World
*WorldsApi* | [**UpdateWorld**](docs/WorldsApi.md#updateworld) | **Put** /worlds/{worldId} | Update World


## Documentation For Models

 - [APIConfig](docs/APIConfig.md)
 - [APIConfigAnnouncement](docs/APIConfigAnnouncement.md)
 - [APIConfigDownloadURLList](docs/APIConfigDownloadURLList.md)
 - [APIConfigEvents](docs/APIConfigEvents.md)
 - [APIHealth](docs/APIHealth.md)
 - [AccountDeletionLog](docs/AccountDeletionLog.md)
 - [AddFavoriteRequest](docs/AddFavoriteRequest.md)
 - [AddGroupGalleryImageRequest](docs/AddGroupGalleryImageRequest.md)
 - [Avatar](docs/Avatar.md)
 - [AvatarUnityPackageUrlObject](docs/AvatarUnityPackageUrlObject.md)
 - [BanGroupMemberRequest](docs/BanGroupMemberRequest.md)
 - [CreateAvatarRequest](docs/CreateAvatarRequest.md)
 - [CreateFileRequest](docs/CreateFileRequest.md)
 - [CreateFileVersionRequest](docs/CreateFileVersionRequest.md)
 - [CreateGroupAnnouncementRequest](docs/CreateGroupAnnouncementRequest.md)
 - [CreateGroupGalleryRequest](docs/CreateGroupGalleryRequest.md)
 - [CreateGroupInviteRequest](docs/CreateGroupInviteRequest.md)
 - [CreateGroupRequest](docs/CreateGroupRequest.md)
 - [CreateGroupRoleRequest](docs/CreateGroupRoleRequest.md)
 - [CreateWorldRequest](docs/CreateWorldRequest.md)
 - [CurrentUser](docs/CurrentUser.md)
 - [DeploymentGroup](docs/DeploymentGroup.md)
 - [DeveloperType](docs/DeveloperType.md)
 - [DynamicContentRow](docs/DynamicContentRow.md)
 - [Error](docs/Error.md)
 - [Favorite](docs/Favorite.md)
 - [FavoriteGroup](docs/FavoriteGroup.md)
 - [FavoriteGroupVisibility](docs/FavoriteGroupVisibility.md)
 - [FavoriteType](docs/FavoriteType.md)
 - [File](docs/File.md)
 - [FileData](docs/FileData.md)
 - [FileStatus](docs/FileStatus.md)
 - [FileUploadURL](docs/FileUploadURL.md)
 - [FileVersion](docs/FileVersion.md)
 - [FileVersionUploadStatus](docs/FileVersionUploadStatus.md)
 - [FinishFileDataUploadRequest](docs/FinishFileDataUploadRequest.md)
 - [FriendStatus](docs/FriendStatus.md)
 - [Group](docs/Group.md)
 - [GroupAnnouncement](docs/GroupAnnouncement.md)
 - [GroupAuditLogEntry](docs/GroupAuditLogEntry.md)
 - [GroupGallery](docs/GroupGallery.md)
 - [GroupGalleryImage](docs/GroupGalleryImage.md)
 - [GroupJoinState](docs/GroupJoinState.md)
 - [GroupLimitedMember](docs/GroupLimitedMember.md)
 - [GroupMember](docs/GroupMember.md)
 - [GroupMemberLimitedUser](docs/GroupMemberLimitedUser.md)
 - [GroupMemberStatus](docs/GroupMemberStatus.md)
 - [GroupMyMember](docs/GroupMyMember.md)
 - [GroupPermission](docs/GroupPermission.md)
 - [GroupPrivacy](docs/GroupPrivacy.md)
 - [GroupRole](docs/GroupRole.md)
 - [GroupRoleTemplate](docs/GroupRoleTemplate.md)
 - [GroupUserVisibility](docs/GroupUserVisibility.md)
 - [InfoPush](docs/InfoPush.md)
 - [InfoPushData](docs/InfoPushData.md)
 - [InfoPushDataArticle](docs/InfoPushDataArticle.md)
 - [InfoPushDataArticleContent](docs/InfoPushDataArticleContent.md)
 - [InfoPushDataClickable](docs/InfoPushDataClickable.md)
 - [Instance](docs/Instance.md)
 - [InstancePlatforms](docs/InstancePlatforms.md)
 - [InstanceShortNameResponse](docs/InstanceShortNameResponse.md)
 - [InstanceType](docs/InstanceType.md)
 - [InviteMessage](docs/InviteMessage.md)
 - [InviteMessageType](docs/InviteMessageType.md)
 - [InviteRequest](docs/InviteRequest.md)
 - [InviteResponse](docs/InviteResponse.md)
 - [License](docs/License.md)
 - [LicenseAction](docs/LicenseAction.md)
 - [LicenseGroup](docs/LicenseGroup.md)
 - [LicenseType](docs/LicenseType.md)
 - [LimitedUnityPackage](docs/LimitedUnityPackage.md)
 - [LimitedUser](docs/LimitedUser.md)
 - [LimitedWorld](docs/LimitedWorld.md)
 - [MIMEType](docs/MIMEType.md)
 - [ModerateUserRequest](docs/ModerateUserRequest.md)
 - [Notification](docs/Notification.md)
 - [NotificationType](docs/NotificationType.md)
 - [OrderOption](docs/OrderOption.md)
 - [PaginatedGroupAuditLogEntryList](docs/PaginatedGroupAuditLogEntryList.md)
 - [PastDisplayName](docs/PastDisplayName.md)
 - [Permission](docs/Permission.md)
 - [PlayerModeration](docs/PlayerModeration.md)
 - [PlayerModerationType](docs/PlayerModerationType.md)
 - [Region](docs/Region.md)
 - [ReleaseStatus](docs/ReleaseStatus.md)
 - [RequestInviteRequest](docs/RequestInviteRequest.md)
 - [RespondGroupJoinRequest](docs/RespondGroupJoinRequest.md)
 - [Response](docs/Response.md)
 - [SentNotification](docs/SentNotification.md)
 - [SortOption](docs/SortOption.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionPeriod](docs/SubscriptionPeriod.md)
 - [Success](docs/Success.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAgreement](docs/TransactionAgreement.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionSteamInfo](docs/TransactionSteamInfo.md)
 - [TransactionSteamWalletInfo](docs/TransactionSteamWalletInfo.md)
 - [TwoFactorAuthCode](docs/TwoFactorAuthCode.md)
 - [TwoFactorEmailCode](docs/TwoFactorEmailCode.md)
 - [UnityPackage](docs/UnityPackage.md)
 - [UpdateAvatarRequest](docs/UpdateAvatarRequest.md)
 - [UpdateFavoriteGroupRequest](docs/UpdateFavoriteGroupRequest.md)
 - [UpdateGroupGalleryRequest](docs/UpdateGroupGalleryRequest.md)
 - [UpdateGroupMemberRequest](docs/UpdateGroupMemberRequest.md)
 - [UpdateGroupRequest](docs/UpdateGroupRequest.md)
 - [UpdateGroupRoleRequest](docs/UpdateGroupRoleRequest.md)
 - [UpdateInviteMessageRequest](docs/UpdateInviteMessageRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateWorldRequest](docs/UpdateWorldRequest.md)
 - [User](docs/User.md)
 - [UserExists](docs/UserExists.md)
 - [UserState](docs/UserState.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserSubscription](docs/UserSubscription.md)
 - [Verify2FAEmailCodeResult](docs/Verify2FAEmailCodeResult.md)
 - [Verify2FAResult](docs/Verify2FAResult.md)
 - [VerifyAuthTokenResult](docs/VerifyAuthTokenResult.md)
 - [World](docs/World.md)
 - [WorldMetadata](docs/WorldMetadata.md)
 - [WorldPublishStatus](docs/WorldPublishStatus.md)


## Documentation For Authorization



### apiKeyCookie

- **Type**: API key
- **API key parameter name**: apiKey
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apiKey and passed in as the auth context for each request.


### apiKeyQuery

- **Type**: API key
- **API key parameter name**: apiKey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apiKey and passed in as the auth context for each request.


### authCookie

- **Type**: API key
- **API key parameter name**: auth
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: auth and passed in as the auth context for each request.


### authHeader

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### twoFactorAuthCookie

- **Type**: API key
- **API key parameter name**: twoFactorAuth
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: twoFactorAuth and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

me@ariesclark.com

