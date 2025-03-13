# nbody-game

This is a multiplayer game in go and godot purely for rendering. Server uses websockets.
A client manages one websocket connection and provides readers and writers in goroutines.
The hub manages all incoming connections and allocates clients to subhubs.
Subhubs are a game server, managing websocket messages from its allocated clients and passes them to the game.

### Backend connection flow
![backend flow diagram](https://github.com/tongshengw/nbody-game/blob/main/imgs/nbody_backend_diagram.svg)

#### TODO: 
- Fix issue where disconnecting doesn't correctly deallocate client and stop game. Close channel when game stops.
- Pass input from godot to game
- Figure out FPS/TPS interactions
- Start implementing gameplay elements

#### Future goals:
- Pixel art shader for procedually generated asteroids
