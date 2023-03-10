package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"todo/logger"
	"todo/model"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	rows, err := h.rp.DB.Query("SELECT * FROM todo ORDER BY id desc")
	if err != nil {
		logger.Error(err)
		return
	}

	var data []model.DBData
	for rows.Next() {
		dbData := model.DBData{}
		if err := rows.Scan(&dbData.ID, &dbData.Name, &dbData.Content, &dbData.Time); err != nil {
			logger.Error(err)
			return
		}
		data = append(data, dbData)
	}

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		logger.Error(err)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//TODO: 一旦このまま進めたい。後ほど確認予定
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ins, err := h.rp.DB.Prepare("INSERT INTO todo(name,content,time) VALUES(?,?,?)")
	if err != nil {
		logger.Error(err)
		return
	}

	if err := ins.Exec(r.Form.Get("name"), r.Form.Get("content"), time.Now()); err != nil {
		logger.Error(err)
		return
	}

	if err := json.NewEncoder(w).Encode(&ins); err != nil {
		logger.Error(err)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	//TODO: 一旦このまま進めたい。後ほど確認予定
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	upd, err := h.rp.DB.Prepare("UPDATE todo SET name=?,content=? WHERE id=?")
	if err != nil {
		logger.Error(err)
		return
	}

	if err := upd.Exec(r.Form.Get("name"), r.Form.Get("content"), r.Form.Get("id")); err != nil {
		logger.Error(err)
		return
	}

	if err := json.NewEncoder(w).Encode(&upd); err != nil {
		logger.Error(err)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO: 一旦このまま進めたい。後ほど確認予定
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	del, err := h.rp.DB.Prepare("DELETE FROM todo WHERE id=?")
	if err != nil {
		logger.Error(err)
		return
	}

	if err := del.Exec(r.Form.Get("id")); err != nil {
		logger.Error(err)
		return
	}

	if err := json.NewEncoder(w).Encode(&del); err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("content delete!")
}
