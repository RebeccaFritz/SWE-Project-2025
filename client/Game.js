import React from 'react'
import p5 from 'p5';

function drawObj(p, obj){
   p.circle(obj.X, obj.Y, obj.Diameter)
}

function drawTargets(p, targets){
    p.fill(0, 255, 0);
    for (let i = 0; i < targets.length; i++){
       if (targets[i].IsEnabled) drawObj(p, targets[i]);
    }
}

function drawProjectiles(p, projectiles){
    p.fill(0, 100, 0);
    for (let i = 0; i < projectiles.length; i++){
        if(projectiles[i].IsEnabled) drawObj(p, projectiles[i]);
    }
}

let socket
let gs

export default class Game extends React.Component{
    constructor(props) {
        super(props)
        this.myRef = React.createRef()
    }

    Sketch = (p) => {
        p.setup = () => {
            p.createCanvas(400, 400);
        }

        p.draw = () => {
           socket = this.props.socket
           gs = this.props.gameState
           if (gs == null) {
              console.log("No gamestate to render")
              return
           }

          // console.log("recieved", this.props.gameState)

            p.background(220);

            // player 1
            p.fill(255, 0 , 0);
            p.circle(gs.Player1.X, gs.Player1.Y, 100);

            // player 2
            p.fill(0, 0 , 255);
            p.circle(gs.Player2.X, gs.Player2.Y, 100);

            drawProjectiles(p, gs.Projectiles);
            drawTargets(p, gs.Targets);
        }

        p.keyPressed = function() {
            let input
            switch (p.keyCode){
                case 65:
                    input = "move_left";
                    break;
                case 68:
                    input = "move_right"
                    break;
                case 32:
                    input = "launch_projectile"
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
