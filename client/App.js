import React, { useEffect, useState } from 'react'
import Game from "./game/Game";

// make the article component
function MenuButton({value}){
    return(
        <article>
            {value}
        </article>
    );
}

// Leaderboard UI receiving live message data from WebSocket from App.js
function Leaderboard({leaderboard}){
    if(leaderboard != null){
        if(leaderboard[9]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[5].Username}</td>
                                <td>{leaderboard[5].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[6].Username}</td>
                                <td>{leaderboard[6].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[7].Username}</td>
                                <td>{leaderboard[7].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[8].Username}</td>
                                <td>{leaderboard[8].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[9].Username}</td>
                                <td>{leaderboard[9].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[8]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[5].Username}</td>
                                <td>{leaderboard[5].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[6].Username}</td>
                                <td>{leaderboard[6].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[7].Username}</td>
                                <td>{leaderboard[7].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[8].Username}</td>
                                <td>{leaderboard[8].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[7]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[5].Username}</td>
                                <td>{leaderboard[5].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[6].Username}</td>
                                <td>{leaderboard[6].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[7].Username}</td>
                                <td>{leaderboard[7].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[6]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[5].Username}</td>
                                <td>{leaderboard[5].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[6].Username}</td>
                                <td>{leaderboard[6].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[5]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[5].Username}</td>
                                <td>{leaderboard[5].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[4]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[4].Username}</td>
                                <td>{leaderboard[4].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[3]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[3].Username}</td>
                                <td>{leaderboard[3].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[2]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[2].Username}</td>
                                <td>{leaderboard[2].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[1]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                            <tr>
                                <td>{leaderboard[1].Username}</td>
                                <td>{leaderboard[1].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else if(leaderboard[0]){
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
                                <td>{leaderboard[0].Username}</td>
                                <td>{leaderboard[0].Wins}</td>
                            </tr>
                        </table>
                    </nav>
                </>
            );
        }
        else{
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
                        </table>
                    </nav>
                </>
            );
        }
    }
}

// Home Screen Component
function HomeScreen({leaderboard}){
    return(
        <div id = "strip">
            <header>
                <h1>Bit Battle 1.0.0</h1>
            </header>
            <br/>
            <section>
                <Leaderboard leaderboard={leaderboard}/>
                <MenuButton value="Start Game"/>
                <MenuButton value="Join Game"/>
            </section>
        </div>
    );
}

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
