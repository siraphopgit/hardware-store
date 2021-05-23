package controllers

import (
	"encoding/json"
	"hardware/api/models"
	"hardware/api/repository"
	"hardware/api/utils"
	"io/ioutil"
	"net/http"
)

type CategoriesController interface {
	PostCategory(http.ResponseWriter, *http.Request)
}

type categoriesControllerImpl struct {
	categoriesRepository repository.CategoriesRepository
}

func NewCategoriesRepository(categoriesRepsitory repository.CategoriesRepository) *categoriesControllerImpl {
	return &categoriesControllerImpl{categoriesRepsitory}
}

func (c *categoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	category, err = c.categoriesRepository.Save(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, category.ID))
	utils.WriteAsJson(w, category)
}
