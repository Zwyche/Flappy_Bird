package main

import rl "github.com/gen2brain/raylib-go/raylib"

type PhysicsBody struct {
	Pos     rl.Vector2
	Vel     rl.Vector2
	Gravity rl.Vector2
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2) PhysicsBody {
	pb := PhysicsBody{Pos: newPos, Vel: newVel, Gravity: rl.Vector2Zero()}
	return pb
}

func (pb *PhysicsBody) PhysicsUpdate() {
	pb.GravityTick()
	pb.VelocityTick()
}

func (pb *PhysicsBody) GravityTick() {
	adjustedGravity := rl.Vector2Scale(pb.Gravity, rl.GetFrameTime())
	pb.Vel = rl.Vector2Add(pb.Vel, adjustedGravity)
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.Vel, rl.GetFrameTime())
	pb.Pos = rl.Vector2Add(pb.Pos, adjustedVel)
}
