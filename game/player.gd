extends CharacterBody3D

func _ready() -> void:
	SignalBus.player1_gamestate_update.connect(handleGamestate)
	

func handleGamestate(s: Dictionary):
	position.x = s["posx"]
	position.y = s["posy"]
	position.z = s["posz"]
	
func _unhandled_input(event: InputEvent) -> void:
	if event is InputEventMouseButton:
		Input.set_mouse_mode(Input.MOUSE_MODE_CAPTURED)
	elif event.is_action_pressed("ui_cancel"):
		Input.set_mouse_mode(Input.MOUSE_MODE_VISIBLE)
	if Input.get_mouse_mode() == Input.MOUSE_MODE_CAPTURED:
		if event is InputEventMouseMotion:
			rotation.y += -event.relative.x * 0.001
			rotation.x += -event.relative.y * 0.001
			var rot_quat = Quaternion.from_euler(rotation)
			

func _physics_process(delta: float) -> void:
	pass
