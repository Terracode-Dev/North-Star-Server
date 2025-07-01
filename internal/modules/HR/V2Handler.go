package hr 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	
	"github.com/labstack/echo/v4"
)

func (S *HRService) CreatePreset(c echo.Context) error {
	var req CreatePresetReqModel
	err := c.Bind(&req)
 	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	params := req.ToCreatePresetParams()
	err = S.q.CreatePreset(c.Request().Context(), params)
	if err != nil {
		log.Printf("Error creating preset: %v", err)
  	return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating preset: %v", err))
 	}
	return c.JSON(http.StatusOK, "Preset created successfully")
}

func (S *HRService) DeletePreset(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	err = S.q.DeletePreset(c.Request().Context(), id)
	if err != nil {
		log.Printf("Error deleting preset: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting preset: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset deleted successfully")
}

func (S *HRService) GetAllPresets(c echo.Context) error {
	rows, err := S.q.SelectAllPresets(c.Request().Context())
	if err != nil {
		log.Printf("Error fetching presets: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching presets: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}

func (S *HRService) GetPresetByID(c echo.Context) error {
	idStr, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	preset, err := S.q.SelectPreset(c.Request().Context(), idStr)
	if err != nil {
		log.Printf("Error fetching preset: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching preset: %v", err))
	}
	return c.JSON(http.StatusOK, preset)
}

func (S *HRService) UpdatePreset(c echo.Context) error {
	id , err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	var req CreatePresetReqModel
 	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	params := req.ToUpdatePresetParams(id)
 	err = S.q.UpdatePreset(c.Request().Context(), params)
	if err != nil {
		log.Printf("Error updating preset: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error updating preset: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset updated successfully")
}

func (S *HRService) GetPresetByName(c echo.Context) error {
	name := c.Param("name")
	search_term := "%" + name + "%"
	rows, err := S.q.SelectPresetByname(c.Request().Context(), search_term)
	if err != nil {
		log.Printf("Error fetching preset by name: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching preset by name: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}

//preset workout
func (S *HRService) CreatePresetWorkout(c echo.Context) error {
	var req CreatePresetWorkoutreqModel
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	params := req.ToCreatePresetWorkoutParams()
	err = S.q.CreatePresetWorkout(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating preset workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset workout created successfully")
}

func (S *HRService) DeletePresetWorkout(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	err = S.q.DeletePresetWorkout(c.Request().Context(), id)
	if err != nil {
		log.Printf("Error deleting preset workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting preset workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset workout deleted successfully")
}

func (S *HRService) GetAllPresetWorkouts(c echo.Context) error {
	rows, err := S.q.SelectAllPresetWorkouts(c.Request().Context())
	if err != nil {
		log.Printf("Error fetching preset workouts: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching preset workouts: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}


func (S *HRService) UpdatePresetWorkout(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	var req CreatePresetWorkoutreqModel
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	params := req.ToUpdatePresetWorkoutParams(id)
	err = S.q.UpdatePresetWorkout(c.Request().Context(), params)
	if err != nil {
		log.Printf("Error updating preset workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error updating preset workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset workout updated successfully")
}



// session 
func (S *HRService) CreateSession(c echo.Context) error {
	var req database.CreateSessionParams
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	err = S.q.CreateSession(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating session: %v", err))
	}
	return c.JSON(http.StatusOK, "Session created successfully")
}

func (S *HRService) DeleteSession(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	err = S.q.DeleteSession(c.Request().Context(), id)
	if err != nil {
		log.Printf("Error deleting session: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting session: %v", err))
	}
	return c.JSON(http.StatusOK, "Session deleted successfully")
}

func (S *HRService) GetAllSessions(c echo.Context) error {
	rows, err := S.q.SelectAllSessions(c.Request().Context())
	if err != nil {
		log.Printf("Error fetching sessions: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching sessions: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}

// preset session workout
func (S *HRService) CreatePresetSession(c echo.Context) error {
	var req database.CreatePresetSessionParams
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	err = S.q.CreatePresetSession(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating session workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Session workout created successfully")
}

func (S *HRService) DeletePresetSession(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	err = S.q.DeletePresetSession(c.Request().Context(), id)
	if err != nil {
		log.Printf("Error deleting session workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting session workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Session workout deleted successfully")
}

func (S *HRService) GetAllPresetSession(c echo.Context) error {
	rows, err := S.q.SelectpresetSessionAll(c.Request().Context())
	if err != nil {
		log.Printf("Error fetching preset sessions: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching preset sessions: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}

func (S *HRService) UpdatePresetSession(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	var req database.UpdatePresetSessionParams
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	req.ID = id
	err = S.q.UpdatePresetSession(c.Request().Context(), req)
	if err != nil {
		log.Printf("Error updating preset session: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error updating preset session: %v", err))
	}
	return c.JSON(http.StatusOK, "Preset session updated successfully")
}

// session workout
func (S *HRService) CreateSessionWorkout(c echo.Context) error {
	var req CreateSessionWorkoutReqModel
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	reqParams := req.ToCreateSessionWorkoutParams()

	err = S.q.CreateSessionWorkout(c.Request().Context(), reqParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating session workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Session workout created successfully")
}

func (S *HRService) DeleteSessionWorkout(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	err = S.q.DeleteSessionWorkout(c.Request().Context(), id)
	if err != nil {
	log.Printf("Error deleting session workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting session workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Session workout deleted successfully")
}

func (S *HRService) GetAllSessionWorkouts(c echo.Context) error {
	rows, err := S.q.SelectAllSessionWorkouts(c.Request().Context())
	if err != nil {
		log.Printf("Error fetching session workouts: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching session workouts: %v", err))
	}
	return c.JSON(http.StatusOK, rows)
}

func (S *HRService) GetSessionWorkoutByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	sessionWorkout, err := S.q.SelectSessionWorkout(c.Request().Context(), id)
	if err != nil {
		log.Printf("Error fetching session workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching session workout: %v", err))
	}
	return c.JSON(http.StatusOK, sessionWorkout)
}

func (S *HRService) UpdateSessionWorkout(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
	}
	var req CreateSessionWorkoutReqModel
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %v", err))
	}
	params := req.ToUpdateSessionWorkoutParams(id)
	params.ID = id
	err = S.q.UpdateSessionWorkout(c.Request().Context(), params)
	if err != nil {
		log.Printf("Error updating session workout: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error updating session workout: %v", err))
	}
	return c.JSON(http.StatusOK, "Session workout updated successfully")
}







