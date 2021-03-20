package dashboard

import (
	"strings"
	"sync"

	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/publisher"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// ChainID -> *sync.Map{} (map of connected clients)
var wsClients = sync.Map{}

func addWsEndpoints(e *echo.Echo) {
	route := e.GET("/chain/:chainid/ws", handleWebSocket)
	route.Name = "chainWs"
}

func handleWebSocket(c echo.Context) error {
	chainID, err := coretypes.NewChainIDFromBase58(c.Param("chainid"))
	if err != nil {
		return err
	}

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		c.Logger().Infof("[WebSocket] opened for %s", c.Request().RemoteAddr)
		defer c.Logger().Infof("[WebSocket] closed for %s", c.Request().RemoteAddr)

		v, _ := wsClients.LoadOrStore(chainID.String(), &sync.Map{})
		chainWsClients := v.(*sync.Map)

		clientCh := make(chan string)
		chainWsClients.Store(clientCh, clientCh)
		defer chainWsClients.Delete(clientCh)

		for {
			msg := <-clientCh
			_, err := ws.Write([]byte(msg))
			if err != nil {
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func startWsForwarder() {
	publisher.Event.Attach(events.NewClosure(func(msgType string, parts []string) {
		if msgType == "state" {
			if len(parts) < 1 {
				return
			}
			chainID := parts[0]

			v, ok := wsClients.Load(chainID)
			if !ok {
				return
			}
			chainWsClients := v.(*sync.Map)

			msg := msgType + " " + strings.Join(parts, " ")
			chainWsClients.Range(func(key interface{}, clientCh interface{}) bool {
				clientCh.(chan string) <- msg
				return true
			})
		}
	}))
}

const tplWs = `
{{define "ws"}}
	<script>
		const url = 'ws://' +  location.host + '{{ uri "chainWs" . }}';
		console.log('opening WebSocket to ' + url);
		const ws = new WebSocket(url);

		ws.addEventListener('error', function (event) {
			console.error('WebSocket error!', event);
		});

		const connectedAt = new Date();
		ws.addEventListener('message', function (event) {
			console.log('Message from server: ', event.data);
			ws.close();
			if (new Date() - connectedAt > 5000) {
				location.reload();
			} else {
				setTimeout(() => location.reload(), 5000);
			}
		});
	</script>
{{end}}
`
