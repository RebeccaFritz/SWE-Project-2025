import React, { useEffect, useState } from 'react'

// make the article component
function MenuButton({value}){
    return(
        <article>
            {value}
        </article>
    );
}

// Basic leaderboard table UI
// How the heck am I supposed to get the db.sqlite data into this???
function Leaderboard(){
    return(
        <>
            <nav>
                <table>
                    <tr>
                        <th colspan="2">Leaderboard</th>
                    </tr>
                    <tr>
                        <th>Username</th>
                        <th width="50%">Wins</th>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                    <tr>
                        <td>Data</td>
                        <td>Data</td>
                    </tr>
                </table>
            </nav>
        </>
    );
}

// make the home screen component
function HomeScreen(){
    return(
        <div id = "strip">
            <header>
                <h1>Bit Battle 1.0.0</h1>
            </header>
            <br/>
            <section>
                <Leaderboard />
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
        // display the client UI
        <HomeScreen />
    );
}
