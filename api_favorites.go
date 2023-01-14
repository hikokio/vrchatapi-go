/*
VRChat API Documentation


API version: 1.10.1
Contact: me@ariesclark.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// FavoritesApiService FavoritesApi service
type FavoritesApiService service

type ApiAddFavoriteRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	addFavoriteRequest *AddFavoriteRequest
}

// 
func (r ApiAddFavoriteRequest) AddFavoriteRequest(addFavoriteRequest AddFavoriteRequest) ApiAddFavoriteRequest {
	r.addFavoriteRequest = &addFavoriteRequest
	return r
}

func (r ApiAddFavoriteRequest) Execute() (*Favorite, *http.Response, error) {
	return r.ApiService.AddFavoriteExecute(r)
}

/*
AddFavorite Add Favorite

Add a new favorite.

Friend groups are named `group_0` through `group_3`. Avatar and World groups are named `avatars1` to `avatars4` and `worlds1` to `worlds4`.

You cannot add people whom you are not friends with to your friends list. Destroying a friendship removes the person as favorite on both sides.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddFavoriteRequest
*/
func (a *FavoritesApiService) AddFavorite(ctx context.Context) ApiAddFavoriteRequest {
	return ApiAddFavoriteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Favorite
func (a *FavoritesApiService) AddFavoriteExecute(r ApiAddFavoriteRequest) (*Favorite, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Favorite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.AddFavorite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addFavoriteRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiClearFavoriteGroupRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	favoriteGroupType string
	favoriteGroupName string
	userId string
}

func (r ApiClearFavoriteGroupRequest) Execute() (*Success, *http.Response, error) {
	return r.ApiService.ClearFavoriteGroupExecute(r)
}

/*
ClearFavoriteGroup Clear Favorite Group

Clear ALL contents of a specific favorite group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param favoriteGroupType The type of group to fetch, must be a valid FavoriteType.
 @param favoriteGroupName The name of the group to fetch, must be a name of a FavoriteGroup.
 @param userId Must be a valid user ID.
 @return ApiClearFavoriteGroupRequest
*/
func (a *FavoritesApiService) ClearFavoriteGroup(ctx context.Context, favoriteGroupType string, favoriteGroupName string, userId string) ApiClearFavoriteGroupRequest {
	return ApiClearFavoriteGroupRequest{
		ApiService: a,
		ctx: ctx,
		favoriteGroupType: favoriteGroupType,
		favoriteGroupName: favoriteGroupName,
		userId: userId,
	}
}

// Execute executes the request
//  @return Success
func (a *FavoritesApiService) ClearFavoriteGroupExecute(r ApiClearFavoriteGroupRequest) (*Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.ClearFavoriteGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupType"+"}", url.PathEscape(parameterToString(r.favoriteGroupType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupName"+"}", url.PathEscape(parameterToString(r.favoriteGroupName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFavoriteRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	favoriteId string
}

func (r ApiGetFavoriteRequest) Execute() (*Favorite, *http.Response, error) {
	return r.ApiService.GetFavoriteExecute(r)
}

/*
GetFavorite Show Favorite

Return information about a specific Favorite.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param favoriteId Must be a valid favorite ID.
 @return ApiGetFavoriteRequest
*/
func (a *FavoritesApiService) GetFavorite(ctx context.Context, favoriteId string) ApiGetFavoriteRequest {
	return ApiGetFavoriteRequest{
		ApiService: a,
		ctx: ctx,
		favoriteId: favoriteId,
	}
}

// Execute executes the request
//  @return Favorite
func (a *FavoritesApiService) GetFavoriteExecute(r ApiGetFavoriteRequest) (*Favorite, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Favorite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.GetFavorite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorites/{favoriteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteId"+"}", url.PathEscape(parameterToString(r.favoriteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFavoriteGroupRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	favoriteGroupType string
	favoriteGroupName string
	userId string
}

func (r ApiGetFavoriteGroupRequest) Execute() (*FavoriteGroup, *http.Response, error) {
	return r.ApiService.GetFavoriteGroupExecute(r)
}

/*
GetFavoriteGroup Show Favorite Group

Fetch information about a specific favorite group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param favoriteGroupType The type of group to fetch, must be a valid FavoriteType.
 @param favoriteGroupName The name of the group to fetch, must be a name of a FavoriteGroup.
 @param userId Must be a valid user ID.
 @return ApiGetFavoriteGroupRequest
*/
func (a *FavoritesApiService) GetFavoriteGroup(ctx context.Context, favoriteGroupType string, favoriteGroupName string, userId string) ApiGetFavoriteGroupRequest {
	return ApiGetFavoriteGroupRequest{
		ApiService: a,
		ctx: ctx,
		favoriteGroupType: favoriteGroupType,
		favoriteGroupName: favoriteGroupName,
		userId: userId,
	}
}

// Execute executes the request
//  @return FavoriteGroup
func (a *FavoritesApiService) GetFavoriteGroupExecute(r ApiGetFavoriteGroupRequest) (*FavoriteGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FavoriteGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.GetFavoriteGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupType"+"}", url.PathEscape(parameterToString(r.favoriteGroupType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupName"+"}", url.PathEscape(parameterToString(r.favoriteGroupName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFavoriteGroupsRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	n *int32
	offset *int32
	ownerId *string
}

// The number of objects to return.
func (r ApiGetFavoriteGroupsRequest) N(n int32) ApiGetFavoriteGroupsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFavoriteGroupsRequest) Offset(offset int32) ApiGetFavoriteGroupsRequest {
	r.offset = &offset
	return r
}

// The owner of whoms favorite groups to return. Must be a UserID.
func (r ApiGetFavoriteGroupsRequest) OwnerId(ownerId string) ApiGetFavoriteGroupsRequest {
	r.ownerId = &ownerId
	return r
}

func (r ApiGetFavoriteGroupsRequest) Execute() ([]FavoriteGroup, *http.Response, error) {
	return r.ApiService.GetFavoriteGroupsExecute(r)
}

/*
GetFavoriteGroups List Favorite Groups

Return a list of favorite groups owned by a user. Returns the same information as `getFavoriteGroups`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFavoriteGroupsRequest
*/
func (a *FavoritesApiService) GetFavoriteGroups(ctx context.Context) ApiGetFavoriteGroupsRequest {
	return ApiGetFavoriteGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []FavoriteGroup
func (a *FavoritesApiService) GetFavoriteGroupsExecute(r ApiGetFavoriteGroupsRequest) ([]FavoriteGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []FavoriteGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.GetFavoriteGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorite/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.n != nil {
		localVarQueryParams.Add("n", parameterToString(*r.n, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ownerId != nil {
		localVarQueryParams.Add("ownerId", parameterToString(*r.ownerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFavoritesRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	n *int32
	offset *int32
	type_ *string
	tag *string
}

// The number of objects to return.
func (r ApiGetFavoritesRequest) N(n int32) ApiGetFavoritesRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFavoritesRequest) Offset(offset int32) ApiGetFavoritesRequest {
	r.offset = &offset
	return r
}

// The type of favorites to return, FavoriteType.
func (r ApiGetFavoritesRequest) Type_(type_ string) ApiGetFavoritesRequest {
	r.type_ = &type_
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiGetFavoritesRequest) Tag(tag string) ApiGetFavoritesRequest {
	r.tag = &tag
	return r
}

func (r ApiGetFavoritesRequest) Execute() ([]Favorite, *http.Response, error) {
	return r.ApiService.GetFavoritesExecute(r)
}

/*
GetFavorites List Favorites

Returns a list of favorites.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFavoritesRequest
*/
func (a *FavoritesApiService) GetFavorites(ctx context.Context) ApiGetFavoritesRequest {
	return ApiGetFavoritesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Favorite
func (a *FavoritesApiService) GetFavoritesExecute(r ApiGetFavoritesRequest) ([]Favorite, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Favorite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.GetFavorites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.n != nil {
		localVarQueryParams.Add("n", parameterToString(*r.n, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.tag != nil {
		localVarQueryParams.Add("tag", parameterToString(*r.tag, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveFavoriteRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	favoriteId string
}

func (r ApiRemoveFavoriteRequest) Execute() (*Success, *http.Response, error) {
	return r.ApiService.RemoveFavoriteExecute(r)
}

/*
RemoveFavorite Remove Favorite

Remove a favorite from your favorites list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param favoriteId Must be a valid favorite ID.
 @return ApiRemoveFavoriteRequest
*/
func (a *FavoritesApiService) RemoveFavorite(ctx context.Context, favoriteId string) ApiRemoveFavoriteRequest {
	return ApiRemoveFavoriteRequest{
		ApiService: a,
		ctx: ctx,
		favoriteId: favoriteId,
	}
}

// Execute executes the request
//  @return Success
func (a *FavoritesApiService) RemoveFavoriteExecute(r ApiRemoveFavoriteRequest) (*Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.RemoveFavorite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorites/{favoriteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteId"+"}", url.PathEscape(parameterToString(r.favoriteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateFavoriteGroupRequest struct {
	ctx context.Context
	ApiService *FavoritesApiService
	favoriteGroupType string
	favoriteGroupName string
	userId string
	updateFavoriteGroupRequest *UpdateFavoriteGroupRequest
}

func (r ApiUpdateFavoriteGroupRequest) UpdateFavoriteGroupRequest(updateFavoriteGroupRequest UpdateFavoriteGroupRequest) ApiUpdateFavoriteGroupRequest {
	r.updateFavoriteGroupRequest = &updateFavoriteGroupRequest
	return r
}

func (r ApiUpdateFavoriteGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateFavoriteGroupExecute(r)
}

/*
UpdateFavoriteGroup Update Favorite Group

Update information about a specific favorite group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param favoriteGroupType The type of group to fetch, must be a valid FavoriteType.
 @param favoriteGroupName The name of the group to fetch, must be a name of a FavoriteGroup.
 @param userId Must be a valid user ID.
 @return ApiUpdateFavoriteGroupRequest
*/
func (a *FavoritesApiService) UpdateFavoriteGroup(ctx context.Context, favoriteGroupType string, favoriteGroupName string, userId string) ApiUpdateFavoriteGroupRequest {
	return ApiUpdateFavoriteGroupRequest{
		ApiService: a,
		ctx: ctx,
		favoriteGroupType: favoriteGroupType,
		favoriteGroupName: favoriteGroupName,
		userId: userId,
	}
}

// Execute executes the request
func (a *FavoritesApiService) UpdateFavoriteGroupExecute(r ApiUpdateFavoriteGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoritesApiService.UpdateFavoriteGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupType"+"}", url.PathEscape(parameterToString(r.favoriteGroupType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"favoriteGroupName"+"}", url.PathEscape(parameterToString(r.favoriteGroupName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateFavoriteGroupRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
