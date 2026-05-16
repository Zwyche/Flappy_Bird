package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	PhysicsBody
	texture rl.Texture2D
	hitBox  rl.Rectangle
	size    float32
}

func NewBird(newTexture rl.Texture2D, newHitBox rl.Rectangle, newSize float32, newPos rl.Vector2, newVel rl.Vector2) Bird {
	pb := NewPhysicsBody(newPos, newVel)
	nb := Bird{texture: newTexture, hitBox: newHitBox, size: newSize, PhysicsBody: pb}

	return nb
}

func (b Bird) Draw() {
	rl.DrawRectangleRec(b.hitBox, rl.Blank)
	rl.DrawTexture(b.texture, int32(b.Pos.X), int32(b.Pos.Y), rl.White)
}

func (b *Bird) PhysicsUpdate() {
	b.PhysicsBody.PhysicsUpdate()
	b.hitBox.X = b.Pos.X
	b.hitBox.Y = b.Pos.Y
	b.hitBox.Width = b.size
	b.hitBox.Height = b.size
}

func (b *Bird) Reset(startPos rl.Vector2) {
	b.Pos = startPos
	b.Vel = rl.NewVector2(0, 0)
	b.hitBox.X = startPos.X
	b.hitBox.Y = startPos.Y
}
