import React, { useEffect, useState } from 'react'
import ReactDom from 'react-dom'

export default function App() {
    const [message, setMessage] = useState('');
    const [ws, setWS] = useState(null);

    useEffect(() => {
        const socket = new WebSocket('ws://localhost:8080/ws');

        socket.onopen = () => {
            console.log('WebSocket connection established');
            socket.send(JSON.stringify({
                message: "Hello!"
            }))
        };

        socket.onmessage = (event) => {
            console.log('Message received: ', event.data);
            setMessage(JSON.parse(event.data).message);
        };

        socket.onclose = () => {
            console.log('Websocket connection closed');
        }

        setWS(socket);

        return () => {
            socket.close();
        };
    }, []);

    return (
        <div>WebSocket Client Received message: {message}</div>
    );
}