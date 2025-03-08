import React, { useEffect, useState } from 'react'

// make the article component
function MenuButton({value}){
    return(
        <article>
            {value}
        </article>
    );
}

// Leaderboard UI receiving live message data from WebSocket from App.js
function Leaderboard({message}){
    if(message.MsgType === 'leaderboard'){
        return(
            <>
                <nav>
                    <table>
                        <tr>
                            <th colSpan="2">Leaderboard</th>
                        </tr>
                        <tr>
                            <th>Username</th>
                            <th>Wins</th>
                        </tr>
                        <tr>
                            <td>{message.Leaderboard[0].Username}</td>
                            <td>{message.Leaderboard[0].Wins}</td>
                        </tr>
                        <tr>
                            <td>{message.Leaderboard[1].Username}</td>
                            <td>{message.Leaderboard[1].Wins}</td>
                        </tr>
                        <tr>
                            <td>{message.Leaderboard[2].Username}</td>
                            <td>{message.Leaderboard[2].Wins}</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                        <tr>
                            <td>Null</td>
                            <td>Null</td>
                        </tr>
                    </table>
                </nav>
            </>
        );
    }
}

// Home Screen Component
function HomeScreen({message}){
    return(
        <div id = "strip">
            <header>
                <h1>Bit Battle 1.0.0</h1>
            </header>
            <br/>
            <section>
                <Leaderboard message = {message}/>
                <MenuButton value="Start Game"/>
                <MenuButton value="Join Game"/>
            </section>
        </div>
    );
}

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
            setMessage(JSON.parse(event.data));
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
        <HomeScreen message={message}/>
    );
}
