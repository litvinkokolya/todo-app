package handler

import (
	"github.com/gin-gonic/gin"
	todo_app "github.com/todo-app"
	"net/http"
	"strconv"
)

// @Summary Create todo_app item
// @Security ApiKeyAuth
// @Tags items
// @Description create todo_app item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body todo_app.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists/:id/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}

	var input todo_app.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	id, err := h.services.TodoItem.Create(userId, listid, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get todo_app items
// @Security ApiKeyAuth
// @Tags items
// @Description get todo_app items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {integer} []todo_app.TodoItem
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/lists/:id/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todo_app.TodoItem
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Item
// @Security ApiKeyAuth
// @Tags items
// @Description update item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param input body todo_app.UpdateItemInput true "update item"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo_app.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Item
// @Security ApiKeyAuth
// @Tags items
// @Description delete item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	{
		userId, err := getUserId(c)
		if err != nil {
			return
		}

		itemId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
			return
		}

		err = h.services.TodoItem.Delete(userId, itemId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{"ok"})
	}
}
