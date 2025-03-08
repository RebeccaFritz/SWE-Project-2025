import './game.css';
import { useEffect } from 'react';

function generate_positions(n, board_width, player_width){
    let positions = []

    let pos = board_width/(n+1) - player_width/2

    for(let i = 1; i <= n; i++){
        positions.push(pos*i);
    }

    return positions
}

export default function Game(){
    useEffect(() =>{
        const player = document.querySelector('.player');
        const board = document.querySelector('.board');

        let positions = generate_positions(3, 500, 50);

        console.log(positions);

        if(!player || !board) return;

        let player_position = player.style.left;

        document.addEventListener('keydown', (e) => {
            if (e.key === 'a' && player_position > 0) {
                player_position--;
                player.style.left = positions[player_position] + 'px';
                console.log(positions[player_position], player_position)
            }

            if (e.key === 'd' && player_position < positions.length - 1) {
                player_position++;
                player.style.left = positions[player_position] + 'px';
                console.log(positions[player_position], player_position)
            }

        });
    }, []);

    return(
        <div className="board">
            <div className="player"></div>
        </div>
    );
}