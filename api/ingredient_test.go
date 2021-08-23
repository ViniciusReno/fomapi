package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ViniciusReno/fomapi/db/mock"
	db "github.com/ViniciusReno/fomapi/db/sqlc"
	"github.com/ViniciusReno/fomapi/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetIngredientRequest(t *testing.T) {
	ing := randomIngrentent()

	testCases := []struct {
		name          string
		ingredientId  int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:         "Ok",
			ingredientId: ing.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetIngreditent(gomock.Any(), gomock.Eq(ing.ID)).
					Times(1).
					Return(ing, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchIngredient(t, recorder.Body, ing)
			},
		},
		{
			name:         "NotFound",
			ingredientId: ing.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetIngreditent(gomock.Any(), gomock.Eq(ing.ID)).
					Times(1).
					Return(db.Ingredient{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:         "InternalError",
			ingredientId: ing.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetIngreditent(gomock.Any(), gomock.Eq(ing.ID)).
					Times(1).
					Return(db.Ingredient{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:         "InvalidID",
			ingredientId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetIngreditent(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer((store))
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/Ingredients/%d", tc.ingredientId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomIngrentent() db.Ingredient {
	return db.Ingredient{
		ID:    util.RandomInit(1, 1000),
		Nome:  util.RandomIngredient(),
		Ativo: true,
	}
}

func requiredBodyMatchIngredient(t *testing.T, body *bytes.Buffer, ing db.Ingredient) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotIng db.Ingredient
	err = json.Unmarshal(data, &gotIng)
	require.NoError(t, err)

	require.Equal(t, ing, gotIng)
}
