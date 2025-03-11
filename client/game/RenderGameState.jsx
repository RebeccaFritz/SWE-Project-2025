import "./GameState";
import { useEffect } from "react";
import "./game.css";

/**
 * Given a game state, renders it.
 * @param {GameState} GameState
 * @constructor
 */
export default function RenderGameState({GameState}){
    useEffect(() => {
        const p1 = document.querySelector(".player");
        p1.style.left = GameState.player1.position[0] + "px";
        p1.style.top = GameState.player1.position[1] + "px";
        console.log(GameState.player1.position[0]);

        const p2 = document.querySelector(".a");
        p2.style.left = GameState.player2.position[0] + "px";
        p2.style.top = GameState.player2.position[1] + "px";
        console.log(GameState.player2.position[0]);
    }, []);

    return (
        <div className="board">
            <div className="player"></div>
            <div className="a"></div>
        </div>
    );
}