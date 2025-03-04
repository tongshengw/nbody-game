extends Node3D

var connected := false
var counter := 0


# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	WebsocketClient.connected_to_server.connect(connectHandler)
	WebsocketClient.connection_closed.connect(disconnectHandler)
	WebsocketClient.message_received.connect(msgHandler)

	print("here")
	WebsocketClient.connect_to_url("ws://127.0.0.1:9999/ws")


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta: float) -> void:
	if connected:
		WebsocketClient.send(str(counter))
		counter += 1

func connectHandler() -> void:
	connected = true
	WebsocketClient.send("connected")
	print("connected")
func disconnectHandler() -> void:
	print("disconnected")
func msgHandler(message: Variant) -> void:
	print(message)
