import "./game.css";
import { useEffect } from "react";

/**
 * Generates a set of evenly spaced x values to be used as the possible x values of the player.
 * @param {int} n
 * @param {int} board_width
 * @param {int} player_width
 * @returns {list}
 */
function GeneratePositions(n, board_width, player_width) {
    let positions = [];

    let pos = board_width / (n + 1) - player_width / 2;

    for (let i = 1; i <= n; i++) {
        positions.push(pos * i);
    }

    return positions;
}

/**
 * Renders a game.
 * @returns {JSX.Element}
 */
export default function Game() {
    useEffect(() => {
        const player = document.querySelector(".player");
        const board = document.querySelector(".board");

        let positions = GeneratePositions(3, 500, 50);

        console.log(positions);

        if (!player || !board) return;

        let player_position = player.style.left;

        document.addEventListener("keydown", (e) => {
            if (e.key === "a" && player_position > 0) {
                player_position--;
                player.style.left = positions[player_position] + "px";
                console.log(positions[player_position], player_position);
            }

            if (e.key === "d" && player_position < positions.length - 1) {
                player_position++;
                player.style.left = positions[player_position] + "px";
                console.log(positions[player_position], player_position);
            }
        });
    }, []);

    return (
        <div className="board">
            <div className="player"></div>
        </div>
    );
}
