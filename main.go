package main

import (
	"math/rand/v2"

	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func genTubeGap(topTubeRect *rl.Rectangle, bottomTubeRect *rl.Rectangle, tubeHeight float32, tubeGap float32) {
	screenHeight := rl.GetScreenHeight()
	margin := float32(25)

	minGap := margin + tubeGap/2
	maxGap := screenHeight - int(margin) - int(tubeGap)/2

	gap := minGap + rand.Float32()*(float32(maxGap)-minGap)

	topTubeRect.Y = gap - tubeGap/2 - tubeHeight
	bottomTubeRect.Y = gap + tubeGap/2
}

func main() {
	var birdSize float32 = 50
	var score int = 0
	var highScore int = score

	var topTubeY float32 = -290
	var bottomTubeY float32 = 260
	var tubeX float32 = 850
	var tubeHeight float32 = 400
	var tubeWidth float32 = 75
	var tubeSpeed float32 = 300
	var tubeGap float32 = 150

	var intersecting bool = false

	gameOverScreen := rl.Blank
	gameOverText := rl.Blank

	var farTriA []rl.Vector2 = make([]rl.Vector2, 10)
	var farTriB []rl.Vector2 = make([]rl.Vector2, 10)
	var farTriC []rl.Vector2 = make([]rl.Vector2, 10)
	var farSpeed float32 = 25

	var closeTriA []rl.Vector2 = make([]rl.Vector2, 10)
	var closeTriB []rl.Vector2 = make([]rl.Vector2, 10)
	var closeTriC []rl.Vector2 = make([]rl.Vector2, 10)
	var closeSpeed float32 = 50

	for x := range 10 {
		farTriA[x] = rl.NewVector2(float32(x*100), 375)
		farTriB[x] = rl.NewVector2(float32(x*100-100), 450)
		farTriC[x] = rl.NewVector2(float32(x*100+100), 450)
	}

	for x := range 6 {
		closeTriA[x] = rl.NewVector2(float32(x*200), 375)
		closeTriB[x] = rl.NewVector2(float32(x*200-200), 450)
		closeTriC[x] = rl.NewVector2(float32(x*200+200), 450)
	}

	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(800, 450, "Flappy Birb")

	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	birdRect := rl.NewRectangle(100, 100, birdSize, birdSize)
	birdTex := rl.LoadTexture("textures/flappyBirdbig.png")
	startBirdPos := rl.NewVector2(100, 100)
	topTubeRect := rl.NewRectangle(tubeX, 0, tubeWidth, tubeHeight)
	bottomTubeRect := rl.NewRectangle(tubeX, 0, tubeWidth, tubeHeight)
	genTubeGap(&topTubeRect, &bottomTubeRect, tubeHeight, tubeGap)
	gameOverRect := rl.NewRectangle(0, 0, 800, 450)

	bird := NewBird(birdTex, birdRect, 50, startBirdPos, rl.NewVector2(0, 0))
	bird.Gravity = rl.NewVector2(0, 800)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)

		for x := range 10 {
			rl.DrawTriangle(farTriA[x], farTriB[x], farTriC[x], rl.Beige)
		}

		for x := range 6 {
			rl.DrawTriangle(closeTriA[x], closeTriB[x], closeTriC[x], rl.Brown)
		}

		if rl.IsKeyPressed(rl.KeyR) {
			topTubeRect.X = tubeX
			topTubeRect.Y = topTubeY
			bottomTubeRect.X = tubeX
			bottomTubeRect.Y = bottomTubeY
			tubeSpeed = 300
			score = 0
			intersecting = false
			gameOverScreen = rl.Blank
			gameOverText = rl.Blank
			farSpeed = 25
			closeSpeed = 50

			bird.Reset(startBirdPos)
		}

		if !intersecting {

			bird.PhysicsUpdate()

			topTubeRect.X -= tubeSpeed * rl.GetFrameTime()
			bottomTubeRect.X -= tubeSpeed * rl.GetFrameTime()

			for x := range 10 {
				farTriA[x].X -= farSpeed * rl.GetFrameTime()
				farTriB[x].X -= farSpeed * rl.GetFrameTime()
				farTriC[x].X -= farSpeed * rl.GetFrameTime()
			}

			for x := range 6 {
				closeTriA[x].X -= closeSpeed * rl.GetFrameTime()
				closeTriB[x].X -= closeSpeed * rl.GetFrameTime()
				closeTriC[x].X -= closeSpeed * rl.GetFrameTime()
			}

		}

		for x := range 10 {
			if farTriC[x].X < 0 {
				farTriA[x].X = 900
				farTriB[x].X = 800
				farTriC[x].X = 1000
			}
		}

		for x := range 6 {
			if closeTriC[x].X < 0 {
				closeTriA[x].X = 1000
				closeTriB[x].X = 800
				closeTriC[x].X = 1200
			}
		}

		if topTubeRect.X < -75 {

			topTubeRect.X = tubeX
			bottomTubeRect.X = tubeX

			genTubeGap(&topTubeRect, &bottomTubeRect, tubeHeight, tubeGap)

			score++
			tubeSpeed += 10
			farSpeed += 2.5
			closeSpeed += 5

			if score > highScore {
				highScore = score
			}
		}

		if rl.IsKeyPressed(rl.KeySpace) && bird.hitBox.Y > 0 {
			bird.Vel = rl.NewVector2(0, -300)
		}

		if rl.CheckCollisionRecs(bird.hitBox, topTubeRect) || rl.CheckCollisionRecs(bird.hitBox, bottomTubeRect) {
			intersecting = true
		}

		if intersecting {
			gameOverScreen = rl.Red
			gameOverText = rl.Black
		}

		scoreStr := fmt.Sprint(score)
		highScoreStr := "High Score: " + fmt.Sprint(highScore)

		bird.Draw()
		rl.DrawRectangleRec(topTubeRect, rl.Green)
		rl.DrawRectangleRec(bottomTubeRect, rl.Green)
		rl.DrawText(scoreStr, 10, 10, 50, rl.Black)
		rl.DrawRectangleRec(gameOverRect, gameOverScreen)
		rl.DrawText("Game Over", 125, 125, 100, gameOverText)
		rl.DrawText("Press 'R' to restart", 123, 225, 50, gameOverText)
		rl.DrawText(highScoreStr, 200, 275, 50, gameOverText)

		rl.EndDrawing()
	}
}
