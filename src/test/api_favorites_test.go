/*
VRChat API Documentation

Testing FavoritesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package vrchatapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_vrchatapi_FavoritesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test FavoritesApiService AddFavorite", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FavoritesApi.AddFavorite(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService ClearFavoriteGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var favoriteGroupType string
        var favoriteGroupName string
        var userId string

        resp, httpRes, err := apiClient.FavoritesApi.ClearFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService GetFavorite", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var favoriteId string

        resp, httpRes, err := apiClient.FavoritesApi.GetFavorite(context.Background(), favoriteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService GetFavoriteGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var favoriteGroupType string
        var favoriteGroupName string
        var userId string

        resp, httpRes, err := apiClient.FavoritesApi.GetFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService GetFavoriteGroups", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FavoritesApi.GetFavoriteGroups(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService GetFavorites", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FavoritesApi.GetFavorites(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService RemoveFavorite", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var favoriteId string

        resp, httpRes, err := apiClient.FavoritesApi.RemoveFavorite(context.Background(), favoriteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FavoritesApiService UpdateFavoriteGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var favoriteGroupType string
        var favoriteGroupName string
        var userId string

        resp, httpRes, err := apiClient.FavoritesApi.UpdateFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}