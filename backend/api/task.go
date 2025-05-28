package api

import (
	"app/pkg/model"
	"app/pkg/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateTask godoc
// @Summary      Создать задачу
// @Description  Метод позволяет записать одну задачу на пользователя
// @Tags         Task
// @Accept		 json
// @Produce      json
// @Param input body model.Task true "Задача"
// @Success      201
// @Router       api/task/ [post]
func CreateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Что-то с запросом", "err": err.Error()})
		return
	}

	res, err := storage.DB.InsertTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Не удалось записать.", "err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"result": res, "message": "Задача успешно записана"})
}

// DeleteTask godoc
// @Summary      Удалить задачу
// @Description  Метод позволяет удалить одну задачу
// @Tags         Task
// @Param id path int true "Задача"
// @Success      200
// @Router       api/task/{id} [delete]
func DeleteTask(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Что-то не так с параметром id.", "err": err.Error()})
	}

	res, err := storage.DB.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Не удалось удалить.", "err": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"result": res, "message": "Задача успешно удалена"})
}
