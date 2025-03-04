extends Node3D


# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	print("mainready")
	WebsocketClient.connect_to_url("localhost:9999/ws")
	WebsocketClient.send("diddy")
	print("msgsent")


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta: float) -> void:
	WebsocketClient.send("loop")
	print("loop")