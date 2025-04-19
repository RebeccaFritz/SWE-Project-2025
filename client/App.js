import React, { useEffect, useState } from 'react'
import Game from "./Game";
import HomeScreen from './homescreen';

export default function App() {
    const [message, setMessage] = useState('');
    const [leaderboard, setLeaderboard] = useState([]);
    var [gameState, setGamestate] = useState(null);
    const [ws, setWS] = useState(null);

    useEffect(() => {
        // create WebSocket at the server port
        const socket = new WebSocket('ws://localhost:8080/ws');
        setWS(socket);

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
            //console.log('Message received: ', event.data);

            var serverMessage = JSON.parse(event.data);
            var msgType = serverMessage.MsgType;

            switch (msgType) {
               case "gamestate":
                  console.log(serverMessage.Gamestate)
                  setGamestate(serverMessage.Gamestate)
                  // socket.send(JSON.stringify({
                  //     MsgType: "pong",
                  //     Message: "recieved gamestate!"
                  // }))
                  return;
                case "test":
                    socket.send(JSON.stringify({
                        MsgType: "client",
                        Message: "We are live!"
                    }))
                    return;
                case "leaderboard":
                    setLeaderboard(serverMessage.Leaderboard);
                    socket.send(JSON.stringify({
                        MsgType: "client",
                        Message: "Leaderboard updated!"
                    }))
                    return;
                case "validate lobby code":
                    console.log('Message received: ', event.data);
                default:
                    setMessage(serverMessage);
                    // socket.send(JSON.stringify({
                    //     MsgType: "client",
                    //     Message: "Carry on"
                    // }))
                    // return;
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
            <HomeScreen socket={ws} leaderboard={leaderboard} />
            <div className="game" >
               <Game gameState={gameState}/>
            </div>
        </div>
    );
}
