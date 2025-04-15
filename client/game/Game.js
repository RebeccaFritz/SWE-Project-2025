import React from 'react'
import p5 from 'p5';
import { GameState, Projectile } from "./gameState.js";
import { Player } from "./gameState.js";
import { Target } from "./gameState.js";

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
           if (this.props.gameState == null) {
              console.log("No gamestate to render")
              return
           }

          // console.log("recieved", this.props.gameState)

          let gs = this.props.gameState
            p.background(220);

            // player 1
            p.fill(0, 0 , 255);
            p.circle(gs.Player1.X, gs.Player1.Y, 100);

            // player 2
            p.fill(255, 0 , 0);
            p.circle(gs.Player2.X, gs.Player2.Y, 100);

            drawProjectiles(p, gs.Projectiles);
            drawTargets(p, gs.Targets);
        }

        // p.keyPressed = function() {
        //     if (p.keyCode === 65 && !(Player1.x - PLAYER_SPEED <= 0)) {
        //         Player1.x -= PLAYER_SPEED;
        //     } else if (p.keyCode === 68 && !(Player1.x + PLAYER_SPEED >= CANVAS_WIDTH)) {
        //         Player1.x += PLAYER_SPEED;
        //     }

        //     if (p.keyCode === 74 && !(Player2.x - PLAYER_SPEED <= 0)) {
        //         Player2.x -= PLAYER_SPEED;
        //     } else if (p.keyCode === 76 && !(Player2.x + PLAYER_SPEED >= CANVAS_WIDTH)) {
        //         Player2.x += PLAYER_SPEED;
        //     }

        //     if (p.keyCode === 83){ // ATTACK!
        //         let projectile = projectilePool.shift();
        //         projectile.isEnabled = true;
        //         projectile.x = Player1.x;
        //         projectile.y = Player1.y;
        //         projectile.velocity = -3;

        //         projectilePool.push(projectile)
        //     }

        //     if (p.keyCode === 75){ // ATTACK!
        //         let projectile = projectilePool.shift();
        //         projectile.isEnabled = true;
        //         projectile.x = Player2.x;
        //         projectile.y = Player2.y;
        //         projectile.velocity = 3;

        //         projectilePool.push(projectile)
        //     }

        //     return false
        // }
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
