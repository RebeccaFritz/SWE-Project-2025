import React from 'react'
import p5 from 'p5';
import pixelfont from 'url:./PressStart2P-Regular.ttf'

function drawObj(p, obj){
   p.circle(obj.X, obj.Y, obj.Diameter)
}

function drawTargets(p, targets){
    p.fill(255, 255, 255);
    for (let i = 0; i < targets.length; i++){
       if (targets[i].IsEnabled){
        p.fill(0)
        p.circle(targets[i].X, targets[i].Y, targets[i].Diameter + 25)
        p.fill(255)
        p.textSize(14);
        let hex = "0x" + targets[i].Convert.toString(16)
        p.text(hex , targets[i].X -  p.textWidth(hex)/2 , targets[i].Y + 5);
       }
    }
}

function drawProjectiles(p, projectiles){
    p.stroke(0)
    p.strokeWeight(10)
    p.fill(255);
    for (let i = 0; i < projectiles.length; i++){
        if(projectiles[i].IsEnabled) {
            p.ellipse(projectiles[i].X + p.random(-5, 5), projectiles[i].Y + p.random(-5, 5), 20, 20);
            p.ellipse(projectiles[i].X + p.random(-5, 5), projectiles[i].Y + p.random(-5, 5), 20, 20);
            p.ellipse(projectiles[i].X + p.random(-5, 5), projectiles[i].Y + p.random(-5, 5), 20, 20);
        }
    }
    p.noStroke()
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
        y = playerY - 15
    }
    let size = 50
    p.noStroke();
    p.fill(0)
    p.beginShape();
    p.vertex(x, y);
    p.bezierVertex(x - size / 2, y - size / 2, x - size, y + size / 3, x, y + size);
    p.bezierVertex(x + size, y + size / 3, x + size / 2, y - size / 2, x, y);
    p.endShape();
    health.toString();
    p.fill(255)
    p.textSize(32);
    p.text(health, x- 9, y + 32);
}


const canvasWidth = 500
const canvasHeight = 600
let socket
let gs
let number = [0, 0, 0, 0, 0, 0, 0, 0]
let font

export default class Game extends React.Component{
    constructor(props) {
        super(props)
        this.myRef = React.createRef()
    }

    Sketch = (p) => {
        p.preload = () => {
            font = p.loadFont(pixelfont)
        }

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
            p.textFont('Futura');
            p.background(255);
            p.fill(255)
            p.rect(canvasWidth - 100, 0, 100, canvasHeight)

            p.stroke(0, 0, 0);
            p.strokeWeight(10); // Thick outline
            p.noFill()
            p.rect(0, 0, canvasWidth, canvasHeight); // Draw a rectangle

            p.stroke(0, 0, 0);
            p.strokeWeight(5);
            p.line(canvasWidth - 100, 0, canvasWidth-100, canvasHeight)

            // conversion
            p.fill(255)
            p.textSize(18);
            p.text(number.join(''), 405, canvasHeight/2);

            // player 1
            p.fill(0);
            p.circle(gs.Player1.X, gs.Player1.Y, gs.Player1.Diameter / 2);

            // player 2
            p.fill(255);
            p.circle(gs.Player2.X, gs.Player2.Y, gs.Player2.Diameter/ 2);

            drawHeart(p, canvasWidth+ 47 - 100, gs.Player1.Y, gs.Player1.Health)
            drawHeart(p, canvasWidth+ 47 - 100, gs.Player2.Y, gs.Player2.Health)

            drawProjectiles(p, gs.Projectiles);
            drawTargets(p, gs.Targets);

        }

        p.keyPressed = function() {
            let input;
            switch (p.keyCode){
                case 65:
                    input = "move_left";
                    number = [0, 0, 0, 0, 0, 0, 0, 0];
                    break;
                case 68:
                    input = "move_right"
                    number = [0, 0, 0, 0, 0, 0, 0, 0];
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
                    number = [0,0,0,0,0,0,0,0,]
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
