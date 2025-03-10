import React, { useEffect, useState } from 'react'
import Game from "./game/Game";

export default function App() {
    const [message, setMessage] = useState('');
    const [ws, setWS] = useState(null);

    useEffect(() => {
        // create WebSocket at the server port
        const socket = new WebSocket('ws://localhost:8080/ws');

        // open WebSocket
        socket.onopen = () => {
            console.log('WebSocket connection established');
            socket.send(JSON.stringify({
                MsgType: "test",
                Message: "Hello!"
            }))
        };

        // if a message is received over WebSocket, parse the JSON and grab the .message
        socket.onmessage = (event) => {
            console.log('Message received: ', event.data);
            setMessage(JSON.parse(event.data).Message);
        };

        // handle severed connection
        socket.onclose = () => {
            console.log('Websocket connection closed');
        }

        setWS(socket); // add the WebSocket to the state

        return () => {
            socket.close();
        };
    }, []);

    return (
        // display the message from the server
        <div id="strip">
            <header>
                <h1>Bit Battle 1.0.0</h1>
            </header>
            <article>Create Game</article>
            <article>Join Game</article>
            <article>Leaderboard</article>
            <div> <Game /> </div>
        </div>

    );
}