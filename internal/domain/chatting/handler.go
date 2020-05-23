package chatting

import (
	"Embassy/internal/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	Chatting(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindByID(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type Client struct {
	Websocket *websocket.Conn
	UserID string
}

type handler struct {
	service Service
	//clients map[*websocket.Conn]bool
	//broadcast chan Chat

}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

var clients = make(map[*websocket.Conn]uuid.UUID)
var broadcast = make(chan Chat)

func HandleMessages(){
	for {
		msg := <- broadcast
		logrus.Println("This is the message", msg)

		// Send it out to every client that is currently connected
		for client := range clients {
			logrus.Println("This is client", clients[client])

			if clients[client] == msg.To {
				err := client.WriteJSON(msg)
				if err != nil {
					logrus.Println("This is the err", err.Error())
					_ = client.Close()
					delete(clients, client)
				}
			}
		}
	}
}


func (s *handler) Chatting(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var chat Chat
	userID := mux.Vars(r)["userID"]
	ids, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		logrus.Println(err.Error())
	}

	// create handshake with client
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		logrus.Println("Upgrade error", err.Error())
	}

	defer ws.Close()

	clients[ws] = ids[0]

	for {
		// Read messages from all clients
		if err := ws.ReadJSON(&chat); err != nil{
			logrus.Println(&chat)
			delete(clients, ws)
			break
		}

		if _, err = s.service.Create(&chat); err != nil {
			logrus.Println(err.Error())
		}

		broadcast <- chat
	}
}

func (s *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	result, err := s.service.FindAll()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

func (s *handler) FindByID(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	// get messages of current user and requested users
	var requestedUserID = mux.Vars(r)["requestedUserID"]
	userDetails, _ := helpers.VerifyToken(r)

	ids, err := helpers.ParseIDs([]string{requestedUserID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	chat := &Chat{
		From:userDetails.ID ,
		To:ids[0],
	}

	result, err := s.service.FindByID(chat)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var chat Chat
	chatID := mux.Vars(r)["chatID"]

	parsedChatID, err := uuid.FromString(chatID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chat.ID = parsedChatID

	result, err := s.service.Update(&chat)
	if err != nil{
		if err.Error() == "is not owner" {
			helpers.ErrorResponse(w, http.StatusForbidden,
				"failed to perform action, please contact administration for help")
			return
		}
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var chat Chat
	chatID := mux.Vars(r)["chatID"]

	parsedChatID, err := uuid.FromString(chatID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	chat.ID = parsedChatID

	if err = s.service.Delete(&chat); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, map[string]string{
		"message": "chat deleted successfully",
	})
	return
}

