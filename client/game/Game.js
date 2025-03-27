import React from 'react'
import p5 from 'p5';
import { GameState, Projectile } from "./gameState.js";
import { Player } from "./gameState.js";
import { Target } from "./gameState.js";


const CANVAS_WIDTH = 400;
const CANVAS_HEIGHT = 400;

const P1_START_X = CANVAS_WIDTH*0.5;
const P2_START_X = CANVAS_WIDTH*0.5;
const P1_START_Y =  CANVAS_HEIGHT*0.9;
const P2_START_Y =  CANVAS_HEIGHT*0.1;
const PLAYER_WIDTH = 25;
const PLAYER_SPEED = 50;

const TARGET_DIAMETER = 20;

const PROJECTILE_DIAMETER = 10;
const PROJECTILE_DIAMETER_FORCE_MULT = 0.04;

let MyGameState = new GameState(new Player, new Player, [], []);

let Player1 = new Player(P1_START_X, P1_START_Y, 100);
let Player2 = new Player(P2_START_X, P2_START_Y, 100);

let projectilePool;
let targetList;


export default class Game extends React.Component{
    constructor(props) {
        super(props)
        this.myRef = React.createRef()
    }

    Sketch = (p) => {

        p.setup = () => {
            p.createCanvas(CANVAS_WIDTH, CANVAS_HEIGHT);
            projectilePool = makeProjectilePool(20);
            targetList = initializeTargets(PLAYER_SPEED);
        }

        p.draw = () => {

            p.background(220);

            if(this.gameState == 1){
               p.circle(30, 40, 100);
            }


            p.fill(0, 0 , 255);
            p.circle(Player1.x, Player1.y, PLAYER_WIDTH);

            p.fill(255, 0 , 0);
            p.circle(Player2.x, Player2.y, PLAYER_WIDTH);

            drawProjectiles(projectilePool);
            drawTargets(targetList);
            checkCollisions(projectilePool, targetList);
        }

        p.keyPressed = function() {
            if (p.keyCode === 65 && !(Player1.x - PLAYER_SPEED <= 0)) {
                Player1.x -= PLAYER_SPEED;
            } else if (p.keyCode === 68 && !(Player1.x + PLAYER_SPEED >= CANVAS_WIDTH)) {
                Player1.x += PLAYER_SPEED;
            }

            if (p.keyCode === 74 && !(Player2.x - PLAYER_SPEED <= 0)) {
                Player2.x -= PLAYER_SPEED;
            } else if (p.keyCode === 76 && !(Player2.x + PLAYER_SPEED >= CANVAS_WIDTH)) {
                Player2.x += PLAYER_SPEED;
            }

            if (p.keyCode === 83){ // ATTACK!
                let projectile = projectilePool.shift();
                projectile.isEnabled = true;
                projectile.x = Player1.x;
                projectile.y = Player1.y;
                projectile.velocity = -3;

                projectilePool.push(projectile)
            }

            if (p.keyCode === 75){ // ATTACK!
                let projectile = projectilePool.shift();
                projectile.isEnabled = true;
                projectile.x = Player2.x;
                projectile.y = Player2.y;
                projectile.velocity = 3;

                projectilePool.push(projectile)
            }

            return false
        }


        /**
         * Returns a queue of n projectiles
         */
        function makeProjectilePool(length){
            let queue = [];

            for (let i = 0; i < length; i++){
                queue.push(new Projectile(0, 0, 0, false, PROJECTILE_DIAMETER));
            }

            return queue;
        }

        function initializeTargets(step){
            let targets = [];
            for(let i = step; i < CANVAS_WIDTH; i+=step){

                targets.push(new Target(i, CANVAS_HEIGHT/2, 0, true, TARGET_DIAMETER));

            }
            return targets;
        }


        function drawTargets(targets){
            p.fill(0, 200);
            for (let i = 0; i < targets.length; i++){
                if(targets[i].isEnabled){
                    p.circle(targets[i].x, targets[i].y, targets[i].diameter);
                    targets[i].y += targets[i].velocity
                }
            }
        }


        function drawProjectiles(projectilePool){
            p.fill(0);
            for (let i = 0; i < projectilePool.length; i++){
                if(projectilePool[i].isEnabled){
                    p.circle(projectilePool[i].x, projectilePool[i].y, projectilePool[i].diameter);
                    projectilePool[i].y += projectilePool[i].velocity
                }
            }
        }


        function checkCollisions(projectiles, targets){

            for(let i = 0; i < targets.length; i++){
                if(targets[i].isEnabled === false) continue;

                for(let j = 0; j < projectiles.length; j++){
                    if(projectiles[j].isEnabled === false) continue;

                    if(p.dist(targets[i].x, targets[i].y, projectiles[j].x, projectiles[j].y) < p.max(targets[i].diameter, projectiles[j].diameter)){

                        projectiles[j].isEnabled = false

                        targets[i].velocity += projectiles[j].velocity * PROJECTILE_DIAMETER_FORCE_MULT
                        projectiles[j].velocity = 0
                    }
                }
            }

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
