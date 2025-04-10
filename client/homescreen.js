import React from "react";

// make the article component
function MenuButton({value}){
    return(
        <article>
            {value}
            <form name="Lobby">
                <input type="text"/><br/>
                <button type="button" onClick={buttonClick}>{value}</button>
            </form>
        </article>
    );
}

function buttonClick(){
    sendCode();
    console.log("Lobby code sent");
}

function sendCode(){
    if(MenuButton.value === "Start Game"){
        socket.send(JSON.stringify({
            MsgType: "create lobby code",
            lobbyCode: ($("#Lobby").serializeArray()),
        }))
        return;
    }
    else if(MenuButton.value === "Join Game"){
        socket.send(JSON.stringify({
            MsgType: "lobby code",
            lobbyCode: ($("#Lobby").serializeArray()),
        }))
    }
}

function leaderboardEntry(username, wins) {
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
            entries.push(leaderboardEntry(leaderboard[i].Username, leaderboard[i].Wins));
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
export default function HomeScreen({leaderboard}){
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