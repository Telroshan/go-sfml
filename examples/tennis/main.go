package main

import (
	"math"
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/telroshan/go-sfml/v2/audio"
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const cTrue = 1
const cFalse = 0

const DegToRad = math.Pi / 180

func makeVector2(x float32, y float32) graphics.SfVector2f {
	v := graphics.NewSfVector2f()
	v.SetX(x)
	v.SetY(y)
	return v
}

func getNullIntRect() graphics.SfIntRect {
	return (graphics.SfIntRect)(graphics.SwigcptrSfIntRect(0))
}

func getNullRenderState() graphics.SfRenderStates {
	return (graphics.SfRenderStates)(graphics.SwigcptrSfRenderStates(0))
}

func init() { runtime.LockOSThread() }

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Define some constants
	const gameWidth = 800
	const gameHeight = 600
	const paddleWidth float32 = 25
	const paddleHeight float32 = 100
	const ballRadius float32 = 10
	const resourcesDir = "resources"

	// Create the window of the application
	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)
	vm.SetWidth(gameWidth)
	vm.SetHeight(gameHeight)
	vm.SetBitsPerPixel(32)
	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)
	w := graphics.SfRenderWindow_create(vm, "SFML Tennis", uint(window.SfTitlebar|window.SfClose), cs)
	defer window.SfWindow_destroy(w)
	window.SfWindow_setVerticalSyncEnabled(w, cTrue)

	// Load the sounds used in the game
	ballSoundBuffer := audio.SfSoundBuffer_createFromFile(path.Join(resourcesDir, "ball.wav"))
	if ballSoundBuffer == nil || ballSoundBuffer.Swigcptr() == cFalse {
		panic("Couldn't load ball.wav")
	}
	defer audio.SfSoundBuffer_destroy(ballSoundBuffer)
	ballSound := audio.SfSound_create()
	defer audio.SfSound_destroy(ballSound)
	audio.SfSound_setBuffer(ballSound, ballSoundBuffer)

	// Create the SFML logo texture:
	sfmlLogoTexture := graphics.SfTexture_createFromFile(path.Join(resourcesDir, "sfml_logo.png"), getNullIntRect())
	defer graphics.SfTexture_destroy(sfmlLogoTexture)
	if sfmlLogoTexture == nil || sfmlLogoTexture.Swigcptr() == cFalse {
		panic("Couldn't load sfml_logo.png")
	}
	sfmlLogo := graphics.SfSprite_create()
	defer graphics.SfSprite_destroy(sfmlLogo)
	graphics.SfSprite_setTexture(sfmlLogo, sfmlLogoTexture, cTrue)
	graphics.SfSprite_setPosition(sfmlLogo, makeVector2(170, 50))

	// Create the left paddle
	leftPaddle := graphics.SfRectangleShape_create()
	defer graphics.SfRectangleShape_destroy(leftPaddle)
	graphics.SfRectangleShape_setSize(leftPaddle, makeVector2(paddleWidth-3, paddleHeight-3))
	graphics.SfRectangleShape_setOutlineThickness(leftPaddle, 3)
	graphics.SfRectangleShape_setOutlineColor(leftPaddle, graphics.GetSfBlack())
	graphics.SfRectangleShape_setFillColor(leftPaddle, graphics.SfColor_fromRGB(100, 100, 200))
	graphics.SfRectangleShape_setOrigin(leftPaddle, makeVector2(paddleWidth/2, paddleHeight/2))

	// Create the right paddle
	rightPaddle := graphics.SfRectangleShape_create()
	defer graphics.SfRectangleShape_destroy(rightPaddle)
	graphics.SfRectangleShape_setSize(rightPaddle, makeVector2(paddleWidth-3, paddleHeight-3))
	graphics.SfRectangleShape_setOutlineThickness(rightPaddle, 3)
	graphics.SfRectangleShape_setOutlineColor(rightPaddle, graphics.GetSfBlack())
	graphics.SfRectangleShape_setFillColor(rightPaddle, graphics.SfColor_fromRGB(200, 100, 100))
	graphics.SfRectangleShape_setOrigin(rightPaddle, makeVector2(paddleWidth/2, paddleHeight/2))

	// Create the ball
	ball := graphics.SfCircleShape_create()
	defer graphics.SfCircleShape_destroy(ball)
	graphics.SfCircleShape_setRadius(ball, ballRadius-3)
	graphics.SfCircleShape_setOutlineThickness(ball, 2)
	graphics.SfCircleShape_setOutlineColor(ball, graphics.GetSfBlack())
	graphics.SfCircleShape_setFillColor(ball, graphics.GetSfWhite())
	graphics.SfCircleShape_setOrigin(ball, makeVector2(ballRadius/2, ballRadius/2))

	// Load the text font
	font := graphics.SfFont_createFromFile(path.Join(resourcesDir, "tuffy.ttf"))
	if font == nil || font.Swigcptr() == cFalse {
		panic("Couldn't load tuffy.ttf")
	}

	// Initialize the pause message
	pauseMessage := graphics.SfText_create()
	graphics.SfText_setFont(pauseMessage, font)
	graphics.SfText_setCharacterSize(pauseMessage, 40)
	graphics.SfText_setPosition(pauseMessage, makeVector2(170, 200))
	graphics.SfText_setFillColor(pauseMessage, graphics.GetSfWhite())
	graphics.SfText_setString(pauseMessage, "Welcome to SFML Tennis!\n\nPress space to start the game.")

	// Define the paddles properties
	const AITime = time.Millisecond * 100
	const paddleSpeed float32 = 400
	var rightPaddleSpeed float32 = 0
	const ballSpeed = 400
	var ballAngle float32 = 0
	const margin = 0.1
	const angleMaxSpread float32 = 20

	isPlaying := false

	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	aiTimer := time.Now()
	timeStamp := time.Now()

	for window.SfWindow_isOpen(w) > 0 {
		// Handle events
		for window.SfWindow_pollEvent(w, ev) > 0 {
			// Window closed or escape key pressed: exit
			if ev.GetEvType() == window.SfEventType(window.SfEvtClosed) ||
				(ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) && ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyEscape)) {
				window.SfWindow_close(w)
				break
			}

			// Space key pressed: play
			if ((ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) && ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeySpace)) ||
				ev.GetEvType() == window.SfEventType(window.SfEvtTouchBegan)) && !isPlaying {
				// (re)start the game
				isPlaying = true
				timeStamp = time.Now()

				// Reset the position of the paddles and ball
				graphics.SfRectangleShape_setPosition(leftPaddle, makeVector2(10+paddleWidth/2, gameHeight/2))
				graphics.SfRectangleShape_setPosition(rightPaddle, makeVector2(gameWidth-10-paddleWidth/2, gameHeight/2))
				graphics.SfCircleShape_setPosition(ball, makeVector2(gameWidth/2, gameHeight/2))

				// Reset the ball angle
				ballAngle = r1.Float32() * 360
				// Make sure the ball initial angle is not too much vertical
				for math.Abs(math.Cos(float64(ballAngle*DegToRad))) < 0.7 {
					ballAngle = r1.Float32() * 360
				}
			}

			// Window size changed, adjust view appropriately
			if ev.GetEvType() == window.SfEventType(window.SfEvtResized) {
				view := graphics.SfView_create()
				defer graphics.SfView_destroy(view)
				graphics.SfView_setSize(view, makeVector2(gameWidth, gameHeight))
				graphics.SfView_setCenter(view, makeVector2(gameWidth/2, gameHeight/2))
				graphics.SfRenderWindow_setView(w, view)
			}
		}

		if isPlaying {
			deltaTime := float32(time.Now().Sub(timeStamp).Seconds())
			timeStamp = time.Now()

			leftPaddlePos := graphics.SfRectangleShape_getPosition(leftPaddle)
			if window.SfKeyboard_isKeyPressed(window.SfKeyCode(window.SfKeyUp)) == cTrue &&
				leftPaddlePos.GetY()-paddleHeight/2 > 5 {
				graphics.SfRectangleShape_move(leftPaddle, makeVector2(0, -paddleSpeed*deltaTime))
			}
			if window.SfKeyboard_isKeyPressed(window.SfKeyCode(window.SfKeyDown)) == cTrue &&
				leftPaddlePos.GetY()+paddleHeight/2 < gameHeight-5 {
				graphics.SfRectangleShape_move(leftPaddle, makeVector2(0, paddleSpeed*deltaTime))
			}
			leftPaddlePos = graphics.SfRectangleShape_getPosition(leftPaddle)

			if window.SfTouch_isDown(0) == cTrue {
				pos := window.SfTouch_getPosition(0, w)
				mappedPos := graphics.SfRenderWindow_mapPixelToCoords(w, pos, graphics.SfRenderWindow_getView(w))
				graphics.SfRectangleShape_setPosition(leftPaddle, makeVector2(leftPaddlePos.GetX(), mappedPos.GetY()))
			}

			// Move the computer's paddle
			rightPaddlePos := graphics.SfRectangleShape_getPosition(rightPaddle)
			if rightPaddleSpeed < 0 && rightPaddlePos.GetY()-paddleHeight/2 > 5 ||
				rightPaddleSpeed > 0 && rightPaddlePos.GetY()+paddleHeight/2 < gameHeight-5 {
				graphics.SfRectangleShape_move(rightPaddle, makeVector2(0, rightPaddleSpeed*deltaTime))
			}
			rightPaddlePos = graphics.SfRectangleShape_getPosition(rightPaddle)

			// Update the computer's paddle direction according to the ball position
			ballPos := graphics.SfCircleShape_getPosition(ball)
			if time.Since(aiTimer) > AITime {
				aiTimer = time.Now()
				if ballPos.GetY()+ballRadius > rightPaddlePos.GetY()+paddleHeight/2 {
					rightPaddleSpeed = paddleSpeed
				} else if ballPos.GetY()-ballRadius < rightPaddlePos.GetY()-paddleHeight/2 {
					rightPaddleSpeed = -paddleSpeed
				} else {
					rightPaddleSpeed = 0
				}
			}

			// Move the ball
			xDirection := float32(math.Cos(float64(ballAngle * DegToRad)))
			yDirection := float32(math.Sin(float64(ballAngle * DegToRad)))
			graphics.SfCircleShape_move(ball, makeVector2(xDirection*ballSpeed*deltaTime, yDirection*ballSpeed*deltaTime))

			const inputString = "Press space to restart or\nescape to exit."

			// Check collisions between the ball and the screen
			if ballPos.GetX()-ballRadius < 0 {
				isPlaying = false
				graphics.SfText_setString(pauseMessage, "You Lost!\n\n"+inputString)
			} else if ballPos.GetX()+ballRadius > gameWidth {
				isPlaying = false
				graphics.SfText_setString(pauseMessage, "You Won!\n\n"+inputString)
			}

			if ballPos.GetY()-ballRadius < 0 {
				audio.SfSound_play(ballSound)
				ballAngle = -ballAngle
				graphics.SfCircleShape_setPosition(ball, makeVector2(ballPos.GetX(), ballRadius+margin))
			} else if ballPos.GetY()+ballRadius > gameHeight {
				audio.SfSound_play(ballSound)
				ballAngle = -ballAngle
				graphics.SfCircleShape_setPosition(ball, makeVector2(ballPos.GetX(), gameHeight-ballRadius-margin))
			}

			// Check the collisions between the ball and the paddles
			// Left Paddle
			if ballPos.GetX()-ballRadius < leftPaddlePos.GetX()+paddleWidth/2 &&
				ballPos.GetX()-ballRadius > leftPaddlePos.GetX() &&
				ballPos.GetY()+ballRadius >= leftPaddlePos.GetY()-paddleHeight/2 &&
				ballPos.GetY()-ballRadius <= leftPaddlePos.GetY()+paddleHeight/2 {
				if ballPos.GetY()+ballRadius > leftPaddlePos.GetY()+paddleHeight/2 {
					ballAngle = 180 - ballAngle + r1.Float32()*angleMaxSpread
				} else {
					ballAngle = 180 - ballAngle - r1.Float32()*angleMaxSpread
				}

				audio.SfSound_play(ballSound)
				graphics.SfCircleShape_setPosition(ball, makeVector2(leftPaddlePos.GetX()+ballRadius+paddleWidth/2+margin, ballPos.GetY()))
			}

			// Right Paddle
			if ballPos.GetX()+ballRadius > rightPaddlePos.GetX()-paddleWidth/2 &&
				ballPos.GetX()+ballRadius < rightPaddlePos.GetX() &&
				ballPos.GetY()+ballRadius >= rightPaddlePos.GetY()-paddleHeight/2 &&
				ballPos.GetY()-ballRadius <= rightPaddlePos.GetY()+paddleHeight/2 {
				if ballPos.GetY()+ballRadius > rightPaddlePos.GetY()+paddleHeight/2 {
					ballAngle = 180 - ballAngle + r1.Float32()*angleMaxSpread
				} else {
					ballAngle = 180 - ballAngle - r1.Float32()*angleMaxSpread
				}

				audio.SfSound_play(ballSound)
				graphics.SfCircleShape_setPosition(ball, makeVector2(rightPaddlePos.GetX()-ballRadius-paddleWidth/2-margin, ballPos.GetY()))
			}

			ballAngle = float32(int(ballAngle)%360) + (ballAngle - float32(int(ballAngle)))
		}

		// Clear the window
		graphics.SfRenderWindow_clear(w, graphics.SfColor_fromRGB(50, 50, 50))

		if isPlaying {
			// Draw the paddles and the ball
			graphics.SfRenderWindow_drawRectangleShape(w, leftPaddle, getNullRenderState())
			graphics.SfRenderWindow_drawRectangleShape(w, rightPaddle, getNullRenderState())
			graphics.SfRenderWindow_drawCircleShape(w, ball, getNullRenderState())
		} else {
			// Draw the pause message
			graphics.SfRenderWindow_drawText(w, pauseMessage, getNullRenderState())
			graphics.SfRenderWindow_drawSprite(w, sfmlLogo, getNullRenderState())
		}

		graphics.SfRenderWindow_display(w)
	}
}
