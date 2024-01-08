package handler

import (
	"github.com/gin-gonic/gin"
	todo_app "github.com/todo-app"
	"net/http"
	"strconv"
)

// @Summary Create todo_app list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo_app list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo_app.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo_app.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo_app.TodoList `json:"data"`
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todo_app.ListsItem
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update List
// @Security ApiKeyAuth
// @Tags lists
// @Description update list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body todo_app.UpdateListInput true "update list"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo_app.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete List
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
