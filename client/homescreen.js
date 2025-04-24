import React from 'react'
import './homescreen.css'

var msgData = '';
var socket = null;

// make the Lobby Button component
function Button({value}){
    function buttonClick(){
        sendData(value);
    }
    function buttonInput(event) {
        msgData = event.target.value;
    }
    return(
        <article>
            {value}
            <form name={value} onChange={buttonInput}>
                <input type="text" name="menuButton"/><br/>
                <button type="button" onClick={buttonClick}>Submit</button>
            </form>
        </article>
    );
}

function sendData(value){
    if(value === "Start Game"){
        socket.send(JSON.stringify({
            MsgType: "create lobby code",
            LobbyCode: msgData,
        }))
        console.log("Lobby code sent");
    } else if(value === "Join Game"){
        socket.send(JSON.stringify({
            MsgType: "lobby code",
            LobbyCode: msgData,
        }))
        console.log("Lobby code sent");
    } else if (value === "Username") {
        socket.send(JSON.stringify({
            MsgType: "create username",
            Username: msgData,
        }))
        console.log("Username sent");
    }
}

function leaderboardEntry(username, wins, i) {
    return (
        <tr>
            <td>{username}</td>
            <td>{wins}</td>
        </tr>
    )
}

// Leaderboard UI receiving live message data from WebSocket from App.js
function Leaderboard({leaderboard}){
    if(leaderboard != null){
        const entries = []; // leaderboard entries
        for(let i = 0; i < leaderboard.length && i < 10; i++) {
            entries.push(leaderboardEntry(leaderboard[i].Username, leaderboard[i].Wins, i));
        }
        return(
            <>
                <nav>
                    <table>
                        <tbody>
                            <tr>
                                <th colSpan="2">Leaderboard</th>
                            </tr>
                            <tr>
                                <th>Username</th>
                                <th>Wins</th>
                            </tr>
                            {entries}
                        </tbody>
                    </table>
                </nav>
            </>
        );
    }
}

const Homescreen = (props) => {
    socket = props.socket
    return(
        <div id = "strip">
            <header>
                <h1 className="headerAccent">Bit Battle 1.0.0</h1>
            </header>
            <br/>
            <div>
                <div className="buttonSet" >
                    <Button value="Username"/>
                    <Button value="Start Game"/>
                    <Button value="Join Game"/>
                </div>
                <div className="leaderboard" >
                    <Leaderboard leaderboard={props.leaderboard}/>
                </div>
            </div>
        </div>
    );
}

export default Homescreen;
