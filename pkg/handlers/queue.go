package handlers

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/juanvillacortac/bank-queue/pkg/auth"
	"github.com/juanvillacortac/bank-queue/pkg/database"
	"github.com/juanvillacortac/bank-queue/pkg/models"
	q "github.com/juanvillacortac/bank-queue/pkg/queue"
	"github.com/juanvillacortac/bank-queue/pkg/repositories"
)

type ClientOnQueue struct {
	models.Client
	AttendedBy         int64     `json:"attendedBy"`
	ArrivalTime        time.Time `json:"arrival"`
	RequiredOperations int       `json:"requiredOperations"`
}

func (c ClientOnQueue) Priority() int {
	return int(c.ClientType)
}

func (c ClientOnQueue) Arrival() time.Time {
	return c.ArrivalTime
}

var queue = &q.PriorityQueue{
	Items:    []q.QueueItem{},
	Capacity: 10,
}

// Upgrader se usa para actualizar la conexión HTTP a WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir todas las conexiones para este ejemplo
	},
}

// WSConnection representa una conexión WebSocket
type WSConnection struct {
	conn *websocket.Conn
	send chan ClientOnQueue // Canal para enviar mensajes al cliente
}

// Mapa para almacenar todos los clientes conectados
var wsConns = make(map[*WSConnection]bool)
var mutex = &sync.Mutex{}

var broadcast = make(chan ClientOnQueue)

func attendClients() {
	for {
		if queue.Len() > 0 {
			client := queue.Items[0].(*ClientOnQueue)

			timeToAttend := time.Duration(client.RequiredOperations) * time.Second

			log.Printf(
				"Atendiendo al cliente %s (DPI %s) con carga de %d operaciones. Tiempo: %v segundos",
				client.Name, client.DPI, client.RequiredOperations, timeToAttend,
			)

			time.Sleep(timeToAttend)

			repo := repositories.NewSQLHistoryRepository(database.Instance)
			if _, err := repo.RegisterEntry(client.ID, client.AttendedBy, client.RequiredOperations); err != nil {
				log.Println("Error ingresando cliente en el historial:", err.Error())
				return
			}

			heap.Pop(queue)

			broadcast <- *client
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		mutex.Lock()
		for conn := range wsConns {
			select {
			case conn.send <- message:
			default:
				close(conn.send)
				delete(wsConns, conn)
			}
		}
		mutex.Unlock()
	}
}

func queueRouter(r chi.Router) {
	go attendClients()
	go handleMessages()

	r.Use(auth.AuthMiddleware)

	r.Get("/", getQueueHandler)
	r.Get("/ws", webSocketsGetQueueHandler)
	r.Post("/", pushClientToQueueHandler)
}

func getQueueHandler(w http.ResponseWriter, r *http.Request) {
	respondWithPayload(w, http.StatusOK, queue)
}

func webSocketsGetQueueHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		respondWithPayload(w, http.StatusNotAcceptable, fmt.Errorf("error actualizando a WebSocket: %s", err.Error()))
		return
	}

	wsConn := &WSConnection{
		conn: conn,
		send: make(chan ClientOnQueue),
	}

	mutex.Lock()
	wsConns[wsConn] = true
	mutex.Unlock()

	if err := wsConn.conn.WriteJSON(queue); err != nil {
		log.Println("Error enviando mensaje al cliente:", err)
		wsConn.conn.Close()
		mutex.Lock()
		delete(wsConns, wsConn)
		mutex.Unlock()
		return
	}

	go func() {
		for {
			select {
			case <-wsConn.send:
				// Enviar el mensaje al cliente
				err := wsConn.conn.WriteJSON(queue)
				if err != nil {
					log.Println("Error enviando mensaje al cliente:", err)
					wsConn.conn.Close()
					mutex.Lock()
					delete(wsConns, wsConn)
					mutex.Unlock()
					return
				}
			}
		}
	}()
}

type PushClientRequest struct {
	DPI                string `json:"dpi"`
	RequiredOperations int    `json:"requiredOperations"`
}

func pushClientToQueueHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := auth.GetUserFromContext(r.Context())
	var request PushClientRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding JSON: ", err)
	}

	repo := repositories.NewSQLClientRepository(database.Instance)

	client, err := repo.GetClient(request.DPI)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}
	if client == nil {
		respondWithError(w, http.StatusUnauthorized, ErrorClientNotFound)
		return
	}

	for _, item := range queue.Items {
		client, ok := item.(*ClientOnQueue)
		if ok {
			if client.DPI == request.DPI {
				respondWithError(w, http.StatusUnauthorized, fmt.Errorf("ya el cliente se encuentra en la cola"))
				return
			}
		}
	}

	payload := ClientOnQueue{
		Client:             *client,
		AttendedBy:         user.ID,
		RequiredOperations: request.RequiredOperations,
		ArrivalTime:        time.Now(),
	}
	heap.Push(queue, &payload)

	broadcast <- payload

	respondWithPayload(w, http.StatusOK, queue)
}
