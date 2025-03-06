extends CharacterBody3D

func _ready() -> void:
	SignalBus.player1_gamestate_update.connect(handleGamestate)
	

func handleGamestate(s: Dictionary):
	position.x = s["posx"]
	position.y = s["posy"]
	position.z = s["posz"]

func _physics_process(delta: float) -> void:
	pass
