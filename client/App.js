import React, { useEffect, useState } from 'react'
import Game from "./game/Game";
import HomeScreen from './homescreen'; 

export default function App() {
    const [message, setMessage] = useState('');
    const [leaderboard, setLeaderboard] = useState(null);
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

            serverMessage = JSON.parse(event.data);
            msgType = serverMessage.MsgType;

            switch (msgType) {
                case "test":
                    return;
                case "leaderboard":
                    setLeaderboard(serverMessage.Leaderboard);
                    return;
                default:
                    setMessage(serverMessage);
                    return;
            }
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
        // display the client UI

        <div>
            <HomeScreen leaderboard={leaderboard} />
            <Game />
        </div>
    );
}
