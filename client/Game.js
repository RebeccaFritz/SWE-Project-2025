import React from 'react'
import p5 from 'p5';

function drawObj(p, obj){
   p.circle(obj.X, obj.Y, obj.Diameter)
}

function drawTargets(p, targets){
    p.fill(0, 0, 0);
    for (let i = 0; i < targets.length; i++){
       if (targets[i].IsEnabled){
        drawObj(p, targets[i]);
        p.text("0x" + targets[i].Convert.toString(16), targets[i].X - 5, targets[i].Y - 8);
       }
    }
}

function drawProjectiles(p, projectiles){
    p.fill(0, 100, 0);
    for (let i = 0; i < projectiles.length; i++){
        if(projectiles[i].IsEnabled) drawObj(p, projectiles[i]);
    }
}

function drawHeart(p, x, playerY, health) {
    // arc parameters: x center, y center, width, height, start, stop, mode, detail
    // start and stop are the angles between which to draw the arc, arcs are drawn clockwise from start
    // angles are given in radians
    // mode is optional. determines the arc's fill style

    let y
    if (playerY > (canvasHeight /2)) {
        y = playerY-25
    } else {
        y = playerY+25
    }

    p.noStroke();
    p.fill(239, 149, 0);
    p.arc(x, y, x-(x+15), 40, p.PI, 0); // top r
    p.arc(x-15, y, x-(x+15), 40, p.PI, 0); // top l
    p.triangle(x-23, y, x+7, y, x-7, y+25); // bottom
    health.toString();
    p.textSize(32);
    p.text(health, x+20, y+15);
}

const canvasWidth = 500
const canvasHeight = 400
let socket
let gs
let number = [0, 0, 0, 0, 0, 0, 0, 0]

export default class Game extends React.Component{
    constructor(props) {
        super(props)
        this.myRef = React.createRef()
    }

    Sketch = (p) => {
        p.setup = () => {
            p.createCanvas(canvasWidth, canvasHeight);
        }

        p.draw = () => {
           socket = this.props.socket
           gs = this.props.gameState

           if (gs == null) {
              console.log("No gamestate to render")
              return
           }

           if (gs.Gameover == true) {
               console.log("Game Over.")
               p.clear()
               return
           }

          // console.log("recieved", this.props.gameState)

            p.background(220);
            p.textSize(10);
            p.text(number.join(''), 100, 100);

            // player 1
            p.fill(255, 0 , 0);
            p.circle(gs.Player1.X, gs.Player1.Y, 100);

            // player 2
            p.fill(0, 0 , 255);
            p.circle(gs.Player2.X, gs.Player2.Y, 100);


            drawProjectiles(p, gs.Projectiles);
            drawTargets(p, gs.Targets);

            drawHeart(p, canvasHeight+35, gs.Player1.Y, gs.Player1.Health)
            drawHeart(p, canvasHeight+35, gs.Player2.Y, gs.Player2.Health)
        }

        p.keyPressed = function() {
            let input;
            switch (p.keyCode){
                case 65:
                    input = "move_left";
                    break;
                case 68:
                    input = "move_right"
                    break;
                case 49:
                    if(number[0] == "0"){
                        number[0] = "1"
                    }
                    else if(number[0] == "1"){
                        number[0] = "0"
                    }
                    break;
                case 50:
                    if(number[1] == "0"){
                        number[1] = "1"
                    }
                    else if(number[1] == "1"){
                        number[1] = "0"
                    }
                    break;
                case 51:
                    if(number[2] == "0"){
                        number[2] = "1"
                    }
                    else if(number[2] == "1"){
                        number[2] = "0"
                    }
                    break;
                case 52:
                    if(number[3] == "0"){
                        number[3] = "1"
                    }
                    else if(number[3] == "1"){
                        number[3] = "0"
                    }
                    break;
                case 53:
                    if(number[4] == "0"){
                        number[4] = "1"
                    }
                    else if(number[4] == "1"){
                        number[4] = "0"
                    }
                    break;
                case 54:
                    if(number[5] == "0"){
                        number[5] = "1"
                    }
                    else if(number[5] == "1"){
                        number[5] = "0"
                    }
                    break;
                case 55:
                    if(number[6] == "0"){
                        number[6] = "1"
                    }
                    else if(number[6] == "1"){
                        number[6] = "0"
                    }
                    break;
                case 56:
                    if(number[7] == "0"){
                        number[7] = "1"
                    }
                    else if(number[7] == "1"){
                        number[7] = "0"
                    }
                    break;
                case 32:
                    input = number.join('')
                    console.log(input)
                    break;
                default: null
            }

            if (input == null || gs == null) return

            socket.send(JSON.stringify({
                MsgType: "input",
                Input: input
            }))
            input = null
        }
    }

    componentDidMount() {
        this.myP5 = new p5(this.Sketch, this.myRef.current)
    }

    render() {
        return (
            <div ref={this.myRef}>
            </div>
        )
    }
}
