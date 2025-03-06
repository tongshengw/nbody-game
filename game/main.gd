extends Node3D

var connected := false
var counter := 0

var subhub_allocated := false
var game_started := false

# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	WebsocketClient.connected_to_server.connect(connectHandler)
	WebsocketClient.connection_closed.connect(disconnectHandler)
	WebsocketClient.message_received.connect(msgHandler)

	WebsocketClient.connect_to_url("ws://127.0.0.1:9999/ws")


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta: float) -> void:
	if connected:
		pass

func _unhandled_key_input(event: InputEvent) -> void:
	if event.is_action_pressed("ui_accept"):
		if !subhub_allocated && connected:
			WebsocketClient.send(title_only_json("request_subhub"))
		elif subhub_allocated && !game_started:
			WebsocketClient.send(title_only_json("game_ready"))
			game_started = true

func title_only_json(input: String) -> String:
	return "{\"title\":\"%s\"}" % input

func receivedJsonHandler(input: Dictionary) -> void:
	match input["title"]:
		"subhub_allocated":
			subhub_allocated = true;
		"gamestate":
			var json = JSON.new()
			var error = json.parse(input["player1"])
			if error == OK:
				if typeof(json.data) == TYPE_DICTIONARY:
					print(json.data)
					SignalBus.player1_gamestate_update.emit(json.data)

func connectHandler() -> void:
	connected = true
	WebsocketClient.send("connected")
	print("connected")
func disconnectHandler() -> void:
	print("disconnected")
func msgHandler(message: Variant) -> void:
	var json = JSON.new()
	var error = json.parse(message)
	if error == OK:
		var data_received = json.data
		if typeof(data_received) == TYPE_DICTIONARY:
			print(data_received)
			receivedJsonHandler(data_received)
		else:
			print("Unexpected data")
	else:
		print("JSON Parse Error: ", json.get_error_message(), " in ", message, " at line ", json.get_error_line())
