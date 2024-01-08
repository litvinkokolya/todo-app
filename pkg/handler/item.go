package handler

import (
	"github.com/gin-gonic/gin"
	todo_app "github.com/todo-app"
	"net/http"
	"strconv"
)

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
