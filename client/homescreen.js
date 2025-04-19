import React from 'react'

// hiiiii

var lobbyData = '';
var socket = null;

// make the article component
function MenuButton({value}){
    function buttonClick(){
        sendCode(value);
    }
    function buttonInput(event) {
        lobbyData = event.target.value;
    }
    return(
        <article>
            {value}
            <form name="Lobby" onChange={buttonInput}>
                <input type="text" name="menuButton"/><br/>
                <button type="button" onClick={buttonClick}>{value}</button>
            </form>
        </article>
    );
}

function sendCode(value){
    if(value === "Start Game"){
        socket.send(JSON.stringify({
            MsgType: "create lobby code",
            LobbyCode: lobbyData,
        }))
    }
    else if(value === "Join Game"){
        socket.send(JSON.stringify({
            MsgType: "lobby code",
            LobbyCode: lobbyData,
        }))
    }
    console.log("Lobby code sent");
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

// Home Screen Component
function renderHomescreen({leaderboard}){
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

const Homescreen = (props) => {
    socket = props.socket
    return(
        <div id = "strip">
            <header>
                <h1>Bit Battle 1.0.0</h1>
            </header>
            <br/>
            <section>
                <Leaderboard leaderboard={props.leaderboard}/>
                <MenuButton value="Start Game"/>
                <MenuButton value="Join Game"/>
            </section>
        </div>
    );
}

export default Homescreen;
